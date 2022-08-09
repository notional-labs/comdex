package rewards

import (
	utils "github.com/comdex-official/comdex/types"
	"github.com/comdex-official/comdex/x/rewards/keeper"
	"github.com/comdex-official/comdex/x/rewards/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func BeginBlocker(ctx sdk.Context, _ abci.RequestBeginBlock, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, ctx.BlockTime(), telemetry.MetricKeyBeginBlocker)

	_ = utils.ApplyFuncIfNoError(ctx, func(ctx sdk.Context) error {
		k.TriggerAndUpdateEpochInfos(ctx)
		err := k.IterateLocker(ctx)
		if err != nil {
			ctx.Logger().Error("error in IterateLocker")
		}

		appIDsVault := k.GetAppIDs(ctx).WhitelistedAppMappingIdsVaults
		for i, _ := range appIDsVault {
			err := k.IterateVaults(ctx, appIDsVault[i])
			if err != nil {
				continue
			}
		}

		err = k.DistributeExtRewardLocker(ctx)
		if err != nil {
			ctx.Logger().Error("error in DistributeExtRewardLocker")
		}
		err = k.DistributeExtRewardVault(ctx)
		if err != nil {
			ctx.Logger().Error("error in DistributeExtRewardVault")
		}

		err = k.SetLastInterestTime(ctx, ctx.BlockTime().Unix())
		if err != nil {
			ctx.Logger().Error("error in SetLastInterestTime")
		}
		return nil
	})
}

// EndBlocker for incentives module.
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {}