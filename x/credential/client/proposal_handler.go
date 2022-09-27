package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/elesto-dao/elesto/v2/x/credential/client/cli"
)

var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitProposePublicCredentialID, nil)
