package keeper

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/tendermint/starport/starport/pkg/cosmoscmd"

	"github.com/elesto-dao/elesto/x/did"
)

// TestGenesisExport tests the genesis export functionality of the did module
func (suite *KeeperTestSuite) TestGenesisExport() {
	moduleBasics := module.NewBasicManager()
	cdc := cosmoscmd.MakeEncodingConfig(moduleBasics)
	didID := "did:cosmos:net:elesto:subject"

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
	cdc := cosmoscmd.MakeEncodingConfig(moduleBasics)
	didID := "did:cosmos:net:elesto:subject"
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
