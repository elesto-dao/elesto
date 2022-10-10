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
		genState.AllowedCredentialIds = append(genState.AllowedCredentialIds, cD.Id)
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
	suite.Require().Len(newGenState.AllowedCredentialIds, 5)

	suite.Require().Equal(newGenState.CredentialDefinitions, genState.CredentialDefinitions)
	suite.Require().Equal(newGenState.PublicVerifiableCredentials, genState.PublicVerifiableCredentials)
	suite.Require().Equal(newGenState.AllowedCredentialIds, genState.AllowedCredentialIds)
}

func (s *KeeperTestSuite) Test_Genesis_AllowedCredentials() {
	testCases := []struct {
		name       string
		reqFn      func() *credential.GenesisState
		verifyFunc func(gs *credential.GenesisState)
		wouldPanic bool
	}{
		{
			name: "PANIC: ID does not exist in store",
			reqFn: func() *credential.GenesisState {
				return credential.NewGenesisState("test")
			},
			wouldPanic: true,
		},
		{
			name: "PASS: Valid id's",
			reqFn: func() *credential.GenesisState {
				genState := credential.DefaultGenesisState()
				for i := 0; i < 4; i++ {
					cd := &credential.CredentialDefinition{
						Id:           fmt.Sprintf("did:cosmos:elesto:cd-%v", i),
						PublisherId:  "did:cosmos:elesto:publisher",
						Schema:       schemaOkCompact,
						Vocab:        vocabOkCompact,
						Name:         "Credential Definition",
						Description:  "This is a sample credential",
						IsPublic:     true,
						SupersededBy: "",
						IsActive:     true,
					}
					s.keeper.SetCredentialDefinition(s.ctx, cd)
					genState.AllowedCredentialIds = append(genState.AllowedCredentialIds, cd.Id)
				}

				return genState
			},
			verifyFunc: func(gs *credential.GenesisState) {
				// check whether the allowed ids are properly set in store
				for _, id := range gs.AllowedCredentialIds {
					s.Require().True(s.keeper.IsPublicCredentialDefinitionAllowed(s.ctx, id))
				}
			},
			wouldPanic: false,
		},
		{
			name: "PANIC: ID already allowed",
			reqFn: func() *credential.GenesisState {
				cd := &credential.CredentialDefinition{
					Id:           "did:cosmos:elesto:cd-4",
					PublisherId:  "did:cosmos:elesto:publisher",
					Schema:       schemaOkCompact,
					Vocab:        vocabOkCompact,
					Name:         "Credential Definition",
					Description:  "This is a sample credential",
					IsPublic:     true,
					SupersededBy: "",
					IsActive:     true,
				}
				s.keeper.SetCredentialDefinition(s.ctx, cd)
				s.keeper.AllowPublicCredential(s.ctx, cd.Id)

				return credential.NewGenesisState(cd.Id)
			},
			wouldPanic: true,
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.name), func() {
			genState := tc.reqFn()
			if tc.wouldPanic {
				s.Require().Panics(func() {
					keeper.InitGenesis(s.ctx, s.keeper, genState)
				})
			} else {
				s.Require().NotPanics(func() {
					keeper.InitGenesis(s.ctx, s.keeper, genState)
				})
				tc.verifyFunc(genState)
			}
		})
	}
}
