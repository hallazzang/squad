package mint

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmosquad-labs/squad/x/mint/keeper"
	"github.com/cosmosquad-labs/squad/x/mint/types"
)

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper, lastBlockTime time.Duration) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// fetch stored params
	params := k.GetParams(ctx)

	InflationSchedules := k.GetInflationSchedules()
	blockInflation := sdk.ZeroInt()
	for _, schedule := range InflationSchedules {
		if !schedule.EndTime.Before(ctx.BlockTime()) && !schedule.StartTime.After(ctx.BlockTime()) {
			// TODO: need to Get/Set LastBlockTime
			//lastBlockTime := ctx.BlockTime().Sub(k.GetLastBlockTime(ctx))
			if lastBlockTime > params.BlockTimeThreshold {
				lastBlockTime = params.BlockTimeThreshold
			}
			// blockInflation = InflationAmountThisPeriod * min(CurrentBlockTime-LastBlockTime,BlockTimeThreshold)/(InflationPeriodEndDate-InflationPeriodStartDate)
			blockInflation = schedule.Amount.MulRaw(lastBlockTime.Nanoseconds()).QuoRaw(schedule.EndTime.Sub(schedule.StartTime).Nanoseconds())
			break
		}
	}
	if blockInflation.IsPositive() {
		mintedCoin := sdk.NewCoin(params.MintDenom, blockInflation)
		mintedCoins := sdk.NewCoins(mintedCoin)
		err := k.MintCoins(ctx, mintedCoins)
		if err != nil {
			panic(err)
		}

		// send the minted coins to the fee collector account
		err = k.AddCollectedFees(ctx, mintedCoins)
		if err != nil {
			panic(err)
		}

		if mintedCoin.Amount.IsInt64() {
			defer telemetry.ModuleSetGauge(types.ModuleName, float32(mintedCoin.Amount.Int64()), "minted_tokens")
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeMint,
				sdk.NewAttribute(sdk.AttributeKeyAmount, mintedCoin.Amount.String()),
			),
		)
	}
}