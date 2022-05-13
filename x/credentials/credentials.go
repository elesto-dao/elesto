package credentials

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	"github.com/elesto-dao/elesto/x/did"
)

// NewCredentialDefinitionFromFile create a credential definition by reading the data from a file
func NewCredentialDefinitionFromFile(did, publisherDID did.DID,
	name, description string,
	isPublic, isActive bool,
	schemaFile, vocabFile string) (*CredentialDefinition, error) {

	def := &CredentialDefinition{
		Id:          did.String(),
		PublisherId: publisherDID.String(),
		Name:        name,
		Description: description,
		IsPublic:    isPublic,
		IsActive:    isActive,
	}

	var err error

	if def.Schema, err = CompactJSON(schemaFile); err != nil {
		err = fmt.Errorf("error reading schema file: %w ", err)
		return nil, err
	}
	if def.Vocab, err = CompactJSON(vocabFile); err != nil {
		err = fmt.Errorf("error reading vocab file: %w ", err)
		return nil, err
	}

	return def, nil
}

// WrappedCredential wraps a PublicVerifiableCredential, this is a workaround
// to deal with the variable content of the credential subject
type WrappedCredential struct {
	*PublicVerifiableCredential
	CredentialSubject map[string]interface{} `json:"credentialSubject"`
}

// NewWrappedCredential wrap a PublicVerifiableCredential to go around serialization
func NewWrappedCredential(pvc *PublicVerifiableCredential) (wc *WrappedCredential, err error) {
	wc = &WrappedCredential{
		PublicVerifiableCredential: pvc,
		CredentialSubject:          map[string]interface{}{},
	}
	err = json.Unmarshal(pvc.CredentialSubject, &wc.CredentialSubject)
	return
}

func NewWrappedPublicCredentialFromFile(credentialFile string) (wc *WrappedCredential, err error) {
	wc = &WrappedCredential{}
	data, err := os.ReadFile(credentialFile)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, wc); err != nil {
		return
	}
	return
}

func (wc *WrappedCredential) GetCredential() (*PublicVerifiableCredential, error) {
	sbj, err := json.Marshal(wc.CredentialSubject)
	if err != nil {
		return nil, err
	}
	wc.PublicVerifiableCredential.CredentialSubject = sbj
	return wc.PublicVerifiableCredential, nil
}

func (wc *WrappedCredential) GetBytes() []byte {
	dAtA, err := json.Marshal(wc)
	if err != nil {
		panic(err) //[(gogoproto.sizer) = true, (gogoproto.marshaler) = true,  (gogoproto.unmarshaler) = true];
	}
	return dAtA
}

func (wc *WrappedCredential) GetSubjectID() (s string, hasSubject bool) {
	v, hasSubject := wc.CredentialSubject["id"]
	if !hasSubject || IsEmpty(v.(string)) {
		return
	}
	id, hasSubject := v.(string)
	if !hasSubject {
		return
	}
	return id, IsEmpty(id)
}

func (pvc PublicVerifiableCredential) GetIssuerDID() did.DID {
	return did.DID(pvc.Issuer)
}

// GetBytes is a helper for serializing
func (pvc PublicVerifiableCredential) GetBytes() []byte {
	dAtA, err := pvc.Marshal()
	if err != nil {
		panic(err) //[(gogoproto.sizer) = true, (gogoproto.marshaler) = true,  (gogoproto.unmarshaler) = true];
	}
	return dAtA
}

// Validate validates a verifiable credential against a provided public key
func (wc WrappedCredential) Validate(
	pk cryptotypes.PubKey,
) bool {
	s, err := base64.StdEncoding.DecodeString(wc.Proof.Signature)
	if err != nil {
		panic(err)
	}

	// reset the proof
	wc.Proof = nil

	// TODO: this is an expensive operation, could lead to DDOS
	// TODO: we can hash this and make this less expensive
	isCorrectPubKey := pk.VerifySignature(
		wc.GetBytes(),
		s,
	)

	return isCorrectPubKey
}

// NewProof create a new proof for a verifiable credential
func NewProof(
	proofType string,
	created string,
	proofPurpose string,
	verificationMethod string,
	signature string,
) Proof {
	return Proof{
		Type:               proofType,
		Created:            created,
		ProofPurpose:       proofPurpose,
		VerificationMethod: verificationMethod,
		Signature:          signature,
	}
}
