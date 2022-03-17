package keeper

import (
	"fmt"
	"time"

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
		didFn   func() (did.DidMetadata, did.DidDocument)
		expPass bool
	}{
		{
			"PASS: did document metadata stored successfully",
			func() (did.DidMetadata, did.DidDocument) {
				dd, _ := did.NewDidDocument("did:cosmos:net:elesto:subject")
				didMeta := did.NewDidMetadata([]byte("xxx"), time.Now())
				suite.keeper.SetDidDocument(suite.ctx, []byte(dd.Id), dd)
				suite.keeper.SetDidMetadata(suite.ctx, []byte(dd.Id), didMeta)
				return didMeta, dd
			},
			true,
		},
		{
			"FAIL: did document does not exist",
			func() (did.DidMetadata, did.DidDocument) {
				dd, _ := did.NewDidDocument("did:cosmos:net:elesto:fail")
				didMeta := did.NewDidMetadata([]byte("xxx"), time.Now().Local())
				suite.keeper.SetDidMetadata(suite.ctx, []byte(dd.Id), didMeta)
				return didMeta, dd
			},
			false,
		},
	}
	for _, tc := range testCases {
		didMeta, dd := tc.didFn()

		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			if tc.expPass {
				didMetaGet, found := suite.keeper.GetDidMetadata(
					suite.ctx,
					[]byte(dd.Id),
				)
				suite.Require().True(found)
				suite.Require().Equal(didMeta.VersionId, didMetaGet.VersionId)
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
