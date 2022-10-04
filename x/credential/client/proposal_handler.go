package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/cosmos/cosmos-sdk/x/gov/client/rest"

	"github.com/elesto-dao/elesto/v2/x/credential/client/cli"
)

var (
	ProposeCredentialIDHandler = govclient.NewProposalHandler(cli.NewSubmitProposePublicCredentialID, func(context client.Context) rest.ProposalRESTHandler {
		return rest.ProposalRESTHandler{}
	})

	ProposeRemoveCredentialIDHandler = govclient.NewProposalHandler(cli.NewSubmitRemoveProposePublicCredentialID, func(context client.Context) rest.ProposalRESTHandler {
		return rest.ProposalRESTHandler{}
	})
)
