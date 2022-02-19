package cli

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/coinbase/kryptology/pkg/accumulator"
	"github.com/coinbase/kryptology/pkg/core/curves"
	"github.com/elesto-dao/elesto/x/credentials"
)

// InitRevocationList initialize a new revocation list by generating a new seed from crypto/rand
// the seed is returned and must be stored securely by the caller
func InitRevocationList() (rl *credentials.RevocationList, hexSeed string, err error) {
	seed, err := generateSeed(32)
	if err != nil {
		return
	}
	hexSeed = hex.EncodeToString(seed)
	rl, err = BuildRevocationList(hex.EncodeToString(seed))
	return
}

// generateSeed generate a random byte seed
func generateSeed(length int) (entropy []byte, err error) {
	entropy = make([]byte, length)
	if _, err = rand.Read(entropy); err != nil {
		return
	}
	return
}

// BuildRevocationList build/updates a revocationList object
func BuildRevocationList(hexSeed string, elements ...string) (rl *credentials.RevocationList, err error) {
	seed, err := hex.DecodeString(hexSeed)
	if err != nil {
		return
	}
	// generate the new revocation
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	key, err := new(accumulator.SecretKey).New(curve, seed[:])
	if err != nil {
		return
	}

	// TODO: unmarshal accumulator
	acc, err := new(accumulator.Accumulator).New(curve)
	if err != nil {
		return
	}
	// add all the elements
	for _, e := range elements {
		element := curve.Scalar.Hash([]byte(e))
		acc, err = acc.Add(key, element)
		if err != nil {
			return
		}
	}
	// marshal binaries
	accBin, err := acc.MarshalBinary()
	if err != nil {
		return
	}
	// get the public key
	pub, err := key.GetPublicKey(curve)
	if err != nil {
		return
	}
	pubBin, err := pub.MarshalBinary()
	if err != nil {
		return
	}
	// build the new revocation list
	rl = &credentials.RevocationList{
		PublicKey:   hex.EncodeToString(accBin),
		Accumulator: hex.EncodeToString(pubBin),
	}
	return
}
