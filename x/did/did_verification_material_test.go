package did

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPublicKeyMultibaseFromHex(t *testing.T) {
	type args struct {
		pubKeyHex string
		vmType    VerificationMaterialType
	}
	tests := []struct {
		name    string
		args    args
		wantPkm PublicKeyMultibase
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"PASS: key match",
			args{
				pubKeyHex: "03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7",
				vmType:    DIDVMethodTypeEcdsaSecp256k1VerificationKey2019,
			},
			PublicKeyMultibase{
				data:   []byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215},
				vmType: DIDVMethodTypeEcdsaSecp256k1VerificationKey2019,
			},
			assert.NoError,
		},
		{
			"FAIL: invalid hex key",
			args{
				pubKeyHex: "not hex string",
				vmType:    DIDVMethodTypeEcdsaSecp256k1VerificationKey2019,
			},
			PublicKeyMultibase{
				data:   nil,
				vmType: "",
			},
			assert.Error, // TODO: check the error message
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPkm, err := NewPublicKeyMultibaseFromHex(tt.args.pubKeyHex, tt.args.vmType)
			if !tt.wantErr(t, err, fmt.Sprintf("NewPublicKeyMultibaseFromHex(%v, %v)", tt.args.pubKeyHex, tt.args.vmType)) {
				return
			}
			assert.Equalf(t, tt.wantPkm, gotPkm, "NewPublicKeyMultibaseFromHex(%v, %v)", tt.args.pubKeyHex, tt.args.vmType)

		})
	}
}

func TestNewPublicKeyHexFromString(t *testing.T) {
	type args struct {
		pubKeyHex string
		vmType    VerificationMaterialType
	}
	tests := []struct {
		name    string
		args    args
		wantPkh PublicKeyHex
		wantErr error
	}{
		{
			"PASS: valid pub key",
			args{
				"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7",
				DIDVMethodTypeX25519KeyAgreementKey2019,
			},
			PublicKeyHex{
				data:   []byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215},
				vmType: DIDVMethodTypeX25519KeyAgreementKey2019,
			},
			nil,
		},
		{
			"FAIL: not hex",
			args{
				"1234&",
				DIDVMethodTypeX25519KeyAgreementKey2019,
			},
			PublicKeyHex{
				data:   []byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215},
				vmType: DIDVMethodTypeX25519KeyAgreementKey2019,
			},
			fmt.Errorf("encoding/hex: invalid byte: U+0026 '&'"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPkh, err := NewPublicKeyHexFromString(tt.args.pubKeyHex, tt.args.vmType)
			if tt.wantErr == nil {
				assert.NoError(t, err)
				assert.Equalf(t, tt.wantPkh, gotPkh, "NewPublicKeyHexFromString(%v, %v)", tt.args.pubKeyHex, tt.args.vmType)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}

		})
	}
}
