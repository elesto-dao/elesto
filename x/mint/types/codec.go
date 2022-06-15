package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
)

//nolint
func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
}

var (
	// ModuleCdc codec used by the module (protobuf)
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
