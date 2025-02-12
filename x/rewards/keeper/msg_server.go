package keeper

import (
	"context"

	esmtypes "github.com/comdex-official/comdex/x/esm/types"
	assettypes "github.com/comdex-official/comdex/x/asset/types"
	"github.com/comdex-official/comdex/x/rewards/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) CreateGauge(goCtx context.Context, msg *types.MsgCreateGauge) (*types.MsgCreateGaugeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := m.Keeper.ValidateMsgCreateGauge(ctx, msg)
	if err != nil {
		return nil, err
	}

	err = m.Keeper.CreateNewGauge(ctx, msg, false)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateGaugeResponse{}, nil
}

func (m msgServer) ExternalRewardsLockers(goCtx context.Context, msg *types.ActivateExternalRewardsLockers) (*types.ActivateExternalRewardsLockersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	klwsParams, _ := m.GetKillSwitchData(ctx, msg.AppMappingId)
	if klwsParams.BreakerEnable {
		return nil, esmtypes.ErrCircuitBreakerEnabled
	}
	esmStatus, found := m.GetESMStatus(ctx, msg.AppMappingId)
	status := false
	if found {
		status = esmStatus.Status
	}
	if status {
		return nil, esmtypes.ErrESMAlreadyExecuted
	}
	Depositor, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		return nil, err
	}
	if err := m.Keeper.ActExternalRewardsLockers(ctx, msg.AppMappingId, msg.AssetId, msg.TotalRewards, msg.DurationDays, Depositor, msg.MinLockupTimeSeconds); err != nil {
		return nil, err
	}
	return &types.ActivateExternalRewardsLockersResponse{}, nil
}

func (m msgServer) ExternalRewardsVault(goCtx context.Context, msg *types.ActivateExternalRewardsVault) (*types.ActivateExternalRewardsVaultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	klwsParams, _ := m.GetKillSwitchData(ctx, msg.AppMappingId)
	if klwsParams.BreakerEnable {
		return nil, esmtypes.ErrCircuitBreakerEnabled
	}
	esmStatus, found := m.GetESMStatus(ctx, msg.AppMappingId)
	status := false
	if found {
		status = esmStatus.Status
	}
	if status {
		return nil, esmtypes.ErrESMAlreadyExecuted
	}
	Depositor, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		return nil, err
	}
	pairVault, found := m.GetPairsVault(ctx, msg.Extended_Pair_Id)
	if !found {
		return nil, assettypes.ErrorExtendedPairDoesNotExistForTheApp
	}
	if pairVault.IsStableMintVault {
		return nil, types.ErrStablemintVaultFound
	}
	if err := m.Keeper.ActExternalRewardsVaults(ctx, msg.AppMappingId, msg.Extended_Pair_Id, msg.DurationDays, msg.MinLockupTimeSeconds, msg.TotalRewards, Depositor); err != nil {
		return nil, err
	}
	return &types.ActivateExternalRewardsVaultResponse{}, nil
}
