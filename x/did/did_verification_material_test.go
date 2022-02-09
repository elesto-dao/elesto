package did

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPublicKeyMultibaseFromHex(t *testing.T) {
	tests := []struct {
		name    string
		pubKeyHex string
		wantPkm *VerificationMethod_PublicKeyMultibase
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"PASS: key match",
			"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7",
			NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
			assert.NoError,
		},
		{
			"FAIL: invalid hex key",
			"not hex string",
			nil,
			assert.Error, // TODO: check the error message
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPkm, err := NewPublicKeyMultibaseFromHex(tt.pubKeyHex)
			if !tt.wantErr(t, err, fmt.Sprintf("NewPublicKeyMultibaseFromHex(%v)", tt.pubKeyHex)) {
				return
			}
			assert.Equalf(t, tt.wantPkm, gotPkm, "NewPublicKeyMultibaseFromHex(%v)", tt.pubKeyHex)

		})
	}
}

func TestNewPublicKeyHexFromString(t *testing.T) {
	tests := []struct {
		name    string
		pubKeyHex string
		wantPkh *VerificationMethod_PublicKeyHex
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"PASS: key match",
			"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7",
			NewPublicKeyHex([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}),
			assert.NoError,
		},
		{
			"FAIL: invalid hex key",
			"not hex string",
			nil,
			assert.Error, // TODO: check the error message
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPkh, err := NewPublicKeyHexFromString(tt.pubKeyHex)
			if !tt.wantErr(t, err, fmt.Sprintf("NewPublicKeyHexFromString(%v)", tt.pubKeyHex)) {
				return
			}
			assert.Equalf(t, tt.wantPkh, gotPkh, "NewPublicKeyHexFromString(%v)", tt.pubKeyHex)
		})
	}
}

func TestNewPublicKeyJwk(t *testing.T) {
	tests := []struct {
		name    string
		pubKey []byte
		wantVm  *VerificationMethod_PublicKeyJwk
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"PASS: can parse pub key",
			[]byte(`{"crv":"secp256k1","kid":"JUvpllMEYUZ2joO59UNui_XYDqxVqiFLLAJ8klWuPBw","kty":"EC","x":"dWCvM4fTdeM0KmloF57zxtBPXTOythHPMm1HCLrdd3A","y":"36uMVGM7hnw-N6GnjFcihWE3SkrhMLzzLCdPMXPEXlA"}`),
			&VerificationMethod_PublicKeyJwk{
				PublicKeyJwk: &PublicKeyJwk{
					Kid: "JUvpllMEYUZ2joO59UNui_XYDqxVqiFLLAJ8klWuPBw",
					Crv: "secp256k1",
					X:   "dWCvM4fTdeM0KmloF57zxtBPXTOythHPMm1HCLrdd3A",
					Y:   "36uMVGM7hnw-N6GnjFcihWE3SkrhMLzzLCdPMXPEXlA",
					Kty: "EC",
				},
			},
			assert.NoError,
		},
		{
			"FAIL: empty kid",
			[]byte(`{"crv":"secp256k1","kid":"","kty":"EC","x":"dWCvM4fTdeM0KmloF57zxtBPXTOythHPMm1HCLrdd3A","y":"36uMVGM7hnw-N6GnjFcihWE3SkrhMLzzLCdPMXPEXlA"}`),
			nil,
			assert.Error,
		},
		{
			"FAIL: empty x and y",
			[]byte(`{"crv":"secp256k1","kid":"JUvpllMEYUZ2joO59UNui_XYDqxVqiFLLAJ8klWuPBw","kty":"EC","x":"","y":"   "}`),
			nil,
			assert.Error,
		},
		{
			"FAIL: empty x and y",
			[]byte(`not a json`),
			nil,
			assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVm, err := NewPublicKeyJwk(tt.pubKey)
			if !tt.wantErr(t, err, fmt.Sprintf("NewPublicKeyJwk(%v)", tt.pubKey)) {
				return
			}
			assert.Equalf(t, tt.wantVm, gotVm, "NewPublicKeyJwk(%v)", tt.pubKey)
		})
	}
}