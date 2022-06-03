package keeper

import (
	_ "embed"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elesto-dao/elesto/x/credential"
	"github.com/elesto-dao/elesto/x/did"
)

const (
//signerAccount = "foochainid1sl48sj2jjed7enrv3lzzplr9wc2f5js5khugy3"
// signerAccount = "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"
)

var (
	//go:embed testdata/dummy.schema.json
	dummySchemaOk string
	//go:embed testdata/dummy.vocab.json
	dummyVocabOk string
	//go:embed testdata/dummy.credential.json
	dummyCredential string
	//go:embed testdata/dummy.credential.signed.json
	dummyCredentialSigned string
)

func (suite *KeeperTestSuite) TestHandleMsgPublishCredentialDefinition() {

	server := NewMsgServerImpl(suite.keeper)

	testCases := []struct {
		name    string
		reqFn   func() credential.MsgPublishCredentialDefinitionRequest
		wantErr error
	}{
		{
			"PASS: can publish definition",
			func() credential.MsgPublishCredentialDefinitionRequest {
				return credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: &credential.CredentialDefinition{
						Id:           "did:cosmos:elesto:cd-1",
						PublisherId:  did.NewKeyDID(suite.GetTestAccount().String()).String(),
						Schema:       []byte(dummyCredential),
						Vocab:        []byte(dummyVocabOk),
						Name:         "CredentialDef1",
						Description:  "",
						IsPublic:     true,
						SupersededBy: "",
						IsActive:     true,
					},
					Signer: suite.GetTestAccount().String(),
				}
			},
			nil,
		},
		{
			"FAIL: credential definition exists",
			func() credential.MsgPublishCredentialDefinitionRequest {

				pcdr := credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: &credential.CredentialDefinition{
						Id:           "did:cosmos:elesto:cd-2",
						PublisherId:  did.NewKeyDID(suite.GetTestAccount().String()).String(),
						Schema:       []byte(dummySchemaOk),
						Vocab:        []byte(dummyVocabOk),
						Name:         "CredentialDef2",
						Description:  "",
						IsPublic:     true,
						SupersededBy: "",
						IsActive:     true,
					},
					Signer: suite.GetTestAccount().String(),
				}
				//create the credential definition
				if _, err := server.PublishCredentialDefinition(sdk.WrapSDKContext(suite.ctx), &pcdr); err != nil {
					suite.Require().FailNowf("expected definition to be created, got:", "%v", err)
				}

				return pcdr
			},
			fmt.Errorf("a credential definition with did did:cosmos:elesto:cd-2 already exists: credential definition found"),
		},
		{
			"FAIL: credential definition publisher cannot be resolved",
			func() credential.MsgPublishCredentialDefinitionRequest {
				pcdr := credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: &credential.CredentialDefinition{
						Id:           "did:cosmos:elesto:cd-3",
						PublisherId:  did.NewChainDID(suite.ctx.ChainID(), "non-existing-did").String(),
						Schema:       []byte(dummySchemaOk),
						Vocab:        []byte(dummyVocabOk),
						Name:         "CredentialDef3",
						Description:  "",
						IsPublic:     true,
						SupersededBy: "",
						IsActive:     true,
					},
					Signer: suite.GetTestAccount().String(),
				}
				return pcdr
			},
			fmt.Errorf("the credential publisher DID did:cosmos:foochainid:non-existing-did cannot be resolved: did document found"),
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			req := tc.reqFn()
			_, err := server.PublishCredentialDefinition(sdk.WrapSDKContext(suite.ctx), &req)
			if tc.wantErr == nil {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(tc.wantErr.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestHandleMsgIssuePublicCredential() {

	server := NewMsgServerImpl(suite.keeper)

	testCases := []struct {
		name    string
		reqFn   func() credential.MsgIssuePublicVerifiableCredentialRequest
		wantErr error
	}{
		{
			"PASS: issue public credential",
			func() credential.MsgIssuePublicVerifiableCredentialRequest {

				var (
					wc  *credential.WrappedCredential
					err error
				)
				//

				// publish the definition
				pcdr := credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: &credential.CredentialDefinition{
						Id:           "did:cosmos:elesto:cd-10",
						PublisherId:  did.NewKeyDID(suite.GetTestAccount().String()).String(),
						Schema:       []byte(dummySchemaOk),
						Vocab:        []byte(dummyVocabOk),
						Name:         "CredentialDef10",
						Description:  "",
						IsPublic:     true,
						SupersededBy: "",
						IsActive:     true,
					},
					Signer: suite.GetTestAccount().String(),
				}
				//create the credential definition
				if _, err := server.PublishCredentialDefinition(sdk.WrapSDKContext(suite.ctx), &pcdr); err != nil {
					suite.Require().FailNowf("expected definition to be created, got:", "%v", err)
				}

				// load the signed credential
				if wc, err = credential.NewWrappedPublicCredentialFromFile("testdata/dummy.credential.signed.json"); err != nil {
					suite.Require().FailNowf("expected wrapped credential, got:", "%v", err)
				}
				// return the message
				return credential.MsgIssuePublicVerifiableCredentialRequest{
					CredentialDefinitionDid: "did:cosmos:elesto:cd-10",
					Credential:              wc.PublicVerifiableCredential,
					Signer:                  suite.GetTestAccount().String(),
				}
			},
			nil,
		},
		{
			"FAIL: invalid signature",
			func() credential.MsgIssuePublicVerifiableCredentialRequest {

				var (
					wc  *credential.WrappedCredential
					err error
				)
				//

				// publish the definition
				pcdr := credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: &credential.CredentialDefinition{
						Id:           "did:cosmos:elesto:cd-11",
						PublisherId:  did.NewKeyDID(suite.GetTestAccount().String()).String(),
						Schema:       []byte(dummySchemaOk),
						Vocab:        []byte(dummyVocabOk),
						Name:         "CredentialDef10",
						Description:  "",
						IsPublic:     true,
						SupersededBy: "",
						IsActive:     true,
					},
					Signer: suite.GetTestAccount().String(),
				}
				//create the credential definition
				if _, err := server.PublishCredentialDefinition(sdk.WrapSDKContext(suite.ctx), &pcdr); err != nil {
					suite.Require().FailNowf("expected definition to be created, got:", "%v", err)
				}

				// load the signed credential
				if wc, err = credential.NewWrappedPublicCredentialFromFile("testdata/dummy.credential.signed.json"); err != nil {
					suite.Require().FailNowf("expected wrapped credential, got:", "%v", err)
				}
				// return the message
				return credential.MsgIssuePublicVerifiableCredentialRequest{
					CredentialDefinitionDid: "did:cosmos:elesto:cd-11",
					Credential:              wc.PublicVerifiableCredential,
					Signer:                  suite.GetRandomAccount().String(),
				}
			},
			fmt.Errorf("should fail for wrong signature"),
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			req := tc.reqFn()
			_, err := server.IssuePublicVerifiableCredential(sdk.WrapSDKContext(suite.ctx), &req)
			if tc.wantErr == nil {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(tc.wantErr.Error(), err.Error())
			}
		})
	}
}
