package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	didmod "github.com/elesto-dao/elesto/v4/x/did"
)

func (suite *KeeperTestSuite) TestHandleMsgCreateDidDocument() {
	var (
		req    didmod.MsgCreateDidDocument
		errExp error
	)

	server := NewMsgServerImpl(suite.keeper)

	testCases := []struct {
		name     string
		malleate func()
	}{
		{
			"Pass: can create a an did",
			func() {
				req = *didmod.NewMsgCreateDidDocument("did:cosmos:elesto:00000000-0000-0000-0000-000000000000", nil, nil, "subject")
				errExp = nil
			},
		},
		{
			"FAIL: did doc validation fails",
			func() {
				req = *didmod.NewMsgCreateDidDocument("invalid did", nil, nil, "subject")
				errExp = sdkerrors.Wrapf(didmod.ErrInvalidDIDFormat, "did %s", "invalid did")
			},
		},
		{
			"FAIL: did already exists",
			func() {
				did := "did:cosmos:elesto:00000000-0000-0000-0000-000000000000"
				didDoc, _ := didmod.NewDidDocument(did)

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *didmod.NewMsgCreateDidDocument(did, nil, nil, "subject")
				errExp = sdkerrors.Wrapf(didmod.ErrDidDocumentFound, "a document with did %s already exists", did)
			},
		},
		{
			"FAIL: did is of type key (1)",
			func() {
				did := "did:cosmos:key:00000000-0000-0000-0000-000000000000"
				req = *didmod.NewMsgCreateDidDocument(did, nil, nil, "subject")
				errExp = sdkerrors.Wrapf(didmod.ErrInvalidInput, "did documents having id with key format cannot be created %s", did)
			},
		},
		{
			"FAIL: did is of type key (2)",
			func() {
				did := "did:cosmos:key:76f3a6c4-e048-4009-bb01-e0668a91ad2f"
				req = *didmod.NewMsgCreateDidDocument(did, nil, nil, "subject")
				errExp = sdkerrors.Wrapf(didmod.ErrInvalidInput, "did documents having id with key format cannot be created %s", did)
			},
		},
		{
			"FAIL: did is not UUID",
			func() {
				did := "did:cosmos:elesto:juno1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"
				req = *didmod.NewMsgCreateDidDocument(did, nil, nil, "subject")
				errExp = sdkerrors.Wrapf(didmod.ErrInvalidInput, "did:cosmos:elesto:juno1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8 is not a valid UUID format")
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			tc.malleate()
			_, err := server.CreateDidDocument(sdk.WrapSDKContext(suite.ctx), &req)
			if errExp == nil {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(errExp.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestHandleMsgUpdateDidDocument() {
	var (
		req    didmod.MsgUpdateDidDocument
		errExp error
	)

	server := NewMsgServerImpl(suite.keeper)

	testCases := []struct {
		name     string
		malleate func()
	}{
		{
			"FAIL: not found",
			func() {
				req = *didmod.NewMsgUpdateDidDocument(&didmod.DidDocument{Id: "did:cosmos:elesto:subject"}, "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
				errExp = sdkerrors.Wrapf(didmod.ErrDidDocumentNotFound, "did document at %s not found", "did:cosmos:elesto:subject")
			},
		},
		{
			"FAIL: unauthorized",
			func() {

				did := "did:cosmos:elesto:subject"
				didDoc, _ := didmod.NewDidDocument(did)
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)

				req = *didmod.NewMsgUpdateDidDocument(&didmod.DidDocument{Id: didDoc.Id, Controller: []string{"did:cosmos:cash:controller"}}, "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
				errExp = sdkerrors.Wrapf(didmod.ErrUnauthorized, "signer account %s not authorized to update the target did document at %s", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8", did)

			},
		},
		{
			"PASS: replace did document",
			func() {

				did := "did:cosmos:elesto:subject"
				didDoc, _ := didmod.NewDidDocument(did, didmod.WithVerifications(
					didmod.NewVerification(
						didmod.NewVerificationMethod(
							"did:cosmos:elesto:subject#key-1",
							"did:cosmos:elesto:subject",
							didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
							didmod.EcdsaSecp256k1VerificationKey2019,
						),
						[]string{didmod.Authentication},
						nil,
					),
				))
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)

				newDidDoc, err := didmod.NewDidDocument(did)
				suite.Require().Nil(err)

				req = *didmod.NewMsgUpdateDidDocument(&newDidDoc, "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
				errExp = nil
			},
		},
		{
			"FAIL: invalid controllers",
			func() {
				didDoc, _ := didmod.NewDidDocument("did:cosmos:elesto:subject", didmod.WithVerifications(
					didmod.NewVerification(
						didmod.NewVerificationMethod(
							"did:cosmos:elesto:subject#key-1",
							"did:cosmos:elesto:subject",
							didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
							didmod.EcdsaSecp256k1VerificationKey2019,
						),
						[]string{didmod.Authentication},
						nil,
					),
				))
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)

				controllers := []string{
					"did:cosmos:cash:controller-1",
					"did:cosmos:cash:controller-2",
					"invalid",
				}

				req = *didmod.NewMsgUpdateDidDocument(&didmod.DidDocument{Id: didDoc.Id, Controller: controllers}, "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
				errExp = sdkerrors.Wrapf(didmod.ErrInvalidDIDFormat, "invalid did document")
			},
		},
		{
			"FAIL: did is of type key (1)",
			func() {
				did := "did:cosmos:key:00000000-0000-0000-0000-000000000000"
				didDoc, _ := didmod.NewDidDocument(did)
				controllers := []string{
					"did:cosmos:cash:controller-1",
					"did:cosmos:cash:controller-2",
				}

				req = *didmod.NewMsgUpdateDidDocument(&didmod.DidDocument{Id: didDoc.Id, Controller: controllers}, "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
				errExp = sdkerrors.Wrapf(didmod.ErrInvalidInput, "did documents having id with key format are read only %s", did)
			},
		},
		{
			"FAIL: did is of type key (2)",
			func() {
				did := "did:cosmos:key:00000000-0000-0000-0000-000000000000"
				didDoc, _ := didmod.NewDidDocument(did)
				controllers := []string{
					"did:cosmos:cash:controller-1",
					"did:cosmos:cash:controller-2",
				}

				req = *didmod.NewMsgUpdateDidDocument(&didmod.DidDocument{Id: didDoc.Id, Controller: controllers}, "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
				errExp = sdkerrors.Wrapf(didmod.ErrInvalidInput, "did documents having id with key format are read only %s", did)
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			tc.malleate()

			_, err := server.UpdateDidDocument(sdk.WrapSDKContext(suite.ctx), &req)

			if errExp == nil {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(errExp.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestHandleMsgAddVerification() {
	var (
		req    didmod.MsgAddVerification
		errExp error
	)

	server := NewMsgServerImpl(suite.keeper)

	testCases := []struct {
		name     string
		malleate func()
	}{
		{
			"FAIL: can not add verification, did does not exist",
			func() {
				req = *didmod.NewMsgAddVerification("did:cosmos:elesto:subject", nil, "subject")
				errExp = sdkerrors.Wrapf(didmod.ErrDidDocumentNotFound, "did document at %s not found", "did:cosmos:elesto:subject")
			},
		},
		{
			"FAIL: can not add verification, unauthorized",
			func() {
				// setup
				didDoc, _ := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.CapabilityInvocation},
							nil,
						),
					),
				)
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				// actual test
				v := didmod.NewVerification(
					didmod.NewVerificationMethod(
						"did:cosmos:elesto:subject#key-2",
						"did:cosmos:elesto:subject",
						didmod.NewBlockchainAccountID("foochainid", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"),
						didmod.CosmosAccountAddress,
					),
					[]string{didmod.Authentication},
					nil,
				)
				req = *didmod.NewMsgAddVerification(didDoc.Id, v, "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
				errExp = sdkerrors.Wrapf(didmod.ErrUnauthorized, "signer account %s not authorized to update the target did document at %s", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8", didDoc.Id)
			},
		},
		{
			"FAIL: can not add verification, unauthorized, key mismatch",
			func() {
				// setup
				didDoc, _ := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.Authentication},
							nil,
						),
					),
				)
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				// actual test
				v := didmod.NewVerification(
					didmod.NewVerificationMethod(
						"did:cosmos:elesto:subject#key-2",
						"did:cosmos:elesto:subject",
						didmod.NewBlockchainAccountID("foochainid", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"),
						didmod.CosmosAccountAddress,
					),
					[]string{didmod.Authentication},
					nil,
				)
				req = *didmod.NewMsgAddVerification(didDoc.Id, v, "cash1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2")
				errExp = sdkerrors.Wrapf(didmod.ErrUnauthorized, "signer account %s not authorized to update the target did document at %s", "cash1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2", didDoc.Id)
			},
		},
		{
			"FAIL: can not add verification, invalid verification",
			func() {
				// setup
				//signer := "subject"
				signer := "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"
				didDoc, _ := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.Authentication},
							nil,
						),
					),
				)
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				// actual test
				v := didmod.NewVerification(
					didmod.NewVerificationMethod(
						"",
						"did:cosmos:elesto:subject",
						didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
						didmod.EcdsaSecp256k1VerificationKey2019,
					),
					[]string{didmod.Authentication},
					nil,
				)
				req = *didmod.NewMsgAddVerification(didDoc.Id, v, signer)
				errExp = sdkerrors.Wrapf(didmod.ErrInvalidDIDURLFormat, "verification method id: %v", "")
			},
		},
		{
			"PASS: can add verification to did document",
			func() {
				didDoc, _ := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.Authentication},
							nil,
						),
					),
				)

				v := didmod.NewVerification(
					didmod.NewVerificationMethod(
						"did:cosmos:elesto:subject#key-2",
						"did:cosmos:elesto:subject",
						didmod.NewBlockchainAccountID("foochainid", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"),
						didmod.CosmosAccountAddress,
					),
					[]string{didmod.Authentication},
					nil,
				)

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *didmod.NewMsgAddVerification(didDoc.Id, v, "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
				errExp = nil
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			tc.malleate()

			_, err := server.AddVerification(sdk.WrapSDKContext(suite.ctx), &req)

			if errExp == nil {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(errExp.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestHandleMsgSetVerificationRelationships() {
	var (
		req    didmod.MsgSetVerificationRelationships
		errExp error
	)

	server := NewMsgServerImpl(suite.keeper)

	testCases := []struct {
		name     string
		malleate func()
	}{
		{
			"FAIL: can not add verification relationship, did does not exist",
			func() {
				req = *didmod.NewMsgSetVerificationRelationships(
					"did:cosmos:elesto:subject",
					"did:cosmos:elesto:subject#key-1",
					[]string{didmod.Authentication},
					"subject",
				)
				errExp = sdkerrors.Wrapf(didmod.ErrDidDocumentNotFound, "did document at %s not found", "did:cosmos:elesto:subject")
			},
		},
		{
			"FAIL: can not add verification relationship, unauthorized",
			func() {
				// setup
				didDoc, _ := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
								"did:cosmos:elesto:subject",
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.CapabilityInvocation},
							nil,
						),
					),
				)
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				// actual test
				req = *didmod.NewMsgSetVerificationRelationships(
					"did:cosmos:elesto:subject",
					"did:cosmos:elesto:subject#key-1",
					[]string{didmod.Authentication},
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)

				errExp = sdkerrors.Wrapf(didmod.ErrUnauthorized, "signer account %s not authorized to update the target did document at %s", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8", "did:cosmos:elesto:subject")
			},
		},
		{
			"FAIL: can not add verification relationship, invalid relationship provided",
			func() {
				// setup
				didDoc, _ := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								//"did:cosmos:elesto:subject#cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.Authentication},
							nil,
						),
					),
				)
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				// actual test
				req = *didmod.NewMsgSetVerificationRelationships(
					"did:cosmos:elesto:subject",
					"did:cosmos:elesto:subject#key-1",
					nil,
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
				errExp = sdkerrors.Wrap(didmod.ErrEmptyRelationships, "at least a verification relationship is required")
			},
		},
		{
			"FAIL: verification method does not exist ",
			func() {
				// setup
				didDoc, _ := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.Authentication},
							nil,
						),
					),
				)
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				// actual test
				req = *didmod.NewMsgSetVerificationRelationships(
					"did:cosmos:elesto:subject",
					"did:cosmos:elesto:subject#key-does-not-exists",
					[]string{didmod.Authentication, didmod.CapabilityInvocation},
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
				errExp = sdkerrors.Wrapf(didmod.ErrVerificationMethodNotFound, "verification method %v not found", "did:cosmos:elesto:subject#key-does-not-exists")
			},
		},
		{
			"PASS: add a new relationship",
			func() {
				// setup
				didDoc, _ := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.Authentication},
							nil,
						),
					),
				)
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				// actual test
				req = *didmod.NewMsgSetVerificationRelationships(
					"did:cosmos:elesto:subject",
					"did:cosmos:elesto:subject#key-1",
					[]string{didmod.Authentication, didmod.CapabilityInvocation},
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
				errExp = nil
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			tc.malleate()

			_, err := server.SetVerificationRelationships(sdk.WrapSDKContext(suite.ctx), &req)

			if errExp == nil {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(errExp.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestHandleMsgRevokeVerification() {
	var (
		req    didmod.MsgRevokeVerification
		errExp error
	)

	server := NewMsgServerImpl(suite.keeper)

	testCases := []struct {
		name     string
		malleate func()
	}{
		{
			"FAIL: can not revoke verification, did does not exist",
			func() {
				req = *didmod.NewMsgRevokeVerification("did:cosmos:cash:2222", "service-id", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
				errExp = sdkerrors.Wrapf(didmod.ErrDidDocumentNotFound, "did document at %s not found", "did:cosmos:cash:2222")
			},
		},
		{
			"FAIL: can not revoke verification, not found",
			func() {
				didDoc, _ := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.Authentication},
							nil,
						),
					),
				)
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *didmod.NewMsgRevokeVerification(didDoc.Id, "did:cosmos:elesto:subject#not-existent", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
				errExp = sdkerrors.Wrapf(didmod.ErrVerificationMethodNotFound, "verification method id: %v", "did:cosmos:elesto:subject#not-existent")
			},
		},
		{
			"FAIL: can not revoke verification, unauthorized",
			func() {
				signer := "controller-1"
				didDoc, _ := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.CapabilityDelegation},
							nil,
						),
					),
				)

				vmID := "did:cosmos:elesto:subject#key-1"

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				// controller-1 does not exists
				req = *didmod.NewMsgRevokeVerification(didDoc.Id, vmID, signer)

				errExp = sdkerrors.Wrapf(didmod.ErrUnauthorized, "signer account %s not authorized to update the target did document at %s", signer, didDoc.Id)
			},
		},
		{
			"PASS: can revoke verification",
			func() {
				didDoc, _ := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.Authentication},
							nil,
						),
					),
				)

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *didmod.NewMsgRevokeVerification(didDoc.Id,
					"did:cosmos:elesto:subject#key-1",
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				)
				errExp = nil
			},
		},
	}
	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("TestHandleMsgRevokeVerification#%v", i), func() {
			tc.malleate()

			_, err := server.RevokeVerification(sdk.WrapSDKContext(suite.ctx), &req)

			if errExp == nil {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(errExp.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestHandleMsgAddService() {
	var (
		req    didmod.MsgAddService
		errExp error
	)

	server := NewMsgServerImpl(suite.keeper)

	testCases := []struct {
		name     string
		malleate func()
	}{
		{
			"FAIL: can not add service, did does not exist",
			func() {
				service := didmod.NewService(
					"did:cosmos:elesto:subject#linked-domain",
					"LinkedDomains",
					"https://elesto.network",
				)
				req = *didmod.NewMsgAddService("did:cosmos:elesto:subject", service, "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
				errExp = sdkerrors.Wrapf(didmod.ErrDidDocumentNotFound, "did document at %s not found", "did:cosmos:elesto:subject")
			},
		},
		{
			"FAIL: can not add service, service not defined",
			func() {
				didDoc, _ := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								didmod.DID("did:cosmos:elesto:subject"),
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.Authentication},
							nil,
						),
					),
				)
				// create the did
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				// try adding a service
				req = *didmod.NewMsgAddService("did:cosmos:elesto:subject", nil, "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
				errExp = sdkerrors.Wrap(didmod.ErrInvalidInput, "service is not defined")
			},
		},
		{
			"FAIL: cannot add service to did document (unauthorized, wrong relationship)",
			func() {
				signer := "subject"
				didDoc, _ := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								didmod.DID("did:cosmos:elesto:subject"),
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.CapabilityInvocation, didmod.CapabilityDelegation},
							nil,
						),
					),
				)

				service := didmod.NewService(
					"did:cosmos:elesto:subject#linked-domain",
					"LinkedDomains",
					"https://elesto.network",
				)

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *didmod.NewMsgAddService(didDoc.Id, service, signer)

				errExp = sdkerrors.Wrapf(didmod.ErrUnauthorized, "signer account %s not authorized to update the target did document at %s", signer, didDoc.Id)
			},
		},
		{
			"FAIL: cannot add service to did document with an empty type",
			func() {
				didDoc, _ := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.Authentication},
							nil,
						),
					),
				)

				service := didmod.NewService(
					"did:cosmos:elesto:subject#linked-domain",
					"",
					"https://elesto.network",
				)

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *didmod.NewMsgAddService(didDoc.Id, service, "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
				errExp = sdkerrors.Wrap(didmod.ErrInvalidInput, "service type cannot be empty")
			},
		},
		{
			"FAIL: duplicated service",
			func() {
				//signer := "subject"
				signer := "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"
				didDoc, _ := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.Authentication},
							nil,
						),
					),
					didmod.WithServices(
						didmod.NewService(
							"did:cosmos:elesto:subject#linked-domain",
							"LinkedDomains",
							"https://elesto.network",
						),
					),
				)

				service := didmod.NewService(
					"did:cosmos:elesto:subject#linked-domain",
					"LinkedDomains",
					"https://elesto.network",
				)

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				if _, found := suite.keeper.GetDidDocument(suite.ctx, []byte(didDoc.Id)); !found {
					suite.FailNow("test setup failed, did not stored ")
				}

				req = *didmod.NewMsgAddService(didDoc.Id, service, signer)
				errExp = sdkerrors.Wrapf(didmod.ErrInvalidInput, "duplicated service id did:cosmos:elesto:subject#linked-domain")
			},
		},
		{
			"PASS: can add service to did document",
			func() {
				signer := "subject"
				didDoc, err := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								didmod.NewBlockchainAccountID("foochainid", signer),
								didmod.CosmosAccountAddress,
							),
							[]string{didmod.Authentication},
							nil,
						),
					),
				)

				if err != nil {
					suite.FailNow("test setup failed: ", err)
				}

				service := didmod.NewService(
					"did:cosmos:elesto:subject#linked-domain",
					"LinkedDomains",
					"https://elesto.network",
				)

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *didmod.NewMsgAddService(didDoc.Id, service, signer)
				errExp = nil
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			tc.malleate()

			_, err := server.AddService(sdk.WrapSDKContext(suite.ctx), &req)

			if errExp == nil {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(errExp.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestHandleMsgDeleteService() {
	var (
		req    didmod.MsgDeleteService
		errExp error
	)

	server := NewMsgServerImpl(suite.keeper)

	testCases := []struct {
		name     string
		malleate func()
	}{
		{
			"FAIL: can not delete service, did does not exist",
			func() {
				req = *didmod.NewMsgDeleteService("did:cosmos:cash:2222", "service-id", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
				errExp = sdkerrors.Wrapf(didmod.ErrDidDocumentNotFound, "did document at %s not found", "did:cosmos:cash:2222")
			},
		},
		{
			"PASS: can delete service from did document",
			func() {
				didDoc, _ := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.Authentication},
							nil,
						),
					),
					didmod.WithServices(
						didmod.NewService(
							"did:cosmos:elesto:subject#linked-domain",
							"LinkedDomains",
							"https://elesto.network",
						),
					),
				)

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *didmod.NewMsgDeleteService(didDoc.Id, "did:cosmos:elesto:subject#linked-domain", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
				errExp = nil
			},
		},
		{
			"FAIL: cannot remove an invalid serviceID",
			func() {

				didDoc, _ := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.Authentication},
							nil,
						),
					),
				)

				serviceID := ""

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *didmod.NewMsgDeleteService(didDoc.Id, serviceID, "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
				errExp = sdkerrors.Wrapf(didmod.ErrInvalidState, "the did document doesn't have services associated")
			},
		},
		{
			"FAIL: unauthorized (wrong relationship)",
			func() {
				didDoc, _ := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.CapabilityInvocation},
							nil,
						),
					),
				)

				serviceID := "did:cosmos:elesto:subject#linked-domain"

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				req = *didmod.NewMsgDeleteService(didDoc.Id, serviceID, "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
				errExp = sdkerrors.Wrapf(didmod.ErrUnauthorized, "signer account %s not authorized to update the target did document at %s", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8", didDoc.Id)
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf(tc.name), func() {
			tc.malleate()

			_, err := server.DeleteService(sdk.WrapSDKContext(suite.ctx), &req)

			if errExp == nil {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(errExp.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestHandleMsgAddController() {
	var (
		req    didmod.MsgAddController
		errExp error
	)

	server := NewMsgServerImpl(suite.keeper)

	testCases := []struct {
		name     string
		malleate func()
	}{
		{
			"FAIL: cannot add controller, did doesn't exist",
			func() {
				req = *didmod.NewMsgAddController(
					"did:cosmos:elesto:subject",
					"did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
				errExp = sdkerrors.Wrapf(didmod.ErrDidDocumentNotFound, "did document at %s not found", "did:cosmos:elesto:subject")
			},
		},
		{
			"FAIL: controller is not a valid did",
			func() {
				didDoc, _ := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								didmod.DID("did:cosmos:elesto:subject"),
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.Authentication},
							nil,
						),
					),
				)
				// create the did
				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)
				// try adding a service
				req = *didmod.NewMsgAddController("did:cosmos:elesto:subject", "", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
				errExp = sdkerrors.Wrap(didmod.ErrInvalidDIDFormat, "did document controller validation error ''")
			},
		},
		{
			"FAIL: signer not authorized to change controller",
			func() {

				didDoc, _ := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								didmod.DID("did:cosmos:elesto:subject"),
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.CapabilityInvocation, didmod.CapabilityDelegation},
							nil,
						),
					),
				)

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)

				req = *didmod.NewMsgAddController(
					"did:cosmos:elesto:subject",
					"did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
					"cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2", // does not match the pub key (it's the new controller)
				)

				errExp = sdkerrors.Wrapf(didmod.ErrUnauthorized, "signer account %s not authorized to update the target did document at %s", "cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2", didDoc.Id)
			},
		},
		{
			"FAIL: controller is not type key",
			func() {

				didDoc, _ := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								didmod.DID("did:cosmos:elesto:subject"),
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.Authentication},
							nil,
						),
					),
				)

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)

				req = *didmod.NewMsgAddController(
					"did:cosmos:elesto:subject",
					"did:cosmos:foochain:whatever",
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8", // does not match the pub key (it's the new controller)
				)

				errExp = sdkerrors.Wrapf(didmod.ErrInvalidInput, "did document controller 'did:cosmos:foochain:whatever' must be of type key")
			},
		},
		{
			"PASS: can add controller (via authentication relationship)",
			func() {
				didDoc, err := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								didmod.DID("did:cosmos:elesto:subject"),
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.Authentication},
							nil,
						),
					),
				)

				if err != nil {
					suite.FailNow("test setup failed: ", err)
				}

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)

				req = *didmod.NewMsgAddController(
					"did:cosmos:elesto:subject",
					"did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
				errExp = nil
			},
		},
		{
			"PASS: can add controller (via controller)",
			func() {
				didDoc, err := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								didmod.DID("did:cosmos:elesto:subject"),
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.AssertionMethod},
							nil,
						),
					),
					didmod.WithControllers("did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2"),
				)

				if err != nil {
					suite.FailNow("test setup failed: ", err)
				}

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)

				req = *didmod.NewMsgAddController(
					"did:cosmos:elesto:subject",
					"did:cosmos:key:cosmos17t8t3t6a6vpgk69perfyq930593sa8dn4kzsdf",
					"cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2") // this is the controller

				errExp = nil
			},
		},
		{
			"PASS: controller already added (duplicated)",
			func() {
				didDoc, err := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								didmod.DID("did:cosmos:elesto:subject"),
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.Authentication},
							nil,
						),
					),
				)

				if err != nil {
					suite.FailNow("test setup failed: ", err)
				}

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)

				req = *didmod.NewMsgAddController(
					"did:cosmos:elesto:subject",
					"did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
				errExp = nil
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			tc.malleate()

			_, err := server.AddController(sdk.WrapSDKContext(suite.ctx), &req)

			if errExp == nil {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(errExp.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestHandleMsgDeleteController() {
	var (
		req    didmod.MsgDeleteController
		errExp error
	)

	server := NewMsgServerImpl(suite.keeper)

	// FAIL: cannot delete controller, did doesn't exist
	// FAIL: signer not authorized to change controller
	// PASS: controller removed
	testCases := []struct {
		name     string
		malleate func()
	}{
		{
			"FAIL: cannot delete controller, did doesn't exist",
			func() {
				req = *didmod.NewMsgDeleteController(
					"did:cosmos:elesto:subject",
					"did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
				errExp = sdkerrors.Wrapf(didmod.ErrDidDocumentNotFound, "did document at %s not found", "did:cosmos:elesto:subject")
			},
		},
		{
			"FAIL: signer not authorized to change controller",
			func() {

				didDoc, _ := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								didmod.DID("did:cosmos:elesto:subject"),
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.CapabilityInvocation, didmod.CapabilityDelegation},
							nil,
						),
					),
				)

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)

				req = *didmod.NewMsgDeleteController(
					"did:cosmos:elesto:subject",
					"did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
					"cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2", // does not match the pub key (it's the new controller)
				)

				errExp = sdkerrors.Wrapf(didmod.ErrUnauthorized, "signer account %s not authorized to update the target did document at %s", "cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2", didDoc.Id)
			},
		},
		{
			"PASS: can delete controller (via authentication relationship)",
			func() {
				didDoc, err := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								didmod.DID("did:cosmos:elesto:subject"),
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.Authentication},
							nil,
						),
					),
				)

				if err != nil {
					suite.FailNow("test setup failed: ", err)
				}

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)

				req = *didmod.NewMsgDeleteController(
					"did:cosmos:elesto:subject",
					"did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
					"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8")
				errExp = nil
			},
		}, {
			"PASS: can delete controller (via controller)",
			func() {
				didDoc, err := didmod.NewDidDocument(
					"did:cosmos:elesto:subject",
					didmod.WithVerifications(
						didmod.NewVerification(
							didmod.NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								didmod.DID("did:cosmos:elesto:subject"),
								didmod.NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								didmod.EcdsaSecp256k1VerificationKey2019,
							),
							[]string{didmod.AssertionMethod},
							nil,
						),
					),
					didmod.WithControllers(
						"did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
						"did:cosmos:key:cosmos17t8t3t6a6vpgk69perfyq930593sa8dn4kzsdf",
					),
				)

				if err != nil {
					suite.FailNow("test setup failed: ", err)
				}

				suite.keeper.SetDidDocument(suite.ctx, []byte(didDoc.Id), didDoc)

				req = *didmod.NewMsgDeleteController(
					"did:cosmos:elesto:subject",
					"did:cosmos:key:cosmos17t8t3t6a6vpgk69perfyq930593sa8dn4kzsdf",
					"cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2", // this is the controller
				)

				errExp = nil
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			tc.malleate()

			_, err := server.DeleteController(sdk.WrapSDKContext(suite.ctx), &req)

			if errExp == nil {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(errExp.Error(), err.Error())
			}
		})
	}
}
