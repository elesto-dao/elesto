package types

type InflationDistribution struct {
	BlockInflation int64
	TeamRewards    int64
	CommunityTax   int64
	StakingRewards int64
}

var (
	// BlockInflationDistribution are the amounts to be minted and distributed for each block in an epoch.
	BlockInflationDistribution = map[int64]InflationDistribution{
		0: {BlockInflation: 31_709_792, TeamRewards: 3_170_979, CommunityTax: 3_170_979, StakingRewards: 25_367_834},
		1: {BlockInflation: 31_709_792, TeamRewards: 3_170_979, CommunityTax: 3_170_979, StakingRewards: 25_367_834},
		2: {BlockInflation: 23_782_344, TeamRewards: 2_378_234, CommunityTax: 2_378_234, StakingRewards: 19_025_876},
		3: {BlockInflation: 14_863_965, TeamRewards: 1_486_396, CommunityTax: 1_486_396, StakingRewards: 11_891_173},
		4: {BlockInflation: 8_360_980, TeamRewards: 836_098, CommunityTax: 836_098, StakingRewards: 6_688_784},
		5: {BlockInflation: 4_441_771, TeamRewards: 444_177, CommunityTax: 444_177, StakingRewards: 3_553_417},
		6: {BlockInflation: 2_931_569, TeamRewards: 293_156, CommunityTax: 293_156, StakingRewards: 2_345_257},
		7: {BlockInflation: 2_990_200, TeamRewards: 299_020, CommunityTax: 299_020, StakingRewards: 2_392_160},
		8: {BlockInflation: 3_050_004, TeamRewards: 305_000, CommunityTax: 305_000, StakingRewards: 2_440_004},
		9: {BlockInflation: 2_998_751, TeamRewards: 299_875, CommunityTax: 299_875, StakingRewards: 2_399_001},
	}
)
