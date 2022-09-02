package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"

	minttypes "github.com/elesto-dao/elesto/v2/x/mint/types"
)

// GenesisParams contains information about the updated genesis attributes
type GenesisParams struct {
	CommunityPool            sdk.DecCoins
	CommunityPoolAccount     []banktypes.Balance
	AirdropAccounts          []banktypes.Balance
	StrategicReserveAccounts []banktypes.Balance
	DevelopmentTeamAccounts  []banktypes.Balance
	PresaleAccounts          []banktypes.Balance
	PartnerAccounts          []banktypes.Balance
	ConsensusParams          *tmproto.ConsensusParams
	GenesisTime              time.Time
	NativeCoinMetadatas      []banktypes.Metadata
	BankGenesis              banktypes.GenesisState
	StakingParams            stakingtypes.Params
	MintParams               minttypes.Params
	DistributionParams       distributiontypes.Params
	GovParams                govtypes.Params
	CrisisConstantFee        sdk.Coin
	SlashingParams           slashingtypes.Params
}

// PrepareGenesisCmd reads a genesis file and updates the genesis with either testnet or mainnet params
func PrepareGenesisCmd(defaultNodeHome string, mbm module.BasicManager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prepare-genesis",
		Short: "Prepare a genesis file with initial setup",
		Long: `Prepare a genesis file with initial setup.
Example:
	elestod prepare-genesis mainnet elesto
	- Check input genesis:
		file is at ~/.elesto/config/genesis.json
`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			depCdc := clientCtx.Codec
			cdc := depCdc
			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			// read genesis file
			genFile := config.GenesisFile()
			appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}

			// get genesis params
			var genesisParams GenesisParams
			network := args[0]
			if network == "testnet" {
				genesisParams = TestnetGenesisParams()
			} else if network == "mainnet" {
				genesisParams = MainnetGenesisParams()
			} else {
				return fmt.Errorf("please choose 'mainnet' or 'testnet'")
			}

			// get genesis params
			chainID := args[1]

			// run Prepare Genesis
			appState, genDoc, err = PrepareGenesis(clientCtx, appState, genDoc, genesisParams, chainID)

			// validate genesis state
			if err = mbm.ValidateGenesis(cdc, clientCtx.TxConfig, appState); err != nil {
				return fmt.Errorf("error validating genesis file: %s", err.Error())
			}

			// save genesis
			appStateJSON, err := json.Marshal(appState)
			if err != nil {
				return fmt.Errorf("failed to marshal application genesis state: %w", err)
			}

			genDoc.AppState = appStateJSON
			err = genutil.ExportGenesisFile(genDoc, genFile)
			return err
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// MainnetGenesisParams updates the genesis file with mainnet params
func MainnetGenesisParams() GenesisParams {
	genParams := GenesisParams{}
	BaseCoinUnit := "uelesto"
	HumanCoinUnit := "elesto"
	// Exponent := 6

	genParams.GenesisTime = time.Date(2022, 11, 04, 17, 0, 0, 0, time.UTC) // Nov 04, 2022 - 17:00 UTC

	genParams.NativeCoinMetadatas = []banktypes.Metadata{
		{
			Description: "The native token of Elesto",
			DenomUnits: []*banktypes.DenomUnit{
				{
					Denom:    BaseCoinUnit,
					Exponent: 0,
					Aliases:  nil,
				},
				{
					Denom:    HumanCoinUnit,
					Exponent: 6,
					Aliases:  nil,
				},
			},
			Base:    BaseCoinUnit,
			Display: HumanCoinUnit,
		},
	}

	// 20 million community pool
	genParams.CommunityPool = sdk.NewDecCoins(sdk.NewDecCoin(genParams.NativeCoinMetadatas[0].Base, sdk.NewInt(20_000_000_000_000)))
	genParams.CommunityPoolAccount = []banktypes.Balance{
		{
			Address: "elesto1laxwuyds8hns5fg44kue9z8xry6355wyywss86",
			Coins:   sdk.NewCoins(sdk.NewCoin(genParams.NativeCoinMetadatas[0].Base, sdk.NewInt(20_000_000_000_000))), // 20 million ELESTO
		},
	}

	// 80 million strategic reserve
	genParams.StrategicReserveAccounts = []banktypes.Balance{
		{
			Address: "elesto1laxwuyds8hns5fg44kue9z8xry6355wyywss86",
			Coins:   sdk.NewCoins(sdk.NewCoin(genParams.NativeCoinMetadatas[0].Base, sdk.NewInt(50_000_000_000_000))), // 50 million ELESTO
		},
		{
			Address: "elesto1hyfeeckq5ely54p36mvxxg5evhkc2znyctp7lc",
			Coins:   sdk.NewCoins(sdk.NewCoin(genParams.NativeCoinMetadatas[0].Base, sdk.NewInt(10_000_000_000_000))), // 10 million ELESTO
		},
		{
			Address: "elesto1ygkzauvssx7a25ldfxa2ehj5jst22y85ddxvu0",
			Coins:   sdk.NewCoins(sdk.NewCoin(genParams.NativeCoinMetadatas[0].Base, sdk.NewInt(10_000_000_000_000))), // 10 million ELESTO
		},
		{
			Address: "elesto18lsqsttnkrw5gs2pc5dx75gpn6uzpvgc2xz32f",
			Coins:   sdk.NewCoins(sdk.NewCoin(genParams.NativeCoinMetadatas[0].Base, sdk.NewInt(5_000_000_000_000))), // 5 million ELESTO
		},
		{
			Address: "elesto1p2nrwmwhh4n8ngdudsnd5uacm7kexxckaaty29",
			Coins:   sdk.NewCoins(sdk.NewCoin(genParams.NativeCoinMetadatas[0].Base, sdk.NewInt(1_000_000_000_000))), // 1 million ELESTO
		},
		{
			Address: "elesto1vvrpupy27h7jywlx07nkw4xgruzmee90flyjkr",
			Coins:   sdk.NewCoins(sdk.NewCoin(genParams.NativeCoinMetadatas[0].Base, sdk.NewInt(1_000_000_000_000))), // 1 million ELESTO
		},
		{
			Address: "elesto1gfnmpzc6th6uejl70ecwwedwmevvrlrfzukjnl",
			Coins:   sdk.NewCoins(sdk.NewCoin(genParams.NativeCoinMetadatas[0].Base, sdk.NewInt(1_000_000_000_000))), // 1 million ELESTO
		},
		{
			Address: "elesto1agxed6zstz79jszrtevfxgxs9vvgg357p73rn5",
			Coins:   sdk.NewCoins(sdk.NewCoin(genParams.NativeCoinMetadatas[0].Base, sdk.NewInt(1_000_000_000_000))), // 1 million ELESTO
		},
		{
			Address: "elesto1mnkwn9adred7kdcpjjnnk77cdv2l4m32u58fne",
			Coins:   sdk.NewCoins(sdk.NewCoin(genParams.NativeCoinMetadatas[0].Base, sdk.NewInt(500_000_000_000))), // 500 thousand ELESTO
		},
		{
			Address: "elesto1a2xzwlcdyzvhuzjqgsus7su3pt7evx0a5swd2r",
			Coins:   sdk.NewCoins(sdk.NewCoin(genParams.NativeCoinMetadatas[0].Base, sdk.NewInt(500_000_000_000))), // 500 thousand ELESTO
		},
	}

	// 20 million airdrop accounts
	genParams.AirdropAccounts = []banktypes.Balance{
		{
			Address: "elesto1zpel8u3ggz3v5emflvjhwysnvzgdqp3c2yk4kh",
			Coins:   sdk.NewCoins(sdk.NewCoin(genParams.NativeCoinMetadatas[0].Base, sdk.NewInt(20_000_000_000_000))), // 20 million ELESTO
		},
	}

	// 20 million development team accounts
	genParams.DevelopmentTeamAccounts = []banktypes.Balance{
		{
			Address: "elesto14ce9ptp65wqknv83cp5l9v67kemqray5n4k08e",
			Coins:   sdk.NewCoins(sdk.NewCoin(genParams.NativeCoinMetadatas[0].Base, sdk.NewInt(20_000_000_000_000))), // 20 million ELESTO
		},
	}

	// 20 million presale accounts
	genParams.PresaleAccounts = []banktypes.Balance{
		{
			Address: "elesto1ztrajkdnceasmxvu5qvzwfftxalvskchcc4754",
			Coins:   sdk.NewCoins(sdk.NewCoin(genParams.NativeCoinMetadatas[0].Base, sdk.NewInt(20_000_000_000_000))), // 20 million ELESTO
		},
	}

	// 20 million partner accounts
	genParams.PartnerAccounts = []banktypes.Balance{
		{
			Address: "elesto1f4eyclp7ehpfzsgrzxmnchqg0x6l8wum4nmwqp",
			Coins:   sdk.NewCoins(sdk.NewCoin(genParams.NativeCoinMetadatas[0].Base, sdk.NewInt(20_000_000_000_000))), // 20 million ELESTO
		},
	}

	genParams.StakingParams = stakingtypes.DefaultParams()
	genParams.StakingParams.UnbondingTime = time.Hour * 24 * 7 * 2 // 2 weeks
	genParams.StakingParams.MaxValidators = 100
	genParams.StakingParams.BondDenom = genParams.NativeCoinMetadatas[0].Base

	genParams.MintParams = minttypes.DefaultParams()
	genParams.MintParams.MintDenom = genParams.NativeCoinMetadatas[0].Base
	genParams.MintParams.BlocksPerYear = 6_308_000
	genParams.MintParams.MaxSupply = 1_000_000_000_000_000
	genParams.MintParams.TeamReward = "0.1"
	genParams.MintParams.TeamAddress = "elesto1ms2wrq8k04cug7ea6ekf60nfke6a8vu8pwm684"

	genParams.DistributionParams = distributiontypes.DefaultParams()
	genParams.DistributionParams.BaseProposerReward = sdk.MustNewDecFromStr("0.01")
	genParams.DistributionParams.BonusProposerReward = sdk.MustNewDecFromStr("0.04")
	genParams.DistributionParams.CommunityTax = sdk.MustNewDecFromStr("0.11") // 11%
	genParams.DistributionParams.WithdrawAddrEnabled = true

	genParams.GovParams = govtypes.DefaultParams()
	genParams.GovParams.DepositParams.MaxDepositPeriod = time.Hour * 24 * 14 // 2 weeks
	genParams.GovParams.DepositParams.MinDeposit = sdk.NewCoins(sdk.NewCoin(
		genParams.NativeCoinMetadatas[0].Base,
		sdk.NewInt(2_500_000_000),
	))
	genParams.GovParams.TallyParams.Quorum = sdk.MustNewDecFromStr("0.2") // 20%
	genParams.GovParams.VotingParams.VotingPeriod = time.Hour * 24 * 3    // 3 days

	genParams.CrisisConstantFee = sdk.NewCoin(
		genParams.NativeCoinMetadatas[0].Base,
		sdk.NewInt(500_000_000_000),
	)

	genParams.SlashingParams = slashingtypes.DefaultParams()
	genParams.SlashingParams.SignedBlocksWindow = int64(60000)                       // 30000 blocks (~41 hr at 5 second blocks)
	genParams.SlashingParams.MinSignedPerWindow = sdk.MustNewDecFromStr("0.05")      // 5% minimum liveness
	genParams.SlashingParams.DowntimeJailDuration = time.Minute                      // 1 minute jail period
	genParams.SlashingParams.SlashFractionDoubleSign = sdk.MustNewDecFromStr("0.05") // 5% double sign slashing
	genParams.SlashingParams.SlashFractionDowntime = sdk.ZeroDec()                   // 0% liveness slashing

	genParams.ConsensusParams = tmtypes.DefaultConsensusParams()
	genParams.ConsensusParams.Block.MaxBytes = 1 * 1024 * 1024 // 1mb
	genParams.ConsensusParams.Block.MaxGas = 6_000_000
	genParams.ConsensusParams.Evidence.MaxAgeDuration = genParams.StakingParams.UnbondingTime
	genParams.ConsensusParams.Evidence.MaxAgeNumBlocks = int64(genParams.StakingParams.UnbondingTime.Seconds()) / 3
	genParams.ConsensusParams.Version.AppVersion = 1

	return genParams
}

// TestnetGenesisParams updates the genesis file with testnet params
func TestnetGenesisParams() GenesisParams {
	genParams := MainnetGenesisParams()

	genParams.GenesisTime = time.Now()

	return genParams
}

// PrepareGenesis adds the new params to the genesis file
func PrepareGenesis(
	clientCtx client.Context,
	appState map[string]json.RawMessage,
	genDoc *tmtypes.GenesisDoc,
	genesisParams GenesisParams,
	chainID string,
) (map[string]json.RawMessage, *tmtypes.GenesisDoc, error) {
	depCdc := clientCtx.Codec
	cdc := depCdc

	// chain params genesis
	genDoc.ChainID = chainID
	genDoc.GenesisTime = genesisParams.GenesisTime

	genDoc.ConsensusParams = genesisParams.ConsensusParams

	// ---
	// auth state genesis
	authGenState := authtypes.GetGenesisStateFromAppState(depCdc, appState)
	bankGenState := banktypes.GetGenesisStateFromAppState(depCdc, appState)

	accs, err := authtypes.UnpackAccounts(authGenState.Accounts)

	// add stragetic reserve accounts
	for _, acc := range genesisParams.StrategicReserveAccounts {
		genAccount, balances, err := prepareGenesisAccount(sdk.MustAccAddressFromBech32(acc.Address), acc.Coins)
		if err != nil {
			panic(err)
		}
		accs = append(accs, genAccount)
		bankGenState.Balances = append(bankGenState.Balances, balances)
		bankGenState.Supply = bankGenState.Supply.Add(balances.Coins...)
	}

	// add airdrop reserve accounts
	for _, acc := range genesisParams.AirdropAccounts {
		genAccount, balances, err := prepareGenesisAccount(sdk.MustAccAddressFromBech32(acc.Address), acc.Coins)
		if err != nil {
			panic(err)
		}
		accs = append(accs, genAccount)
		bankGenState.Balances = append(bankGenState.Balances, balances)
		bankGenState.Supply = bankGenState.Supply.Add(balances.Coins...)
	}

	// add development team accounts
	for _, acc := range genesisParams.DevelopmentTeamAccounts {
		genAccount, balances, err := prepareGenesisAccount(sdk.MustAccAddressFromBech32(acc.Address), acc.Coins)
		if err != nil {
			panic(err)
		}
		accs = append(accs, genAccount)
		bankGenState.Balances = append(bankGenState.Balances, balances)
		bankGenState.Supply = bankGenState.Supply.Add(balances.Coins...)
	}

	// add presale accounts
	for _, acc := range genesisParams.PresaleAccounts {
		genAccount, balances, err := prepareGenesisAccount(sdk.MustAccAddressFromBech32(acc.Address), acc.Coins)
		if err != nil {
			panic(err)
		}
		accs = append(accs, genAccount)
		bankGenState.Balances = append(bankGenState.Balances, balances)
		bankGenState.Supply = bankGenState.Supply.Add(balances.Coins...)
	}

	// add partner accounts
	for _, acc := range genesisParams.PartnerAccounts {
		genAccount, balances, err := prepareGenesisAccount(sdk.MustAccAddressFromBech32(acc.Address), acc.Coins)
		if err != nil {
			panic(err)
		}
		accs = append(accs, genAccount)
		bankGenState.Balances = append(bankGenState.Balances, balances)
		bankGenState.Supply = bankGenState.Supply.Add(balances.Coins...)
	}

	// add community pool balance to distribution account
	distrInitialBalances := banktypes.Balance{
		Address: authtypes.NewModuleAddress(distributiontypes.ModuleName).String(),
		Coins:   genesisParams.CommunityPoolAccount[0].Coins.Sort(),
	}
	distrModuleAccount := authtypes.NewEmptyModuleAccount(distributiontypes.ModuleName)
	accs = append(accs, distrModuleAccount)
	bankGenState.Balances = append(bankGenState.Balances, distrInitialBalances)
	bankGenState.Supply = bankGenState.Supply.Add(distrInitialBalances.Coins...)

	// TODO: validate total supply === 180_000_000_000_000

	accs = authtypes.SanitizeGenesisAccounts(accs)
	genAccs, err := authtypes.PackAccounts(accs)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to convert accounts into any's: %w", err)
	}
	authGenState.Accounts = genAccs

	authGenStateBz, err := cdc.MarshalJSON(&authGenState)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal auth genesis state: %w", err)
	}
	appState[authtypes.ModuleName] = authGenStateBz

	// ---
	// bank state genesis
	bankGenState.Balances = banktypes.SanitizeGenesisBalances(bankGenState.Balances)

	bankGenStateBz, err := cdc.MarshalJSON(bankGenState)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal bank genesis state: %w", err)
	}

	appState[banktypes.ModuleName] = bankGenStateBz

	// ---
	// staking module genesis
	stakingGenState := stakingtypes.GetGenesisStateFromAppState(depCdc, appState)
	stakingGenState.Params = genesisParams.StakingParams
	stakingGenStateBz, err := cdc.MarshalJSON(stakingGenState)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal staking genesis state: %w", err)
	}
	appState[stakingtypes.ModuleName] = stakingGenStateBz

	// mint module genesis
	mintGenState := minttypes.DefaultGenesisState()
	mintGenState.Params = genesisParams.MintParams
	mintGenStateBz, err := cdc.MarshalJSON(mintGenState)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal mint genesis state: %w", err)
	}
	appState[minttypes.ModuleName] = mintGenStateBz

	// distribution module genesis
	distributionGenState := distributiontypes.DefaultGenesisState()
	distributionGenState.Params = genesisParams.DistributionParams
	distributionGenState.FeePool.CommunityPool = genesisParams.CommunityPool
	distributionGenStateBz, err := cdc.MarshalJSON(distributionGenState)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal distribution genesis state: %w", err)
	}
	appState[distributiontypes.ModuleName] = distributionGenStateBz

	// gov module genesis
	govGenState := govtypes.DefaultGenesisState()
	govGenState.DepositParams = genesisParams.GovParams.DepositParams
	govGenState.TallyParams = genesisParams.GovParams.TallyParams
	govGenState.VotingParams = genesisParams.GovParams.VotingParams
	govGenStateBz, err := cdc.MarshalJSON(govGenState)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal gov genesis state: %w", err)
	}
	appState[govtypes.ModuleName] = govGenStateBz

	// crisis module genesis
	crisisGenState := crisistypes.DefaultGenesisState()
	crisisGenState.ConstantFee = genesisParams.CrisisConstantFee
	crisisGenStateBz, err := cdc.MarshalJSON(crisisGenState)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal crisis genesis state: %w", err)
	}
	appState[crisistypes.ModuleName] = crisisGenStateBz

	// slashing module genesis
	slashingGenState := slashingtypes.DefaultGenesisState()
	slashingGenState.Params = genesisParams.SlashingParams
	slashingGenStateBz, err := cdc.MarshalJSON(slashingGenState)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal slashing genesis state: %w", err)
	}
	appState[slashingtypes.ModuleName] = slashingGenStateBz

	// return appState and genDoc
	return appState, genDoc, nil
}

func prepareGenesisAccount(addr sdk.AccAddress, coins sdk.Coins) (authtypes.GenesisAccount, banktypes.Balance, error) {
	// create concrete account type based on input parameters
	var genAccount authtypes.GenesisAccount

	balances := banktypes.Balance{Address: addr.String(), Coins: coins.Sort()}
	genAccount = authtypes.NewBaseAccount(addr, nil, 0, 0)

	// TODO: vesting schedules

	if genAccountValidateErr := genAccount.Validate(); genAccountValidateErr != nil {
		return nil, banktypes.Balance{}, fmt.Errorf("failed to validate new genesis account")
	}

	return genAccount, balances, nil
}
