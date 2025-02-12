package keeper

import (
	"context"

	collectortypes "github.com/comdex-official/comdex/x/collector/types"
	esmtypes "github.com/comdex-official/comdex/x/esm/types"
	rewardstypes "github.com/comdex-official/comdex/x/rewards/types"
	"github.com/comdex-official/comdex/x/vault/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ types.MsgServer = msgServer{}
)

type msgServer struct {
	Keeper
}

func NewMsgServer(keeper Keeper) types.MsgServer {
	return &msgServer{
		Keeper: keeper,
	}
}

// MsgCreate Creating a new CDP.
func (k msgServer) MsgCreate(c context.Context, msg *types.MsgCreateRequest) (*types.MsgCreateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	esmStatus, found := k.GetESMStatus(ctx, msg.AppId)
	status := false
	if found {
		status = esmStatus.Status
	}
	if status {
		return nil, esmtypes.ErrESMAlreadyExecuted
	}
	killSwitchParams, _ := k.GetKillSwitchData(ctx, msg.AppId)
	if killSwitchParams.BreakerEnable {
		return nil, esmtypes.ErrCircuitBreakerEnabled
	}
	extendedPairVault, found := k.GetPairsVault(ctx, msg.ExtendedPairVaultId)
	if !found {
		return nil, types.ErrorExtendedPairVaultDoesNotExists
	}
	pairData, found := k.GetPair(ctx, extendedPairVault.PairId)
	if !found {
		return nil, types.ErrorPairDoesNotExist
	}
	assetInData, found := k.GetAsset(ctx, pairData.AssetIn)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}
	assetOutData, found := k.GetAsset(ctx, pairData.AssetOut)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}

	appMapping, found := k.GetApp(ctx, msg.AppId)
	if !found {
		return nil, types.ErrorAppMappingDoesNotExist
	}

	if appMapping.Id != extendedPairVault.AppId {
		return nil, types.ErrorAppMappingIDMismatch
	}

	depositorAddress, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Checking if this is a stableMint pair or not  -- stableMintPair == psmPair
	if extendedPairVault.IsStableMintVault {
		return nil, types.ErrorCannotCreateStableMintVault
	}
	//Checking
	if !extendedPairVault.IsVaultActive {
		return nil, types.ErrorVaultCreationInactive
	}
	//if does then check app to extendedPair mapping has any vault key
	//if it does throw error
	_, userExists := k.GetUserAppExtendedPairMappingData(ctx, msg.From, msg.AppId, msg.ExtendedPairVaultId)

	if userExists {
		// _, alreadyExists := k.CheckUserAppToExtendedPairMapping(ctx, userVaultExtendedPairMapping, extendedPairVault.Id, appMapping.Id)
		return nil, types.ErrorUserVaultAlreadyExists
	}
	//Call CheckAppExtendedPairVaultMapping function to get counter - it also initialised the kv store if appMapping_id does not exists, or extendedPairVault_id does not exists.
	tokenMintedStatistics, _ := k.CheckAppExtendedPairVaultMapping(ctx, appMapping.Id, extendedPairVault.Id)
	//Check debt Floor
	if !msg.AmountOut.GTE(extendedPairVault.DebtFloor) {
		return nil, types.ErrorAmountOutLessThanDebtFloor
	}
	//Check Debt Ceil
	currentMintedStatistics := tokenMintedStatistics.Add(msg.AmountOut)

	if currentMintedStatistics.GT(extendedPairVault.DebtCeiling) {
		return nil, types.ErrorAmountOutGreaterThanDebtCeiling
	}

	//Calculate CR - make necessary changes to calculate collateralization function
	if err := k.VerifyCollaterlizationRatio(ctx, extendedPairVault.Id, msg.AmountIn, msg.AmountOut, extendedPairVault.MinCr, status); err != nil {
		return nil, err
	}
	//Take amount from user
	if msg.AmountIn.GT(sdk.ZeroInt()) {
		if err := k.SendCoinFromAccountToModule(ctx, depositorAddress, types.ModuleName, sdk.NewCoin(assetInData.Denom, msg.AmountIn)); err != nil {
			return nil, err
		}
	}

	//Mint Tokens for user
	if err := k.MintCoin(ctx, types.ModuleName, sdk.NewCoin(assetOutData.Denom, msg.AmountOut)); err != nil {
		return nil, err
	}

	//Send Fees to Accumulator
	//Deducting Opening Fee if 0 opening fee then act accordingly
	if extendedPairVault.DrawDownFee.IsZero() && msg.AmountOut.GT(sdk.ZeroInt()) { //Send Rest to user
		if err := k.SendCoinFromModuleToAccount(ctx, types.ModuleName, depositorAddress, sdk.NewCoin(assetOutData.Denom, msg.AmountOut)); err != nil {
			return nil, err
		}
	} else {
		//If not zero deduct send to collector//////////
		//one approach could be
		collectorShare := (msg.AmountOut.Mul(sdk.Int(extendedPairVault.DrawDownFee))).Quo(sdk.Int(sdk.OneDec()))

		if collectorShare.GT(sdk.ZeroInt()) {
			if err := k.SendCoinFromModuleToModule(ctx, types.ModuleName, collectortypes.ModuleName, sdk.NewCoins(sdk.NewCoin(assetOutData.Denom, collectorShare))); err != nil {
				return nil, err
			}

			err := k.UpdateCollector(ctx, appMapping.Id, pairData.AssetOut, sdk.ZeroInt(), sdk.ZeroInt(), collectorShare, sdk.ZeroInt())
			if err != nil {
				return nil, err
			}
		}

		// and send the rest to the user
		amountToUser := msg.AmountOut.Sub(collectorShare)
		if amountToUser.GT(sdk.ZeroInt()) {
			if err := k.SendCoinFromModuleToAccount(ctx, types.ModuleName, depositorAddress, sdk.NewCoin(assetOutData.Denom, amountToUser)); err != nil {
				return nil, err
			}
		}
	}
	blockHeight := ctx.BlockHeight()
	blockTime := ctx.BlockTime()
	if extendedPairVault.StabilityFee.IsZero() {
		blockHeight = 0
	}

	//If all correct  create vault
	oldID := k.GetIDForVault(ctx)
	zeroVal := sdk.ZeroInt()
	var newVault types.Vault
	updatedID := oldID + 1
	newVault.Id = updatedID
	newVault.AmountIn = msg.AmountIn

	closingFeeVal := msg.AmountOut.Mul(sdk.Int(extendedPairVault.ClosingFee)).Quo(sdk.Int(sdk.OneDec()))

	newVault.ClosingFeeAccumulated = closingFeeVal
	newVault.AmountOut = msg.AmountOut
	newVault.AppId = appMapping.Id
	newVault.InterestAccumulated = zeroVal
	newVault.Owner = msg.From
	newVault.CreatedAt = ctx.BlockTime()
	newVault.BlockHeight = blockHeight
	newVault.BlockTime = blockTime
	newVault.ExtendedPairVaultID = extendedPairVault.Id

	k.SetVault(ctx, newVault)
	k.SetIDForVault(ctx, updatedID)

	//Update mapping data - take proper approach
	// lookup table already exists
	//only need to update counter and token statistics value
	k.UpdateAppExtendedPairVaultMappingDataOnMsgCreate(ctx, newVault)

	var mappingData types.OwnerAppExtendedPairVaultMappingData
	mappingData.Owner = msg.From
	mappingData.AppId = msg.AppId
	mappingData.ExtendedPairId = msg.ExtendedPairVaultId
	mappingData.VaultId = newVault.Id

	k.SetUserAppExtendedPairMappingData(ctx, mappingData)

	ctx.GasMeter().ConsumeGas(types.CreateVaultGas, "CreateVaultGas")

	return &types.MsgCreateResponse{}, nil
}

// MsgDeposit Only for depositing new collateral.
func (k msgServer) MsgDeposit(c context.Context, msg *types.MsgDepositRequest) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	esmStatus, found := k.GetESMStatus(ctx, msg.AppId)
	status := false
	if found {
		status = esmStatus.Status
	}
	if status {
		return nil, esmtypes.ErrESMAlreadyExecuted
	}
	klwsParams, _ := k.GetKillSwitchData(ctx, msg.AppId)
	if klwsParams.BreakerEnable {
		return nil, esmtypes.ErrCircuitBreakerEnabled
	}
	depositor, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	//checks if extended pair exists
	extendedPairVault, found := k.GetPairsVault(ctx, msg.ExtendedPairVaultId)
	if !found {
		return nil, types.ErrorExtendedPairVaultDoesNotExists
	}
	pairData, found := k.GetPair(ctx, extendedPairVault.PairId)
	if !found {
		return nil, types.ErrorPairDoesNotExist
	}
	assetInData, found := k.GetAsset(ctx, pairData.AssetIn)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}

	//Checking if appMapping_id exists
	appMapping, found := k.GetApp(ctx, msg.AppId)
	if !found {
		return nil, types.ErrorAppMappingDoesNotExist
	}
	//Checking if vault access disabled
	if !extendedPairVault.IsVaultActive {
		return nil, types.ErrorVaultInactive
	}

	//Checking if the appMapping_id in the msg_create & extendedPairVault_are same or not
	if appMapping.Id != extendedPairVault.AppId {
		return nil, types.ErrorAppMappingIDMismatch
	}

	userVault, found := k.GetVault(ctx, msg.UserVaultId)
	if !found {
		return nil, types.ErrorVaultDoesNotExist
	}
	if userVault.Owner != msg.From {
		return nil, types.ErrVaultAccessUnauthorised
	}

	if appMapping.Id != userVault.AppId {
		return nil, types.ErrorInvalidAppMappingData
	}
	if extendedPairVault.Id != userVault.ExtendedPairVaultID {
		return nil, types.ErrorInvalidExtendedPairMappingData
	}

	totalDebt := userVault.AmountOut.Add(userVault.InterestAccumulated)
	err1 := k.CalculateVaultInterest(ctx, appMapping.Id, msg.ExtendedPairVaultId, msg.UserVaultId, totalDebt, userVault.BlockHeight, userVault.BlockTime.Unix())
	if err1 != nil {
		return nil, err1
	}
	userVault, found1 := k.GetVault(ctx, msg.UserVaultId)
	if !found1 {
		return nil, types.ErrorVaultDoesNotExist
	}
	userVault.AmountIn = userVault.AmountIn.Add(msg.Amount)
	if !userVault.AmountIn.IsPositive() {
		return nil, types.ErrorInvalidAmount
	}

	if msg.Amount.GT(sdk.ZeroInt()) {
		if err := k.SendCoinFromAccountToModule(ctx, depositor, types.ModuleName, sdk.NewCoin(assetInData.Denom, msg.Amount)); err != nil {
			return nil, err
		}
	}
	userVault.BlockHeight = ctx.BlockHeight()
	userVault.BlockTime = ctx.BlockTime()

	k.SetVault(ctx, userVault)
	//Updating appExtendedPairvaultMappingData data -
	appExtendedPairVaultData, _ := k.GetAppExtendedPairVaultMappingData(ctx, appMapping.Id, msg.ExtendedPairVaultId)
	k.UpdateCollateralLockedAmountLockerMapping(ctx, appExtendedPairVaultData.AppId, appExtendedPairVaultData.ExtendedPairId, msg.Amount, true)

	ctx.GasMeter().ConsumeGas(types.DepositVaultGas, "DepositVaultGas")

	return &types.MsgDepositResponse{}, nil
}

// MsgWithdraw Withdrawing collateral.
func (k msgServer) MsgWithdraw(c context.Context, msg *types.MsgWithdrawRequest) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	killSwitchParams, _ := k.GetKillSwitchData(ctx, msg.AppId)
	if killSwitchParams.BreakerEnable {
		return nil, esmtypes.ErrCircuitBreakerEnabled
	}
	esmStatus, found := k.GetESMStatus(ctx, msg.AppId)
	status := false
	if found {
		status = esmStatus.Status
	}

	if ctx.BlockTime().After(esmStatus.EndTime) && status {
		return nil, esmtypes.ErrCoolOffPeriodPassed
	}

	depositor, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	//checks if extended pair exists
	extendedPairVault, found := k.GetPairsVault(ctx, msg.ExtendedPairVaultId)
	if !found {
		return nil, types.ErrorExtendedPairVaultDoesNotExists
	}
	pairData, found := k.GetPair(ctx, extendedPairVault.PairId)
	if !found {
		return nil, types.ErrorPairDoesNotExist
	}
	assetInData, found := k.GetAsset(ctx, pairData.AssetIn)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}

	//Checking if appMapping_id exists
	appMapping, found := k.GetApp(ctx, msg.AppId)
	if !found {
		return nil, types.ErrorAppMappingDoesNotExist
	}
	//Checking if vault access disabled
	if !extendedPairVault.IsVaultActive {
		return nil, types.ErrorVaultInactive
	}
	//Checking if the appMapping_id in the msg_create & extendedPairVault_are same or not
	if appMapping.Id != extendedPairVault.AppId {
		return nil, types.ErrorAppMappingIDMismatch
	}

	userVault, found := k.GetVault(ctx, msg.UserVaultId)
	if !found {
		return nil, types.ErrorVaultDoesNotExist
	}
	if userVault.Owner != msg.From {
		return nil, types.ErrVaultAccessUnauthorised
	}

	if appMapping.Id != userVault.AppId {
		return nil, types.ErrorInvalidAppMappingData
	}
	if extendedPairVault.Id != userVault.ExtendedPairVaultID {
		return nil, types.ErrorInvalidExtendedPairMappingData
	}

	totalDebt := userVault.AmountOut.Add(userVault.InterestAccumulated)
	err1 := k.CalculateVaultInterest(ctx, appMapping.Id, msg.ExtendedPairVaultId, msg.UserVaultId, totalDebt, userVault.BlockHeight, userVault.BlockTime.Unix())
	if err1 != nil {
		return nil, err1
	}

	userVault, found1 := k.GetVault(ctx, msg.UserVaultId)
	if !found1 {
		return nil, types.ErrorVaultDoesNotExist
	}
	userVault.AmountIn = userVault.AmountIn.Sub(msg.Amount)
	if !userVault.AmountIn.IsPositive() {
		return nil, types.ErrorInvalidAmount
	}

	totalDebtCalculation := userVault.AmountOut.Add(userVault.InterestAccumulated)
	totalDebtCalculation = totalDebtCalculation.Add(userVault.ClosingFeeAccumulated)

	//Calculate CR - make necessary changes to the calculate collateralization function
	if err := k.VerifyCollaterlizationRatio(ctx, extendedPairVault.Id, userVault.AmountIn, totalDebtCalculation, extendedPairVault.MinCr, status); err != nil {
		return nil, err
	}
	if msg.Amount.GT(sdk.ZeroInt()) {
		if err := k.SendCoinFromModuleToAccount(ctx, types.ModuleName, depositor, sdk.NewCoin(assetInData.Denom, msg.Amount)); err != nil {
			return nil, err
		}
	}
	userVault.BlockHeight = ctx.BlockHeight()
	userVault.BlockTime = ctx.BlockTime()
	k.SetVault(ctx, userVault)

	//Updating appExtendedPairVaultMappingData
	appExtendedPairVaultData, _ := k.GetAppExtendedPairVaultMappingData(ctx, appMapping.Id, msg.ExtendedPairVaultId)
	k.UpdateCollateralLockedAmountLockerMapping(ctx, appExtendedPairVaultData.AppId, appExtendedPairVaultData.ExtendedPairId, msg.Amount, false)

	ctx.GasMeter().ConsumeGas(types.WithdrawVaultGas, "WithdrawVaultGas")

	return &types.MsgWithdrawResponse{}, nil
}

// MsgDraw To borrow more amount.
func (k msgServer) MsgDraw(c context.Context, msg *types.MsgDrawRequest) (*types.MsgDrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	esmStatus, found := k.GetESMStatus(ctx, msg.AppId)
	status := false
	if found {
		status = esmStatus.Status
	}
	if status {
		return nil, esmtypes.ErrESMAlreadyExecuted
	}
	killSwitchParams, _ := k.GetKillSwitchData(ctx, msg.AppId)
	if killSwitchParams.BreakerEnable {
		return nil, esmtypes.ErrCircuitBreakerEnabled
	}
	depositor, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	//checks if extended pair exists
	extendedPairVault, found := k.GetPairsVault(ctx, msg.ExtendedPairVaultId)
	if !found {
		return nil, types.ErrorExtendedPairVaultDoesNotExists
	}
	pairData, found := k.GetPair(ctx, extendedPairVault.PairId)
	if !found {
		return nil, types.ErrorPairDoesNotExist
	}
	assetOutData, found := k.GetAsset(ctx, pairData.AssetOut)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}

	//Checking if appMapping_id exists
	appMapping, found := k.GetApp(ctx, msg.AppId)
	if !found {
		return nil, types.ErrorAppMappingDoesNotExist
	}
	//Checking if vault access disabled
	if !extendedPairVault.IsVaultActive {
		return nil, types.ErrorVaultInactive
	}
	//Checking if the appMapping_id in the msg_create & extendedPairVault_are same or not
	if appMapping.Id != extendedPairVault.AppId {
		return nil, types.ErrorAppMappingIDMismatch
	}

	userVault, found := k.GetVault(ctx, msg.UserVaultId)
	if !found {
		return nil, types.ErrorVaultDoesNotExist
	}
	if userVault.Owner != msg.From {
		return nil, types.ErrVaultAccessUnauthorised
	}

	if appMapping.Id != userVault.AppId {
		return nil, types.ErrorInvalidAppMappingData
	}
	if extendedPairVault.Id != userVault.ExtendedPairVaultID {
		return nil, types.ErrorInvalidExtendedPairMappingData
	}
	if msg.Amount.LTE(sdk.NewInt(0)) {
		return nil, types.ErrorInvalidAmount
	}

	totalCalDebt := userVault.AmountOut.Add(userVault.InterestAccumulated)
	err1 := k.CalculateVaultInterest(ctx, appMapping.Id, msg.ExtendedPairVaultId, msg.UserVaultId, totalCalDebt, userVault.BlockHeight, userVault.BlockTime.Unix())
	if err1 != nil {
		return nil, err1
	}

	userVault, found1 := k.GetVault(ctx, msg.UserVaultId)
	if !found1 {
		return nil, types.ErrorVaultDoesNotExist
	}

	newUpdatedAmountOut := userVault.AmountOut.Add(msg.Amount)
	totalDebt := newUpdatedAmountOut.Add(userVault.InterestAccumulated)
	totalDebt = totalDebt.Add(userVault.ClosingFeeAccumulated)

	tokenMintedStatistics, _ := k.CheckAppExtendedPairVaultMapping(ctx, appMapping.Id, extendedPairVault.Id)

	//Check Debt Ceil
	currentMintedStatistics := tokenMintedStatistics.Add(msg.Amount)

	if currentMintedStatistics.GTE(extendedPairVault.DebtCeiling) {
		return nil, types.ErrorAmountOutGreaterThanDebtCeiling
	}

	if err := k.VerifyCollaterlizationRatio(ctx, extendedPairVault.Id, userVault.AmountIn, totalDebt, extendedPairVault.MinCr, status); err != nil {
		return nil, err
	}

	if err := k.MintCoin(ctx, types.ModuleName, sdk.NewCoin(assetOutData.Denom, msg.Amount)); err != nil {
		return nil, err
	}

	if extendedPairVault.DrawDownFee.IsZero() && msg.Amount.GT(sdk.ZeroInt()) {
		//Send Rest to user
		if err := k.SendCoinFromModuleToAccount(ctx, types.ModuleName, depositor, sdk.NewCoin(assetOutData.Denom, msg.Amount)); err != nil {
			return nil, err
		}
	} else {
		//If not zero deduct send to collector//////////
		//one approach could be
		collectorShare := (msg.Amount.Mul(sdk.Int(extendedPairVault.DrawDownFee))).Quo(sdk.Int(sdk.OneDec()))

		if collectorShare.GT(sdk.ZeroInt()) {
			if err := k.SendCoinFromModuleToModule(ctx, types.ModuleName, collectortypes.ModuleName, sdk.NewCoins(sdk.NewCoin(assetOutData.Denom, collectorShare))); err != nil {
				return nil, err
			}

			err := k.UpdateCollector(ctx, appMapping.Id, pairData.AssetOut, sdk.ZeroInt(), sdk.ZeroInt(), collectorShare, sdk.ZeroInt())
			if err != nil {
				return nil, err
			}
		}
		// and send the rest to the user
		amountToUser := msg.Amount.Sub(collectorShare)
		if amountToUser.GT(sdk.ZeroInt()) {
			if err := k.SendCoinFromModuleToAccount(ctx, types.ModuleName, depositor, sdk.NewCoin(assetOutData.Denom, amountToUser)); err != nil {
				return nil, err
			}
		}
	}

	userVault.AmountOut = userVault.AmountOut.Add(msg.Amount)
	userVault.BlockHeight = ctx.BlockHeight()
	userVault.BlockTime = ctx.BlockTime()
	k.SetVault(ctx, userVault)

	//Updating appExtendedPairVaultMappingData
	appExtendedPairVaultData, _ := k.GetAppExtendedPairVaultMappingData(ctx, appMapping.Id, msg.ExtendedPairVaultId)
	k.UpdateTokenMintedAmountLockerMapping(ctx, appExtendedPairVaultData.AppId, appExtendedPairVaultData.ExtendedPairId, msg.Amount, true)

	ctx.GasMeter().ConsumeGas(types.DrawVaultGas, "DrawVaultGas")

	return &types.MsgDrawResponse{}, nil
}

func (k msgServer) MsgRepay(c context.Context, msg *types.MsgRepayRequest) (*types.MsgRepayResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	esmStatus, found := k.GetESMStatus(ctx, msg.AppId)
	status := false
	if found {
		status = esmStatus.Status
	}
	if status {
		return nil, esmtypes.ErrESMAlreadyExecuted
	}
	killSwitchParams, _ := k.GetKillSwitchData(ctx, msg.AppId)
	if killSwitchParams.BreakerEnable {
		return nil, esmtypes.ErrCircuitBreakerEnabled
	}
	depositor, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	//checks if extended pair exists
	extendedPairVault, found := k.GetPairsVault(ctx, msg.ExtendedPairVaultId)
	if !found {
		return nil, types.ErrorExtendedPairVaultDoesNotExists
	}
	pairData, found := k.GetPair(ctx, extendedPairVault.PairId)
	if !found {
		return nil, types.ErrorPairDoesNotExist
	}
	assetOutData, found := k.GetAsset(ctx, pairData.AssetOut)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}

	//Checking if appMapping_id exists
	appMapping, found := k.GetApp(ctx, msg.AppId)
	if !found {
		return nil, types.ErrorAppMappingDoesNotExist
	}
	//Checking if vault acccess disabled

	//Checking if the appMapping_id in the msg_create & extendedPairVault_are same or not
	if appMapping.Id != extendedPairVault.AppId {
		return nil, types.ErrorAppMappingIDMismatch
	}

	userVault, found := k.GetVault(ctx, msg.UserVaultId)
	if !found {
		return nil, types.ErrorVaultDoesNotExist
	}
	if userVault.Owner != msg.From {
		return nil, types.ErrVaultAccessUnauthorised
	}

	if appMapping.Id != userVault.AppId {
		return nil, types.ErrorInvalidAppMappingData
	}
	if extendedPairVault.Id != userVault.ExtendedPairVaultID {
		return nil, types.ErrorInvalidExtendedPairMappingData
	}
	if msg.Amount.LTE(sdk.NewInt(0)) {
		return nil, types.ErrorInvalidAmount
	}

	totalDebt := userVault.AmountOut.Add(userVault.InterestAccumulated)
	err1 := k.CalculateVaultInterest(ctx, appMapping.Id, msg.ExtendedPairVaultId, msg.UserVaultId, totalDebt, userVault.BlockHeight, userVault.BlockTime.Unix())
	if err1 != nil {
		return nil, err1
	}

	userVault, found1 := k.GetVault(ctx, msg.UserVaultId)
	if !found1 {
		return nil, types.ErrorVaultDoesNotExist
	}

	newAmount := userVault.AmountOut.Add(userVault.InterestAccumulated)
	newAmount = newAmount.Sub(msg.Amount)
	if newAmount.LT(sdk.NewInt(0)) {
		return nil, types.ErrorInvalidAmount
	}

	if msg.Amount.LTE(userVault.InterestAccumulated) {
		//Amount is less than equal to the interest accumulated
		//subtract that as interest
		reducedFees := userVault.InterestAccumulated.Sub(msg.Amount)
		userVault.InterestAccumulated = reducedFees
		//and send it to the collector module
		if msg.Amount.GT(sdk.ZeroInt()) {
			if err := k.SendCoinFromAccountToModule(ctx, depositor, types.ModuleName, sdk.NewCoin(assetOutData.Denom, msg.Amount)); err != nil {
				return nil, err
			}
			//			SEND TO COLLECTOR- msg.Amount
			if err := k.SendCoinFromModuleToModule(ctx, types.ModuleName, collectortypes.ModuleName, sdk.NewCoins(sdk.NewCoin(assetOutData.Denom, msg.Amount))); err != nil {
				return nil, err
			}
			err := k.UpdateCollector(ctx, appMapping.Id, pairData.AssetOut, msg.Amount, sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt())
			if err != nil {
				return nil, err
			}
		}
		userVault.BlockHeight = ctx.BlockHeight()
		userVault.BlockTime = ctx.BlockTime()
		k.SetVault(ctx, userVault)
	} else {
		updatedUserSentAmountAfterFeesDeduction := msg.Amount.Sub(userVault.InterestAccumulated)

		updatedUserDebt := userVault.AmountOut.Sub(updatedUserSentAmountAfterFeesDeduction)

		// //If user's closing fees is a bigger amount than the debt floor, user will not close the debt floor

		if !updatedUserDebt.GTE(extendedPairVault.DebtFloor) {
			return nil, types.ErrorAmountOutLessThanDebtFloor
		}
		if msg.Amount.GT(sdk.ZeroInt()) {
			if err := k.SendCoinFromAccountToModule(ctx, depositor, types.ModuleName, sdk.NewCoin(assetOutData.Denom, msg.Amount)); err != nil {
				return nil, err
			}
		}
		if updatedUserSentAmountAfterFeesDeduction.GT(sdk.ZeroInt()) {
			if err := k.BurnCoin(ctx, types.ModuleName, sdk.NewCoin(assetOutData.Denom, updatedUserSentAmountAfterFeesDeduction)); err != nil {
				return nil, err
			}
		}
		//			SEND TO COLLECTOR----userVault.InterestAccumulated
		if userVault.InterestAccumulated.GT(sdk.ZeroInt()) {
			if err := k.SendCoinFromModuleToModule(ctx, types.ModuleName, collectortypes.ModuleName, sdk.NewCoins(sdk.NewCoin(assetOutData.Denom, userVault.InterestAccumulated))); err != nil {
				return nil, err
			}
			err := k.UpdateCollector(ctx, appMapping.Id, pairData.AssetOut, userVault.InterestAccumulated, sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt())
			if err != nil {
				return nil, err
			}
		}

		userVault.AmountOut = updatedUserDebt
		zeroVal := sdk.ZeroInt()
		userVault.InterestAccumulated = zeroVal
		userVault.BlockHeight = ctx.BlockHeight()
		userVault.BlockTime = ctx.BlockTime()
		k.SetVault(ctx, userVault)
		appExtendedPairVaultData, _ := k.GetAppExtendedPairVaultMappingData(ctx, appMapping.Id, msg.ExtendedPairVaultId)
		k.UpdateTokenMintedAmountLockerMapping(ctx, appExtendedPairVaultData.AppId, appExtendedPairVaultData.ExtendedPairId, updatedUserSentAmountAfterFeesDeduction, false)
	}

	ctx.GasMeter().ConsumeGas(types.RepayVaultGas, "RepayVaultGas")

	return &types.MsgRepayResponse{}, nil
}

func (k msgServer) MsgClose(c context.Context, msg *types.MsgCloseRequest) (*types.MsgCloseResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	esmStatus, found := k.GetESMStatus(ctx, msg.AppId)
	status := false
	if found {
		status = esmStatus.Status
	}
	if status {
		return nil, esmtypes.ErrESMAlreadyExecuted
	}
	killSwitchParams, _ := k.GetKillSwitchData(ctx, msg.AppId)
	if killSwitchParams.BreakerEnable {
		return nil, esmtypes.ErrCircuitBreakerEnabled
	}
	depositor, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	//checks if extended pair exists
	extendedPairVault, found := k.GetPairsVault(ctx, msg.ExtendedPairVaultId)
	if !found {
		return nil, types.ErrorExtendedPairVaultDoesNotExists
	}
	pairData, found := k.GetPair(ctx, extendedPairVault.PairId)
	if !found {
		return nil, types.ErrorPairDoesNotExist
	}
	assetInData, found := k.GetAsset(ctx, pairData.AssetIn)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}
	assetOutData, found := k.GetAsset(ctx, pairData.AssetOut)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}

	//Checking if appMapping_id exists
	appMapping, found := k.GetApp(ctx, msg.AppId)
	if !found {
		return nil, types.ErrorAppMappingDoesNotExist
	}
	// //Checking if vault acccess disabled

	//Checking if the appMapping_id in the msg_create & extendedPairVault_are same or not
	if appMapping.Id != extendedPairVault.AppId {
		return nil, types.ErrorAppMappingIDMismatch
	}

	userVault, found := k.GetVault(ctx, msg.UserVaultId)
	if !found {
		return nil, types.ErrorVaultDoesNotExist
	}
	if userVault.Owner != msg.From {
		return nil, types.ErrVaultAccessUnauthorised
	}

	if appMapping.Id != userVault.AppId {
		return nil, types.ErrorInvalidAppMappingData
	}
	if extendedPairVault.Id != userVault.ExtendedPairVaultID {
		return nil, types.ErrorInvalidExtendedPairMappingData
	}

	totalDebt := userVault.AmountOut.Add(userVault.InterestAccumulated)
	err1 := k.CalculateVaultInterest(ctx, appMapping.Id, msg.ExtendedPairVaultId, msg.UserVaultId, totalDebt, userVault.BlockHeight, userVault.BlockTime.Unix())
	if err1 != nil {
		return nil, err1
	}

	userVault, found1 := k.GetVault(ctx, msg.UserVaultId)
	if !found1 {
		return nil, types.ErrorVaultDoesNotExist
	}

	totalUserDebt := userVault.AmountOut.Add(userVault.InterestAccumulated)
	totalUserDebt = totalUserDebt.Add(userVault.ClosingFeeAccumulated)
	if totalUserDebt.GT(sdk.ZeroInt()) {
		if err := k.SendCoinFromAccountToModule(ctx, depositor, types.ModuleName, sdk.NewCoin(assetOutData.Denom, totalUserDebt)); err != nil {
			return nil, err
		}
	}

	//			SEND TO COLLECTOR----userVault.InterestAccumulated & userVault.ClosingFees

	err = k.UpdateCollector(ctx, appMapping.Id, pairData.AssetOut, userVault.InterestAccumulated, userVault.ClosingFeeAccumulated, sdk.ZeroInt(), sdk.ZeroInt())
	if err != nil {
		return nil, err
	}
	if userVault.InterestAccumulated.GT(sdk.ZeroInt()) {
		if err := k.SendCoinFromModuleToModule(ctx, types.ModuleName, collectortypes.ModuleName, sdk.NewCoins(sdk.NewCoin(assetOutData.Denom, userVault.InterestAccumulated))); err != nil {
			return nil, err
		}
	}
	if userVault.ClosingFeeAccumulated.GT(sdk.ZeroInt()) {
		if err := k.SendCoinFromModuleToModule(ctx, types.ModuleName, collectortypes.ModuleName, sdk.NewCoins(sdk.NewCoin(assetOutData.Denom, userVault.ClosingFeeAccumulated))); err != nil {
			return nil, err
		}
	}
	if userVault.AmountOut.GT(sdk.ZeroInt()) {
		if err := k.BurnCoin(ctx, types.ModuleName, sdk.NewCoin(assetOutData.Denom, userVault.AmountOut)); err != nil {
			return nil, err
		}
	}
	if userVault.AmountIn.GT(sdk.ZeroInt()) {
		if err := k.SendCoinFromModuleToAccount(ctx, types.ModuleName, depositor, sdk.NewCoin(assetInData.Denom, userVault.AmountIn)); err != nil {
			return nil, err
		}
	}

	//Update LookupTable minting Status
	appExtendedPairVaultData, _ := k.GetAppExtendedPairVaultMappingData(ctx, appMapping.Id, msg.ExtendedPairVaultId)

	k.UpdateCollateralLockedAmountLockerMapping(ctx, appExtendedPairVaultData.AppId, appExtendedPairVaultData.ExtendedPairId, userVault.AmountIn, false)
	k.UpdateTokenMintedAmountLockerMapping(ctx, appExtendedPairVaultData.AppId, appExtendedPairVaultData.ExtendedPairId, userVault.AmountOut, false)

	//Remove address from lookup table
	k.DeleteAddressFromAppExtendedPairVaultMapping(ctx, extendedPairVault.Id, userVault.Id, appMapping.Id)

	//Remove user extendedPair to address field in UserLookupStruct
	k.DeleteUserVaultExtendedPairMapping(ctx, msg.From, appMapping.Id, extendedPairVault.Id)

	//Delete Vault
	k.DeleteVault(ctx, userVault.Id)

	var rewards rewardstypes.VaultInterestTracker
	rewards.AppMappingId = appMapping.Id
	rewards.VaultId = userVault.Id
	k.DeleteVaultInterestTracker(ctx, rewards)

	ctx.GasMeter().ConsumeGas(types.CloseVaultGas, "CloseVaultGas")

	return &types.MsgCloseResponse{}, nil
}

// MsgDepositAndDraw.
func (k msgServer) MsgDepositAndDraw(c context.Context, msg *types.MsgDepositAndDrawRequest) (*types.MsgDepositAndDrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	userVault, found := k.GetVault(ctx, msg.UserVaultId)
	if !found {
		return nil, types.ErrorVaultDoesNotExist
	}
	newAmt := k.calculateUserToken(userVault, msg.Amount)
	msgDepositReq := types.MsgDepositRequest{
		From:                msg.From,
		AppId:               msg.AppId,
		ExtendedPairVaultId: msg.ExtendedPairVaultId,
		UserVaultId:         msg.UserVaultId,
		Amount:              msg.Amount,
	}
	_, err := k.MsgDeposit(c, &msgDepositReq)
	if err != nil {
		return nil, err
	}
	msgDrawReq := types.MsgDrawRequest{
		From:                msg.From,
		AppId:               msg.AppId,
		ExtendedPairVaultId: msg.ExtendedPairVaultId,
		UserVaultId:         msg.UserVaultId,
		Amount:              newAmt,
	}
	_, err = k.MsgDraw(c, &msgDrawReq)
	if err != nil {
		return nil, err
	}
	return &types.MsgDepositAndDrawResponse{}, nil
}

func (k msgServer) MsgCreateStableMint(c context.Context, msg *types.MsgCreateStableMintRequest) (*types.MsgCreateStableMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	esmStatus, found := k.GetESMStatus(ctx, msg.AppId)
	status := false
	if found {
		status = esmStatus.Status
	}
	if status {
		return nil, esmtypes.ErrESMAlreadyExecuted
	}
	killSwitchParams, _ := k.GetKillSwitchData(ctx, msg.AppId)
	if killSwitchParams.BreakerEnable {
		return nil, esmtypes.ErrCircuitBreakerEnabled
	}
	//Checking if extended pair exists
	extendedPairVault, found := k.GetPairsVault(ctx, msg.ExtendedPairVaultId)
	if !found {
		return nil, types.ErrorExtendedPairVaultDoesNotExists
	}
	pairData, found := k.GetPair(ctx, extendedPairVault.PairId)
	if !found {
		return nil, types.ErrorPairDoesNotExist
	}
	assetInData, found := k.GetAsset(ctx, pairData.AssetIn)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}
	assetOutData, found := k.GetAsset(ctx, pairData.AssetOut)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}
	//Checking if appMapping_id exists
	appMapping, found := k.GetApp(ctx, msg.AppId)
	if !found {
		return nil, types.ErrorAppMappingDoesNotExist
	}

	//Checking if the appMapping_id in the msg_create & extendedPairVault_are same or not
	if appMapping.Id != extendedPairVault.AppId {
		return nil, types.ErrorAppMappingIDMismatch
	}

	//Converting user address for bank transaction
	depositorAddress, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Checking if this is a stableMint pair or not  -- stableMintPair == psmPair
	if !extendedPairVault.IsStableMintVault {
		return nil, types.ErrorCannotCreateStableMintVault
	}
	//Checking
	if !extendedPairVault.IsVaultActive {
		return nil, types.ErrorVaultCreationInactive
	}
	//Call CheckAppExtendedPairVaultMapping function to get counter - it also initialised the kv store if appMapping_id does not exists, or extendedPairVault_id does not exists.

	tokenMintedStatistics, _ := k.CheckAppExtendedPairVaultMapping(ctx, appMapping.Id, extendedPairVault.Id)

	extPairData, _ := k.GetAppExtendedPairVaultMappingData(ctx, appMapping.Id, msg.ExtendedPairVaultId)
	if len(extPairData.VaultIds) >= 1 {
		return nil, types.ErrorStableMintVaultAlreadyCreated
	}

	//Check Debt Ceil
	currentMintedStatistics := tokenMintedStatistics.Add(msg.Amount)

	if currentMintedStatistics.GTE(extendedPairVault.DebtCeiling) {
		return nil, types.ErrorAmountOutGreaterThanDebtCeiling
	}

	if msg.Amount.GT(sdk.ZeroInt()) {
		//Take amount from user
		if err := k.SendCoinFromAccountToModule(ctx, depositorAddress, types.ModuleName, sdk.NewCoin(assetInData.Denom, msg.Amount)); err != nil {
			return nil, err
		}
		//Mint Tokens for user

		if err := k.MintCoin(ctx, types.ModuleName, sdk.NewCoin(assetOutData.Denom, msg.Amount)); err != nil {
			return nil, err
		}
	}

	if extendedPairVault.DrawDownFee.IsZero() && msg.Amount.GT(sdk.ZeroInt()) {
		//Send Rest to user
		if err := k.SendCoinFromModuleToAccount(ctx, types.ModuleName, depositorAddress, sdk.NewCoin(assetOutData.Denom, msg.Amount)); err != nil {
			return nil, err
		}
	} else {
		//If not zero deduct send to collector//////////
		//			COLLECTOR FUNCTION
		collectorShare := (msg.Amount.Mul(sdk.Int(extendedPairVault.DrawDownFee))).Quo(sdk.Int(sdk.OneDec()))
		if collectorShare.GT(sdk.ZeroInt()) {
			if err := k.SendCoinFromModuleToModule(ctx, types.ModuleName, collectortypes.ModuleName, sdk.NewCoins(sdk.NewCoin(assetOutData.Denom, collectorShare))); err != nil {
				return nil, err
			}
			err := k.UpdateCollector(ctx, appMapping.Id, pairData.AssetOut, sdk.ZeroInt(), sdk.ZeroInt(), collectorShare, sdk.ZeroInt())
			if err != nil {
				return nil, err
			}
		}

		// and send the rest to the user
		amountToUser := msg.Amount.Sub(collectorShare)
		if amountToUser.GT(sdk.ZeroInt()) {
			if err := k.SendCoinFromModuleToAccount(ctx, types.ModuleName, depositorAddress, sdk.NewCoin(assetOutData.Denom, amountToUser)); err != nil {
				return nil, err
			}
		}
	}
	//Create Mint Vault

	oldID := k.GetIDForStableVault(ctx)
	var stableVault types.StableMintVault
	newID := oldID + 1

	stableVault.Id = newID
	stableVault.AmountIn = msg.Amount
	stableVault.AmountOut = msg.Amount
	stableVault.AppId = appMapping.Id
	stableVault.CreatedAt = ctx.BlockTime()
	stableVault.ExtendedPairVaultID = extendedPairVault.Id
	k.SetStableMintVault(ctx, stableVault)
	k.SetIDForStableVault(ctx, newID)
	//update Locker Data 	//Update Amount
	k.UpdateAppExtendedPairVaultMappingDataOnMsgCreateStableMintVault(ctx, stableVault)

	ctx.GasMeter().ConsumeGas(types.CreateStableVaultGas, "CreateStableVaultGas")

	return &types.MsgCreateStableMintResponse{}, nil
}

func (k msgServer) MsgDepositStableMint(c context.Context, msg *types.MsgDepositStableMintRequest) (*types.MsgDepositStableMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	esmStatus, found := k.GetESMStatus(ctx, msg.AppId)
	status := false
	if found {
		status = esmStatus.Status
	}
	if status {
		return nil, esmtypes.ErrESMAlreadyExecuted
	}
	killSwitchParams, _ := k.GetKillSwitchData(ctx, msg.AppId)
	if killSwitchParams.BreakerEnable {
		return nil, esmtypes.ErrCircuitBreakerEnabled
	}
	depositorAddress, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	//checks if extended pair exists
	extendedPairVault, found := k.GetPairsVault(ctx, msg.ExtendedPairVaultId)
	if !found {
		return nil, types.ErrorExtendedPairVaultDoesNotExists
	}
	pairData, found := k.GetPair(ctx, extendedPairVault.PairId)
	if !found {
		return nil, types.ErrorPairDoesNotExist
	}
	assetInData, found := k.GetAsset(ctx, pairData.AssetIn)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}
	assetOutData, found := k.GetAsset(ctx, pairData.AssetOut)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}

	//Checking if appMapping_id exists
	appMapping, found := k.GetApp(ctx, msg.AppId)
	if !found {
		return nil, types.ErrorAppMappingDoesNotExist
	}
	//Checking if vault access disabled
	if !extendedPairVault.IsVaultActive {
		return nil, types.ErrorVaultInactive
	}
	if !extendedPairVault.IsStableMintVault {
		return nil, types.ErrorCannotCreateStableMintVault
	}
	//Checking if the appMapping_id in the msg_create & extendedPairVault_are same or not
	if appMapping.Id != extendedPairVault.AppId {
		return nil, types.ErrorAppMappingIDMismatch
	}

	stableVault, found := k.GetStableMintVault(ctx, msg.StableVaultId)
	if !found {
		return nil, types.ErrorVaultDoesNotExist
	}
	if appMapping.Id != stableVault.AppId {
		return nil, types.ErrorInvalidAppMappingData
	}
	if extendedPairVault.Id != stableVault.ExtendedPairVaultID {
		return nil, types.ErrorInvalidExtendedPairMappingData
	}

	stableAmountIn := stableVault.AmountIn.Add(msg.Amount)
	if !stableAmountIn.IsPositive() {
		return nil, types.ErrorInvalidAmount
	}
	//Looking for a case where apart from create function , this function creates new vaults and its data.
	tokenMintedStatistics, _ := k.CheckAppExtendedPairVaultMapping(ctx, appMapping.Id, extendedPairVault.Id)

	//Check Debt Ceil
	currentMintedStatistics := tokenMintedStatistics.Add(msg.Amount)

	if currentMintedStatistics.GTE(extendedPairVault.DebtCeiling) {
		return nil, types.ErrorAmountOutGreaterThanDebtCeiling
	}

	if msg.Amount.GT(sdk.ZeroInt()) {
		//Take amount from user
		if err := k.SendCoinFromAccountToModule(ctx, depositorAddress, types.ModuleName, sdk.NewCoin(assetInData.Denom, msg.Amount)); err != nil {
			return nil, err
		}
		//Mint Tokens for user

		if err := k.MintCoin(ctx, types.ModuleName, sdk.NewCoin(assetOutData.Denom, msg.Amount)); err != nil {
			return nil, err
		}
	}
	if extendedPairVault.DrawDownFee.IsZero() && msg.Amount.GT(sdk.ZeroInt()) {
		//Send Rest to user
		if err := k.SendCoinFromModuleToAccount(ctx, types.ModuleName, depositorAddress, sdk.NewCoin(assetOutData.Denom, msg.Amount)); err != nil {
			return nil, err
		}
	} else {
		//If not zero deduct send to collector//////////
		//			COLLECTOR FUNCTION
		/////////////////////////////////////////////////

		collectorShare := (msg.Amount.Mul(sdk.Int(extendedPairVault.DrawDownFee))).Quo(sdk.Int(sdk.OneDec()))
		if collectorShare.GT(sdk.ZeroInt()) {
			if err := k.SendCoinFromModuleToModule(ctx, types.ModuleName, collectortypes.ModuleName, sdk.NewCoins(sdk.NewCoin(assetOutData.Denom, collectorShare))); err != nil {
				return nil, err
			}
			err := k.UpdateCollector(ctx, appMapping.Id, pairData.AssetOut, sdk.ZeroInt(), sdk.ZeroInt(), collectorShare, sdk.ZeroInt())
			if err != nil {
				return nil, err
			}
		}

		// and send the rest to the user
		amountToUser := msg.Amount.Sub(collectorShare)
		if amountToUser.GT(sdk.ZeroInt()) {
			if err := k.SendCoinFromModuleToAccount(ctx, types.ModuleName, depositorAddress, sdk.NewCoin(assetOutData.Denom, amountToUser)); err != nil {
				return nil, err
			}
		}
	}
	stableVault.AmountIn = stableVault.AmountIn.Add(msg.Amount)
	stableVault.AmountOut = stableVault.AmountOut.Add(msg.Amount)

	k.SetStableMintVault(ctx, stableVault)
	appExtendedPairVaultData, _ := k.GetAppExtendedPairVaultMappingData(ctx, appMapping.Id, msg.ExtendedPairVaultId)
	k.UpdateCollateralLockedAmountLockerMapping(ctx, appExtendedPairVaultData.AppId, appExtendedPairVaultData.ExtendedPairId, stableVault.AmountIn, true)
	k.UpdateTokenMintedAmountLockerMapping(ctx, appExtendedPairVaultData.AppId, appExtendedPairVaultData.ExtendedPairId, stableVault.AmountOut, true)

	ctx.GasMeter().ConsumeGas(types.DepositStableVaultGas, "DepositStableVaultGas")
	return &types.MsgDepositStableMintResponse{}, nil
}

func (k msgServer) MsgWithdrawStableMint(c context.Context, msg *types.MsgWithdrawStableMintRequest) (*types.MsgWithdrawStableMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	esmStatus, found := k.GetESMStatus(ctx, msg.AppId)
	status := false
	if found {
		status = esmStatus.Status
	}
	if status {
		return nil, esmtypes.ErrESMAlreadyExecuted
	}
	killSwitchParams, _ := k.GetKillSwitchData(ctx, msg.AppId)
	if killSwitchParams.BreakerEnable {
		return nil, esmtypes.ErrCircuitBreakerEnabled
	}
	depositorAddress, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	//checks if extended pair exists
	extendedPairVault, found := k.GetPairsVault(ctx, msg.ExtendedPairVaultId)
	if !found {
		return nil, types.ErrorExtendedPairVaultDoesNotExists
	}
	pairData, found := k.GetPair(ctx, extendedPairVault.PairId)
	if !found {
		return nil, types.ErrorPairDoesNotExist
	}
	assetInData, found := k.GetAsset(ctx, pairData.AssetIn)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}
	assetOutData, found := k.GetAsset(ctx, pairData.AssetOut)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}

	//Checking if appMapping_id exists
	appMapping, found := k.GetApp(ctx, msg.AppId)
	if !found {
		return nil, types.ErrorAppMappingDoesNotExist
	}
	//Checking if vault access disabled

	if !extendedPairVault.IsStableMintVault {
		return nil, types.ErrorCannotCreateStableMintVault
	}
	//Checking if the appMapping_id in the msg_create & extendedPairVault_are same or not
	if appMapping.Id != extendedPairVault.AppId {
		return nil, types.ErrorAppMappingIDMismatch
	}

	stableVault, found := k.GetStableMintVault(ctx, msg.StableVaultId)
	if !found {
		return nil, types.ErrorVaultDoesNotExist
	}
	if appMapping.Id != stableVault.AppId {
		return nil, types.ErrorInvalidAppMappingData
	}
	if extendedPairVault.Id != stableVault.ExtendedPairVaultID {
		return nil, types.ErrorInvalidExtendedPairMappingData
	}

	stableAmountIn := stableVault.AmountIn.Sub(msg.Amount)
	if stableAmountIn.LT(sdk.NewInt(0)) {
		return nil, types.ErrorInvalidAmount
	}
	var updatedAmount sdk.Int
	//Take amount from user
	if msg.Amount.GT(sdk.ZeroInt()) {
		if err := k.SendCoinFromAccountToModule(ctx, depositorAddress, types.ModuleName, sdk.NewCoin(assetOutData.Denom, msg.Amount)); err != nil {
			return nil, err
		}
	}

	if extendedPairVault.DrawDownFee.IsZero() && msg.Amount.GT(sdk.ZeroInt()) {
		//BurnTokens for user
		if err := k.BurnCoin(ctx, types.ModuleName, sdk.NewCoin(assetOutData.Denom, msg.Amount)); err != nil {
			return nil, err
		}

		//Send Rest to user
		if err := k.SendCoinFromModuleToAccount(ctx, types.ModuleName, depositorAddress, sdk.NewCoin(assetInData.Denom, msg.Amount)); err != nil {
			return nil, err
		}
		updatedAmount = msg.Amount
	} else {
		//If not zero deduct send to collector//////////
		//			COLLECTOR FUNCTION
		/////////////////////////////////////////////////
		collectorShare := (msg.Amount.Mul(sdk.Int(extendedPairVault.DrawDownFee))).Quo(sdk.Int(sdk.OneDec()))
		if collectorShare.GT(sdk.ZeroInt()) {
			if err := k.SendCoinFromModuleToModule(ctx, types.ModuleName, collectortypes.ModuleName, sdk.NewCoins(sdk.NewCoin(assetOutData.Denom, collectorShare))); err != nil {
				return nil, err
			}
			err := k.UpdateCollector(ctx, appMapping.Id, pairData.AssetOut, sdk.ZeroInt(), sdk.ZeroInt(), collectorShare, sdk.ZeroInt())
			if err != nil {
				return nil, err
			}
		}

		updatedAmount = msg.Amount.Sub(collectorShare)

		if updatedAmount.GT(sdk.ZeroInt()) {
			//BurnTokens for user
			if err := k.BurnCoin(ctx, types.ModuleName, sdk.NewCoin(assetOutData.Denom, updatedAmount)); err != nil {
				return nil, err
			}

			// and send the rest to the user

			if err := k.SendCoinFromModuleToAccount(ctx, types.ModuleName, depositorAddress, sdk.NewCoin(assetInData.Denom, updatedAmount)); err != nil {
				return nil, err
			}
		}
	}
	stableVault.AmountIn = stableVault.AmountIn.Sub(updatedAmount)
	stableVault.AmountOut = stableVault.AmountOut.Sub(updatedAmount)
	k.SetStableMintVault(ctx, stableVault)
	appExtendedPairVaultData, _ := k.GetAppExtendedPairVaultMappingData(ctx, appMapping.Id, msg.ExtendedPairVaultId)
	k.UpdateCollateralLockedAmountLockerMapping(ctx, appExtendedPairVaultData.AppId, appExtendedPairVaultData.ExtendedPairId, stableVault.AmountIn, false)
	k.UpdateTokenMintedAmountLockerMapping(ctx, appExtendedPairVaultData.AppId, appExtendedPairVaultData.ExtendedPairId, stableVault.AmountOut, false)

	ctx.GasMeter().ConsumeGas(types.WithdrawStableVaultGas, "WithdrawStableVaultGas")

	return &types.MsgWithdrawStableMintResponse{}, nil
}

//take app id
//check app id
//take vault id
// check vault id
//calculate total debt
//call function
//exit function

func (k msgServer) MsgVaultInterestCalc(c context.Context, msg *types.MsgVaultInterestCalcRequest) (*types.MsgVaultInterestCalcResponse, error) {

	ctx := sdk.UnwrapSDKContext(c)
	appMapping, found := k.GetApp(ctx, msg.AppId)
	if !found {
		return nil, types.ErrorAppMappingDoesNotExist
	}
	userVault, found := k.GetVault(ctx, msg.UserVaultId)
	if !found {
		return nil, types.ErrorVaultDoesNotExist
	}

	totalDebt := userVault.AmountOut.Add(userVault.InterestAccumulated)
	err1 := k.CalculateVaultInterest(ctx, appMapping.Id, userVault.ExtendedPairVaultID, msg.UserVaultId, totalDebt, userVault.BlockHeight, userVault.BlockTime.Unix())
	if err1 != nil {
		return nil, err1
	}

	return &types.MsgVaultInterestCalcResponse{}, nil
}
