package did

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

// This assigment are here to make sure that the VerificationMaterial implement
// the validable interface. If a VerificationMaterial does not implement the interface
// then it will not be recognized and accepted at runtime
var (
	_ Validable = (*VerificationMethod_PublicKeyHex)(nil)
	_ Validable = (*VerificationMethod_PublicKeyJwk)(nil)
	_ Validable = (*VerificationMethod_PublicKeyMultibase)(nil)
	_ Validable = (*VerificationMethod_BlockchainAccountID)(nil)
)

// VerificationMethodType encode the verification material type
type VerificationMethodType string

// Verification methods supported types
const (
	EcdsaSecp256k1VerificationKey2019 VerificationMethodType = "EcdsaSecp256k1VerificationKey2019"
	Ed25519VerificationKey2018        VerificationMethodType = "Ed25519VerificationKey2018"
	X25519KeyAgreementKey2019         VerificationMethodType = "X25519KeyAgreementKey2019"
	Bls12381G1Key2020                 VerificationMethodType = "Bls12381G1Key2020"
	Bls12381G2Key2020                 VerificationMethodType = "Bls12381G2Key2020"
	// CosmosAccountAddress is a custom implementation for ephemeral dids
	CosmosAccountAddress VerificationMethodType = "CosmosAccountAddress"
)

// String return string name for the Verification Method type
func (p VerificationMethodType) String() string {
	return string(p)
}

// MatchAddress check if a blockchain id address matches another address
// the match ignore the chain ID
func (baID VerificationMethod_BlockchainAccountID) MatchAddress(address string) bool {
	return baID.GetAddress() == address
}

// GetAddress get the address from a blockchain account id
func (baID VerificationMethod_BlockchainAccountID) GetAddress() string {
	addrStart := strings.LastIndex(baID.BlockchainAccountID, ":")
	if addrStart < 0 {
		return ""
	}
	return baID.BlockchainAccountID[addrStart+1:]
}

// NewBlockchainAccountID build a new blockchain account ID struct
func NewBlockchainAccountID(chainID, account string) *VerificationMethod_BlockchainAccountID {
	return &VerificationMethod_BlockchainAccountID{
		BlockchainAccountID: fmt.Sprint("cosmos:", chainID, ":", account),
	}
}

// NewBlockchainAccountIDFromString build a new blockchain account ID struct
func NewBlockchainAccountIDFromString(baID string) *VerificationMethod_BlockchainAccountID {
	return &VerificationMethod_BlockchainAccountID{
		BlockchainAccountID: baID,
	}
}

// NewPublicKeyMultibase formats an account address as per the CAIP-10 Account ID specification.
// https://w3c.github.io/did-spec-registries/#publickeymultibase
// https://datatracker.ietf.org/doc/html/draft-multiformats-multibase-03#appendix-B.1
func NewPublicKeyMultibase(pubKey []byte) *VerificationMethod_PublicKeyMultibase {
	return &VerificationMethod_PublicKeyMultibase{
		PublicKeyMultibase: fmt.Sprint("F", strings.ToUpper(hex.EncodeToString(pubKey))),
	}
}

// NewPublicKeyMultibaseFromHex encode a hex string in a pub key multibase format
func NewPublicKeyMultibaseFromHex(pubKeyHex string) (pkm *VerificationMethod_PublicKeyMultibase, err error) {
	pkb, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		return
	}
	pkm = NewPublicKeyMultibase(pkb)

	if err = pkm.Validate(); err != nil {
		return nil, err
	}
	return
}

// NewPublicKeyHex build a new public key hex struct
// https://w3c.github.io/did-spec-registries/#publickeyhex
func NewPublicKeyHex(pubKey []byte) *VerificationMethod_PublicKeyHex {
	return &VerificationMethod_PublicKeyHex{
		PublicKeyHex: hex.EncodeToString(pubKey),
	}
}

// NewPublicKeyHexFromString build a new public key hex struct from a string
// https://w3c.github.io/did-spec-registries/#publickeyhex
func NewPublicKeyHexFromString(pubKeyHex string) (pkh *VerificationMethod_PublicKeyHex, err error) {
	pkb, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		return
	}
	pkh = NewPublicKeyHex(pkb)

	if err = pkh.Validate(); err != nil {
		return nil, err
	}

	return
}

// NewPublicKeyJwk build a PublicKeyJwk struct from a json string encoded as a byte sequence.
func NewPublicKeyJwk(pubKey []byte) (vm *VerificationMethod_PublicKeyJwk, err error) {
	var pkj PublicKeyJwk
	if err = json.Unmarshal(pubKey, &pkj); err != nil {
		return
	}
	vm = &VerificationMethod_PublicKeyJwk{
		PublicKeyJwk: &pkj,
	}
	if err = vm.Validate(); err != nil {
		return nil, err
	}
	return
}

// NewPublicKeyJwkFromJSON build a PublicKeyJwk struct from a json string
func NewPublicKeyJwkFromJSON(pubKeyJSON string) (vm *VerificationMethod_PublicKeyJwk, err error) {
	return NewPublicKeyJwk([]byte(pubKeyJSON))
}

// Validable interface requires implementation of a validate function for a
// verification VerificationMaterial
type Validable interface {
	isVerificationMethod_VerificationMaterial
	Validate() error
}

// Validate verify that the PublicKeyJwk is not empty and contains mandatory fields
func (vm VerificationMethod_PublicKeyJwk) Validate() (err error) {
	if IsEmpty(vm.PublicKeyJwk.Kid) {
		err = fmt.Errorf("publicKeyJwk.kid cannot be empty")
		return
	}
	if IsEmpty(vm.PublicKeyJwk.X + vm.PublicKeyJwk.Y) {
		err = fmt.Errorf("publicKeyJwk.X or publicKeyJwk.Y cannot be empty")
		return
	}
	return nil
}

// Validate verify that the PublicKeyHex is not empty
func (vm VerificationMethod_PublicKeyHex) Validate() error {
	if IsEmpty(vm.PublicKeyHex) {
		return fmt.Errorf("publicKeyHex is empty")
	}
	return nil
}

// Validate verify that the PublicKeyMultibase is not empty
func (vm VerificationMethod_PublicKeyMultibase) Validate() error {
	if IsEmpty(vm.PublicKeyMultibase) {
		return fmt.Errorf("publicKeyMultibase is empty")
	}
	return nil
}

// Validate verify that the BlockchainAccountID is not empty
func (baID VerificationMethod_BlockchainAccountID) Validate() error {
	if IsEmpty(baID.BlockchainAccountID) {
		return fmt.Errorf("blockchainAccountId is empty")
	}
	return nil
}
