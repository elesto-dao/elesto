package credential

import ( // this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

//nolint
func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&ProposePublicCredentialID{}, "credential/propose-public-credential", nil)
	cdc.RegisterConcrete(&ProposeRemovePublicCredentialID{}, "credential/propose-remove-public-credential", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPublishCredentialDefinitionRequest{},
		&MsgUpdateCredentialDefinitionRequest{},
	)

	registry.RegisterImplementations((*govtypes.Content)(nil),
		&ProposePublicCredentialID{},
		&ProposeRemovePublicCredentialID{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	// ModuleCdc codec used by the module (protobuf)
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
