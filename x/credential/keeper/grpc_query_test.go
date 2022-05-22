package keeper

import (
	"context"
	"fmt"
	"github.com/elesto-dao/elesto/x/credential"
)

func (suite *KeeperTestSuite) TestGRPCQueryCredentialDefinition() {
	queryClient := suite.queryClient

	testCases := []struct {
		msg     string
		reqFn   func() *credential.QueryCredentialDefinitionRequest
		expPass bool
	}{
		{
			"FAIL: will fail because no id is provided",
			func() *credential.QueryCredentialDefinitionRequest {
				return &credential.QueryCredentialDefinitionRequest{}
			},
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			req := tc.reqFn()
			didsResp, err := queryClient.CredentialDefinition(context.Background(), req)
			if tc.expPass {
				suite.NoError(err)
				suite.NotNil(didsResp)

			} else {
				suite.Require().Error(err)
			}
		})
	}
}
