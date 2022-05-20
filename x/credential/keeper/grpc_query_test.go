package keeper

import (
	"context"
	"fmt"

	"github.com/elesto-dao/elesto/x/did"
)

func (suite *KeeperTestSuite) TestGRPCQueryDidDocument() {
	queryClient := suite.queryClient
	var req *did.QueryDidDocumentRequest
	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"Fail: will fail because no id is provided",
			func() {
				req = &did.QueryDidDocumentRequest{}
			},
			false,
		},
		{
			"Fail: will fail because no did is found",
			func() {
				req = &did.QueryDidDocumentRequest{
					Id: "did:cosmos:cash:1234",
				}
			},
			false,
		},
		{
			"Pass: will pass because a address did is autoresolved",
			func() {
				req = &did.QueryDidDocumentRequest{
					Id: "did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				}
			},
			true,
		},
		{
			"Fail: will fail because the only cosmos based address are supported",
			func() {
				req = &did.QueryDidDocumentRequest{
					Id: "did:cosmos:key:0xB88F61E6FbdA83fbfffAbE364112137480398018",
				}
			},
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()
			didsResp, err := queryClient.DidDocument(context.Background(), req)
			if tc.expPass {
				suite.NoError(err)
				suite.NotNil(didsResp)

			} else {
				suite.Require().Error(err)
			}
		})
	}
}
