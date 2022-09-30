package keeper_test

import (
	"fmt"

	"github.com/elesto-dao/elesto/v2/x/credential"
	"github.com/elesto-dao/elesto/v2/x/credential/keeper"
	"github.com/elesto-dao/elesto/v2/x/did"
)

func (s *KeeperTestSuite) TestGenesis() {
	//default genesis on export
	genState := keeper.ExportGenesis(s.ctx, s.keeper)
	s.Require().Equal(genState, &credential.GenesisState{})

	// add credential definitions and pvcs in genesis state
	for i := 0; i < 5; i++ {
		cD := credential.CredentialDefinition{
			Id:           fmt.Sprintf("did:cosmos:elesto:cd-%v", i),
			PublisherId:  did.NewKeyDID(s.GetTestAccount().String()).String(),
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
			s.Require().FailNowf("expected wrapped credential, got:", "%v", err)
		}

		wc.PublicVerifiableCredential.Id = fmt.Sprintf("https://test.xyz/credential/%v", i)
		genState.CredentialDefinitions = append(genState.CredentialDefinitions, cD)
		genState.PublicVerifiableCredentials = append(genState.PublicVerifiableCredentials, *wc.PublicVerifiableCredential)
	}

	// init genesis
	keeper.InitGenesis(s.ctx, s.keeper, genState)

	// check credentials and pvcs in keeper
	for i := 0; i < 5; i++ {
		cd, found := s.keeper.GetCredentialDefinition(s.ctx, fmt.Sprintf("did:cosmos:elesto:cd-%v", i))
		s.Require().True(found)
		s.Require().Equal(genState.CredentialDefinitions[i], cd)

		pvc, found := s.keeper.GetPublicCredential(s.ctx, fmt.Sprintf("https://test.xyz/credential/%v", i))
		s.Require().True(found)
		s.Require().Equal(genState.PublicVerifiableCredentials[i], pvc)
	}

	// check export genesis
	newGenState := keeper.ExportGenesis(s.ctx, s.keeper)
	s.Require().Len(newGenState.PublicVerifiableCredentials, 5)
	s.Require().Len(newGenState.CredentialDefinitions, 5)
	s.Require().Equal(newGenState.CredentialDefinitions, genState.CredentialDefinitions)
	s.Require().Equal(newGenState.PublicVerifiableCredentials, genState.PublicVerifiableCredentials)
}
