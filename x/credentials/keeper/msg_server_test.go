package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elesto-dao/elesto/x/credentials"
)

func (suite *KeeperTestSuite) TestHandleMsgPublishCredentialDefinition() {
	var (
		req    credentials.MsgPublishCredentialDefinitionRequest
		errExp error
	)

	server := NewMsgServerImpl(suite.keeper, nil)

	testCases := []struct {
		name     string
		malleate func()
	}{}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			tc.malleate()
			_, err := server.PublishCredentialDefinition(sdk.WrapSDKContext(suite.ctx), &req)
			if errExp == nil {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(errExp.Error(), err.Error())
			}
		})
	}
}
