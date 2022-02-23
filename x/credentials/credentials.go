package credentials

import "github.com/elesto-dao/elesto/x/did"

func NewCredentialIssuer(did did.DID, options ...IssuerOption) (*CredentialIssuer, error) {
	issuer := &CredentialIssuer{
		Did:         did.String(),
		Revocations: &RevocationList{},
		Issues:      make([]*CredentialIssuance, 0),
		Accepts:     make([]*CredentialConstraint, 0),
	}
	for _, fn := range options {
		if err := fn(issuer); err != nil {
			return nil, err
		}
	}
	return issuer, nil
}

// IssuerOption implements variadic pattern for optional did document fields
type IssuerOption func(issuer *CredentialIssuer) error

// WithRevocationList add optional verifications
func WithRevocationList(list RevocationList) IssuerOption {
	return func(issuer *CredentialIssuer) error {
		issuer.Revocations = &list
		return nil
	}
}