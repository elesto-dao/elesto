package credential

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/noandrea/rl2020"

	"github.com/elesto-dao/elesto/v4/x/did"
)

const (
	ProposePublicCredentialIDType       = "ProposePublicCredentialID"
	ProposeRemovePublicCredentialIDType = "ProposeRemovePublicCredentialID"
	// ProofTypeCosmosADR036 is a proof whose value is computed by signing a credential with kepler following ADR036.
	// An example would be credential signed with kepler
	ProofTypeCosmosADR036 = "CosmosADR036EcdsaSecp256k1Signature"
)

func init() {
	govtypes.RegisterProposalType(ProposePublicCredentialIDType)
	govtypes.RegisterProposalType(ProposeRemovePublicCredentialIDType)
	govtypes.RegisterProposalTypeCodec(&ProposePublicCredentialID{}, "credential/ProposePublicCredential")
	govtypes.RegisterProposalTypeCodec(&ProposeRemovePublicCredentialID{}, "credential/ProposeRemovePublicCredential")
}

// NewCredentialDefinitionFromFile create a credential definition by reading the data from a file
func NewCredentialDefinitionFromFile(id string, publisherDID did.DID,
	name, description string,
	isActive bool,
	schemaFile, vocabFile string) (*CredentialDefinition, error) {

	def := &CredentialDefinition{
		Id:          id,
		PublisherId: publisherDID.String(),
		Name:        name,
		Description: description,
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

// NewPublicVerifiableCredential creates a new Public Verifiable Credential
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
func (wc *WrappedCredential) GetBytes() ([]byte, error) {
	return json.Marshal(wc)
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
	if !hasSubject {
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

// Validate validates a verifiable credential against a provided public key
func (wc WrappedCredential) Validate(
	pk cryptotypes.PubKey,
) (err error) {
	sig, err := base64.StdEncoding.DecodeString(wc.Proof.ProofValue)
	if err != nil {
		return
	}
	// create a copy to reset the proof
	wcCopy := wc.Copy()
	wcCopy.Proof = nil
	wcData, err := wcCopy.GetBytes()
	if err != nil {
		return
	}

	if wc.Proof.Type == ProofTypeCosmosADR036 {
		//NOTE: we rely completely on the proof to get the account address
		sp := strings.Split(wc.Proof.VerificationMethod, "#")
		if len(sp) != 2 {
			err = fmt.Errorf("cannot retrieve account address from proof verification method")
			return
		}
		if wcData, err = toCosmosADR036Message(sp[1], wcData); err != nil {
			return
		}
	}

	if !pk.VerifySignature(wcData, sig) {
		err = fmt.Errorf("signature cannot be verified")
	}
	return
}

func toCosmosADR036Message(cosmosAddress string, vc []byte) (tx []byte, err error) {
	base := `{
		"chain_id": "",
		"account_number": "0",
		"sequence": "0",
		"fee": {
		  "gas": "0",
		  "amount": []
		},
		"msgs": [
		  {
			"type": "sign/MsgSignData",
			"value": {
			  "signer": "%s",
			  "data": "%s"
			}
		  }
		],
		"memo": ""
	  }`

	var c interface{}
	if err = json.Unmarshal([]byte(fmt.Sprintf(base, cosmosAddress, base64.StdEncoding.EncodeToString(vc))), &c); err != nil {
		return
	}
	tx, err = json.Marshal(c)
	return
}

// NewProof create a new proof for a verifiable credential
func NewProof(
	proofType string,
	created string,
	proofPurpose string,
	verificationMethod string,
	signature string,
) *Proof {
	return &Proof{
		Type:               proofType,
		Created:            created,
		ProofPurpose:       proofPurpose,
		VerificationMethod: verificationMethod,
		ProofValue:         signature,
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

func NewProposePublicCredentialID(Title, Description, CredentialDefinitionID string) *ProposePublicCredentialID {
	return &ProposePublicCredentialID{
		Title:                  Title,
		Description:            Description,
		CredentialDefinitionID: CredentialDefinitionID,
	}
}

func (m *ProposePublicCredentialID) ProposalRoute() string {
	return RouterKey
}

func (m *ProposePublicCredentialID) ProposalType() string {
	return ProposePublicCredentialIDType
}

func (m *ProposePublicCredentialID) ValidateBasic() error {
	if m.CredentialDefinitionID == "" {
		return fmt.Errorf("empty credential definition id")
	}

	return govtypes.ValidateAbstract(m)
}

func NewProposeRemovePublicCredentialID(Title, Description, CredentialDefinitionID string) *ProposeRemovePublicCredentialID {
	return &ProposeRemovePublicCredentialID{
		Title:                  Title,
		Description:            Description,
		CredentialDefinitionID: CredentialDefinitionID,
	}
}

func (m *ProposeRemovePublicCredentialID) ProposalRoute() string {
	return RouterKey
}

func (m *ProposeRemovePublicCredentialID) ProposalType() string {
	return ProposeRemovePublicCredentialIDType
}

func (m *ProposeRemovePublicCredentialID) ValidateBasic() error {
	if m.CredentialDefinitionID == "" {
		return fmt.Errorf("empty credential definition id")
	}

	return govtypes.ValidateAbstract(m)
}
