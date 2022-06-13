package credential

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"time"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/noandrea/rl2020"

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

func NewPublicVerifiableCredential(id string, opts ...PublicVerifiableCredentialOption) *PublicVerifiableCredential {
	pvc := &PublicVerifiableCredential{
		Context: []string{"https://www.w3.org/2018/credentials/v1"},
		Id:      id,
		Type:    []string{"VerifiableCredential"},
	}
	for _, opt := range opts {
		opt(pvc)
	}
	return pvc
}

type PublicVerifiableCredentialOption func(credential *PublicVerifiableCredential)

func WithContext(context ...string) PublicVerifiableCredentialOption {
	return func(pvc *PublicVerifiableCredential) {
		pvc.Context = append(pvc.Context, context...)
	}
}

func WithType(typ ...string) PublicVerifiableCredentialOption {
	return func(pvc *PublicVerifiableCredential) {
		pvc.Type = append(pvc.Type, typ...)
	}
}

func WithIssuerDID(issuer did.DID) PublicVerifiableCredentialOption {
	return func(pvc *PublicVerifiableCredential) {
		pvc.Issuer = issuer.String()
	}
}

func WithIssuanceDate(date time.Time) PublicVerifiableCredentialOption {
	return func(pvc *PublicVerifiableCredential) {
		// this is to avoid
		utc := date.Truncate(time.Minute).UTC()
		pvc.IssuanceDate = &utc
	}
}

func WithExpirationDate(date time.Time) PublicVerifiableCredentialOption {
	return func(pvc *PublicVerifiableCredential) {
		utc := date.Truncate(time.Minute).UTC()
		pvc.ExpirationDate = &utc
	}
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
	if len(pvc.CredentialSubject) > 0 {
		err = json.Unmarshal(pvc.CredentialSubject, &wc.CredentialSubject)
	}
	return
}

// NewWrappedPublicCredentialFromFile read a credential from file
func NewWrappedPublicCredentialFromFile(credentialFile string) (wc *WrappedCredential, err error) {
	wc = &WrappedCredential{}
	data, err := os.ReadFile(credentialFile)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, wc); err != nil {
		return
	}
	if wc.PublicVerifiableCredential.CredentialSubject, err = json.Marshal(wc.CredentialSubject); err != nil {
		return
	}
	return
}

// GetBytes returns the JSON encoded byte slice of the credential
func (wc *WrappedCredential) GetBytes() (d []byte, err error) {
	dAtA, err := json.Marshal(wc)
	if err != nil {
		return nil, err
	}
	return dAtA, nil
}

// Copy create a deep copy of the WrappedCredential
func (wc *WrappedCredential) Copy() WrappedCredential {
	pvc := *wc.PublicVerifiableCredential
	wcCopy := *wc
	wcCopy.PublicVerifiableCredential = &pvc
	return wcCopy
}

func (wc *WrappedCredential) GetSubjectID() (s string, isDID bool) {
	v, hasSubject := wc.CredentialSubject["id"]
	if !hasSubject || IsEmpty(v.(string)) {
		return
	}
	id, hasSubject := v.(string)
	if !hasSubject {
		return
	}
	return id, did.IsValidDID(id)
}

// GetIssuerDID returns the DID of the issuer
func (pvc PublicVerifiableCredential) GetIssuerDID() did.DID {
	return did.DID(pvc.Issuer)
}

// HasType check if the credential has the type in input
func (wc *WrappedCredential) HasType(credentialType string) bool {
	for _, t := range wc.Type {
		if t == credentialType {
			return true
		}
	}
	return false
}

// SetSubject set the credential subject of the credential, it must be json serializable
func (wc *WrappedCredential) SetSubject(val interface{}) (err error) {
	if wc.PublicVerifiableCredential.CredentialSubject, err = json.Marshal(val); err != nil {
		return
	}
	// now unmarshal the wc
	return json.Unmarshal(wc.PublicVerifiableCredential.CredentialSubject, &wc.CredentialSubject)
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
) (err error) {
	sig, err := base64.StdEncoding.DecodeString(wc.Proof.Signature)
	if err != nil {
		return
	}
	// create a copy to reset the proof
	wcCopy := wc.Copy()
	wcCopy.Proof = nil
	// TODO: this is an expensive operation, could lead to DDOS
	// TODO: we can hash this and make this less expensive
	if !pk.VerifySignature(wcCopy.GetBytes(), sig) {
		err = fmt.Errorf("signature cannot be verified")
	}
	return
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

// implement the credential status interface required by the rl2020 lib

// NewCredentialStatus returns a new credential status
func NewCredentialStatus(credentialList string, index int) *CredentialStatus {
	return &CredentialStatus{
		Id:                       fmt.Sprint(credentialList, "/", index),
		Type:                     rl2020.TypeRevocationList2020Status,
		RevocationListIndex:      int32(index),
		RevocationListCredential: credentialList,
	}
}

// Coordinates retun the revocation list id and credential index within the list
func (m CredentialStatus) Coordinates() (string, int) {
	return m.RevocationListCredential, int(m.RevocationListIndex)
}

// TypeDef returns the credential status ID and type for correctness check
func (m CredentialStatus) TypeDef() (string, string) {
	return m.Id, m.Type
}
