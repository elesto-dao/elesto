package keeper

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"

	"github.com/elesto-dao/elesto/v4/x/did"
)

// NOTE: we recreate the encoding config here as we cannot import
//	 "github.com/elesto-dao/elesto/v4/app" into this test

// EncodingConfig specifies the concrete encoding types to use for a given
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

// makeEncodingConfig creates an EncodingConfig for an amino based test configuration.
func makeEncodingConfig() EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(marshaler, tx.DefaultSignModes)

	return EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             amino,
	}
}

// MakeEncodingConfig creates an EncodingConfig for testing
func MakeEncodingConfig(moduleBasics module.BasicManager) EncodingConfig {
	encodingConfig := makeEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	moduleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	moduleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}

// TestGenesisExport tests the genesis export functionality of the did module
func (suite *KeeperTestSuite) TestGenesisExport() {
	moduleBasics := module.NewBasicManager()
	cdc := MakeEncodingConfig(moduleBasics)
	didID := "did:cosmos:elesto:subject"

	dd, _ := did.NewDidDocument(didID)
	suite.keeper.SetDidDocument(suite.ctx, []byte(dd.Id), dd)
	dd2, _ := did.NewDidDocument(didID + "2")
	suite.keeper.SetDidDocument(suite.ctx, []byte(dd2.Id), dd)
	dd3, _ := did.NewDidDocument(didID + "3")
	suite.keeper.SetDidDocument(suite.ctx, []byte(dd3.Id), dd)
	dd4, _ := did.NewDidDocument(didID + "4")
	suite.keeper.SetDidDocument(suite.ctx, []byte(dd4.Id), dd)

	gz := suite.keeper.ExportGenesis(suite.ctx, cdc.Marshaler)

	allEntities := suite.keeper.GetAllDidDocuments(
		suite.ctx,
	)

	suite.Require().Equal(len(gz.DidDocuments), len(allEntities))
}

// TestGenesisExport tests the genesis import functionality of the did module
func (suite *KeeperTestSuite) TestGenesisInit() {
	moduleBasics := module.NewBasicManager()
	cdc := MakeEncodingConfig(moduleBasics)
	didID := "did:cosmos:elesto:subject"
	dd, _ := did.NewDidDocument(didID)
	suite.keeper.SetDidDocument(suite.ctx, []byte(dd.Id), dd)
	dd2, _ := did.NewDidDocument(didID + "2")
	suite.keeper.SetDidDocument(suite.ctx, []byte(dd2.Id), dd)

	gz := suite.keeper.ExportGenesis(suite.ctx, cdc.Marshaler)

	dd, _ = did.NewDidDocument(didID)
	suite.keeper.Delete(
		suite.ctx,
		[]byte(dd.Id),
		[]byte{0x01},
	)

	dd2, _ = did.NewDidDocument(didID + "2")
	suite.keeper.Delete(
		suite.ctx,
		[]byte(dd2.Id),
		[]byte{0x01},
	)
	suite.keeper.InitGenesis(suite.ctx, cdc.Marshaler, cdc.Marshaler.MustMarshalJSON(gz))

	allEntities := suite.keeper.GetAllDidDocuments(
		suite.ctx,
	)

	hasDd := suite.keeper.HasDidDocument(
		suite.ctx,
		[]byte(didID),
	)
	hasDd2 := suite.keeper.HasDidDocument(
		suite.ctx,
		[]byte(didID+"2"),
	)
	suite.Require().True(hasDd)
	suite.Require().True(hasDd2)
	suite.Require().Equal(len(gz.DidDocuments), len(allEntities))

}
