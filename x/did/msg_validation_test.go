package did

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestMsgCreateDidDocument(t *testing.T) {
	tests := []struct {
		id            string
		verifications Verifications
		services      Services
		owner         string
		expectPass    bool
	}{
		{
			"did:cosmos:net:elesto:whatever",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:cosmos:net:elesto:whatever#1",
						EcdsaSecp256k1VerificationKey2019.String(),
						"did:cosmos:net:elesto:whatever",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					[]string{},
				},
			},
			Services{},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			true,
		},
		{
			"did:cosmos:net:elesto:whatever",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:cosmos:net:elesto:whatever#1",
						CosmosAccountAddress.String(),
						"did:cosmos:net:elesto:whatever",
						&VerificationMethod_BlockchainAccountID{""},
					},
					[]string{},
				},
			},
			Services{},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // empty pub key
		},
		{
			"did:cosmos:net:elesto:whatever",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:cosmos:net:elesto:whatever#1",
						"",
						"did:cosmos:net:elesto:whatever",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					[]string{},
				},
			},
			Services{},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // emtpy verification method type
		},
		{
			"did:cosmos:net:elesto:whatever",
			Verifications{
				&Verification{
					[]string{},
					&VerificationMethod{
						"did:cosmos:net:elesto:whatever#1",
						EcdsaSecp256k1VerificationKey2019.String(),
						"did:cosmos:net:elesto:whatever",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					[]string{},
				},
			},
			Services{},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // empty relationships
		},
		{
			"did:cosmos:net:elesto:whatever",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:cosmos:net:elesto:whatever#/asd 123",
						EcdsaSecp256k1VerificationKey2019.String(),
						"did:cosmos:net:elesto:whatever",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					[]string{},
				},
			},
			Services{},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // invalid method id
		},
		{
			"did:cosmos:net:elesto:whatever",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:cosmos:net:elesto:whatever#1",
						EcdsaSecp256k1VerificationKey2019.String(),
						"did:cosmos:net:elesto:whatever",
						&VerificationMethod_PublicKeyHex{""},
					},
					[]string{},
				},
			},
			Services{},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // empty verification key
		},
		{
			"did:cosmos:net:elesto:whatever",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:cosmos:net:elesto:whatever#1",
						"",
						"did:cosmos:net:elesto:whatever",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					[]string{},
				},
			},
			Services{},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // empty verification method type
		},
		{
			"did:cosmos:net:elesto:whatever",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:cosmos:net:elesto:whatever#1",
						CosmosAccountAddress.String(),
						"",
						&VerificationMethod_BlockchainAccountID{"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"},
					},
					[]string{},
				},
			},
			Services{},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // invalid verification method controller
		},
		{
			"did:cosmos:net:elesto:whatever",
			Verifications{},
			Services{},
			"owner",
			false, // empty verifications
		},

		{
			"invalid did",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:cosmos:net:elesto:whatever#1",
						EcdsaSecp256k1VerificationKey2019.String(),
						"cont",
						&VerificationMethod_BlockchainAccountID{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					[]string{},
				},
			},
			Services{},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // invalid did
		},
		{
			"did:cosmos:net:elesto:whatever",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:cosmos:net:elesto:whatever#1",
						EcdsaSecp256k1VerificationKey2019.String(),
						"did:cosmos:net:elesto:whatever",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					[]string{},
				},
			},
			Services{
				&Service{
					"the:agent:service",
					"DIDCommMessaging",
					"https://agent.xyz/agent/123",
				},
			},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			true,
		},
		{
			"did:cosmos:net:elesto:whatever",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:cosmos:net:elesto:whatever#1",
						EcdsaSecp256k1VerificationKey2019.String(),
						"did:cosmos:net:elesto:whatever",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					[]string{},
				},
			},
			Services{
				&Service{
					"the:agent:service",
					"",
					"https://agent.xyz/agent/123",
				},
			},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // empty service type
		},
		{
			"did:cosmos:net:elesto:whatever",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:cosmos:net:elesto:whatever#1",
						EcdsaSecp256k1VerificationKey2019.String(),
						"did:cosmos:net:elesto:whatever",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					[]string{},
				},
			},
			Services{
				&Service{
					"",
					"DIDCommMessaging",
					"https://agent.xyz/agent/123",
				},
			},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // service id is not valid
		},
		{
			"did:cosmos:net:elesto:whatever",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:cosmos:net:elesto:whatever#1",
						EcdsaSecp256k1VerificationKey2019.String(),
						"did:cosmos:net:elesto:whatever",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					[]string{},
				},
			},
			Services{
				&Service{
					"this:is:fine",
					"DIDCommMessaging",
					"",
				},
			},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // service id is not valid
		},
	}

	for i, tc := range tests {
		msg := NewMsgCreateDidDocument(
			tc.id,
			tc.verifications,
			tc.services,
			tc.owner,
		)

		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: TestMsgCreateDidDocument#%v", i)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: TestMsgCreateDidDocument#%v", i)
		}
	}
}

func TestMsgAddVerification(t *testing.T) {
	tests := []struct {
		id         string
		auth       Verification
		owner      string
		expectPass bool
	}{
		{
			"did:cosmos:net:elesto:subject",
			Verification{
				[]string{string(Authentication)},
				&VerificationMethod{
					"did:cosmos:net:elesto:subject#1",
					EcdsaSecp256k1VerificationKey2019.String(),
					"did:cosmos:net:elesto:subject",
					&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
				},
				[]string{},
			},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			true,
		},
		{
			"something not right",
			Verification{
				[]string{string(Authentication)},
				&VerificationMethod{
					"did:cosmos:net:elesto:subject#1",
					EcdsaSecp256k1VerificationKey2019.String(),
					"did:cosmos:net:elesto:subject",
					&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
				},
				[]string{},
			},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // invalid did
		},
	}

	for i, tc := range tests {
		msg := NewMsgAddVerification(
			tc.id,
			&tc.auth,
			tc.owner,
		)

		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: TestMsgAddVerification#%v", i)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: TestMsgAddVerification#%v", i)
		}
	}
}

func TestMsgRevokeVerification(t *testing.T) {
	tests := []struct {
		id         string
		key        string
		signer     string
		expectPass bool
	}{
		{
			"did:cosmos:net:elesto:subject",
			"did:cosmos:net:elesto:subject#key-method-1",
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			true,
		},
		{
			"invalid did",
			"did:cosmos:net:elesto:subject#key-method-1",
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, // invalid did
		},
		{
			"did:cosmos:net:elesto:subject",
			"did:cosmos:net:elesto:subject  #   key-method-1",
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, // invalid method id
		},
		{
			"did:cosmos:net:elesto:subject",
			"did:cosmos:net:elesto:subject#key-method-1",
			"",
			true, // empty signer
		},
	}

	for i, tc := range tests {
		msg := NewMsgRevokeVerification(
			tc.id,
			tc.key,
			tc.signer,
		)

		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: TestMsgRevokeVerification#%v", i)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: TestMsgRevokeVerification#%v", i)
		}
	}
}

func TestMsgSetVerificationRelationships(t *testing.T) {
	tests := []struct {
		id            string
		key           string
		relationships []string
		signer        string
		expectPass    bool
	}{
		{
			"did:cosmos:net:elesto:subject",
			"did:cosmos:net:elesto:subject#key-method-1",
			[]string{"authorization", "keyExchange"},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			true,
		},
		{
			"did:cosmos:net:elesto:subject",
			"did:cosmos:net:elesto:subject#key-method-1",
			[]string{"authorization"},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			true,
		},
		{
			"did:cosmos:net:elesto:subject",
			"did:cosmos:net:elesto:subject  #   key-method-1",
			[]string{"authorization", "keyExchange"},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, // invalid method id
		},
		{
			"invalid did",
			"did:cosmos:net:elesto:subject#key-method-1",
			[]string{"authorization", "keyExchange"},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, // invalid did
		},
		{
			"did:cosmos:net:elesto:subject",
			"did:cosmos:net:elesto:subject#key-method-1",
			[]string{},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, // empty relationship
		},
		{
			"did:cosmos:net:elesto:subject",
			"did:cosmos:net:elesto:subject#key-method-1",
			[]string{"authorization", "keyExchange"},
			"",
			true, // empty signer
		},
	}

	for i, tc := range tests {
		t.Logf("TestMsgRevokeVerification#%d", i)
		msg := NewMsgSetVerificationRelationships(
			tc.id,
			tc.key,
			tc.relationships,
			tc.signer,
		)

		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: TestMsgSetVerificationRelationships#%v", i)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: TestMsgSetVerificationRelationships#%v", i)
		}
	}
}

func TestMsgAddService(t *testing.T) {
	tests := []struct {
		id         string
		service    *Service
		signer     string
		expectPass bool
	}{
		{
			"did:cosmos:net:elesto:subject",
			&Service{
				Id:              "a:valid:url",
				Type:            "DIDCommMessaging",
				ServiceEndpoint: "https://agent.xyz/validate",
			},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			true,
		},
		{
			"invalid did",
			&Service{
				Id:              "my:agent",
				Type:            "DIDCommMessaging",
				ServiceEndpoint: "https://agent.xyz/validate",
			},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, // invalid did
		},
		{
			"did:cosmos:net:elesto:subject",
			&Service{
				Id:              "",
				Type:            "DIDCommMessaging",
				ServiceEndpoint: "https://agent.xyz/validate",
			},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, // invalid agent id
		},
		{
			"did:cosmos:net:elesto:subject",
			&Service{
				Id:              "my:agent",
				Type:            "",
				ServiceEndpoint: "https://agent.xyz/validate",
			},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, // empty type
		},
		{
			"did:cosmos:net:elesto:subject",
			&Service{
				Id:              "my:agent",
				Type:            "DIDCommMessaging",
				ServiceEndpoint: "",
			},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, // empty service endpoint
		},
	}

	for i, tc := range tests {
		t.Logf("TestMsgRevokeVerification#%d", i)
		msg := NewMsgAddService(
			tc.id,
			tc.service,
			tc.signer,
		)

		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: TestMsgAddService#%v", i)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: TestMsgAddService#%v", i)
		}
	}
}

func TestMsgDeleteService(t *testing.T) {
	tests := []struct {
		id         string
		serviceID  string
		signer     string
		expectPass bool
	}{
		{
			"did:cosmos:net:elesto:subject",
			"my:service:uri",
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			true,
		},
		{
			"invalid did",
			"my:service:uri",
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, //invalid did
		},
		{
			"did:cosmos:net:elesto:subject",
			"",
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, // empty service id
		},
	}

	for i, tc := range tests {
		t.Logf("TestMsgRevokeVerification#%d", i)
		msg := NewMsgDeleteService(
			tc.id,
			tc.serviceID,
			tc.signer,
		)

		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: TestMsgDeleteService#%v", i)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: TestMsgDeleteService#%v", i)
		}
	}
}

func TestMsgAddController_ValidateBasic(t *testing.T) {
	type fields struct {
		Id            string
		ControllerDid string
		Signer        string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			"PASS: controller is valid",
			fields{
				Id:            "did:cosmos:net:foochain:12345",
				ControllerDid: "did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				Signer:        "cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
			},
			nil,
		},
		{
			"FAIL: invalid did",
			fields{
				Id:            "not a did",
				ControllerDid: "did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				Signer:        "cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
			},
			sdkerrors.Wrap(ErrInvalidDIDFormat, "not a did"),
		},
		{
			"FAIL: invalid controller did",
			fields{
				Id:            "did:cosmos:net:foochain:12345",
				ControllerDid: "did:cosmos:net:foochain:whatever",
				Signer:        "cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
			},
			sdkerrors.Wrap(ErrInvalidDIDFormat, "did:cosmos:net:foochain:whatever"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgAddController{
				Id:            tt.fields.Id,
				ControllerDid: tt.fields.ControllerDid,
				Signer:        tt.fields.Signer,
			}

			err := msg.ValidateBasic()
			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
		})
	}
}

func TestMsgDeleteController_ValidateBasic(t *testing.T) {
	type fields struct {
		Id            string
		ControllerDid string
		Signer        string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			"PASS: controller is valid",
			fields{
				Id:            "did:cosmos:net:foochain:12345",
				ControllerDid: "did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				Signer:        "cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
			},
			nil,
		},
		{
			"FAIL: invalid did",
			fields{
				Id:            "not a did",
				ControllerDid: "did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				Signer:        "cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
			},
			sdkerrors.Wrap(ErrInvalidDIDFormat, "not a did"),
		},
		{
			"FAIL: invalid controller did",
			fields{
				Id:            "did:cosmos:net:foochain:12345",
				ControllerDid: "not a did",
				Signer:        "cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
			},
			sdkerrors.Wrap(ErrInvalidDIDFormat, "not a did"),
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgDeleteController{
				Id:            tt.fields.Id,
				ControllerDid: tt.fields.ControllerDid,
				Signer:        tt.fields.Signer,
			}
			err := msg.ValidateBasic()
			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
		})
	}
}
