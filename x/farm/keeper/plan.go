package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmosquad-labs/squad/v3/x/farm/types"
	liquiditytypes "github.com/cosmosquad-labs/squad/v3/x/liquidity/types"
)

// CreatePrivatePlan creates a new private farming plan.
func (k Keeper) CreatePrivatePlan(
	ctx sdk.Context, creatorAddr sdk.AccAddress, description string,
	rewardAllocs []types.RewardAllocation, startTime, endTime time.Time,
) (types.Plan, error) {
	// Check if end time > block time
	if !endTime.After(ctx.BlockTime()) {
		return types.Plan{}, sdkerrors.Wrap(
			sdkerrors.ErrInvalidRequest, "end time is past")
	}

	// Check if the number of non-terminated private plans is not greater than
	// the MaxNumPrivatePlans param.
	// TODO: store the counter separately to optimize gas usage?
	numPrivatePlans := 0
	k.IterateAllPlans(ctx, func(plan types.Plan) (stop bool) {
		if plan.IsPrivate && !plan.IsTerminated {
			numPrivatePlans++
		}
		return false
	})
	if maxNum := k.GetMaxNumPrivatePlans(ctx); uint32(numPrivatePlans) >= maxNum {
		return types.Plan{}, sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"maximum number of active private plans reached: %d", maxNum)
	}

	for _, rewardAlloc := range rewardAllocs {
		_, found := k.liquidityKeeper.GetPair(ctx, rewardAlloc.PairId)
		if !found {
			return types.Plan{}, sdkerrors.Wrapf(
				sdkerrors.ErrNotFound, "pair %d not found", rewardAlloc.PairId)
		}
	}

	fee := k.GetPrivatePlanCreationFee(ctx)
	feeCollectorAddr, err := sdk.AccAddressFromBech32(k.GetFeeCollector(ctx))
	if err != nil {
		return types.Plan{}, err
	}
	if err := k.bankKeeper.SendCoins(ctx, creatorAddr, feeCollectorAddr, fee); err != nil {
		return types.Plan{}, err
	}

	// Generate the next plan id and update the last plan id.
	id, _ := k.GetLastPlanId(ctx)
	id++
	k.SetLastPlanId(ctx, id)

	farmingPoolAddr := types.DeriveFarmingPoolAddress(id)
	plan := types.NewPlan(
		id, description, farmingPoolAddr, creatorAddr, rewardAllocs,
		startTime, endTime, true)
	k.SetPlan(ctx, plan)

	return plan, nil
}

// TerminateEndedPlans iterates through all plans and terminate the plans
// which should be ended by the current block time.
func (k Keeper) TerminateEndedPlans(ctx sdk.Context) (err error) {
	k.IterateAllPlans(ctx, func(plan types.Plan) (stop bool) {
		if plan.IsTerminated {
			return false
		}
		if !ctx.BlockTime().Before(plan.EndTime) {
			if err = k.TerminatePlan(ctx, plan); err != nil {
				return true
			}
		}
		return false
	})
	return err
}

// TerminatePlan mark the plan as terminated and send remaining balances
// in the farming pool to the termination address.
func (k Keeper) TerminatePlan(ctx sdk.Context, plan types.Plan) error {
	if plan.IsTerminated {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "already terminated plan")
	}
	farmingPoolAddr := plan.GetFarmingPoolAddress()
	balances := k.bankKeeper.SpendableCoins(ctx, farmingPoolAddr)
	if err := k.bankKeeper.SendCoins(
		ctx, farmingPoolAddr, plan.GetTerminationAddress(), balances); err != nil {
		return err
	}
	plan.IsTerminated = true
	k.SetPlan(ctx, plan)
	return nil
}

func (k Keeper) AllocateRewards(ctx sdk.Context) error {
	lastBlockTime, found := k.GetLastBlockTime(ctx)
	if !found {
		// Skip this block.
		k.SetLastBlockTime(ctx, ctx.BlockTime())
		return nil
	}

	blockDuration := ctx.BlockTime().Sub(lastBlockTime)
	if maxBlockDuration := k.GetMaxBlockDuration(ctx); blockDuration > maxBlockDuration {
		// Constrain the block duration to the max block duration param.
		blockDuration = maxBlockDuration
	}

	var activePlans []types.Plan
	k.IterateAllPlans(ctx, func(plan types.Plan) (stop bool) {
		if plan.IsTerminated || !plan.IsActiveAt(ctx.BlockTime()) {
			return false // Skip
		}
		activePlans = append(activePlans, plan)
		return false
	})

	cache := NewCachedKeeper(k)
	totalRewardsByFarmingPool := map[string]sdk.Coins{}
	rewardsByPair := map[uint64]sdk.Coins{}
	eligiblePoolsByPair := map[uint64][]liquiditytypes.Pool{}
	for _, plan := range activePlans {
		for _, rewardAlloc := range plan.RewardAllocations {
			pair, found := cache.GetPair(ctx, rewardAlloc.PairId)
			if !found { // It should never happen
				panic("pair not found")
			}
			if pair.LastPrice == nil { // If the pair doesn't have the last price, skip.
				continue
			}

			// Collect pools eligible for reward allocation.
			_ = k.liquidityKeeper.IteratePoolsByPair(
				ctx, pair.Id, func(pool liquiditytypes.Pool) (stop bool, err error) {
					if pool.Disabled {
						return false, nil
					}
					// If the pool is a ranged pool and its pair's last price is out of
					// its price range, skip the pool.
					// This is because the amplification factor would be zero
					// so its reward weight would eventually be zero, too.
					if pool.Type == liquiditytypes.PoolTypeRanged &&
						(pair.LastPrice.LT(*pool.MinPrice) || pair.LastPrice.GT(*pool.MaxPrice)) {
						return false, nil
					}

					farm, found := cache.GetFarm(ctx, pool.PoolCoinDenom)
					if !found || farm.TotalFarmingAmount.IsZero() {
						return false, nil
					}

					eligiblePoolsByPair[pair.Id] = append(eligiblePoolsByPair[pair.Id], pool)
					return false, nil
				},
			)

			if len(eligiblePoolsByPair[pair.Id]) > 0 {
				rewards := types.RewardsForBlock(rewardAlloc.RewardsPerDay, blockDuration)
				// TODO: allocate sdk.DecCoins instead of sdk.Coins in future
				truncatedRewards, _ := rewards.TruncateDecimal()
				rewardsByPair[pair.Id] = rewardsByPair[pair.Id].Add(truncatedRewards...)
				totalRewardsByFarmingPool[plan.FarmingPoolAddress] =
					totalRewardsByFarmingPool[plan.FarmingPoolAddress].Add(truncatedRewards...)
			}
		}
	}

	rewardsByDenom := map[string]sdk.DecCoins{}
	// We keep this slice for deterministic iteration over the rewardsByDenom map.
	var denomsWithRewards []string
	for _, plan := range activePlans {
		totalRewards := totalRewardsByFarmingPool[plan.FarmingPoolAddress]
		balances := cache.SpendableCoins(ctx, plan.GetFarmingPoolAddress())
		if !balances.IsAllGTE(totalRewardsByFarmingPool[plan.FarmingPoolAddress]) {
			continue
		}
		if err := k.bankKeeper.SendCoins(
			ctx, plan.GetFarmingPoolAddress(), types.RewardsPoolAddress, totalRewards); err != nil {
			return err
		}

		for _, rewardAlloc := range plan.RewardAllocations {
			pair, _ := cache.GetPair(ctx, rewardAlloc.PairId)

			rewardWeightByPool := map[uint64]sdk.Dec{}
			totalRewardWeight := sdk.ZeroDec()
			for _, pool := range eligiblePoolsByPair[pair.Id] {
				rewardWeight := k.PoolRewardWeight(ctx, pool, pair)
				rewardWeightByPool[pool.Id] = rewardWeight
				totalRewardWeight = totalRewardWeight.Add(rewardWeight)
			}

			for _, pool := range eligiblePoolsByPair[pair.Id] {
				rewardProportion := rewardWeightByPool[pool.Id].QuoTruncate(totalRewardWeight)
				rewards := sdk.NewDecCoinsFromCoins(rewardsByPair[pair.Id]...).
					MulDecTruncate(rewardProportion)

				if _, ok := rewardsByDenom[pool.PoolCoinDenom]; !ok {
					denomsWithRewards = append(denomsWithRewards, pool.PoolCoinDenom)
				}
				rewardsByDenom[pool.PoolCoinDenom] =
					rewardsByDenom[pool.PoolCoinDenom].Add(rewards...)
			}
		}
	}

	for _, denom := range denomsWithRewards {
		farm, _ := cache.GetFarm(ctx, denom)
		farm.CurrentRewards = farm.CurrentRewards.Add(rewardsByDenom[denom]...)
		farm.OutstandingRewards = farm.OutstandingRewards.Add(rewardsByDenom[denom]...)
		k.SetFarm(ctx, denom, farm)
	}

	return nil
}
