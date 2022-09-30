package keeper_test

import (
	"fmt"

	"github.com/elesto-dao/elesto/v3/x/credential"
	"github.com/elesto-dao/elesto/v3/x/credential/keeper"
	"github.com/elesto-dao/elesto/v3/x/did"
)

func (suite *KeeperTestSuite) TestGenesis() {
	//default genesis on export
	genState := keeper.ExportGenesis(suite.ctx, suite.keeper)
	suite.Require().Equal(genState, &credential.GenesisState{})

	// add credential definitions and pvcs in genesis state
	for i := 0; i < 5; i++ {
		cD := credential.CredentialDefinition{
			Id:           fmt.Sprintf("did:cosmos:elesto:cd-%v", i),
			PublisherId:  did.NewKeyDID(suite.GetTestAccount().String()).String(),
			Schema:       []byte(dummyCredential),
			Vocab:        []byte(dummyVocabOk),
			Name:         "CredentialDef1",
			Description:  "",
			IsPublic:     true,
			SupersededBy: "",
			IsActive:     true,
		}

		// load the signed credential
		wc, err := credential.NewWrappedPublicCredentialFromFile("testdata/dummy.credential.signed.json")
		if err != nil {
			suite.Require().FailNowf("expected wrapped credential, got:", "%v", err)
		}

		wc.PublicVerifiableCredential.Id = fmt.Sprintf("https://test.xyz/credential/%v", i)
		genState.CredentialDefinitions = append(genState.CredentialDefinitions, cD)
		genState.PublicVerifiableCredentials = append(genState.PublicVerifiableCredentials, *wc.PublicVerifiableCredential)
	}

	// init genesis
	keeper.InitGenesis(suite.ctx, suite.keeper, genState)

	// check credentials and pvcs in keeper
	for i := 0; i < 5; i++ {
		cd, found := suite.keeper.GetCredentialDefinition(suite.ctx, fmt.Sprintf("did:cosmos:elesto:cd-%v", i))
		suite.Require().True(found)
		suite.Require().Equal(genState.CredentialDefinitions[i], cd)

		pvc, found := suite.keeper.GetPublicCredential(suite.ctx, fmt.Sprintf("https://test.xyz/credential/%v", i))
		suite.Require().True(found)
		suite.Require().Equal(genState.PublicVerifiableCredentials[i], pvc)
	}

	// check export genesis
	newGenState := keeper.ExportGenesis(suite.ctx, suite.keeper)
	suite.Require().Len(newGenState.PublicVerifiableCredentials, 5)
	suite.Require().Len(newGenState.CredentialDefinitions, 5)
	suite.Require().Equal(newGenState.CredentialDefinitions, genState.CredentialDefinitions)
	suite.Require().Equal(newGenState.PublicVerifiableCredentials, genState.PublicVerifiableCredentials)
}
