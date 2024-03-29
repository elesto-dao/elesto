package keeper_test

import (
	"context"
	_ "embed"
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elesto-dao/elesto/v4/x/credential"
	"github.com/elesto-dao/elesto/v4/x/credential/keeper"
	"github.com/elesto-dao/elesto/v4/x/did"
)

var (
	//go:embed testdata/dummy.schema.json
	dummySchemaOk string
	//go:embed testdata/dummy.vocab.json
	dummyVocabOk string
	//go:embed testdata/dummy.credential.json
	dummyCredential string
)

func (suite *KeeperTestSuite) TestHandleMsgPublishCredentialDefinition() {

	server := keeper.NewMsgServerImpl(suite.keeper)

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
						SupersededBy: "",
						IsActive:     true,
					},
					Signer: suite.GetTestAccount().String(),
				}
				return pcdr
			},
			fmt.Errorf("the credential publisher DID did:cosmos:foochainid:non-existing-did cannot be resolved: did document found"),
		},
		{
			"FAIL: credential definition not set",
			func() credential.MsgPublishCredentialDefinitionRequest {
				pcdr := credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: nil,
					Signer:               suite.GetTestAccount().String(),
				}
				return pcdr
			},
			fmt.Errorf("credential definition not set: input is invalid"),
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			req := tc.reqFn()
			req.ValidateBasic()

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

func (suite *KeeperTestSuite) TestHandleMsgUpdateCredentialDefinition() {

	server := keeper.NewMsgServerImpl(suite.keeper)
	query := suite.queryClient

	testCases := []struct {
		name         string
		reqFn        func() (*credential.MsgUpdateCredentialDefinitionRequest, *credential.CredentialDefinition)
		wantBasicErr error
		wantErr      error
	}{
		{
			"PASS: can update definition active",
			func() (*credential.MsgUpdateCredentialDefinitionRequest, *credential.CredentialDefinition) {
				cd := &credential.CredentialDefinition{
					Id:           "did:cosmos:elesto:cd-update-01",
					PublisherId:  did.NewKeyDID(suite.GetTestAccount().String()).String(),
					Schema:       []byte(dummySchemaOk),
					Vocab:        []byte(dummyVocabOk),
					Name:         "CredentialDef10",
					Description:  "",
					SupersededBy: "",
					IsActive:     true,
				}
				// publish the definition
				pcdr := credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: cd,
					Signer:               suite.GetTestAccount().String(),
				}
				//create the credential definition
				if _, err := server.PublishCredentialDefinition(sdk.WrapSDKContext(suite.ctx), &pcdr); err != nil {
					suite.Require().FailNowf("expected definition to be created, got:", "%v", err)
				}
				// message to update the credential definition
				mucdr := &credential.MsgUpdateCredentialDefinitionRequest{
					CredentialDefinitionID: "did:cosmos:elesto:cd-update-01",
					Active:                 false,
					SupersededBy:           "",
					Signer:                 suite.GetTestAccount().String(),
				}
				// expected credential definition
				cd.IsActive = false

				return mucdr, cd
			},
			nil,
			nil,
		},
		{
			"PASS: can update definition active/superseded by",
			func() (*credential.MsgUpdateCredentialDefinitionRequest, *credential.CredentialDefinition) {
				cd := &credential.CredentialDefinition{
					Id:           "did:cosmos:elesto:cd-update-02",
					PublisherId:  did.NewKeyDID(suite.GetTestAccount().String()).String(),
					Schema:       []byte(dummySchemaOk),
					Vocab:        []byte(dummyVocabOk),
					Name:         "CredentialDef10",
					Description:  "",
					SupersededBy: "",
					IsActive:     true,
				}
				// publish the definition
				pcdr := credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: cd,
					Signer:               suite.GetTestAccount().String(),
				}
				//create the credential definition
				if _, err := server.PublishCredentialDefinition(sdk.WrapSDKContext(suite.ctx), &pcdr); err != nil {
					suite.Require().FailNowf("expected definition to be created, got:", "%v", err)
				}
				// message to update the credential definition
				mucdr := &credential.MsgUpdateCredentialDefinitionRequest{
					CredentialDefinitionID: "did:cosmos:elesto:cd-update-02",
					Active:                 false,
					SupersededBy:           "did:cosmos:elesto:cd-update-01",
					Signer:                 suite.GetTestAccount().String(),
				}
				// expected credential definition
				cd.IsActive = false
				cd.SupersededBy = "did:cosmos:elesto:cd-update-01"

				return mucdr, cd
			},
			nil,
			nil,
		},
		{
			"FAIL: SupersededBy definition not found",
			func() (*credential.MsgUpdateCredentialDefinitionRequest, *credential.CredentialDefinition) {
				cd := &credential.CredentialDefinition{
					Id:           "did:cosmos:elesto:cd-update-03",
					PublisherId:  did.NewKeyDID(suite.GetTestAccount().String()).String(),
					Schema:       []byte(dummySchemaOk),
					Vocab:        []byte(dummyVocabOk),
					Name:         "CredentialDef10",
					Description:  "",
					SupersededBy: "",
					IsActive:     true,
				}
				// publish the definition
				pcdr := credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: cd,
					Signer:               suite.GetTestAccount().String(),
				}
				//create the credential definition
				if _, err := server.PublishCredentialDefinition(sdk.WrapSDKContext(suite.ctx), &pcdr); err != nil {
					suite.Require().FailNowf("expected definition to be created, got:", "%v", err)
				}
				// message to update the credential definition
				mucdr := &credential.MsgUpdateCredentialDefinitionRequest{
					CredentialDefinitionID: "did:cosmos:elesto:cd-update-03",
					Active:                 false,
					SupersededBy:           "did:cosmos:elesto:cd-update-non-existing",
					Signer:                 suite.GetTestAccount().String(),
				}
				// expected credential definition
				cd.IsActive = false
				cd.SupersededBy = "did:cosmos:elesto:cd-update-01"

				return mucdr, cd
			},
			nil,
			errors.New("credential definition did:cosmos:elesto:cd-update-non-existing not found: credential definition not found"),
		},
		{
			"FAIL: cd not found",
			func() (*credential.MsgUpdateCredentialDefinitionRequest, *credential.CredentialDefinition) {
				// message to update the credential definition
				mucdr := &credential.MsgUpdateCredentialDefinitionRequest{
					CredentialDefinitionID: "did:cosmos:elesto:cd-update-non-existing",
					Active:                 false,
					SupersededBy:           "did:cosmos:elesto:cd-update-01",
					Signer:                 suite.GetTestAccount().String(),
				}
				// expected credential definition
				return mucdr, nil
			},
			nil,
			errors.New("credential definition did:cosmos:elesto:cd-update-non-existing does not exists: credential definition not found"),
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			req, expectedCD := tc.reqFn()

			err := req.ValidateBasic()
			if tc.wantBasicErr == nil {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(tc.wantBasicErr.Error(), err.Error())
				return
			}

			_, err = server.UpdateCredentialDefinition(sdk.WrapSDKContext(suite.ctx), req)
			if tc.wantErr == nil {
				suite.Require().NoError(err)
				r, qErr := query.CredentialDefinition(context.Background(), &credential.QueryCredentialDefinitionRequest{Id: expectedCD.Id})
				suite.Require().NoError(qErr)
				suite.Require().Equal(expectedCD, r.Definition)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(tc.wantErr.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestHandleMsgIssuePublicCredential() {

	server := keeper.NewMsgServerImpl(suite.keeper)

	testCases := []struct {
		name         string
		reqFn        func() credential.MsgIssuePublicVerifiableCredentialRequest
		wantBasicErr error
		wantErr      error
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
					CredentialDefinitionID: "did:cosmos:elesto:cd-10",
					Credential:             wc.PublicVerifiableCredential,
					Signer:                 suite.GetTestAccount().String(),
				}
			},
			nil,
			nil,
		},
		{
			"PASS: tx signer != credential issuer",
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
					CredentialDefinitionID: "did:cosmos:elesto:cd-11",
					Credential:             wc.PublicVerifiableCredential,
					Signer:                 suite.GetRandomAccount().String(),
				}
			},
			nil,
			nil,
		},
		{
			"FAIL: credential definition not found",
			func() credential.MsgIssuePublicVerifiableCredentialRequest {

				var (
					wc  *credential.WrappedCredential
					err error
				)
				// load the signed credential
				if wc, err = credential.NewWrappedPublicCredentialFromFile("testdata/dummy.credential.signed.json"); err != nil {
					suite.Require().FailNowf("expected wrapped credential, got:", "%v", err)
				}
				// return the message
				return credential.MsgIssuePublicVerifiableCredentialRequest{
					CredentialDefinitionID: "did:cosmos:elesto:cd-non-existing",
					Credential:             wc.PublicVerifiableCredential,
					Signer:                 suite.GetRandomAccount().String(),
				}
			},
			nil,
			fmt.Errorf("credential definition did:cosmos:elesto:cd-non-existing not found: credential definition not found"),
		},
		{
			"FAIL: credential definition is not active",
			func() credential.MsgIssuePublicVerifiableCredentialRequest {

				var (
					wc  *credential.WrappedCredential
					err error
				)
				//

				// publish the definition
				pcdr := credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: &credential.CredentialDefinition{
						Id:           "did:cosmos:elesto:cd-13",
						PublisherId:  did.NewKeyDID(suite.GetTestAccount().String()).String(),
						Schema:       []byte(dummySchemaOk),
						Vocab:        []byte(dummyVocabOk),
						Name:         "CredentialDef13",
						Description:  "",
						SupersededBy: "",
						IsActive:     false,
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
					CredentialDefinitionID: "did:cosmos:elesto:cd-13",
					Credential:             wc.PublicVerifiableCredential,
					Signer:                 suite.GetTestAccount().String(),
				}
			},
			nil,
			fmt.Errorf("the credential definition did:cosmos:elesto:cd-13 issuance is suspended: credential cannot be issued on-chain"),
		},
		{
			"FAIL: credential signature invalid",
			func() credential.MsgIssuePublicVerifiableCredentialRequest {

				var (
					wc  *credential.WrappedCredential
					err error
				)
				//

				// publish the definition
				pcdr := credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: &credential.CredentialDefinition{
						Id:           "did:cosmos:elesto:cd-14",
						PublisherId:  did.NewKeyDID(suite.GetTestAccount().String()).String(),
						Schema:       []byte(dummySchemaOk),
						Vocab:        []byte(dummyVocabOk),
						Name:         "CredentialDef14",
						Description:  "",
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
				if wc, err = credential.NewWrappedPublicCredentialFromFile("testdata/dummy.credential.signed.wrong.json"); err != nil {
					suite.Require().FailNowf("expected wrapped credential, got:", "%v", err)
				}
				// return the message
				return credential.MsgIssuePublicVerifiableCredentialRequest{
					CredentialDefinitionID: "did:cosmos:elesto:cd-14",
					Credential:             wc.PublicVerifiableCredential,
					Signer:                 suite.GetTestAccount().String(),
				}
			},
			nil,
			fmt.Errorf("signature cannot be verified: credential proof validation error"),
		},
		{
			"FAIL: credential proof missing",
			func() credential.MsgIssuePublicVerifiableCredentialRequest {

				var (
					wc  *credential.WrappedCredential
					err error
				)
				//

				// publish the definition
				pcdr := credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: &credential.CredentialDefinition{
						Id:           "did:cosmos:elesto:cd-15",
						PublisherId:  did.NewKeyDID(suite.GetTestAccount().String()).String(),
						Schema:       []byte(dummySchemaOk),
						Vocab:        []byte(dummyVocabOk),
						Name:         "CredentialDef15",
						Description:  "",
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
				if wc, err = credential.NewWrappedPublicCredentialFromFile("testdata/dummy.credential.json"); err != nil {
					suite.Require().FailNowf("expected wrapped credential, got:", "%v", err)
				}
				// return the message
				return credential.MsgIssuePublicVerifiableCredentialRequest{
					CredentialDefinitionID: "did:cosmos:elesto:cd-15",
					Credential:             wc.PublicVerifiableCredential,
					Signer:                 suite.GetTestAccount().String(),
				}
			},
			nil,
			fmt.Errorf("missing credential proof: credential proof validation error"),
		},
		{
			"FAIL: credential not set",
			func() credential.MsgIssuePublicVerifiableCredentialRequest {

				// publish the definition
				pcdr := credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: &credential.CredentialDefinition{
						Id:           "did:cosmos:elesto:cd-16",
						PublisherId:  did.NewKeyDID(suite.GetTestAccount().String()).String(),
						Schema:       []byte(dummySchemaOk),
						Vocab:        []byte(dummyVocabOk),
						Name:         "CredentialDef16",
						Description:  "",
						SupersededBy: "",
						IsActive:     true,
					},
					Signer: suite.GetTestAccount().String(),
				}
				//create the credential definition
				if _, err := server.PublishCredentialDefinition(sdk.WrapSDKContext(suite.ctx), &pcdr); err != nil {
					suite.Require().FailNowf("expected definition to be created, got:", "%v", err)
				}

				// return the message
				return credential.MsgIssuePublicVerifiableCredentialRequest{
					CredentialDefinitionID: "did:cosmos:elesto:cd-16",
					Credential:             nil,
					Signer:                 suite.GetTestAccount().String(),
				}
			},
			fmt.Errorf("credential must be set"),
			nil,
		},
		{
			"FAIL: credential invalid",
			func() credential.MsgIssuePublicVerifiableCredentialRequest {

				// publish the definition
				pcdr := credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: &credential.CredentialDefinition{
						Id:           "did:cosmos:elesto:cd-17",
						PublisherId:  did.NewKeyDID(suite.GetTestAccount().String()).String(),
						Schema:       []byte(dummySchemaOk),
						Vocab:        []byte(dummyVocabOk),
						Name:         "CredentialDef17",
						Description:  "",
						SupersededBy: "",
						IsActive:     true,
					},
					Signer: suite.GetTestAccount().String(),
				}
				//create the credential definition
				if _, err := server.PublishCredentialDefinition(sdk.WrapSDKContext(suite.ctx), &pcdr); err != nil {
					suite.Require().FailNowf("expected definition to be created, got:", "%v", err)
				}

				// return the message
				return credential.MsgIssuePublicVerifiableCredentialRequest{
					CredentialDefinitionID: "did:cosmos:elesto:cd-17",
					Credential:             &credential.PublicVerifiableCredential{},
					Signer:                 suite.GetTestAccount().String(),
				}
			},
			nil,
			fmt.Errorf("schema: did:cosmos:elesto:cd-17, errors: [(root): @context is required (root): type is required (root): issuer is required (root): issuanceDate is required credentialSubject: id is required]: the credential doesn't match the definition schema"),
		},
		{
			"PASS: issue public credential (CosmosADR0036 signature type)",
			func() credential.MsgIssuePublicVerifiableCredentialRequest {

				var (
					wc  *credential.WrappedCredential
					err error
				)
				//

				// publish the definition
				pcdr := credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: &credential.CredentialDefinition{
						Id:           "did:cosmos:elesto:cd-18",
						PublisherId:  did.NewKeyDID(suite.GetTestAccount().String()).String(),
						Schema:       []byte(dummySchemaOk),
						Vocab:        []byte(dummyVocabOk),
						Name:         "CredentialDef18",
						Description:  "",
						SupersededBy: "",
						IsActive:     true,
					},
					Signer: suite.GetTestAccount().String(),
				}
				//create the credential definition
				_, err = server.PublishCredentialDefinition(sdk.WrapSDKContext(suite.ctx), &pcdr)
				suite.Require().NoError(err, "expected definition to be created")

				// load the signed credential
				wc, err = credential.NewWrappedPublicCredentialFromFile("testdata/dummy.credential.signed.cosmosadr036.json")
				suite.Require().NoError(err, "expected wrapped credential")

				// return the message
				return credential.MsgIssuePublicVerifiableCredentialRequest{
					CredentialDefinitionID: "did:cosmos:elesto:cd-18",
					Credential:             wc.PublicVerifiableCredential,
					Signer:                 suite.GetTestAccount().String(),
				}
			},
			nil,
			nil,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			req := tc.reqFn()

			err := req.ValidateBasic()
			if tc.wantBasicErr == nil {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(tc.wantBasicErr.Error(), err.Error())
				return
			}

			suite.keeper.AllowPublicCredential(suite.ctx, req.CredentialDefinitionID)
			_, err = server.IssuePublicVerifiableCredential(sdk.WrapSDKContext(suite.ctx), &req)
			if tc.wantErr == nil {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(tc.wantErr.Error(), err.Error())
			}
		})
	}
}
