package keeper

import (
	"fmt"

	"github.com/elesto-dao/elesto/v2/x/did"
)

func (suite *KeeperTestSuite) TestDidDocumentKeeper() {
	testCases := []struct {
		msg     string
		didFn   func() did.DidDocument
		expPass bool
	}{
		{
			"PASS: did document stored successfully",
			func() did.DidDocument {
				dd, _ := did.NewDidDocument("did:cosmos:elesto:subject")
				suite.keeper.SetDidDocument(suite.ctx, []byte(dd.Id), dd)
				return dd
			},
			true,
		},
		{
			"FAIL: did document does not exist",
			func() did.DidDocument {
				dd, _ := did.NewDidDocument("did:cosmos:elesto:not")
				return dd
			},
			false,
		},
	}
	for _, tc := range testCases {
		dd := tc.didFn()

		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			if tc.expPass {
				found := suite.keeper.HasDidDocument(
					suite.ctx,
					[]byte(dd.Id),
				)
				suite.Require().True(found)

				ddGet, found := suite.keeper.GetDidDocument(
					suite.ctx,
					[]byte(dd.Id),
				)
				suite.Require().Equal(dd, ddGet)

				allEntities := suite.keeper.GetAllDidDocuments(
					suite.ctx,
				)
				suite.Require().Equal(1, len(allEntities))
			} else {
				found := suite.keeper.HasDidDocument(
					suite.ctx,
					[]byte(dd.Id),
				)
				suite.Require().False(found)
			}
		})
	}
}
