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
	signerAccount = "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"
)

var (
	//go:embed testdata/schema.json
	schemaOk string
	//go:embed testdata/vocab.json
	vocabOk string
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
						PublisherId:  did.NewKeyDID(signerAccount).String(),
						Schema:       schemaOk,
						Vocab:        vocabOk,
						Name:         "CredentialDef1",
						Description:  "",
						IsPublic:     true,
						SupersededBy: "",
						IsActive:     true,
					},
					Signer: signerAccount,
				}
			},
			nil,
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
