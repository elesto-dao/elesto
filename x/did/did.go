package did

import (
	"encoding/hex"
	"fmt"
	"net/url"
	"regexp"
	"sort"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// VerificationRelationship defines the enum that represents the VerificationRelationship
type VerificationRelationship int

// A verification relationship expresses the relationship between the DID subject and a verification method.
// This enum is used to
// cfr. https://www.w3.org/TR/did-core/#verification-relationships
const (
	Authentication       = "authentication"       // https://www.w3.org/TR/did-core/#authentication
	AssertionMethod      = "assertionMethod"      // https://www.w3.org/TR/did-core/#assertion
	KeyAgreement         = "keyAgreement"         // https://www.w3.org/TR/did-core/#key-agreement
	CapabilityInvocation = "capabilityInvocation" // https://www.w3.org/TR/did-core/#capability-invocation
	CapabilityDelegation = "capabilityDelegation" // https://www.w3.org/TR/did-core/#capability-delegation
)

const (
	authentication VerificationRelationship = iota
	assertionMethod
	keyAgreement
	capabilityInvocation
	capabilityDelegation
)

// VerificationRelationships are the supported list of verification relationships
var supportedRelationships = map[string]VerificationRelationship{
	Authentication:       authentication,
	AssertionMethod:      assertionMethod,
	KeyAgreement:         keyAgreement,
	CapabilityInvocation: capabilityInvocation,
	CapabilityDelegation: capabilityDelegation,
}

// verificationRelationships retrieve the pointer to the verification relationship
// if it exists, otherwise returns nil
func (didDoc *DidDocument) getRelationships(rel VerificationRelationship) *[]string {
	switch rel {
	case authentication:
		return &didDoc.Authentication
	case assertionMethod:
		return &didDoc.AssertionMethod
	case keyAgreement:
		return &didDoc.KeyAgreement
	case capabilityInvocation:
		return &didDoc.CapabilityInvocation
	case capabilityDelegation:
		return &didDoc.CapabilityDelegation
	default:
		return nil
	}
}

// parseRelationshipLabels parse relationships labels to a slice of VerificationRelationship
// making sure that the relationships are not repeated
func parseRelationshipLabels(relNames ...string) (vrs []VerificationRelationship, err error) {
	names := distinct(relNames)
	vrs = make([]VerificationRelationship, len(names))
	for i, vrn := range names {
		vr, validName := supportedRelationships[vrn]
		if !validName {
			err = sdkerrors.Wrapf(ErrInvalidInput, "unsupported verification relationship %s", vrn)
			return
		}
		vrs[i] = vr
	}
	return
}

/**
Regexp generated using this ABNF specs and using https://abnf.msweet.org/index.php

did-url            = did path-abempty [ "?" query ] [ "#" fragment ]
did                = "did:" method-name ":" method-specific-id
method-name        = 1*method-char
method-char        = %x61-7A / DIGIT
method-specific-id = *( *idchar ":" ) 1*idchar
idchar             = ALPHA / DIGIT / "." / "-" / "_" / pct-encoded
pct-encoded        = "%" HEXDIG HEXDIG
query              = *( pchar / "/" / "?" )
fragment           = *( pchar / "/" / "?" )
path-abempty       = *( "/" segment )
segment            = *pchar
unreserved         = ALPHA / DIGIT / "-" / "." / "_" / "~"
pchar              = unreserved / pct-encoded / sub-delims / ":" / "@"
sub-delims         = "!" / "$" / "&" / "'" / "(" / ")"
                 / "*" / "+" / "," / ";" / "="
*/

const (
	contextDIDBase            = "https://www.w3.org/ns/did/v1"
	didValidationRegexpStr    = `^did\:cosmos\:(([A-Z.a-z0-9]|\-|_|%[0-9A-Fa-f][0-9A-Fa-f])*\:)*([A-Z.a-z0-9]|\-|_|%[0-9A-Fa-f][0-9A-Fa-f])+$`
	didURLValidationRegexpStr = `^did\:cosmos\:(([A-Z.a-z0-9]|\-|_|%[0-9A-Fa-f][0-9A-Fa-f])*\:)*([A-Z.a-z0-9]|\-|_|%[0-9A-Fa-f][0-9A-Fa-f])+(/(([-A-Z._a-z0-9]|~)|%[0-9A-Fa-f][0-9A-Fa-f]|(\!|\$|&|'|\(|\)|\*|\+|,|;|\=)|\:|@)*)*(\?(((([-A-Z._a-z0-9]|~)|%[0-9A-Fa-f][0-9A-Fa-f]|(\!|\$|&|'|\(|\)|\*|\+|,|;|\=)|\:|@)|/|\?)*))?(#(((([-A-Z._a-z0-9]|~)|%[0-9A-Fa-f][0-9A-Fa-f]|(\!|\$|&|'|\(|\)|\*|\+|,|;|\=)|\:|@)|/|\?)*))?$`
)

var (
	didValidationRegexp    = regexp.MustCompile(didValidationRegexpStr)
	didURLValidationRegexp = regexp.MustCompile(didURLValidationRegexpStr)
)

// DID as typed string
type DID string

// VerificationRelationships for did document manipulation
type VerificationRelationships []string

// Contexts for json-ld contects
type Contexts []string

// NewChainDID format a DID from a method specific did
// cfr.https://www.w3.org/TR/did-core/#did
func NewChainDID(chainName, didID string) DID {
	return DID(fmt.Sprint(DidChainPrefix, chainName, ":", didID))
}

// NewKeyDID format a DID of type key
func NewKeyDID(account string) DID {
	return DID(fmt.Sprint(DidKeyPrefix, account))
}

// String return the string representation of the did
func (did DID) String() string {
	return string(did)
}

// NewVerificationMethodID compose a verification method id from an account address
func (did DID) NewVerificationMethodID(vmID string) string {
	return fmt.Sprint(did, "#", strings.TrimSpace(vmID))
}

// NewServiceID compose a new service id URI from the did and the provided id
func (did DID) NewServiceID(serviceID string) string {
	return fmt.Sprint(did, "#", strings.TrimSpace(serviceID))
}

// IsValidDID validate the input string according to the
// did specification (cfr. https://www.w3.org/TR/did-core/#did-syntax ).
func IsValidDID(input string) bool {
	return didValidationRegexp.MatchString(input)
}

// IsValidDIDURL validate the input string according to the
// did url specification (cfr. https://www.w3.org/TR/did-core/#did-url-syntax  ).
func IsValidDIDURL(input string) bool {
	return didURLValidationRegexp.MatchString(input)
}

// IsValidRFC3986Uri checks if the input string is a valid RFC3986 URI
// (cfr https://datatracker.ietf.org/doc/html/rfc3986#page-50)
// it uses the algorithm from https://github.com/xeipuuv/gojsonschema/blob/v1.2.0/format_checkers.go#L275
// that is used in hyperledger aries go https://github.com/hyperledger/aries-framework-go/tree/v0.1.7
func IsValidRFC3986Uri(input string) bool {
	u, err := url.Parse(input)
	if err != nil || u.Scheme == "" {
		return false
	}
	return !strings.Contains(input, `\`)
}

// IsValidDIDDocument tells if a DID document is valid,
// that is if it has the default context and a valid subject
func IsValidDIDDocument(didDoc *DidDocument) bool {
	if didDoc == nil {
		return false
	}

	if !IsValidDID(didDoc.Id) {
		return false
	}

	for _, element := range didDoc.Controller {
		if !IsValidDIDKeyFormat(element) {
			return false
		}
	}

	//check that in the relationships there are only listed vm ids
	// list all the vmIds from relationships
	relVMIDs := map[string]struct{}{}
	for _, rK := range supportedRelationships {
		for _, vmID := range *didDoc.getRelationships(rK) {
			relVMIDs[vmID] = struct{}{}
		}
	}
	// verify all the vm
	for _, element := range didDoc.VerificationMethod {
		vm, _ := element.VerificationMaterial.(Validable)
		if vErr := vm.Validate(); vErr != nil {
			return false
		}
		// remove the vmId for the list of vmIds from relationships
		delete(relVMIDs, element.Id)
	}
	// if there are items left in the list of vmIds from relationships
	// then there are invalid relationships vmIds
	if len(relVMIDs) > 0 {
		return false
	}
	// validate all services
	for _, s := range didDoc.Service {
		if err := ValidateService(s); err != nil {
			return false
		}
	}
	// check that there is at least the did base context
	hasBaseContext := false
	for _, c := range didDoc.Context {
		if c == contextDIDBase {
			hasBaseContext = true
		}
		if !IsValidRFC3986Uri(c) {
			return false
		}
	}
	// last check about the presence of the base context
	return hasBaseContext
}

// IsValidDIDKeyFormat verify that a did is compliant with the did:cosmos:key format
// that is the ID must be a bech32 address no longer than 255 bytes
func IsValidDIDKeyFormat(did string) bool {
	if _, err := sdk.AccAddressFromBech32(strings.TrimPrefix(did, DidKeyPrefix)); err != nil {
		return false
	}
	return true
}

// ValidateVerification perform basic validation on a verification struct
func ValidateVerification(v *Verification) (err error) {
	if v == nil {
		err = sdkerrors.Wrap(ErrInvalidInput, "verification is not defined")
		return
	}

	// check the method exists before accessing
	if v.Method == nil {
		err = sdkerrors.Wrap(ErrInvalidInput, "verification.Method is not defined")
		return
	}

	// verify that the method id is correct
	if !IsValidDIDURL(v.Method.Id) {
		err = sdkerrors.Wrapf(ErrInvalidDIDURLFormat, "verification method id: %v", v.Method.Id)
		return
	}

	// if the controller is not set return error
	if !IsValidDID(v.Method.Controller) {
		err = sdkerrors.Wrapf(ErrInvalidDIDFormat, "verification method controller %v", v.Method.Controller)
		return
	}

	// check for empty method type
	if IsEmpty(v.Method.Type) {
		err = sdkerrors.Wrapf(ErrInvalidInput, "verification method type not set for verification method %s", v.Method.Id)
		return
	}

	// check that there is at least a relationship
	if len(v.Relationships) == 0 {
		err = sdkerrors.Wrap(ErrEmptyRelationships, "at least a verification relationship is required")
		return
	}

	// check the verification material
	vm, ok := v.Method.VerificationMaterial.(Validable)
	if !ok {
		err = sdkerrors.Wrap(
			ErrInvalidInput,
			fmt.Sprintf(
				"verification material '%v' unknown for verification method id %s",
				v.Method.VerificationMaterial,
				v.Method.Id,
			),
		)
		return
	}
	if vErr := vm.Validate(); vErr != nil {
		err = sdkerrors.Wrap(
			ErrInvalidInput,
			fmt.Sprintf(
				"verification material %s for verification method id %s",
				vErr.Error(),
				v.Method.Id,
			),
		)
	}
	return
}

// ValidateService performs basic on a service struct
func ValidateService(s *Service) (err error) {
	if s == nil {
		err = sdkerrors.Wrap(ErrInvalidInput, "service is not defined")
		return
	}
	// verify that the id is not empty and is a valid url (according to RFC3986)
	if IsEmpty(s.Id) {
		err = sdkerrors.Wrap(ErrInvalidInput, "service id cannot be empty")
		return
	}

	if !IsValidRFC3986Uri(s.Id) {
		err = sdkerrors.Wrapf(ErrInvalidRFC3986UriFormat, "service id %s is not a valid RFC3986 uri", s.Id)
		return
	}

	// verify that the endpoint is not empty and is a valid url (according to RFC3986)
	if IsEmpty(s.ServiceEndpoint) {
		err = sdkerrors.Wrap(ErrInvalidInput, "service endpoint cannot be empty;")
		return
	}

	if !IsValidRFC3986Uri(s.ServiceEndpoint) {
		err = sdkerrors.Wrapf(ErrInvalidRFC3986UriFormat, "service endpoint %s is not a valid RFC3986 uri", s.ServiceEndpoint)
		return
	}

	// check that the service type is not empty
	if IsEmpty(s.Type) {
		err = sdkerrors.Wrap(ErrInvalidInput, "service type cannot be empty")
		return
	}

	return
}

// IsEmpty tells if the trimmed input is empty
func IsEmpty(input string) bool {
	return strings.TrimSpace(input) == ""
}

// DocumentOption implements variadic pattern for optional did document fields
type DocumentOption func(*DidDocument) error

// WithVerifications add optional verifications
func WithVerifications(verifications ...*Verification) DocumentOption {
	return func(did *DidDocument) error {
		return did.AddVerifications(verifications...)
	}
}

// WithServices add optional services
func WithServices(services ...*Service) DocumentOption {
	return func(did *DidDocument) error {
		return did.AddServices(services...)
	}
}

// WithControllers add optional did controller
func WithControllers(controllers ...string) DocumentOption {
	return func(did *DidDocument) (err error) {
		return did.AddControllers(controllers...)
	}
}

// NewDidDocument constructs a new DidDocument
func NewDidDocument(id string, options ...DocumentOption) (did DidDocument, err error) {

	if !IsValidDID(id) {
		err = sdkerrors.Wrapf(ErrInvalidDIDFormat, "did %s", id)
		return
	}

	did = DidDocument{
		Context: []string{contextDIDBase},
		Id:      id,
	}
	// apply all the options
	for _, fn := range options {
		if err = fn(&did); err != nil {
			return
		}
	}
	return
}

// AddControllers add a controller to a did document if not exists
func (didDoc *DidDocument) AddControllers(controllers ...string) error {
	if len(controllers) == 0 {
		return nil
	}
	// join the exiting controllers with the new ones
	dc := distinct(append(didDoc.Controller, controllers...))
	for _, c := range dc {
		if !IsValidDID(c) {
			return sdkerrors.Wrapf(ErrInvalidDIDFormat, "did document controller validation error '%s'", c)
		}
		if !IsValidDIDKeyFormat(c) {
			return sdkerrors.Wrapf(ErrInvalidInput, "did document controller '%s' must be of type key", c)
		}
	}

	// remove duplicates
	didDoc.Controller = dc
	return nil
}

// DeleteControllers delete controllers from a did document
func (didDoc *DidDocument) DeleteControllers(controllers ...string) error {
	if len(controllers) == 0 {
		return nil
	}
	dc := distinct(controllers)
	for _, c := range dc {
		if !IsValidDID(c) {
			return sdkerrors.Wrapf(ErrInvalidDIDFormat, "did document controller validation error '%s'", c)
		}
	}
	// remove existing
	didDoc.Controller = subtraction(didDoc.Controller, controllers)
	return nil
}

// AddVerifications add one or more verification method and relations to a did document
func (didDoc *DidDocument) AddVerifications(verifications ...*Verification) (err error) {
	// verify that there are no duplicates in method ids
	index := make(map[string]struct{}, len(didDoc.VerificationMethod))
	// load existing verifications if any
	for _, v := range didDoc.VerificationMethod {
		index[v.Id] = struct{}{}
	}

	// loop through the verifications and look for problems
	for _, v := range verifications {
		// perform base validation checks
		if err = ValidateVerification(v); err != nil {
			return
		}

		// verify that there are no duplicates in method ids
		if _, found := index[v.Method.Id]; found {
			err = sdkerrors.Wrapf(ErrInvalidInput, "duplicated verification method id %s", v.Method.Id)
			return
		}
		index[v.Method.Id] = struct{}{}

		// first add the method to the list of methods
		didDoc.VerificationMethod = append(didDoc.VerificationMethod, v.GetMethod())

		// now add the relationships
		vrs, err := parseRelationshipLabels(v.Relationships...)
		if err != nil {
			return err
		}
		didDoc.setRelationships(v.Method.Id, vrs...)

		// update context
		didDoc.Context = union(didDoc.Context, v.Context)

	}
	return
}

// RevokeVerification revoke a verification method
// and all relationships associated with it
func (didDoc *DidDocument) RevokeVerification(methodID string) error {

	del := func(x int) {
		lastIdx := len(didDoc.VerificationMethod) - 1
		switch lastIdx {
		case 0:
			didDoc.VerificationMethod = nil
		case x:
			didDoc.VerificationMethod = didDoc.VerificationMethod[:lastIdx]
		default:
			didDoc.VerificationMethod[x] = didDoc.VerificationMethod[lastIdx]
			didDoc.VerificationMethod = didDoc.VerificationMethod[:lastIdx]
		}
	}

	// remove relationships
	didDoc.setRelationships(methodID)

	// now remove the method
	for i, vm := range didDoc.VerificationMethod {
		if vm.Id == methodID {
			del(i)
			return nil
		}
	}
	return sdkerrors.Wrapf(ErrVerificationMethodNotFound, "verification method id: %v", methodID)
}

// SetVerificationRelationships for a did document
func (didDoc *DidDocument) SetVerificationRelationships(methodID string, relationships ...string) error {
	// verify that the method id is correct
	if !IsValidDIDURL(methodID) {
		return sdkerrors.Wrapf(ErrInvalidDIDURLFormat, "verification method id: %v", methodID)
	}
	// check that the methodID exists
	hasVM := false
	for _, vm := range didDoc.VerificationMethod {
		if vm.Id == methodID {
			hasVM = true
			break
		}
	}
	if !hasVM {
		return sdkerrors.Wrapf(ErrVerificationMethodNotFound, "verification method %v not found", methodID)
	}
	// check that there is at least a relationship
	if len(relationships) == 0 {
		return sdkerrors.Wrap(ErrEmptyRelationships, "at least a verification relationship is required")
	}
	// check that the provided relationships are valid
	vrs, err := parseRelationshipLabels(relationships...)
	if err != nil {
		return err
	}
	// update the relationships
	didDoc.setRelationships(methodID, vrs...)
	return nil
}

// setRelationships overwrite relationships for a did document
func (didDoc *DidDocument) setRelationships(methodID string, relationships ...VerificationRelationship) {
	// first remove existing relationships
	for _, vr := range supportedRelationships {
		vrs := didDoc.getRelationships(vr)
		for i, vmID := range *vrs {
			if vmID == methodID {
				lastIdx := len(*vrs) - 1 // get the last index of the current relationship list
				switch lastIdx {
				case 0: // remove the relationships since there is no elements left
					*vrs = nil
				case i: // if it's at the last position, just drop the last position
					*vrs = (*vrs)[:lastIdx]
				default: // swap and drop last position
					(*vrs)[i] = (*vrs)[lastIdx]
					(*vrs) = (*vrs)[:lastIdx]
				}
			}
		}
	}

	// then assign the new ones
	for _, vr := range relationships {
		vrs := didDoc.getRelationships(vr)
		*vrs = append(*vrs, methodID)
	}
}

// GetVerificationMethodBlockchainAddress returns the verification method cosmos blockchain address of a verification method.
// it fails if the verification method is not supported or if the verification method is not found
func (didDoc DidDocument) GetVerificationMethodBlockchainAddress(methodID string) (address string, err error) {
	for _, vm := range didDoc.VerificationMethod {
		if vm.Id == methodID {
			switch k := vm.VerificationMaterial.(type) {
			case *VerificationMethod_BlockchainAccountID:
				address = k.GetAddress()
			case *VerificationMethod_PublicKeyMultibase:
				address, err = toAddress(k.PublicKeyMultibase[1:])
			case *VerificationMethod_PublicKeyHex:
				address, err = toAddress(k.PublicKeyHex)
			}
			return
		}
	}
	err = ErrVerificationMethodNotFound
	return
}

// GetVerificationRelationships returns the relationships associated with the
// verification method id.
func (didDoc DidDocument) GetVerificationRelationships(methodID string) []string {
	relationships := []string{}
	for vrn, vr := range supportedRelationships {
		for _, vmID := range *didDoc.getRelationships(vr) {
			if vmID == methodID {
				relationships = append(relationships, vrn)
			}
		}
	}
	return relationships
}

// HasRelationship verifies if a controller did
// exist for at least one of the relationships in the did document
func (didDoc DidDocument) HasRelationship(
	signer *VerificationMethod_BlockchainAccountID,
	relationships ...string,
) bool {
	// first check if the controller exists
	for _, vm := range didDoc.VerificationMethod {
		switch k := vm.VerificationMaterial.(type) {
		case *VerificationMethod_BlockchainAccountID:
			if k.BlockchainAccountID != signer.BlockchainAccountID {
				continue
			}
		case *VerificationMethod_PublicKeyMultibase:
			addr, err := toAddress(k.PublicKeyMultibase[1:])
			if err != nil || !signer.MatchAddress(addr) {
				continue
			}
		case *VerificationMethod_PublicKeyHex:
			addr, err := toAddress(k.PublicKeyHex)
			if err != nil || !signer.MatchAddress(addr) {
				continue
			}
		}
		vrs := didDoc.GetVerificationRelationships(vm.Id)
		if len(intersection(vrs, relationships)) > 0 {
			return true
		}
	}
	return false
}

// HasPublicKey validates if a public key is contained in a DidDocument
func (didDoc DidDocument) HasPublicKey(pubkey cryptotypes.PubKey) bool {
	for _, vm := range didDoc.VerificationMethod {
		switch key := vm.VerificationMaterial.(type) {
		case *VerificationMethod_BlockchainAccountID:
			address := sdk.MustBech32ifyAddressBytes(
				sdk.GetConfig().GetBech32AccountAddrPrefix(),
				pubkey.Address().Bytes(),
			)
			if key.MatchAddress(address) {
				return true
			}
		case *VerificationMethod_PublicKeyMultibase:
			if strings.EqualFold(key.PublicKeyMultibase, fmt.Sprint("F", hex.EncodeToString(pubkey.Bytes()))) {
				return true
			}

		case *VerificationMethod_PublicKeyHex:
			if key.PublicKeyHex == hex.EncodeToString(pubkey.Bytes()) {
				return true
			}
		}
	}
	return false
}

// HasController returns true if the DID document has the input DID as a controller, false otherwise
func (didDoc *DidDocument) HasController(controller DID) bool {
	ctrl := controller.String()
	for _, c := range didDoc.Controller {
		if c == ctrl {
			return true
		}
	}
	return false
}

// AddServices add services to a did document
func (didDoc *DidDocument) AddServices(services ...*Service) (err error) {
	if didDoc.Service == nil {
		didDoc.Service = []*Service{}
	}

	// used to check duplicates
	index := make(map[string]struct{}, len(didDoc.Service))

	// load existing services
	for _, s := range didDoc.Service {
		index[s.Id] = struct{}{}
	}

	// services must be unique
	for _, s := range services {
		if err = ValidateService(s); err != nil {
			return
		}

		// verify that there are no duplicates in method ids
		if _, found := index[s.Id]; found {
			err = sdkerrors.Wrapf(ErrInvalidInput, "duplicated service id %s", s.Id)
			return
		}
		index[s.Id] = struct{}{}

		didDoc.Service = append(didDoc.Service, s)
	}
	return
}

// DeleteService delete an existing service from a did document
func (didDoc *DidDocument) DeleteService(serviceID string) {
	del := func(x int) {
		lastIdx := len(didDoc.Service) - 1
		switch lastIdx {
		case 0: // remove the relationships since there is no elements left
			didDoc.Service = nil
		case x: // if it's at the last position, just drop the last position
			didDoc.Service = didDoc.Service[:lastIdx]
		default: // swap and drop last position
			didDoc.Service[x] = didDoc.Service[lastIdx]
			didDoc.Service = didDoc.Service[:lastIdx]
		}
	}

	for i, s := range didDoc.Service {
		if s.Id == serviceID {
			del(i)
			break
		}
	}
}

// GetBytes is a helper for serializing
func (didDoc DidDocument) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&didDoc))
}

// Verifications is a list of verification
type Verifications []*Verification

// NewVerification build a new verification to be
// attached to a did document
func NewVerification(
	method VerificationMethod,
	relationships VerificationRelationships,
	contexts Contexts,
) *Verification {
	return &Verification{
		Context:       contexts,
		Method:        &method,
		Relationships: relationships,
	}
}

// NewAccountVerification is a shortcut to create a verification based on cosmos address
func NewAccountVerification(did DID, chainID, accountAddress string, verificationRelationships ...string) *Verification {
	return NewVerification(
		NewVerificationMethod(
			fmt.Sprint(did.String(), "#", accountAddress),
			did,
			NewBlockchainAccountID(chainID, accountAddress),
			CosmosAccountAddress,
		),
		verificationRelationships,
		nil,
	)
}

// NewVerificationMethod build a new verification method
func NewVerificationMethod(
	id string,
	controller DID,
	vmr isVerificationMethod_VerificationMaterial,
	vmt VerificationMethodType,
) VerificationMethod {
	vm := VerificationMethod{
		Id:                   id,
		Controller:           controller.String(),
		Type:                 string(vmt),
		VerificationMaterial: vmr,
	}
	return vm
}

// GetBytes is a helper for serializing
func (did Verification) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&did))
}

// Services are a list of services
type Services []*Service

// NewService creates a new service
func NewService(id, serviceType, serviceEndpoint string) *Service {
	return &Service{
		Id:              id,
		Type:            serviceType,
		ServiceEndpoint: serviceEndpoint,
	}
}

// ResolveAccountDID generates a DID document from an address
func ResolveAccountDID(did, chainID string) (didDoc DidDocument, err error) {
	if !IsValidDIDKeyFormat(did) {
		err = ErrInvalidDidMethodFormat
		return
	}
	account := strings.TrimPrefix(did, DidKeyPrefix)
	accountDID := DID(did)
	// compose the did document
	didDoc, err = NewDidDocument(did, WithVerifications(
		NewVerification(
			NewVerificationMethod(
				accountDID.NewVerificationMethodID(account),
				accountDID, // the controller is the same as the did subject
				NewBlockchainAccountID(chainID, account),
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
	))
	return
}

// toAddress encode a kexKey string to cosmos based address
func toAddress(hexKey string) (addr string, err error) {
	// decode the hex string
	pkb, err := hex.DecodeString(hexKey)
	if err != nil {
		return
	}
	// check the size of the decoded byte slice, otherwise the pk.Address will panic
	if len(pkb) != secp256k1.PubKeySize {
		err = fmt.Errorf("invalid public key size")
		return
	}
	// load the public key
	pk := &secp256k1.PubKey{Key: pkb}
	// generate the address
	addr, err = sdk.Bech32ifyAddressBytes(sdk.GetConfig().GetBech32AccountAddrPrefix(), pk.Address())
	return
}

// union perform union, distinct amd sort operation between two slices
// duplicated element in list are removed
func union(a, b []string) []string {
	if len(b) == 0 {
		return a
	}
	m := make(map[string]struct{})
	for _, item := range a {
		m[item] = struct{}{}
	}
	for _, item := range b {
		if _, ok := m[item]; !ok {
			m[item] = struct{}{}
		}
	}
	u := make([]string, 0, len(m))
	for k := range m {
		u = append(u, k)
	}
	sort.Strings(u)
	return u
}

// intersection - keep only strings that are in both slices
func intersection(a, b []string) []string {
	m := make(map[string]struct{})
	for _, item := range a {
		m[item] = struct{}{}
	}
	var i []string
	for _, item := range distinct(b) {
		if _, ok := m[item]; ok {
			i = append(i, item)
		}
	}
	sort.Strings(i)
	return i
}

// distinct remove duplicates and sorts from a list of strings
func distinct(a []string) []string {
	m := make(map[string]struct{})
	for _, item := range a {
		m[item] = struct{}{}
	}
	d := make([]string, 0, len(m))
	for k := range m {
		d = append(d, k)
	}
	sort.Strings(d)
	return d
}

// subtraction remove set b from a
func subtraction(a, b []string) []string {
	m := make(map[string]struct{})
	for _, item := range a {
		m[item] = struct{}{}
	}
	var s []string
	for _, item := range distinct(b) {
		if _, ok := m[item]; !ok {
			s = append(s, item)
		}
	}
	sort.Strings(s)
	return s
}
