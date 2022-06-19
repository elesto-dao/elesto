package credential

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/elesto-dao/elesto/v2/x/did"
)

// AccountKeeper defines the expected account keeper (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	GetPubKey(ctx sdk.Context, addr sdk.AccAddress) (cryptotypes.PubKey, error)
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	IsSendEnabledCoins(ctx sdk.Context, coins ...sdk.Coin) error
}

// DidKeeper defines the expected interfaces needed to manipulate did document
type DidKeeper interface {
	GetDidDocument(ctx sdk.Context, key []byte) (did.DidDocument, bool)
	SetDidDocument(ctx sdk.Context, key []byte, document did.DidDocument)
	ResolveDid(ctx sdk.Context, didDoc did.DID) (doc did.DidDocument, err error)
}
