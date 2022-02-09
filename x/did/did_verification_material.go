package did

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
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
// TODO: this function shall return an error for invalid addresses
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
		PublicKeyMultibase: fmt.Sprint("F", hex.EncodeToString(pubKey)),
	}
}

// NewPublicKeyMultibaseFromHex encode a hex string in a pub key multibase format
func NewPublicKeyMultibaseFromHex(pubKeyHex string) (pkm *VerificationMethod_PublicKeyMultibase, err error) {
	pkb, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		return
	}
	// TODO: shall we check if it is conform to the verification material? probably
	pkm = NewPublicKeyMultibase(pkb)
	return
}

// NewPublicKeyHex build a new public key hex struct
// https://w3c.github.io/did-spec-registries/#publickeyhex
func NewPublicKeyHex(pubKey []byte) *VerificationMethod_PublicKeyHex {
	return &VerificationMethod_PublicKeyHex{
		PublicKeyHex: hex.EncodeToString(pubKey),
	}
}

// NewPublicKeyHexFromString build a new blockchain account ID struct
// https://w3c.github.io/did-spec-registries/#publickeyhex
func NewPublicKeyHexFromString(pubKeyHex string) (pkh *VerificationMethod_PublicKeyHex, err error) {
	pkb, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		return
	}
	// TODO: shall we check if it is conform to the verification material? probably
	pkh = NewPublicKeyHex(pkb)
	return
}

// NewPublicKeyJwk formats an account public key as hex string.
func NewPublicKeyJwk(pubKey []byte) (vm *VerificationMethod_PublicKeyJwk, err error) {
	var pkj PublicKeyJwk
	if err = json.Unmarshal(pubKey, &pkj); err != nil {
		return
	}
	if IsEmpty(pkj.Kid) {
		err = fmt.Errorf("publicKeyJwk.kid cannot be empty")
		return
	}
	if IsEmpty(pkj.X + pkj.Y) {
		err = fmt.Errorf("publicKeyJwk.X or publicKeyJwk.Y cannot be empty")
		return
	}
	vm = &VerificationMethod_PublicKeyJwk{
		PublicKeyJwk: &pkj,
	}
	return
}

// NewPublicKeyJwkFromJSON build a new blockchain account ID struct
func NewPublicKeyJwkFromJSON(pubKeyJSON string) (vm *VerificationMethod_PublicKeyJwk, err error) {
	return NewPublicKeyJwk([]byte(pubKeyJSON))
}
