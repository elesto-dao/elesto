package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/elesto-dao/elesto/v4/x/credential/client/cli"
	"github.com/elesto-dao/elesto/v4/x/credential/client/rest"
)

var (
	ProposeCredentialIDHandler       = govclient.NewProposalHandler(cli.NewSubmitProposePublicCredentialID, rest.ProposePublicCredentialRESTHandler)
	ProposeRemoveCredentialIDHandler = govclient.NewProposalHandler(cli.NewSubmitRemoveProposePublicCredentialID, rest.ProposeRemovePublicCredentialRESTHandler)
)
