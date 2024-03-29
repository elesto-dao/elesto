package did

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewChainDID(t *testing.T) {
	tests := []struct {
		did   string
		chain string
		want  DID
	}{
		{
			"subject",
			"cash",
			DID("did:cosmos:cash:subject"),
		},
		{
			"",
			"cash",
			DID("did:cosmos:cash:"),
		},
		{
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			"cosmoshub",
			DID("did:cosmos:cosmoshub:cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75"),
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestDID#", i), func(t *testing.T) {
			if got := NewChainDID(tt.chain, tt.did); got != tt.want {
				t.Errorf("DID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewKeyDID(t *testing.T) {
	tests := []struct {
		name    string
		account string
		want    DID
	}{
		{
			"PASS: valid account",
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			"did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewKeyDID(tt.account), "NewKeyDID(%v)", tt.account)
		})
	}
}

func TestDID_NewVerificationMethodID(t *testing.T) {
	tests := []struct {
		name string
		did  DID
		vmID string
		want string
	}{
		{
			"PASS: generated vmId for network DID",
			DID("did:cosmos:foochain:whatever"),
			"123456",
			"did:cosmos:foochain:whatever#123456",
		},
		{
			"PASS: generated vmId for key DID",
			DID("did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"),
			"123456",
			"did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8#123456",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.did.NewVerificationMethodID(tt.vmID), "NewVerificationMethodID(%v)", tt.vmID)
		})
	}
}

func TestIsValidDID(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"did:cosmos:elesto:00000000-0000-0000-0000-000000000000", true},
		{"did:cosmos:cash:subject", true},
		{"did:cosmos:key:cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75", true},
		{"did:cosmos:key:cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75#key-1", false},
		{"did:subject", false},
		{"DID:cosmos:elesto:subject", false},
		{"d1d:cash:subject", false},
		{"d1d:#:subject", false},
		{"", false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestIsValidDID#", i), func(t *testing.T) {
			if got := IsValidDID(tt.input); got != tt.want {
				t.Errorf("IsValidDID(%s) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsValidDIDURL(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"did:cosmos:subject", true},
		{"did:cosmos:key:cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75", true},
		{"did:cosmos:key:cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75#key-1", true},
		{"did:cosmos:key:cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75?key=1", true},
		{"did:cosmos:cosmoscash-testnet:575d062c-d110-42a9-9c04-cb1ff8c01f06#Z46DAL1MrJlVW_WmJ19WY8AeIpGeFOWl49Qwhvsnn2M", true},
		{"did:subject", false},
		{"DID:cosmos:subject", false},
		{"d1d:cash:subject", false},
		{"d1d:#:subject", false},
		{"", false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestIsValidDIDURL#", i), func(t *testing.T) {
			if got := IsValidDIDURL(tt.input); got != tt.want {
				t.Errorf("IsValidDIDURL(%s) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsValidRFC3986Uri(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			"[][àséf",
			false,
		},
		{
			"# \u007e // / / ///// //// // / / ??? ?? 不",
			false,
		},
		{`ftp://ftp.is.co.za/rfc/rfc1808.txt`, true},
		{`http://www.ietf.org/rfc/rfc2396.txt`, true},
		{`ldap://[2001:db8::7]/c=GB?objectClass?one`, true},
		{`mailto:John.Doe@example.com`, true},
		{`news:comp.infosystems.www.servers.unix`, true},
		{`tel:+1-816-555-1212`, true},
		{`telnet://192.0.2.16:80/`, true},
		{`urn:oasis:names:specification:docbook:dtd:xml:4.1.2`, true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestIsValidRFC3986Uri#", i), func(t *testing.T) {
			if got := IsValidRFC3986Uri(tt.input); got != tt.want {
				t.Errorf("IsValidRFC3986Uri() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidDIDDocument(t *testing.T) {
	tests := []struct {
		name  string
		didFn func() *DidDocument
		want  bool
	}{
		{
			"PASS: document is valid",
			func() *DidDocument {
				return &DidDocument{
					Context: []string{contextDIDBase},
					Id:      "did:cosmos:cash:1",
				}
			},
			true, // all good
		},
		{
			"FAIL: empty context",
			func() *DidDocument {
				return &DidDocument{
					Context: []string{},
					Id:      "did:cosmos:cash:1",
				}
			},
			false, // missing context
		},
		{
			"PASS: minimal did document",
			func() *DidDocument {
				dd, err := NewDidDocument("did:cosmos:cash:1")
				assert.NoError(t, err)
				return &dd
			},
			true, // all good
		},
		{
			"FAIL: empty did",
			func() *DidDocument {
				return &DidDocument{
					Context: []string{contextDIDBase},
					Id:      "",
				}
			},
			false, // empty id
		},
		{
			"FAIL: nil did document",
			func() *DidDocument {
				return nil
			},
			false, // nil pointer
		},
		{
			"PASS: did with valid controller",
			func() *DidDocument {
				dd, err := NewDidDocument("did:cosmos:key:cas:1", WithControllers(
					"did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
				))
				assert.NoError(t, err)
				return &dd
			},
			true,
		},
		{
			"FAIL: invalid controller",
			func() *DidDocument {
				return &DidDocument{
					Context: []string{contextDIDBase},
					Id:      "did:cosmos:foochain:1",
					Controller: []string{
						"did:cosmos:foochain:whatever",
					},
				}
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint("TestIsValidDIDDocument#", tt.name), func(t *testing.T) {
			if got := IsValidDIDDocument(tt.didFn()); got != tt.want {
				t.Errorf("TestIsValidDIDDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateVerification(t *testing.T) {
	tests := []struct {
		name    string
		v       *Verification
		wantErr error
	}{
		{
			name:    "FAIL: verification is nil",
			v:       nil,
			wantErr: sdkerrors.Wrap(ErrInvalidInput, "verification is not defined"),
		},
		{
			name: "FAIL: invalid verification method id",
			v: NewVerification(
				NewVerificationMethod(
					"not:a:did",
					"did:cosmos:elesto:subject",
					NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
					EcdsaSecp256k1VerificationKey2019,
				),
				[]string{string(Authentication)},
				nil,
			),
			wantErr: sdkerrors.Wrapf(ErrInvalidDIDURLFormat, "verification method id: %v", "not:a:did"),
		},
		{
			name: "FAIL: invalid method controller",
			v: NewVerification(
				NewVerificationMethod(
					"did:cosmos:elesto:subject#key-1",
					"not:a:did",
					NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
					EcdsaSecp256k1VerificationKey2019,
				),
				[]string{string(Authentication)},
				nil,
			),
			wantErr: sdkerrors.Wrapf(ErrInvalidDIDFormat, "verification method controller %v", "not:a:did"),
		},
		{
			name: "FAIL: empty method type",
			v: NewVerification(
				NewVerificationMethod(
					"did:cosmos:elesto:subject#key-1",
					"did:cosmos:elesto:subject",
					NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
					"",
				),
				[]string{Authentication},
				nil,
			),
			wantErr: sdkerrors.Wrapf(ErrInvalidInput, "verification method type not set for verification method %s", "did:cosmos:elesto:subject#key-1"),
		},
		{
			name: "FAIL: invalid blockchain id",
			v: NewVerification(
				NewVerificationMethod(
					"did:cosmos:elesto:subject#key-1",
					"did:cosmos:elesto:subject",
					&VerificationMethod_BlockchainAccountID{BlockchainAccountID: ""},
					CosmosAccountAddress,
				),
				[]string{Authentication},
				nil,
			),
			wantErr: sdkerrors.Wrapf(ErrInvalidInput, "verification material blockchainAccountId is empty for verification method id did:cosmos:elesto:subject#key-1"),
		},
		{
			name: "FAIL: invalid multibase key",
			v: NewVerification(
				NewVerificationMethod(
					"did:cosmos:elesto:subject#key-1",
					"did:cosmos:elesto:subject",
					&VerificationMethod_PublicKeyMultibase{PublicKeyMultibase: ""},
					EcdsaSecp256k1VerificationKey2019,
				),
				[]string{Authentication},
				nil,
			),
			wantErr: sdkerrors.Wrapf(ErrInvalidInput, "verification material publicKeyMultibase is empty for verification method id did:cosmos:elesto:subject#key-1"),
		},
		{
			name: "FAIL: invalid hex key",
			v: NewVerification(
				NewVerificationMethod(
					"did:cosmos:elesto:subject#key-1",
					"did:cosmos:elesto:subject",
					&VerificationMethod_PublicKeyHex{PublicKeyHex: ""},
					EcdsaSecp256k1VerificationKey2019,
				),
				[]string{string(Authentication)},
				nil,
			),
			wantErr: sdkerrors.Wrapf(ErrInvalidInput, "verification material publicKeyHex is empty for verification method id did:cosmos:elesto:subject#key-1"),
		},
		{
			name: "FAIL: no relationships defined",
			v: NewVerification(
				NewVerificationMethod(
					"did:cosmos:elesto:subject#key-1",
					"did:cosmos:elesto:subject",
					NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
					EcdsaSecp256k1VerificationKey2019,
				),
				nil,
				nil,
			),
			wantErr: sdkerrors.Wrap(ErrEmptyRelationships, "at least a verification relationship is required"), // relationships are nil
		},
		{
			name: "FAIL: undefined verification material",
			v: NewVerification(
				NewVerificationMethod(
					"did:cosmos:elesto:subject#key-1",
					DID("did:cosmos:elesto:subject"),
					nil,
					EcdsaSecp256k1VerificationKey2019,
				),
				[]string{string(AssertionMethod)},
				nil,
			),
			wantErr: sdkerrors.Wrapf(ErrInvalidInput, "verification material '<nil>' unknown for verification method id did:cosmos:elesto:subject#key-1"),
		},
		{
			name: "PASS: can add verification",
			v: NewVerification(
				NewVerificationMethod(
					"did:cosmos:elesto:subject#key-1",
					DID("did:cosmos:elesto:subject"),
					NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
					EcdsaSecp256k1VerificationKey2019,
				),
				[]string{string(AssertionMethod)},
				nil,
			),
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.name), func(t *testing.T) {
			err := ValidateVerification(tt.v)
			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
		})
	}
}

func TestValidateService(t *testing.T) {

	tests := []struct {
		name    string
		s       *Service
		wantErr error
	}{
		{
			name:    "PASS: valid service",
			s:       NewService("did:cosmos:elesto:subject#did-comm", "DIDCommMessaging", "https://agent.abc/abc"),
			wantErr: nil,
		},
		{
			name:    "FAIL: service is nilS",
			s:       nil,
			wantErr: sdkerrors.Wrap(ErrInvalidInput, "service is not defined"),
		},
		{
			name:    "FAIL: invalid service id",
			s:       NewService("service", "DIDCommMessaging", "https://agent.abc/abc"),
			wantErr: sdkerrors.Wrapf(ErrInvalidRFC3986UriFormat, "service id %s is not a valid RFC3986 uri", "service"),
		},
		{
			name:    "FAIL: invalid service endpoint",
			s:       NewService("did:cosmos:elesto:subject#did-comm", "DIDCommMessaging", "not-an-uri"),
			wantErr: sdkerrors.Wrapf(ErrInvalidRFC3986UriFormat, "service endpoint %s is not a valid RFC3986 uri", "not-an-uri"),
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestValidateService#", i), func(t *testing.T) {
			err := ValidateService(tt.s)
			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"    a    ", false},
		{"\t", true},
		{"\n", true},
		{"   ", true},
		{"  \t \n", true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestIsEmpty#", i), func(t *testing.T) {
			if got := IsEmpty(tt.input); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDidDocument(t *testing.T) {
	type params struct {
		id      string
		options []DocumentOption
	}
	tests := []struct {
		params  params
		wantDid DidDocument
		wantErr bool
	}{
		{
			params: params{
				"did:cosmos:elesto:subject",
				[]DocumentOption{
					WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								EcdsaSecp256k1VerificationKey2019,
							),
							[]string{
								string(Authentication),
								string(KeyAgreement),
								string(KeyAgreement), // test duplicated relationship
							},
							[]string{
								"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
							},
						),
					),
					WithVerifications( // multiple verifications in separate entity
						NewVerification(
							NewVerificationMethod(
								"did:cosmos:elesto:subject#key-2",
								"did:cosmos:elesto:subject",
								NewBlockchainAccountID("cash", "cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2"),
								CosmosAccountAddress,
							),
							[]string{
								string(Authentication),
							},
							[]string{
								"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
							},
						),
					),
					WithServices(&Service{
						"agent:xyz",
						"DIDCommMessaging",
						"https://agent.xyz/1234",
					}),
					WithControllers("did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2"),
				},
			},
			wantDid: DidDocument{
				Context: []string{
					"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
					contextDIDBase,
				},
				Id:         "did:cosmos:elesto:subject",
				Controller: []string{"did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2"},
				VerificationMethod: []*VerificationMethod{
					{
						"did:cosmos:elesto:subject#key-1",
						EcdsaSecp256k1VerificationKey2019.String(),
						"did:cosmos:elesto:subject",
						&VerificationMethod_PublicKeyMultibase{"F03DFD0A469806D66A23C7C948F55C129467D6D0974A222EF6E24A538FA6882F3D7"},
					},
					{
						"did:cosmos:elesto:subject#key-2",
						CosmosAccountAddress.String(),
						"did:cosmos:elesto:subject",
						&VerificationMethod_BlockchainAccountID{"cosmos:cash:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2"},
					},
				},
				Service: []*Service{
					{
						"agent:xyz",
						"DIDCommMessaging",
						"https://agent.xyz/1234",
					},
				},
				Authentication: []string{"did:cosmos:elesto:subject#key-1", "did:cosmos:elesto:subject#key-2"},
				KeyAgreement:   []string{"did:cosmos:elesto:subject#key-1"},
			},
			wantErr: false,
		},
		{
			params: params{
				"did:cosmos:elesto:subject",
				[]DocumentOption{
					WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								EcdsaSecp256k1VerificationKey2019,
							),
							[]string{
								Authentication,
								KeyAgreement,
							},
							[]string{
								"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
							},
						),
						NewVerification(
							NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1", // duplicate key
								"did:cosmos:elesto:subject",
								NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								EcdsaSecp256k1VerificationKey2019,
							),
							[]string{
								Authentication,
								KeyAgreement,
							},
							[]string{},
						),
					),
					WithServices(&Service{
						"agent:xyz",
						"DIDCommMessaging",
						"https://agent.xyz/1234",
					}),
				},
			},
			wantDid: DidDocument{},
			wantErr: true, // the error is caused by duplicated verification method id
		},
		{
			params: params{
				"did:cosmos:elesto:subject",
				[]DocumentOption{
					WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								NewPublicKeyMultibase([]byte("02503c8ace59c085b15c5f9c2474e9235bcb9694f07516bdc06f7caec788c3dd2c")),
								EcdsaSecp256k1VerificationKey2019,
							),
							[]string{
								Authentication,
								KeyAgreement,
							},
							[]string{
								"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
							},
						),
					),
					WithServices(
						&Service{
							"agent:xyz",
							"DIDCommMessaging",
							"https://agent.xyz/1234",
						},
						&Service{
							"agent:xyz",
							"DIDCommMessaging",
							"https://agent.xyz/1234",
						},
					),
				},
			},
			wantDid: DidDocument{},
			wantErr: true, //duplicated service id
		},
		{
			wantErr: true, // invalid did
			params: params{
				id:      "something not right",
				options: []DocumentOption{},
			},
			wantDid: DidDocument{},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestNewDidDocument#", i), func(t *testing.T) {
			gotDid, err := NewDidDocument(tt.params.id, tt.params.options...)

			if tt.wantErr {
				require.NotNil(t, err, "test: TestNewDidDocument#%v", i)
				return
			}

			require.Nil(t, err, "test: TestNewDidDocument#%v", i)
			assert.Equal(t, tt.wantDid, gotDid)
		})
	}
}

func TestDidDocument_AddControllers(t *testing.T) {

	tests := []struct {
		name                string
		malleate            func() DidDocument
		controllers         []string
		expectedControllers []string
		wantErr             error
	}{
		{
			"PASS: controllers added",
			func() DidDocument {
				dd, _ := NewDidDocument("did:cosmos:elesto:subject",
					WithControllers(
						"did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
						"did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8", // duplicate controllers
					),
				)
				return dd
			},
			[]string{
				"did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
			},
			[]string{
				"did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
				"did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			},
			nil,
		},
		{
			"FAIL: invalid controller did",
			func() DidDocument {
				dd, _ := NewDidDocument("did:cosmos:elesto:subject",
					WithControllers(
						"did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
					),
				)
				return dd
			},
			[]string{
				"did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
				"not a did:cosmos:key:cosmos100000000000000000000000000000000000004",
			},
			[]string{},
			sdkerrors.Wrapf(ErrInvalidDIDFormat, "did document controller validation error 'not a did:cosmos:key:cosmos100000000000000000000000000000000000004'"),
		},
		{
			"FAIL: controller did is not type key",
			func() DidDocument {
				dd, _ := NewDidDocument("did:cosmos:elesto:subject",
					WithControllers(
						"did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
					),
				)
				return dd
			},
			[]string{
				"did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
				"did:cosmos:foochain:1234",
			},
			[]string{},
			sdkerrors.Wrapf(ErrInvalidInput, "did document controller 'did:cosmos:foochain:1234' must be of type key"),
		},
		{
			"PASS: controllers empty",
			func() DidDocument {
				dd, _ := NewDidDocument("did:cosmos:elesto:subject",
					WithControllers(
						"did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
						"did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
					),
				)
				return dd
			},
			nil,
			[]string{
				"did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
				"did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			},
			nil,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestDidDocument_AddControllers#", i), func(t *testing.T) {
			didDoc := tt.malleate()
			err := didDoc.AddControllers(tt.controllers...)

			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
		})
	}
}

func TestDidDocument_DeleteControllers(t *testing.T) {

	tests := []struct {
		name                string
		malleate            func() DidDocument
		controllers         []string
		expectedControllers []string
		wantErr             error
	}{
		{
			"PASS: controllers deleted",
			func() DidDocument {
				dd, _ := NewDidDocument("did:cosmos:elesto:subject",
					WithControllers(
						"did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
						"did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
					),
				)
				return dd
			},
			[]string{
				"did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
			},
			[]string{
				"did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			},
			nil,
		},
		{
			"FAIL: invalid controller did",
			func() DidDocument {
				dd, _ := NewDidDocument("did:cosmos:elesto:subject",
					WithControllers(
						"did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
					),
				)
				return dd
			},
			[]string{
				"not a did:cosmos:key:cosmos100000000000000000000000000000000000004",
			},
			[]string{},
			sdkerrors.Wrapf(ErrInvalidDIDFormat, "did document controller validation error 'not a did:cosmos:key:cosmos100000000000000000000000000000000000004'"),
		},
		{
			"PASS: controllers empty",
			func() DidDocument {
				dd, _ := NewDidDocument("did:cosmos:elesto:subject",
					WithControllers(
						"did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
						"did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
					),
				)
				return dd
			},
			nil,
			[]string{
				"did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
				"did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			},
			nil,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestDidDocument_AddControllers#", i), func(t *testing.T) {
			didDoc := tt.malleate()
			err := didDoc.DeleteControllers(tt.controllers...)

			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
		})
	}
}

func TestDidDocument_AddVerifications(t *testing.T) {
	type params struct {
		malleate      func() DidDocument // build a did document
		verifications []*Verification    // input list of verifications
	}
	tests := []struct {
		params  params
		wantDid DidDocument // expected result
		wantErr bool
	}{
		{
			wantErr: false,
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cosmos:elesto:subject")
					return d
				},
				[]*Verification{
					NewVerification(
						NewVerificationMethod(
							"did:cosmos:elesto:subject#key-1",
							"did:cosmos:elesto:subject",
							NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
							EcdsaSecp256k1VerificationKey2019,
						),
						[]string{
							Authentication,
							KeyAgreement,
						},
						nil,
					),
					NewVerification(
						NewVerificationMethod(
							"did:cosmos:elesto:subject#key-2",
							"did:cosmos:elesto:subject",
							NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
							EcdsaSecp256k1VerificationKey2019,
						),
						[]string{
							Authentication,
							CapabilityInvocation,
						},
						[]string{
							"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
						},
					),
				},
			},
			wantDid: DidDocument{
				Context: []string{
					"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
					contextDIDBase,
				},
				Id:         "did:cosmos:elesto:subject",
				Controller: nil,
				VerificationMethod: []*VerificationMethod{
					{
						"did:cosmos:elesto:subject#key-1",
						EcdsaSecp256k1VerificationKey2019.String(),
						"did:cosmos:elesto:subject",
						&VerificationMethod_PublicKeyMultibase{"F03DFD0A469806D66A23C7C948F55C129467D6D0974A222EF6E24A538FA6882F3D7"},
					},
					{
						"did:cosmos:elesto:subject#key-2",
						EcdsaSecp256k1VerificationKey2019.String(),
						"did:cosmos:elesto:subject",
						&VerificationMethod_PublicKeyMultibase{"F03DFD0A469806D66A23C7C948F55C129467D6D0974A222EF6E24A538FA6882F3D7"},
					},
				},
				Service:              nil,
				Authentication:       []string{"did:cosmos:elesto:subject#key-1", "did:cosmos:elesto:subject#key-2"},
				KeyAgreement:         []string{"did:cosmos:elesto:subject#key-1"},
				CapabilityInvocation: []string{"did:cosmos:elesto:subject#key-2"},
			},
		},
		{
			wantErr: true, // duplicated existing method id
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cosmos:elesto:subject", WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								EcdsaSecp256k1VerificationKey2019,
							),
							[]string{
								Authentication,
								KeyAgreement,
								KeyAgreement, // test duplicated relationship
							},
							[]string{
								"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
							},
						),
					))
					return d
				},
				[]*Verification{
					NewVerification(
						NewVerificationMethod(
							"did:cosmos:elesto:subject#key-1",
							"did:cosmos:elesto:subject",
							NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
							EcdsaSecp256k1VerificationKey2019,
						),
						[]string{
							string(CapabilityDelegation),
						},
						[]string{
							"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
						},
					),
				},
			},
			wantDid: DidDocument{},
		},
		{
			wantErr: true, // duplicated new method id
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cosmos:elesto:subject", WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								EcdsaSecp256k1VerificationKey2019,
							),
							[]string{
								Authentication,
								KeyAgreement,
								KeyAgreement, // test duplicated relationship
							},
							[]string{
								"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
							},
						),
					))
					return d
				},
				[]*Verification{
					NewVerification(
						NewVerificationMethod(
							"did:cosmos:elesto:subject#key-2",
							"did:cosmos:elesto:subject",
							NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
							EcdsaSecp256k1VerificationKey2019,
						),
						[]string{
							KeyAgreement,
						},
						[]string{
							"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
						},
					),
					NewVerification(
						NewVerificationMethod(
							"did:cosmos:elesto:subject#key-2",
							"did:cosmos:elesto:subject",
							NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
							EcdsaSecp256k1VerificationKey2019,
						),
						[]string{
							Authentication,
						},
						[]string{
							"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
						},
					),
				},
			},
			wantDid: DidDocument{},
		},
		{
			wantErr: true, // fail validation
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cosmos:elesto:subject", WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								EcdsaSecp256k1VerificationKey2019,
							),
							[]string{
								Authentication,
								KeyAgreement,
								KeyAgreement, // test duplicated relationship
							},
							[]string{
								"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
							},
						),
					))
					return d
				},
				[]*Verification{
					{
						[]string{
							string(Authentication),
							string(KeyAgreement),
							string(KeyAgreement), // test duplicated relationship
						},
						&VerificationMethod{
							"invalid method url",
							EcdsaSecp256k1VerificationKey2019.String(),
							"did:cosmos:elesto:subject",
							&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
						},
						[]string{
							"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
						},
					},
				},
			},
			wantDid: DidDocument{},
		},
		{
			wantErr: true, // verification relationship does not exists
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cosmos:elesto:subject")
					return d
				},
				[]*Verification{
					{
						[]string{
							Authentication,
							"UNSUPPORTED RELATIONSHIP",
							KeyAgreement,
						},
						&VerificationMethod{
							"did:cosmos:elesto:subject#key1",
							EcdsaSecp256k1VerificationKey2019.String(),
							"did:cosmos:elesto:subject",
							&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
						},
						[]string{
							"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
						},
					},
				},
			},
			wantDid: DidDocument{},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestDidDocument_AddVerifications#", i), func(t *testing.T) {
			gotDid := tt.params.malleate()

			err := gotDid.AddVerifications(tt.params.verifications...)

			if tt.wantErr {
				require.NotNil(t, err, "test: TestDidDocument_AddVerifications#%v", i)
				return
			}

			require.Nil(t, err, "test: TestDidDocument_AddVerifications#%v", i)
			assert.Equal(t, tt.wantDid, gotDid)
		})
	}
}

func TestDidDocument_RevokeVerification(t *testing.T) {
	type params struct {
		malleate func() DidDocument // build a did document
		methodID string             // input list of verifications
	}
	tests := []struct {
		params  params
		wantDid DidDocument // expected result
		wantErr bool
	}{
		{
			wantErr: false,
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cosmos:elesto:subject",
						WithVerifications(
							NewVerification(
								NewVerificationMethod(
									"did:cosmos:elesto:subject#key-1",
									"did:cosmos:elesto:subject",
									NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
									EcdsaSecp256k1VerificationKey2019,
								),
								[]string{
									Authentication,
									KeyAgreement,
								},
								nil,
							),
							NewVerification(
								NewVerificationMethod(
									"did:cosmos:elesto:subject#key-2",
									"did:cosmos:elesto:subject",
									NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
									EcdsaSecp256k1VerificationKey2019,
								),
								[]string{
									Authentication,
									CapabilityInvocation,
								},
								[]string{
									"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
								},
							),
						),
					)
					return d
				},
				"did:cosmos:elesto:subject#key-2",
			},
			wantDid: DidDocument{
				Context: []string{
					"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
					contextDIDBase,
				},
				Id:         "did:cosmos:elesto:subject",
				Controller: nil,
				VerificationMethod: []*VerificationMethod{
					{
						"did:cosmos:elesto:subject#key-1",
						EcdsaSecp256k1VerificationKey2019.String(),
						"did:cosmos:elesto:subject",
						&VerificationMethod_PublicKeyMultibase{"F03DFD0A469806D66A23C7C948F55C129467D6D0974A222EF6E24A538FA6882F3D7"},
					},
				},
				Service:        nil,
				Authentication: []string{"did:cosmos:elesto:subject#key-1"},
				KeyAgreement:   []string{"did:cosmos:elesto:subject#key-1"},
			},
		},
		{
			wantErr: false,
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cosmos:elesto:subject",
						WithVerifications(
							NewVerification(
								VerificationMethod{
									"did:cosmos:elesto:subject#key-1",
									EcdsaSecp256k1VerificationKey2019.String(),
									"did:cosmos:elesto:subject",
									&VerificationMethod_PublicKeyMultibase{"F03DFD0A469806D66A23C7C948F55C129467D6D0974A222EF6E24A538FA6882F3D7"},
								},
								[]string{
									Authentication,
									KeyAgreement,
								},
								nil,
							),
						),
					)
					return d
				},
				"did:cosmos:elesto:subject#key-1",
			},
			wantDid: DidDocument{
				Context: []string{
					contextDIDBase,
				},
				Id: "did:cosmos:elesto:subject",
			},
		},
		{
			wantErr: false,
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cosmos:elesto:subject",
						WithVerifications(
							NewVerification(
								VerificationMethod{
									"did:cosmos:elesto:subject#key-1",
									EcdsaSecp256k1VerificationKey2019.String(),
									"did:cosmos:elesto:subject",
									&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
								},
								[]string{
									Authentication,
									KeyAgreement,
								},
								nil,
							),
							NewVerification(
								VerificationMethod{
									"did:cosmos:elesto:subject#key-2",
									EcdsaSecp256k1VerificationKey2019.String(),
									"did:cosmos:elesto:subject",
									&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
								},
								[]string{
									Authentication,
									CapabilityInvocation,
								},
								nil,
							),
							NewVerification(
								VerificationMethod{
									"did:cosmos:elesto:subject#key-3",
									EcdsaSecp256k1VerificationKey2019.String(),
									"did:cosmos:elesto:subject",
									&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
								},
								[]string{
									Authentication,
									KeyAgreement,
									AssertionMethod,
								},
								nil,
							),
						),
					)
					return d
				},
				"did:cosmos:elesto:subject#key-2",
			},
			wantDid: DidDocument{
				Context: []string{
					contextDIDBase,
				},
				Id:         "did:cosmos:elesto:subject",
				Controller: nil,
				VerificationMethod: []*VerificationMethod{
					{
						"did:cosmos:elesto:subject#key-1",
						EcdsaSecp256k1VerificationKey2019.String(),
						"did:cosmos:elesto:subject",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					{
						"did:cosmos:elesto:subject#key-3",
						EcdsaSecp256k1VerificationKey2019.String(),
						"did:cosmos:elesto:subject",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
				},
				Service:         nil,
				Authentication:  []string{"did:cosmos:elesto:subject#key-1", "did:cosmos:elesto:subject#key-3"},
				KeyAgreement:    []string{"did:cosmos:elesto:subject#key-1", "did:cosmos:elesto:subject#key-3"},
				AssertionMethod: []string{"did:cosmos:elesto:subject#key-3"},
			},
		},
		{
			wantErr: true, // verification method not found
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cosmos:elesto:subject",
						WithVerifications(
							NewVerification(
								NewVerificationMethod(
									"did:cosmos:elesto:subject#key-1",
									"did:cosmos:elesto:subject",
									NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
									EcdsaSecp256k1VerificationKey2019,
								),
								[]string{
									Authentication,
									KeyAgreement,
								},
								nil,
							),
							NewVerification(
								NewVerificationMethod(
									"did:cosmos:elesto:subject#key-2",
									"did:cosmos:elesto:subject",
									NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
									EcdsaSecp256k1VerificationKey2019,
								),
								[]string{
									Authentication,
									CapabilityInvocation,
								},
								[]string{
									"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
								},
							),
						),
					)
					return d
				},
				"did:cosmos:elesto:subject#key-3",
			},
			wantDid: DidDocument{},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestDidDocument_RevokeVerification#", i), func(t *testing.T) {
			gotDid := tt.params.malleate()

			err := gotDid.RevokeVerification(tt.params.methodID)

			if tt.wantErr {
				require.NotNil(t, err, "test: TestDidDocument_RevokeVerification#%v", i)
				return
			}

			require.Nil(t, err, "test: TestDidDocument_RevokeVerification#%v", i)

			assert.Equal(t, tt.wantDid, gotDid)
		})
	}
}

func TestDidDocument_SetVerificationRelationships(t *testing.T) {
	type params struct {
		malleate      func() DidDocument
		methodID      string
		relationships []string
	}
	tests := []struct {
		params  params
		wantDid DidDocument // expected result
		wantErr bool
	}{
		{
			wantErr: true, // empty relationships
			params: params{
				malleate: func() DidDocument {
					dd, _ := NewDidDocument("did:cosmos:elesto:subject")
					return dd
				},
				methodID:      "did:cosmos:elesto:subject#key-1",
				relationships: []string{},
			},
			wantDid: DidDocument{
				Context: []string{contextDIDBase},
				Id:      "did:cosmos:elesto:subject",
			},
		},
		{
			wantErr: true, //invalid method id
			params: params{
				malleate: func() DidDocument {
					dd, _ := NewDidDocument("did:cosmos:elesto:subject")
					return dd
				},
				methodID:      "did:cosmos:elesto:subject#key-1 invalid ",
				relationships: []string{},
			},
			wantDid: DidDocument{},
		},
		{
			wantErr: false,
			params: params{
				malleate: func() DidDocument {
					dd, _ := NewDidDocument("did:cosmos:elesto:subject", WithVerifications(
						NewVerification(
							VerificationMethod{
								"did:cosmos:elesto:subject#key-1",
								EcdsaSecp256k1VerificationKey2019.String(),
								"did:cosmos:elesto:subject",
								&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
							},
							[]string{
								Authentication,
								KeyAgreement,
							},
							[]string{},
						),
					))
					return dd
				},
				methodID: "did:cosmos:elesto:subject#key-1",
				relationships: []string{
					string(AssertionMethod),
					string(AssertionMethod), // test duplicated relationship
					string(AssertionMethod), // test duplicated relationship
					string(AssertionMethod), // test duplicated relationship
				},
			},

			wantDid: DidDocument{
				Context: []string{contextDIDBase},
				Id:      "did:cosmos:elesto:subject",
				VerificationMethod: []*VerificationMethod{
					{
						"did:cosmos:elesto:subject#key-1",
						EcdsaSecp256k1VerificationKey2019.String(),
						"did:cosmos:elesto:subject",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
				},
				AssertionMethod: []string{"did:cosmos:elesto:subject#key-1"},
			},
		},
		{
			wantErr: false, // different delete scenarios
			params: params{
				malleate: func() DidDocument {
					dd, _ := NewDidDocument("did:cosmos:elesto:subject", WithVerifications(
						NewVerification(
							VerificationMethod{
								"did:cosmos:elesto:subject#key-1",
								EcdsaSecp256k1VerificationKey2019.String(),
								"did:cosmos:elesto:subject",
								&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
							},
							[]string{
								Authentication,
								KeyAgreement,
							},
							[]string{},
						),
						NewVerification(
							VerificationMethod{
								"did:cosmos:elesto:subject#key-2",
								EcdsaSecp256k1VerificationKey2019.String(),
								"did:cosmos:elesto:subject",
								&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
							},
							[]string{
								Authentication,
							},
							[]string{},
						),
					))
					return dd
				},
				methodID:      "did:cosmos:elesto:subject#key-1",
				relationships: []string{string(AssertionMethod)},
			},
			wantDid: DidDocument{
				Context: []string{contextDIDBase},
				Id:      "did:cosmos:elesto:subject",
				VerificationMethod: []*VerificationMethod{
					{
						"did:cosmos:elesto:subject#key-1",
						EcdsaSecp256k1VerificationKey2019.String(),
						"did:cosmos:elesto:subject",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					{
						"did:cosmos:elesto:subject#key-2",
						EcdsaSecp256k1VerificationKey2019.String(),
						"did:cosmos:elesto:subject",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
				},
				Authentication:  []string{"did:cosmos:elesto:subject#key-2"},
				AssertionMethod: []string{"did:cosmos:elesto:subject#key-1"},
			},
		},
		{
			wantErr: false, // different delete scenarios
			params: params{
				malleate: func() DidDocument {
					dd, _ := NewDidDocument("did:cosmos:elesto:subject", WithVerifications(
						NewVerification(
							VerificationMethod{
								"did:cosmos:elesto:subject#key-2",
								EcdsaSecp256k1VerificationKey2019.String(),
								"did:cosmos:elesto:subject",
								&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
							},
							[]string{
								Authentication,
							},
							[]string{},
						),
						NewVerification(
							VerificationMethod{
								"did:cosmos:elesto:subject#key-3",
								EcdsaSecp256k1VerificationKey2019.String(),
								"did:cosmos:elesto:subject",
								&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
							},
							[]string{
								Authentication,
							},
							[]string{},
						),
						NewVerification(
							VerificationMethod{
								"did:cosmos:elesto:subject#key-1",
								EcdsaSecp256k1VerificationKey2019.String(),
								"did:cosmos:elesto:subject",
								&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
							},
							[]string{
								Authentication,
								KeyAgreement,
							},
							[]string{},
						),
					))
					return dd
				},
				methodID:      "did:cosmos:elesto:subject#key-1",
				relationships: []string{string(AssertionMethod)},
			},
			wantDid: DidDocument{
				Context: []string{contextDIDBase},
				Id:      "did:cosmos:elesto:subject",
				VerificationMethod: []*VerificationMethod{
					{
						"did:cosmos:elesto:subject#key-2",
						EcdsaSecp256k1VerificationKey2019.String(),
						"did:cosmos:elesto:subject",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					{
						"did:cosmos:elesto:subject#key-3",
						EcdsaSecp256k1VerificationKey2019.String(),
						"did:cosmos:elesto:subject",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					{
						"did:cosmos:elesto:subject#key-1",
						EcdsaSecp256k1VerificationKey2019.String(),
						"did:cosmos:elesto:subject",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
				},

				Authentication:  []string{"did:cosmos:elesto:subject#key-2", "did:cosmos:elesto:subject#key-3"},
				AssertionMethod: []string{"did:cosmos:elesto:subject#key-1"},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestDidDocument_SetVerificationRelationships#", i), func(t *testing.T) {
			didDoc := tt.params.malleate()
			err := didDoc.SetVerificationRelationships(tt.params.methodID, tt.params.relationships...)

			if tt.wantErr {
				require.NotNil(t, err, "test: TestDidDocument_SetVerificationRelationships#%v", i)
				return
			}

			require.Nil(t, err, "test: TestDidDocument_SetVerificationRelationships#%v", i)
			assert.Equal(t, tt.wantDid, didDoc)

		})
	}
}

func TestDidDocument_HasRelationship(t *testing.T) {

	type params struct {
		didFn         func() DidDocument
		signer        *VerificationMethod_BlockchainAccountID
		relationships []string
	}
	tests := []struct {
		name                    string
		params                  params
		expectedHasRelationship bool
	}{
		{
			name:                    "PASS: has relationships (multibase)",
			expectedHasRelationship: true,
			params: params{
				didFn: func() DidDocument {
					dd, err := NewDidDocument("did:cosmos:elesto:subject", WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								EcdsaSecp256k1VerificationKey2019,
							),
							[]string{
								string(Authentication),
								string(KeyAgreement),
							},
							nil,
						),
					))
					assert.NoError(t, err)
					return dd
				},
				signer: NewBlockchainAccountID("cash", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"),
				relationships: []string{
					string(AssertionMethod),
					string(Authentication),
				},
			},
		},
		{
			name:                    "PASS: relationships missing (multibase, blockchainaccountid, hex)",
			expectedHasRelationship: false,
			params: params{
				didFn: func() DidDocument {
					dd, err := NewDidDocument("did:cosmos:elesto:subject", WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								EcdsaSecp256k1VerificationKey2019,
							),
							[]string{
								Authentication,
								KeyAgreement,
							},
							nil,
						),
						NewVerification(
							NewVerificationMethod(
								"did:cosmos:elesto:controller-1#key-2",
								"did:cosmos:elesto:controller-1",
								NewBlockchainAccountID("cash", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"),
								CosmosAccountAddress,
							),
							[]string{
								CapabilityDelegation,
							},
							nil,
						),
						NewVerification(
							NewVerificationMethod(
								"did:cosmos:elesto:subject#key-3",
								"did:cosmos:elesto:subject",
								NewPublicKeyHex([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								EcdsaSecp256k1VerificationKey2019,
							),
							[]string{
								Authentication,
								KeyAgreement,
							},
							nil,
						),
					))
					assert.NoError(t, err)
					return dd
				},
				signer: NewBlockchainAccountID("cash", "subject"),
				relationships: []string{
					string(CapabilityDelegation),
				},
			},
		},
		{
			name:                    "PASS: relationships missing (blockchainaccountid)",
			expectedHasRelationship: false,
			params: params{
				didFn: func() DidDocument {
					dd, err := NewDidDocument("did:cosmos:elesto:subject")
					assert.NoError(t, err)
					return dd
				},
				signer: NewBlockchainAccountID("cash", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"),
				relationships: []string{
					string(CapabilityDelegation),
				},
			},
		},
		{
			name:                    "PASS: relationships missing (Multibase)",
			expectedHasRelationship: false,
			params: params{
				didFn: func() DidDocument {
					dd, err := NewDidDocument("did:cosmos:elesto:subject", WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								EcdsaSecp256k1VerificationKey2019,
							),
							[]string{
								Authentication,
								KeyAgreement,
							},
							nil,
						),
					))
					assert.NoError(t, err)
					return dd
				},
				signer:        NewBlockchainAccountID("cash", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"),
				relationships: nil,
			},
		},
		{
			name:                    "PASS: has relationship (BlockchainAccountID)",
			expectedHasRelationship: true,
			params: params{
				didFn: func() DidDocument {
					dd, err := NewDidDocument("did:cosmos:elesto:subject", WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								NewPublicKeyMultibase([]byte("00dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7")),
								EcdsaSecp256k1VerificationKey2019,
							),
							[]string{
								Authentication,
							},
							nil,
						),
						NewVerification(
							NewVerificationMethod(
								"did:cosmos:elesto:subject#key-2",
								"did:cosmos:elesto:subject",
								NewBlockchainAccountID("cash", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"),
								CosmosAccountAddress,
							),
							[]string{
								KeyAgreement,
							},
							nil,
						),
					))
					assert.NoError(t, err)
					return dd
				},
				signer: NewBlockchainAccountID("cash", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"),
				relationships: []string{
					string(KeyAgreement),
				},
			},
		},
		{
			name:                    "PASS:  missing relationship (PublicKeyHex)",
			expectedHasRelationship: false,
			params: params{
				didFn: func() DidDocument {
					dd, err := NewDidDocument("did:cosmos:elesto:subject", WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cosmos:elesto:subject#key-1",
								"did:cosmos:elesto:subject",
								NewPublicKeyHex([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
								EcdsaSecp256k1VerificationKey2019,
							),
							[]string{
								Authentication,
								KeyAgreement,
							},
							nil,
						),
					))
					assert.NoError(t, err)
					return dd
				},
				signer:        NewBlockchainAccountID("cash", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"),
				relationships: nil,
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestDidDocument_SetVerificationRelationships#", i), func(t *testing.T) {
			didDoc := tt.params.didFn()
			gotHasRelationship := didDoc.HasRelationship(tt.params.signer, tt.params.relationships...)
			assert.Equal(t, tt.expectedHasRelationship, gotHasRelationship)
		})
	}
}

func TestDidDocument_AddServices(t *testing.T) {
	type params struct {
		malleate func() DidDocument // build a did document
		services []*Service         // input list of verifications
	}
	tests := []struct {
		params  params
		wantDid DidDocument // expected result
		wantErr bool
	}{
		{
			wantErr: false,
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cosmos:elesto:subject")
					return d
				},
				[]*Service{
					NewService(
						"agent:abc",
						"DIDCommMessaging",
						"https://agent.abc/1234",
					),
					NewService(
						"agent:xyz",
						"DIDCommMessaging",
						"https://agent.xyz/1234",
					),
				},
			},
			wantDid: DidDocument{
				Context: []string{contextDIDBase},
				Id:      "did:cosmos:elesto:subject",
				Service: []*Service{
					NewService(
						"agent:abc",
						"DIDCommMessaging",
						"https://agent.abc/1234",
					),
					NewService(
						"agent:xyz",
						"DIDCommMessaging",
						"https://agent.xyz/1234",
					),
				},
			},
		},
		{
			wantErr: true, // duplicated existing service id
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument(
						"did:cosmos:elesto:subject",
						WithServices(
							NewService(
								"agent:xyz",
								"DIDCommMessaging",
								"https://agent.xyz/1234",
							),
						),
					)
					return d
				},
				[]*Service{
					{
						"agent:abc",
						"DIDCommMessaging",
						"https://agent.abc/1234",
					}, {
						"agent:xyz",
						"DIDCommMessaging",
						"https://agent.xyz/1234",
					},
				},
			},
			wantDid: DidDocument{},
		},
		{
			wantErr: true, // duplicated new service id
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cosmos:elesto:subject")
					return d
				},
				[]*Service{
					{
						"agent:xyz",
						"DIDCommMessaging",
						"https://agent.xyz/1234",
					}, {
						"agent:xyz",
						"DIDCommMessaging",
						"https://agent.xyz/1234",
					},
				},
			},
			wantDid: DidDocument{},
		},
		{
			wantErr: true, // fail validation
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cosmos:elesto:subject")
					return d
				},
				[]*Service{
					{
						"agent:abc",
						"DIDCommMessaging",
						"https://agent.abc/1234",
					}, {
						"",
						"DIDCommMessaging",
						"https://agent.xyz/1234",
					},
				},
			},
			wantDid: DidDocument{},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestDidDocument_AddServices#", i), func(t *testing.T) {
			gotDid := tt.params.malleate()

			err := gotDid.AddServices(tt.params.services...)

			if tt.wantErr {
				require.NotNil(t, err, "test: TestDidDocument_AddServices#%v", i)
				return
			}

			require.Nil(t, err, "test: TestDidDocument_AddServices#%v", i)
			assert.Equal(t, tt.wantDid, gotDid)
		})
	}
}

func TestDidDocument_DeleteService(t *testing.T) {
	type params struct {
		didFn    func() DidDocument // build a did document
		methodID string             // input list of verifications
	}
	tests := []struct {
		params  params
		wantDid DidDocument // expected result
		wantErr bool
	}{
		{
			wantErr: false,
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cosmos:elesto:subject",
						WithServices(
							&Service{
								"agent:abc",
								"DIDCommMessaging",
								"https://agent.abc/1234",
							},
						),
					)
					return d
				},
				"agent:abc",
			},
			wantDid: DidDocument{
				Context: []string{contextDIDBase},
				Id:      "did:cosmos:elesto:subject",
			},
		},
		{
			wantErr: false,
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cosmos:elesto:subject",
						WithServices(
							&Service{
								"agent:zyz",
								"DIDCommMessaging",
								"https://agent.abc/1234",
							},
							&Service{
								"agent:abc",
								"DIDCommMessaging",
								"https://agent.abc/1234",
							},
						),
					)
					return d
				},
				"agent:abc",
			},
			wantDid: DidDocument{
				Context: []string{contextDIDBase},
				Id:      "did:cosmos:elesto:subject",
				Service: []*Service{
					{
						"agent:zyz",
						"DIDCommMessaging",
						"https://agent.abc/1234",
					},
				},
			},
		},
		{
			wantErr: false,
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cosmos:elesto:subject",
						WithServices(
							&Service{
								"agent:zyz",
								"DIDCommMessaging",
								"https://agent.abc/1234",
							},
							&Service{
								"agent:abc",
								"DIDCommMessaging",
								"https://agent.abc/1234",
							},
							&Service{
								"agent:007",
								"DIDCommMessaging",
								"https://agent.abc/007",
							},
						),
					)
					return d
				},
				"agent:abc",
			},
			wantDid: DidDocument{
				Context: []string{contextDIDBase},
				Id:      "did:cosmos:elesto:subject",
				Service: []*Service{
					{
						"agent:zyz",
						"DIDCommMessaging",
						"https://agent.abc/1234",
					}, {
						"agent:007",
						"DIDCommMessaging",
						"https://agent.abc/007",
					},
				},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestDidDocument_DeleteService#", i), func(t *testing.T) {
			gotDid := tt.params.didFn()

			gotDid.DeleteService(tt.params.methodID)

			assert.Equal(t, tt.wantDid, gotDid)
		})
	}
}

func TestBlockchainAccountID_GetAddress(t *testing.T) {
	tests := []struct {
		name string
		baID *VerificationMethod_BlockchainAccountID
		want string
	}{
		{
			"PASS: can get address",
			NewBlockchainAccountIDFromString("cosmos:foochain:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"),
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
		},
		{
			"PASS: address is empty",
			NewBlockchainAccountIDFromString("cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"),
			"",
		},
		{
			"PASS: can get address (but address is wrong)",
			NewBlockchainAccountIDFromString("cosmos:foochain:whatever"),
			"whatever",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.baID.GetAddress(), "GetAddress()")
		})
	}
}

func TestDidDocument_HasPublicKey(t *testing.T) {

	tests := []struct {
		name   string
		didFn  func() DidDocument
		pubkey func() types.PubKey
		want   bool
	}{
		{
			"PASS: has public key (multibase)",
			func() DidDocument {
				dd, err := NewDidDocument("did:cosmos:elesto:subject", WithVerifications(
					NewVerification(
						NewVerificationMethod(
							"did:cosmos:elesto:subject#key-1",
							"did:cosmos:elesto:subject",
							NewPublicKeyMultibase([]byte{2, 201, 95, 248, 187, 133, 206, 97, 166, 70, 229, 226, 88, 124, 29, 43, 70, 3, 244, 225, 19, 128, 44, 132, 110, 15, 15, 35, 40, 189, 237, 71, 245}),
							EcdsaSecp256k1VerificationKey2019,
						),
						[]string{
							Authentication,
							KeyAgreement,
						},
						nil,
					),
				))
				assert.NoError(t, err)
				return dd
			},
			func() types.PubKey {
				var pk types.PubKey
				c := simapp.MakeTestEncodingConfig().Marshaler
				err := c.UnmarshalInterfaceJSON([]byte(`{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Aslf+LuFzmGmRuXiWHwdK0YD9OETgCyEbg8PIyi97Uf1"}`), &pk)
				assert.NoError(t, err)

				return pk

			},
			true,
		},
		{
			"PASS: doesn't have public key (multibase)",
			func() DidDocument {
				dd, err := NewDidDocument("did:cosmos:elesto:subject", WithVerifications(
					NewVerification(
						NewVerificationMethod(
							"did:cosmos:elesto:subject#key-1",
							"did:cosmos:elesto:subject",
							NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
							EcdsaSecp256k1VerificationKey2019,
						),
						[]string{
							Authentication,
							KeyAgreement,
						},
						nil,
					),
				))
				assert.NoError(t, err)
				return dd
			},
			func() types.PubKey {
				var pk types.PubKey
				c := simapp.MakeTestEncodingConfig().Marshaler
				err := c.UnmarshalInterfaceJSON([]byte(`{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Aslf+LuFzmGmRuXiWHwdK0YD9OETgCyEbg8PIyi97Uf1"}`), &pk)
				assert.NoError(t, err)
				return pk

			},
			false,
		},
		{
			"PASS: has public key (blockchainAccount)",
			func() DidDocument {
				dd, err := NewDidDocument("did:cosmos:elesto:subject", WithVerifications(
					NewVerification(
						NewVerificationMethod(
							"did:cosmos:elesto:subject#key-1",
							"did:cosmos:elesto:subject",
							NewBlockchainAccountID("foochain", "cosmos17t8t3t6a6vpgk69perfyq930593sa8dn4kzsdf"),
							CosmosAccountAddress,
						),
						[]string{
							Authentication,
							KeyAgreement,
						},
						nil,
					),
				))
				assert.NoError(t, err)
				return dd
			},
			func() types.PubKey {
				var pk types.PubKey
				c := simapp.MakeTestEncodingConfig().Marshaler
				err := c.UnmarshalInterfaceJSON([]byte(`{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Aslf+LuFzmGmRuXiWHwdK0YD9OETgCyEbg8PIyi97Uf1"}`), &pk)
				assert.NoError(t, err)
				return pk

			},
			true,
		},
		{
			"PASS: doesn't have public key (blockchainAccountId)",
			func() DidDocument {
				dd, err := NewDidDocument("did:cosmos:elesto:subject", WithVerifications(
					NewVerification(
						NewVerificationMethod(
							"did:cosmos:elesto:subject#key-1",
							"did:cosmos:elesto:subject",
							NewBlockchainAccountID("foochain", "cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2"),
							CosmosAccountAddress,
						),
						[]string{
							Authentication,
							KeyAgreement,
						},
						nil,
					),
				))
				assert.NoError(t, err)
				return dd
			},
			func() types.PubKey {
				var pk types.PubKey
				c := simapp.MakeTestEncodingConfig().Marshaler
				err := c.UnmarshalInterfaceJSON([]byte(`{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Aslf+LuFzmGmRuXiWHwdK0YD9OETgCyEbg8PIyi97Uf1"}`), &pk)
				assert.NoError(t, err)
				return pk

			},
			false,
		},
		{
			"PASS: has public key (publicKeyHex)",
			func() DidDocument {
				dd, err := NewDidDocument("did:cosmos:elesto:subject", WithVerifications(
					NewVerification(
						NewVerificationMethod(
							"did:cosmos:elesto:subject#key-1",
							"did:cosmos:elesto:subject",
							NewPublicKeyHex([]byte{2, 201, 95, 248, 187, 133, 206, 97, 166, 70, 229, 226, 88, 124, 29, 43, 70, 3, 244, 225, 19, 128, 44, 132, 110, 15, 15, 35, 40, 189, 237, 71, 245}),
							EcdsaSecp256k1VerificationKey2019,
						),
						[]string{
							Authentication,
							KeyAgreement,
						},
						nil,
					),
				))
				assert.NoError(t, err)
				return dd
			},
			func() types.PubKey {
				var pk types.PubKey
				c := simapp.MakeTestEncodingConfig().Marshaler
				err := c.UnmarshalInterfaceJSON([]byte(`{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Aslf+LuFzmGmRuXiWHwdK0YD9OETgCyEbg8PIyi97Uf1"}`), &pk)
				assert.NoError(t, err)
				return pk

			},
			true,
		},
		{
			"PASS: doesn't have public key (pubKeyHex)",
			func() DidDocument {
				dd, err := NewDidDocument("did:cosmos:elesto:subject", WithVerifications(
					NewVerification(
						NewVerificationMethod(
							"did:cosmos:elesto:subject#key-1",
							"did:cosmos:elesto:subject",
							NewPublicKeyHex([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
							EcdsaSecp256k1VerificationKey2019,
						),
						[]string{
							Authentication,
							KeyAgreement,
						},
						nil,
					),
				))
				assert.NoError(t, err)
				return dd
			},
			func() types.PubKey {
				var pk types.PubKey
				c := simapp.MakeTestEncodingConfig().Marshaler
				err := c.UnmarshalInterfaceJSON([]byte(`{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Aslf+LuFzmGmRuXiWHwdK0YD9OETgCyEbg8PIyi97Uf1"}`), &pk)
				assert.NoError(t, err)
				return pk

			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			didDoc := tt.didFn()
			pubKey := tt.pubkey()
			assert.Equalf(t, tt.want, didDoc.HasPublicKey(pubKey), "HasPublicKey(%v)", pubKey)
		})
	}
}

func TestDidDocument_GetVerificationMethodBlockchainAddress(t *testing.T) {
	tests := []struct {
		name        string
		didFn       func() DidDocument
		methodID    string
		wantAddress string
		wantErr     error
	}{
		{
			"PASS: get address (PublicKeyMultibase)",
			func() DidDocument {
				dd, err := NewDidDocument("did:cosmos:elesto:subject", WithVerifications(
					NewVerification(
						NewVerificationMethod(
							"did:cosmos:elesto:subject#key-1",
							"did:cosmos:elesto:subject",
							NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
							EcdsaSecp256k1VerificationKey2019,
						),
						[]string{
							Authentication,
							KeyAgreement,
						},
						nil,
					),
				))
				assert.NoError(t, err)
				return dd
			},
			"did:cosmos:elesto:subject#key-1",
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			nil,
		},
		{
			"PASS: get address (PublicKeyHex)",
			func() DidDocument {
				dd, err := NewDidDocument("did:cosmos:elesto:subject", WithVerifications(
					NewVerification(
						NewVerificationMethod(
							"did:cosmos:elesto:subject#key-1",
							"did:cosmos:elesto:subject",
							NewPublicKeyHex([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
							EcdsaSecp256k1VerificationKey2019,
						),
						[]string{
							Authentication,
							KeyAgreement,
						},
						nil,
					),
				))
				assert.NoError(t, err)
				return dd
			},
			"did:cosmos:elesto:subject#key-1",
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			nil,
		},
		{
			"PASS: get address (BlockchainAccountID)",
			func() DidDocument {
				dd, err := NewDidDocument("did:cosmos:elesto:subject", WithVerifications(
					NewVerification(
						NewVerificationMethod(
							"did:cosmos:elesto:subject#key-1",
							"did:cosmos:elesto:subject",
							NewBlockchainAccountID("foochain", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"),
							CosmosAccountAddress,
						),
						[]string{
							Authentication,
							KeyAgreement,
						},
						nil,
					),
				))
				assert.NoError(t, err)
				return dd
			},
			"did:cosmos:elesto:subject#key-1",
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			nil,
		},
		{
			"PASS: get address (BlockchainAccountID)",
			func() DidDocument {
				dd, err := NewDidDocument("did:cosmos:elesto:subject", WithVerifications(
					NewVerification(
						NewVerificationMethod(
							"did:cosmos:elesto:subject#key-1",
							"did:cosmos:elesto:subject",
							NewBlockchainAccountID("foochain", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"),
							CosmosAccountAddress,
						),
						[]string{
							Authentication,
							KeyAgreement,
						},
						nil,
					),
				))
				assert.NoError(t, err)
				return dd
			},
			"did:cosmos:elesto:subject#key-2",
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			ErrVerificationMethodNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			didDoc := tt.didFn()
			gotAddress, err := didDoc.GetVerificationMethodBlockchainAddress(tt.methodID)
			if tt.wantErr == nil {
				assert.NoError(t, err)
				assert.Equalf(t, tt.wantAddress, gotAddress, "GetVerificationMethodBlockchainAddress(%v)", tt.methodID)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
		})
	}
}

func TestDidDocument_HasController(t *testing.T) {

	tests := []struct {
		name          string
		didFn         func() DidDocument
		controllerDID string
		want          bool
	}{
		{
			"PASS: controller found",
			func() DidDocument {
				dd, err := NewDidDocument(
					"did:cosmos:elesto:subject",
					WithControllers(
						"did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
						"did:cosmos:key:cosmos17t8t3t6a6vpgk69perfyq930593sa8dn4kzsdf",
					),
				)
				assert.NoError(t, err)
				return dd
			},
			"did:cosmos:key:cosmos17t8t3t6a6vpgk69perfyq930593sa8dn4kzsdf",
			true,
		},
		{
			"PASS: controller not found",
			func() DidDocument {
				dd, err := NewDidDocument(
					"did:cosmos:elesto:subject",
					WithControllers(
						"did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
						"did:cosmos:key:cosmos17t8t3t6a6vpgk69perfyq930593sa8dn4kzsdf",
					),
				)
				assert.NoError(t, err)
				return dd
			},
			"did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			didDoc := tt.didFn()
			assert.Equalf(t, tt.want, didDoc.HasController(DID(tt.controllerDID)), "HasController(%v)", tt.controllerDID)
		})
	}
}

func TestResolveAccountDID(t *testing.T) {
	type args struct {
		did     string
		chainID string
	}
	tests := []struct {
		name       string
		args       args
		wantDidDoc func() DidDocument
		wantErr    error
	}{
		{
			"FAIL: not a key did",
			args{
				"did:cosmos:elesto:1234",
				"elesto",
			},
			func() DidDocument {
				return DidDocument{}
			},
			ErrInvalidDidMethodFormat,
		},
		{
			"PASS: key did",
			args{
				"did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				"elesto",
			},
			func() DidDocument {
				dd, err := NewDidDocument(
					"did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
					WithVerifications(
						NewVerification(
							NewVerificationMethod(
								DID("did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8").NewVerificationMethodID("cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"),
								DID("did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"), // the controller is the same as the did subject
								NewBlockchainAccountID("elesto", "cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"),
								CosmosAccountAddress,
							),
							[]string{
								Authentication,
								KeyAgreement,
								AssertionMethod,
								CapabilityInvocation,
								CapabilityDelegation,
							},
							nil,
						),
					),
				)
				assert.NoError(t, err)
				return dd
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDidDoc, err := ResolveAccountDID(tt.args.did, tt.args.chainID)
			if tt.wantErr == nil {
				assert.NoError(t, err)
				assert.Equalf(t, tt.wantDidDoc(), gotDidDoc, "ResolveAccountDID(%v, %v)", tt.args.did, tt.args.chainID)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}

		})
	}
}

func TestNewAccountVerification(t *testing.T) {
	type args struct {
		did                 DID
		chainID             string
		accountAddress      string
		verificationMethods []string
	}
	tests := []struct {
		name string
		args args
		want *Verification
	}{
		{
			"PASS: net did",
			args{
				DID("did:cosmos:elesto:1234"),
				"elesto",
				"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
				[]string{Authentication},
			},
			&Verification{
				Method: &VerificationMethod{
					Id:                   "did:cosmos:elesto:1234#cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
					Controller:           "did:cosmos:elesto:1234",
					VerificationMaterial: &VerificationMethod_BlockchainAccountID{BlockchainAccountID: "cosmos:elesto:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"},
					Type:                 CosmosAccountAddress.String(),
				},
				Relationships: []string{Authentication},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewAccountVerification(tt.args.did, tt.args.chainID, tt.args.accountAddress, tt.args.verificationMethods...), "NewAccountVerification(%v, %v, %v, %v)", tt.args.did, tt.args.chainID, tt.args.accountAddress, tt.args.verificationMethods)
		})
	}
}
