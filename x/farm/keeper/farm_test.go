package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	utils "github.com/cosmosquad-labs/squad/v3/types"
	"github.com/cosmosquad-labs/squad/v3/x/farm/types"
)

func (s *KeeperTestSuite) TestFarm() {
	s.createPair("denom1", "denom2")

	s.createPool(1, utils.ParseCoins("100_000000denom1,100_000000denom2"))

	farmerAddr := utils.TestAddress(0)
	s.farm(farmerAddr, utils.ParseCoin("1_000000pool1"))

	farm, found := s.keeper.GetFarm(s.ctx, "pool1")
	s.Require().True(found)
	s.assertEq(sdk.NewInt(1_000000), farm.TotalFarmingAmount)
	s.assertEq(sdk.DecCoins{}, farm.CurrentRewards)
	s.assertEq(sdk.DecCoins{}, farm.OutstandingRewards)
	s.Require().EqualValues(2, farm.Period)

	position, found := s.keeper.GetPosition(s.ctx, farmerAddr, "pool1")
	s.Require().True(found)
	s.Require().Equal(farmerAddr.String(), position.Farmer)
	s.Require().Equal("pool1", position.Denom)
	s.assertEq(sdk.NewInt(1_000000), position.FarmingAmount)
	s.Require().EqualValues(1, position.PreviousPeriod)

	s.assertHistoricalRewards(map[string]map[uint64]types.HistoricalRewards{
		"pool1": {
			1: {
				CumulativeUnitRewards: sdk.DecCoins{},
				ReferenceCount:        2,
			},
		},
	})

	farmerAddr2 := utils.TestAddress(1)
	s.farm(farmerAddr2, utils.ParseCoin("1_000000pool1"))

	farm, _ = s.keeper.GetFarm(s.ctx, "pool1")
	s.assertEq(sdk.NewInt(2_000000), farm.TotalFarmingAmount)
	s.Require().EqualValues(3, farm.Period)

	position, found = s.keeper.GetPosition(s.ctx, farmerAddr2, "pool1")
	s.Require().True(found)
	s.Require().EqualValues(2, position.PreviousPeriod)

	s.assertHistoricalRewards(map[string]map[uint64]types.HistoricalRewards{
		"pool1": {
			1: {
				CumulativeUnitRewards: sdk.DecCoins{},
				ReferenceCount:        1,
			},
			2: {
				CumulativeUnitRewards: sdk.DecCoins{},
				ReferenceCount:        2,
			},
		},
	})
}

func (s *KeeperTestSuite) TestFarm_ImmediateUnfarm() {
	s.createPairWithLastPrice("denom1", "denom2", sdk.NewDec(1))
	s.createPrivatePlan([]types.RewardAllocation{
		types.NewRewardAllocation(1, utils.ParseCoins("100_000000stake")),
	}, utils.ParseCoins("10000_000000stake"))

	s.createPool(1, utils.ParseCoins("100_000000denom1,100_000000denom2"))
	farmerAddr := utils.TestAddress(0)
	s.farm(farmerAddr, utils.ParseCoin("1_000000pool1"))
	s.assertHistoricalRewards(map[string]map[uint64]types.HistoricalRewards{
		"pool1": {
			1: {
				CumulativeUnitRewards: sdk.DecCoins{},
				ReferenceCount:        2,
			},
		},
	})

	s.nextBlock()
	s.assertEq(utils.ParseDecCoins("5787stake"), s.rewards(farmerAddr, "pool1"))

	farmerAddr2 := utils.TestAddress(1)
	s.farm(farmerAddr2, utils.ParseCoin("1_000000pool1"))
	s.unfarm(farmerAddr2, utils.ParseCoin("1_000000pool1"))
	s.assertEq(utils.ParseCoins("1_000000pool1"), s.getBalances(farmerAddr2))
	s.assertEq(utils.ParseDecCoins("5787stake"), s.rewards(farmerAddr, "pool1"))
	s.assertHistoricalRewards(map[string]map[uint64]types.HistoricalRewards{
		"pool1": {
			1: {
				CumulativeUnitRewards: sdk.DecCoins{},
				ReferenceCount:        1,
			},
			3: {
				CumulativeUnitRewards: utils.ParseDecCoins("0.005787stake"),
				ReferenceCount:        1,
			},
		},
	})
}
