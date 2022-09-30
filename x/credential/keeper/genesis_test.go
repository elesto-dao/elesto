package keeper

import (
	"fmt"

	"github.com/elesto-dao/elesto/v2/x/credential"
)

func (s *KeeperTestSuite) Test_Genesis() {
	testCases := []struct {
		name       string
		reqFn      func() *credential.GenesisState
		verifyFunc func(gs *credential.GenesisState)
		wouldPanic bool
	}{
		//{
		//	name: "PANIC: ID does not exist in store",
		//	reqFn: func() *credential.GenesisState {
		//		return credential.NewGenesisState("test")
		//	},
		//	wouldPanic: true,
		//},
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
				gsFromKeeper := ExportGenesis(s.ctx, s.keeper)

				// check export genesis
				s.Require().Equal(gs, gsFromKeeper)

				// check keeper
				for _, id := range gs.AllowedCredentialIds {
					s.Require().True(s.keeper.IsPublicCredentialDefinitionAllowed(s.ctx, id))
				}
			},
			wouldPanic: false,
		},
		//{
		//	name: "PANIC: ID already allowed",
		//	reqFn: func() *credential.GenesisState {
		//		cd := &credential.CredentialDefinition{
		//			Id:           "did:cosmos:elesto:cd-4",
		//			PublisherId:  "did:cosmos:elesto:publisher",
		//			Schema:       schemaOkCompact,
		//			Vocab:        vocabOkCompact,
		//			Name:         "Credential Definition",
		//			Description:  "This is a sample credential",
		//			IsPublic:     true,
		//			SupersededBy: "",
		//			IsActive:     true,
		//		}
		//		s.keeper.SetCredentialDefinition(s.ctx, cd)
		//		s.keeper.SetAllowedPublicCredential(s.ctx, cd.Id)
		//
		//		return credential.NewGenesisState(cd.Id)
		//	},
		//	wouldPanic: true,
		//},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.name), func() {
			genState := tc.reqFn()
			if tc.wouldPanic {
				s.Require().Panics(func() {
					InitGenesis(s.ctx, s.keeper, genState)
				})
			} else {
				s.Require().NotPanics(func() {
					InitGenesis(s.ctx, s.keeper, genState)
				})
				tc.verifyFunc(genState)
			}
		})
	}
}
