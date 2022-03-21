package keeper

import (
	"fmt"

	"github.com/elesto-dao/elesto/x/did"
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
				dd, _ := did.NewDidDocument("did:cosmos:net:elesto:subject")
				suite.keeper.SetDidDocument(suite.ctx, []byte(dd.Id), dd)
				return dd
			},
			true,
		},
		{
			"FAIL: did document does not exist",
			func() did.DidDocument {
				dd, _ := did.NewDidDocument("did:cosmos:net:elesto:not")
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

func (suite *KeeperTestSuite) TestDidDocumentMetadataKeeper() {
	testCases := []struct {
		msg     string
		didFn   func() did.DidDocument
		expPass bool
	}{
		{
			"PASS: did document metadata stored successfully",
			func() did.DidDocument {
				dd, _ := did.NewDidDocument("did:cosmos:net:elesto:subject")
				suite.keeper.SetDidDocument(suite.ctx, []byte(dd.Id), dd)
				return dd
			},
			true,
		},
		{
			"FAIL: did document does not exist",
			func() did.DidDocument {
				dd, _ := did.NewDidDocument("did:cosmos:net:elesto:fail")
				suite.keeper.SetDidMetadata(suite.ctx, []byte(dd.Id))
				return dd
			},
			false,
		},
	}
	for _, tc := range testCases {
		dd := tc.didFn()

		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			if tc.expPass {
				didMetaGet, _ := suite.keeper.GetDidMetadata(
					suite.ctx,
					[]byte(dd.Id),
				)

				suite.Require().Equal(
					uint64(0x0),
					didMetaGet.VersionId,
				)
			} else {
				_, _, err := suite.keeper.ResolveDid(
					suite.ctx,
					did.NewChainDID("elesto", "fail"),
				)
				suite.Require().Error(err)
			}
		})
	}
}
