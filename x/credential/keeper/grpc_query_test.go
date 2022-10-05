package keeper_test

import (
	"context"
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/elesto-dao/elesto/v3/x/credential"
	"github.com/elesto-dao/elesto/v3/x/credential/keeper"
	"github.com/elesto-dao/elesto/v3/x/did"
)

func (suite *KeeperTestSuite) TestKeeper_CredentialDefinition() {
	queryClient := suite.queryClient
	server := keeper.NewMsgServerImpl(suite.keeper)

	testCases := []struct {
		msg     string
		reqFn   func() (*credential.QueryCredentialDefinitionRequest, *credential.QueryCredentialDefinitionResponse)
		wantErr error
	}{
		{
			"PASS: can get the credential definition",
			func() (*credential.QueryCredentialDefinitionRequest, *credential.QueryCredentialDefinitionResponse) {
				// publish the definition
				pcdr := credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: &credential.CredentialDefinition{
						Id:           "did:cosmos:elesto:cd-query-001",
						PublisherId:  did.NewKeyDID(suite.GetTestAccount().String()).String(),
						Schema:       []byte(dummySchemaOk),
						Vocab:        []byte(dummyVocabOk),
						Name:         "CredentialDef001",
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
				return &credential.QueryCredentialDefinitionRequest{Id: "did:cosmos:elesto:cd-query-001"}, &credential.QueryCredentialDefinitionResponse{Definition: pcdr.CredentialDefinition}
			},
			nil,
		},
		{
			"FAIL: credential definition not found",
			func() (*credential.QueryCredentialDefinitionRequest, *credential.QueryCredentialDefinitionResponse) {
				return &credential.QueryCredentialDefinitionRequest{Id: "did:cosmos:elesto:cd-not-found"}, nil
			},
			errors.New("rpc error: code = NotFound desc = credential definition not found"),
		},
		{
			"FAIL: will fail because no id is provided",
			func() (*credential.QueryCredentialDefinitionRequest, *credential.QueryCredentialDefinitionResponse) {
				return &credential.QueryCredentialDefinitionRequest{}, nil
			},
			errors.New("rpc error: code = InvalidArgument desc = credential definition id must not be empty"),
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			req, expectedResp := tc.reqFn()
			gotResp, err := queryClient.CredentialDefinition(context.Background(), req)
			if tc.wantErr == nil {
				suite.Require().NoError(err)
				suite.Require().Equal(expectedResp, gotResp)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(tc.wantErr.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_CredentialDefinitionsByPublisher() {
	queryClient := suite.queryClient
	server := keeper.NewMsgServerImpl(suite.keeper)

	testCases := []struct {
		msg     string
		reqFn   func() (*credential.QueryCredentialDefinitionsByPublisherRequest, *credential.QueryCredentialDefinitionsByPublisherResponse)
		wantErr error
	}{
		{
			"PASS: can get the credential definition",
			func() (*credential.QueryCredentialDefinitionsByPublisherRequest, *credential.QueryCredentialDefinitionsByPublisherResponse) {
				// publish the definition
				pcdr := credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: &credential.CredentialDefinition{
						Id:           "did:cosmos:elesto:cdsbi-query-001",
						PublisherId:  did.NewKeyDID(suite.GetTestAccount().String()).String(),
						Schema:       []byte(dummySchemaOk),
						Vocab:        []byte(dummyVocabOk),
						Name:         "CredentialDef001",
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
				// publish another one with a different publisher
				if _, err := server.PublishCredentialDefinition(sdk.WrapSDKContext(suite.ctx), &credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: &credential.CredentialDefinition{
						Id:           "did:cosmos:elesto:cdsbi-query-002",
						PublisherId:  did.NewKeyDID(suite.GetTestAccountByIndex(1).String()).String(),
						Schema:       []byte(dummySchemaOk),
						Vocab:        []byte(dummyVocabOk),
						Name:         "CredentialDef002",
						Description:  "",
						SupersededBy: "",
						IsActive:     true,
					},
					Signer: suite.GetTestAccount().String(),
				}); err != nil {
					suite.Require().FailNowf("expected definition to be created, got:", "%v", err)
				}

				return &credential.QueryCredentialDefinitionsByPublisherRequest{Did: did.NewKeyDID(suite.GetTestAccount().String()).String()},
					&credential.QueryCredentialDefinitionsByPublisherResponse{
						Definitions: []*credential.CredentialDefinition{
							pcdr.CredentialDefinition,
						},
						Pagination: &query.PageResponse{Total: 2},
					}
			},
			nil,
		},
		{
			"FAIL: credential definition not found",
			func() (*credential.QueryCredentialDefinitionsByPublisherRequest, *credential.QueryCredentialDefinitionsByPublisherResponse) {
				return &credential.QueryCredentialDefinitionsByPublisherRequest{Did: "did:cosmos:elesto:cd-not-found"}, &credential.QueryCredentialDefinitionsByPublisherResponse{Definitions: nil, Pagination: &query.PageResponse{Total: 2}}
			},
			nil,
		},
		{
			"FAIL: will fail because no id is provided",
			func() (*credential.QueryCredentialDefinitionsByPublisherRequest, *credential.QueryCredentialDefinitionsByPublisherResponse) {
				return &credential.QueryCredentialDefinitionsByPublisherRequest{}, nil
			},
			errors.New("rpc error: code = InvalidArgument desc = publisher DID must be a valid DID"),
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			req, expectedResp := tc.reqFn()
			gotResp, err := queryClient.CredentialDefinitionsByPublisher(context.Background(), req)
			if tc.wantErr == nil {
				suite.Require().NoError(err)
				suite.Require().Equal(expectedResp, gotResp)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(tc.wantErr.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_CredentialDefinitions() {
	queryClient := suite.queryClient
	server := keeper.NewMsgServerImpl(suite.keeper)

	testCases := []struct {
		msg     string
		reqFn   func() (*credential.QueryCredentialDefinitionsRequest, *credential.QueryCredentialDefinitionsResponse)
		wantErr error
	}{
		{
			"PASS: no credential definitions",
			func() (*credential.QueryCredentialDefinitionsRequest, *credential.QueryCredentialDefinitionsResponse) {
				return &credential.QueryCredentialDefinitionsRequest{}, &credential.QueryCredentialDefinitionsResponse{Definitions: nil, Pagination: &query.PageResponse{}}
			},
			nil,
		},
		{
			"PASS: can get the credential definition",
			func() (*credential.QueryCredentialDefinitionsRequest, *credential.QueryCredentialDefinitionsResponse) {
				// publish the definition
				pcdr := credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: &credential.CredentialDefinition{
						Id:           "did:cosmos:elesto:cds-query-001",
						PublisherId:  did.NewKeyDID(suite.GetTestAccount().String()).String(),
						Schema:       []byte(dummySchemaOk),
						Vocab:        []byte(dummyVocabOk),
						Name:         "CredentialDef001",
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
				return &credential.QueryCredentialDefinitionsRequest{},
					&credential.QueryCredentialDefinitionsResponse{
						Definitions: []*credential.CredentialDefinition{
							pcdr.CredentialDefinition,
						},
						Pagination: &query.PageResponse{Total: 1},
					}
			},
			nil,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			req, expectedResp := tc.reqFn()
			gotResp, err := queryClient.CredentialDefinitions(context.Background(), req)
			if tc.wantErr == nil {
				suite.Require().NoError(err)
				suite.Require().Equal(expectedResp, gotResp)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(tc.wantErr.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_PublicCredential() {
	queryClient := suite.queryClient
	server := keeper.NewMsgServerImpl(suite.keeper)

	testCases := []struct {
		msg     string
		reqFn   func() (*credential.QueryPublicCredentialRequest, *credential.QueryPublicCredentialResponse)
		wantErr error
	}{
		{
			"PASS: can get the credential",
			func() (*credential.QueryPublicCredentialRequest, *credential.QueryPublicCredentialResponse) {
				var (
					id  = "001"
					wc  *credential.WrappedCredential
					err error
				)
				//

				// publish the definition
				pcdr := credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: &credential.CredentialDefinition{
						Id:           "did:cosmos:elesto:ipcq-" + id,
						PublisherId:  did.NewKeyDID(suite.GetTestAccount().String()).String(),
						Schema:       []byte(dummySchemaOk),
						Vocab:        []byte(dummyVocabOk),
						Name:         "CredentialDef" + id,
						Description:  "",
						SupersededBy: "",
						IsActive:     true,
					},
					Signer: suite.GetTestAccount().String(),
				}
				suite.keeper.AllowPublicCredential(suite.ctx, pcdr.CredentialDefinition.Id)
				//create the credential definition
				_, err = server.PublishCredentialDefinition(sdk.WrapSDKContext(suite.ctx), &pcdr)
				suite.Require().NoError(err)
				// load the signed credential
				if wc, err = credential.NewWrappedPublicCredentialFromFile("testdata/dummy.credential.signed.json"); err != nil {
					suite.Require().FailNowf("expected wrapped credential, got:", "%v", err)
				}
				// publish the credential
				_, err = server.IssuePublicVerifiableCredential(sdk.WrapSDKContext(suite.ctx), &credential.MsgIssuePublicVerifiableCredentialRequest{
					CredentialDefinitionID: "did:cosmos:elesto:ipcq-" + id,
					Credential:             wc.PublicVerifiableCredential,
					Signer:                 suite.GetTestAccount().String(),
				})
				suite.Require().NoError(err)

				return &credential.QueryPublicCredentialRequest{Id: wc.Id}, &credential.QueryPublicCredentialResponse{Credential: wc.PublicVerifiableCredential}
			},
			nil,
		},
		{
			"FAIL: credential not found",
			func() (*credential.QueryPublicCredentialRequest, *credential.QueryPublicCredentialResponse) {
				return &credential.QueryPublicCredentialRequest{Id: "https://does.not.exists"}, &credential.QueryPublicCredentialResponse{Credential: nil}
			},
			errors.New("rpc error: code = NotFound desc = public verifiable credential not found"),
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			req, expectedResp := tc.reqFn()
			gotResp, err := queryClient.PublicCredential(context.Background(), req)
			if tc.wantErr == nil {
				suite.Require().NoError(err)
				suite.Require().Equal(expectedResp, gotResp)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(tc.wantErr.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_PublicCredentials() {
	queryClient := suite.queryClient
	server := keeper.NewMsgServerImpl(suite.keeper)
	_ = server

	testCases := []struct {
		msg     string
		reqFn   func() (*credential.QueryPublicCredentialsRequest, *credential.QueryPublicCredentialsResponse)
		wantErr error
	}{
		{
			"PASS: no credentials",
			func() (*credential.QueryPublicCredentialsRequest, *credential.QueryPublicCredentialsResponse) {
				return &credential.QueryPublicCredentialsRequest{}, &credential.QueryPublicCredentialsResponse{Credential: nil, Pagination: &query.PageResponse{Total: 0}}
			},
			nil,
		},
		{
			"PASS: can get the credential",
			func() (*credential.QueryPublicCredentialsRequest, *credential.QueryPublicCredentialsResponse) {
				var (
					id  = "001"
					wc  *credential.WrappedCredential
					err error
				)
				//

				// publish the definition
				pcdr := credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: &credential.CredentialDefinition{
						Id:           "did:cosmos:elesto:ipcq-" + id,
						PublisherId:  did.NewKeyDID(suite.GetTestAccount().String()).String(),
						Schema:       []byte(dummySchemaOk),
						Vocab:        []byte(dummyVocabOk),
						Name:         "CredentialDef" + id,
						Description:  "",
						SupersededBy: "",
						IsActive:     true,
					},
					Signer: suite.GetTestAccount().String(),
				}
				suite.keeper.AllowPublicCredential(suite.ctx, pcdr.CredentialDefinition.Id)
				//create the credential definition
				_, err = server.PublishCredentialDefinition(sdk.WrapSDKContext(suite.ctx), &pcdr)
				suite.Require().NoError(err)
				// load the signed credential
				if wc, err = credential.NewWrappedPublicCredentialFromFile("testdata/dummy.credential.signed.json"); err != nil {
					suite.Require().FailNowf("expected wrapped credential, got:", "%v", err)
				}
				// publish the credential
				_, err = server.IssuePublicVerifiableCredential(sdk.WrapSDKContext(suite.ctx), &credential.MsgIssuePublicVerifiableCredentialRequest{
					CredentialDefinitionID: "did:cosmos:elesto:ipcq-" + id,
					Credential:             wc.PublicVerifiableCredential,
					Signer:                 suite.GetTestAccount().String(),
				})
				suite.Require().NoError(err)

				return &credential.QueryPublicCredentialsRequest{}, &credential.QueryPublicCredentialsResponse{Credential: []*credential.PublicVerifiableCredential{wc.PublicVerifiableCredential}, Pagination: &query.PageResponse{Total: 1}}
			},
			nil,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			req, expectedResp := tc.reqFn()
			gotResp, err := queryClient.PublicCredentials(context.Background(), req)
			if tc.wantErr == nil {
				suite.Require().NoError(err)
				suite.Require().Equal(expectedResp, gotResp)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(tc.wantErr.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_PublicCredentialsByHolder() {
	queryClient := suite.queryClient
	server := keeper.NewMsgServerImpl(suite.keeper)
	_ = server

	testCases := []struct {
		msg     string
		reqFn   func() (*credential.QueryPublicCredentialsByHolderRequest, *credential.QueryPublicCredentialsByHolderResponse)
		wantErr error
	}{
		{
			"PASS: no credentials",
			func() (*credential.QueryPublicCredentialsByHolderRequest, *credential.QueryPublicCredentialsByHolderResponse) {
				return &credential.QueryPublicCredentialsByHolderRequest{Did: did.NewKeyDID(suite.GetTestAccount().String()).String()}, &credential.QueryPublicCredentialsByHolderResponse{Credential: nil, Pagination: &query.PageResponse{}}
			},
			nil,
		},
		{
			"PASS: can get the credential",
			func() (*credential.QueryPublicCredentialsByHolderRequest, *credential.QueryPublicCredentialsByHolderResponse) {
				var (
					id  = "001"
					wc  *credential.WrappedCredential
					err error
				)

				// publish the definition
				pcdr := credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: &credential.CredentialDefinition{
						Id:           "did:cosmos:elesto:ipcq-" + id,
						PublisherId:  did.NewKeyDID(suite.GetTestAccount().String()).String(),
						Schema:       []byte(dummySchemaOk),
						Vocab:        []byte(dummyVocabOk),
						Name:         "CredentialDef" + id,
						Description:  "",
						SupersededBy: "",
						IsActive:     true,
					},
					Signer: suite.GetTestAccount().String(),
				}
				//create the credential definition
				_, err = server.PublishCredentialDefinition(sdk.WrapSDKContext(suite.ctx), &pcdr)
				suite.Require().NoError(err)

				// allowing the credential definition id for publishing
				suite.keeper.AllowPublicCredential(suite.ctx, pcdr.CredentialDefinition.Id)

				// load the signed credential
				if wc, err = credential.NewWrappedPublicCredentialFromFile("testdata/dummy.credential.signed.json"); err != nil {
					suite.Require().FailNowf("expected wrapped credential, got:", "%v", err)
				}
				// publish the credential
				_, err = server.IssuePublicVerifiableCredential(sdk.WrapSDKContext(suite.ctx), &credential.MsgIssuePublicVerifiableCredentialRequest{
					CredentialDefinitionID: "did:cosmos:elesto:ipcq-" + id,
					Credential:             wc.PublicVerifiableCredential,
					Signer:                 suite.GetTestAccount().String(),
				})
				suite.Require().NoError(err)

				return &credential.QueryPublicCredentialsByHolderRequest{Did: did.NewKeyDID(suite.GetTestAccount().String()).String()}, &credential.QueryPublicCredentialsByHolderResponse{Credential: []*credential.PublicVerifiableCredential{wc.PublicVerifiableCredential}, Pagination: &query.PageResponse{Total: 1}}
			},
			nil,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			req, expectedResp := tc.reqFn()
			gotResp, err := queryClient.PublicCredentialsByHolder(context.Background(), req)
			if tc.wantErr == nil {
				suite.Require().NoError(err)
				suite.Require().Equal(expectedResp, gotResp)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(tc.wantErr.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_PublicCredentialsByIssuer() {
	queryClient := suite.queryClient
	server := keeper.NewMsgServerImpl(suite.keeper)
	_ = server

	testCases := []struct {
		msg     string
		reqFn   func() (*credential.QueryPublicCredentialsByIssuerRequest, *credential.QueryPublicCredentialsByIssuerResponse)
		wantErr error
	}{
		{
			"PASS: no credentials",
			func() (*credential.QueryPublicCredentialsByIssuerRequest, *credential.QueryPublicCredentialsByIssuerResponse) {
				return &credential.QueryPublicCredentialsByIssuerRequest{Did: did.NewKeyDID(suite.GetTestAccount().String()).String()}, &credential.QueryPublicCredentialsByIssuerResponse{Credential: nil, Pagination: &query.PageResponse{}}
			},
			nil,
		},
		{
			"PASS: can get the credential",
			func() (*credential.QueryPublicCredentialsByIssuerRequest, *credential.QueryPublicCredentialsByIssuerResponse) {
				var (
					id  = "001"
					wc  *credential.WrappedCredential
					err error
				)

				// publish the definition
				pcdr := credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: &credential.CredentialDefinition{
						Id:           "did:cosmos:elesto:ipcq-" + id,
						PublisherId:  did.NewKeyDID(suite.GetTestAccount().String()).String(),
						Schema:       []byte(dummySchemaOk),
						Vocab:        []byte(dummyVocabOk),
						Name:         "CredentialDef" + id,
						Description:  "",
						SupersededBy: "",
						IsActive:     true,
					},
					Signer: suite.GetTestAccount().String(),
				}
				//create the credential definition
				_, err = server.PublishCredentialDefinition(sdk.WrapSDKContext(suite.ctx), &pcdr)
				suite.Require().NoError(err)

				// allowing the credential definition id for publishing
				suite.keeper.AllowPublicCredential(suite.ctx, pcdr.CredentialDefinition.Id)

				// load the signed credential
				if wc, err = credential.NewWrappedPublicCredentialFromFile("testdata/dummy.credential.signed.json"); err != nil {
					suite.Require().FailNowf("expected wrapped credential, got:", "%v", err)
				}
				// publish the credential
				_, err = server.IssuePublicVerifiableCredential(sdk.WrapSDKContext(suite.ctx), &credential.MsgIssuePublicVerifiableCredentialRequest{
					CredentialDefinitionID: "did:cosmos:elesto:ipcq-" + id,
					Credential:             wc.PublicVerifiableCredential,
					Signer:                 suite.GetTestAccount().String(),
				})
				suite.Require().NoError(err)

				return &credential.QueryPublicCredentialsByIssuerRequest{Did: did.NewKeyDID(suite.GetTestAccount().String()).String()}, &credential.QueryPublicCredentialsByIssuerResponse{Credential: []*credential.PublicVerifiableCredential{wc.PublicVerifiableCredential}, Pagination: &query.PageResponse{Total: 1}}
			},
			nil,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			req, expectedResp := tc.reqFn()
			gotResp, err := queryClient.PublicCredentialsByIssuer(context.Background(), req)
			if tc.wantErr == nil {
				suite.Require().NoError(err)
				suite.Require().Equal(expectedResp, gotResp)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(tc.wantErr.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_AllowedPublicCredentials() {
	queryClient := suite.queryClient
	server := keeper.NewMsgServerImpl(suite.keeper)

	testCases := []struct {
		msg     string
		reqFn   func() (*credential.QueryAllowedPublicCredentialsRequest, *credential.QueryAllowedPublicCredentialsResponse)
		wantErr error
	}{
		{
			"PASS: no credentials",
			func() (*credential.QueryAllowedPublicCredentialsRequest, *credential.QueryAllowedPublicCredentialsResponse) {
				return &credential.QueryAllowedPublicCredentialsRequest{}, &credential.QueryAllowedPublicCredentialsResponse{Credentials: nil, Pagination: &query.PageResponse{}}
			},
			nil,
		},
		{
			"PASS: no credentials in allow list",
			func() (*credential.QueryAllowedPublicCredentialsRequest, *credential.QueryAllowedPublicCredentialsResponse) {
				var (
					id  = "001"
					err error
				)
				//

				// publish the definition
				pcdr := credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: &credential.CredentialDefinition{
						Id:           "did:cosmos:elesto:ipcq-" + id,
						PublisherId:  did.NewKeyDID(suite.GetTestAccount().String()).String(),
						Schema:       []byte(dummySchemaOk),
						Vocab:        []byte(dummyVocabOk),
						Name:         "CredentialDef" + id,
						Description:  "",
						SupersededBy: "",
						IsActive:     true,
					},
					Signer: suite.GetTestAccount().String(),
				}

				//create the credential definition
				_, err = server.PublishCredentialDefinition(sdk.WrapSDKContext(suite.ctx), &pcdr)
				suite.Require().NoError(err)

				return &credential.QueryAllowedPublicCredentialsRequest{}, &credential.QueryAllowedPublicCredentialsResponse{Pagination: &query.PageResponse{Total: 0}}
			},
			nil,
		},
		{
			"PASS: can get the credential",
			func() (*credential.QueryAllowedPublicCredentialsRequest, *credential.QueryAllowedPublicCredentialsResponse) {
				var (
					id  = "002"
					err error
				)
				//

				// publish the definition
				pcdr := credential.MsgPublishCredentialDefinitionRequest{
					CredentialDefinition: &credential.CredentialDefinition{
						Id:           "did:cosmos:elesto:ipcq-" + id,
						PublisherId:  did.NewKeyDID(suite.GetTestAccount().String()).String(),
						Schema:       []byte(dummySchemaOk),
						Vocab:        []byte(dummyVocabOk),
						Name:         "CredentialDef" + id,
						Description:  "",
						SupersededBy: "",
						IsActive:     true,
					},
					Signer: suite.GetTestAccount().String(),
				}
				suite.keeper.AllowPublicCredential(suite.ctx, pcdr.CredentialDefinition.Id)
				//create the credential definition
				_, err = server.PublishCredentialDefinition(sdk.WrapSDKContext(suite.ctx), &pcdr)
				suite.Require().NoError(err)

				return &credential.QueryAllowedPublicCredentialsRequest{}, &credential.QueryAllowedPublicCredentialsResponse{Credentials: []*credential.CredentialDefinition{pcdr.CredentialDefinition}, Pagination: &query.PageResponse{Total: 1}}
			},
			nil,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			req, expectedResp := tc.reqFn()
			gotResp, err := queryClient.AllowedPublicCredentials(context.Background(), req)
			if tc.wantErr == nil {
				suite.Require().NoError(err)
				suite.Require().Equal(expectedResp, gotResp)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(tc.wantErr.Error(), err.Error())
			}
		})
	}
}
