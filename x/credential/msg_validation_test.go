package credential

import (
	_ "embed"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	//go:embed keeper/testdata/dummy.schema.json
	dummySchemaOk string
	//go:embed keeper/testdata/dummy.vocab.json
	dummyVocabOk string
)

func TestMsgPublishCredentialDefinitionRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		CredentialDefinition *CredentialDefinition
		Signer               string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			"PASS: valid request",
			fields{
				&CredentialDefinition{
					Id:           "did:cosmos:elesto:cd-00001",
					PublisherId:  "did:cosmos:key:elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr",
					Schema:       []byte(dummySchemaOk),
					Vocab:        []byte(dummyVocabOk),
					Name:         "CredentialDef00001",
					Description:  "",
					SupersededBy: "",
					IsActive:     true,
				},
				"elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr",
			},
			nil,
		},
		{
			"FAIL: definition not set",
			fields{
				nil,
				"elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr",
			},
			errors.New("credential definition must be set"),
		},
		{
			"FAIL: definition ID is empty",
			fields{
				&CredentialDefinition{
					Id:           "",
					PublisherId:  "did:cosmos:key:elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr",
					Schema:       []byte(dummySchemaOk),
					Vocab:        []byte(dummyVocabOk),
					Name:         "CredentialDef00001",
					Description:  "",
					SupersededBy: "",
					IsActive:     true,
				},
				"elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr",
			},
			errors.New("credential definition id must be set"),
		},
		{
			"FAIL: name is empty",
			fields{
				&CredentialDefinition{
					Id:           "did:cosmos:elesto:cd-00001",
					PublisherId:  "did:cosmos:key:elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr",
					Schema:       []byte(dummySchemaOk),
					Vocab:        []byte(dummyVocabOk),
					Name:         "",
					Description:  "",
					SupersededBy: "",
					IsActive:     true,
				},
				"elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr",
			},
			errors.New("credential definition name must be set"),
		},
		{
			"FAIL: publisher id not a DID",
			fields{
				&CredentialDefinition{
					Id:           "did:cosmos:elesto:cd-00001",
					PublisherId:  "not a did",
					Schema:       []byte(dummySchemaOk),
					Vocab:        []byte(dummyVocabOk),
					Name:         "CredentialDef00001",
					Description:  "",
					SupersededBy: "",
					IsActive:     true,
				},
				"elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr",
			},
			errors.New("publisher id must be a valid DID"),
		},
		{
			"FAIL: schema cannot be empty",
			fields{
				&CredentialDefinition{
					Id:           "did:cosmos:elesto:cd-00001",
					PublisherId:  "did:cosmos:key:elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr",
					Schema:       []byte{},
					Vocab:        []byte(dummyVocabOk),
					Name:         "CredentialDef00001",
					Description:  "",
					SupersededBy: "",
					IsActive:     true,
				},
				"elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr",
			},
			errors.New("schema cannot be empty"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MsgPublishCredentialDefinitionRequest{
				CredentialDefinition: tt.fields.CredentialDefinition,
				Signer:               tt.fields.Signer,
			}
			err := m.ValidateBasic()

			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
		})
	}
}

func TestMsgUpdateCredentialDefinitionRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		Active       bool
		SupersededBy string
		Signer       string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			"PASS: msg is valid",
			fields{
				Active:       true,
				SupersededBy: "did:cosmos:elesto:another-cd",
				Signer:       "elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr",
			},
			nil,
		},
		{
			"PASS: msg is valid",
			fields{
				Active:       true,
				SupersededBy: "",
				Signer:       "elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr",
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MsgUpdateCredentialDefinitionRequest{
				Active:       tt.fields.Active,
				SupersededBy: tt.fields.SupersededBy,
				Signer:       tt.fields.Signer,
			}
			err := m.ValidateBasic()

			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
		})
	}
}

func TestMsgIssuePublicVerifiableCredentialRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		Credential             *PublicVerifiableCredential
		CredentialDefinitionID string
		Signer                 string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			"PASS: msg is valid",
			fields{
				Credential: &PublicVerifiableCredential{
					Id: "https://test.xyz/credential/1",
					Context: []string{
						"https://www.w3.org/2018/credentials/v1",
						"https://resolver.cc/context/did:cosmos:elesto:dummy",
					},
					Type: []string{
						"VerifiableCredential",
						"DummyCredential",
					},
					Issuer:            "did:cosmos:key:elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr",
					IssuanceDate:      func() *time.Time { v := time.Date(2022, 6, 2, 14, 13, 0, 0, time.UTC); return &v }(),
					CredentialSubject: []byte{123, 34, 97, 103, 101, 34, 58, 34, 52, 50, 34, 44, 34, 105, 100, 34, 58, 34, 100, 105, 100, 58, 99, 111, 115, 109, 111, 115, 58, 107, 101, 121, 58, 101, 108, 101, 115, 116, 111, 49, 55, 116, 56, 116, 51, 116, 54, 97, 54, 118, 112, 103, 107, 54, 57, 112, 101, 114, 102, 121, 113, 57, 51, 48, 53, 57, 51, 115, 97, 56, 100, 110, 102, 108, 57, 56, 109, 114, 34, 44, 34, 110, 97, 109, 101, 34, 58, 34, 65, 114, 116, 104, 117, 114, 32, 68, 101, 110, 116, 34, 125},
				},
				CredentialDefinitionID: "did:cosmos:elesto:cd-0001",
				Signer:                 "elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr",
			},
			nil,
		},
		{
			"FAIL: empty CredentialDefinitionID",
			fields{
				Credential: &PublicVerifiableCredential{
					Id: "https://test.xyz/credential/1",
					Context: []string{
						"https://www.w3.org/2018/credentials/v1",
						"https://resolver.cc/context/did:cosmos:elesto:dummy",
					},
					Type: []string{
						"VerifiableCredential",
						"DummyCredential",
					},
					Issuer:            "did:cosmos:key:elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr",
					IssuanceDate:      func() *time.Time { v := time.Date(2022, 6, 2, 14, 13, 0, 0, time.UTC); return &v }(),
					CredentialSubject: []byte{123, 34, 97, 103, 101, 34, 58, 34, 52, 50, 34, 44, 34, 105, 100, 34, 58, 34, 100, 105, 100, 58, 99, 111, 115, 109, 111, 115, 58, 107, 101, 121, 58, 101, 108, 101, 115, 116, 111, 49, 55, 116, 56, 116, 51, 116, 54, 97, 54, 118, 112, 103, 107, 54, 57, 112, 101, 114, 102, 121, 113, 57, 51, 48, 53, 57, 51, 115, 97, 56, 100, 110, 102, 108, 57, 56, 109, 114, 34, 44, 34, 110, 97, 109, 101, 34, 58, 34, 65, 114, 116, 104, 117, 114, 32, 68, 101, 110, 116, 34, 125},
				},
				CredentialDefinitionID: "",
				Signer:                 "elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr",
			},
			errors.New("credential definition DID must be set"),
		},
		{
			"FAIL: empty credential",
			fields{
				Credential:             nil,
				CredentialDefinitionID: "did:cosmos:elesto:cd-0001",
				Signer:                 "elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr",
			},
			errors.New("credential must be set"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MsgIssuePublicVerifiableCredentialRequest{
				Credential:             tt.fields.Credential,
				CredentialDefinitionID: tt.fields.CredentialDefinitionID,
				Signer:                 tt.fields.Signer,
			}
			err := m.ValidateBasic()

			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
		})
	}
}
