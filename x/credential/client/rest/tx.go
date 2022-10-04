package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	govrest "github.com/cosmos/cosmos-sdk/x/gov/client/rest"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/elesto-dao/elesto/v2/x/credential"
)

type ProposePublicCredentialID struct {
	BaseReq                rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title                  string       `json:"title" yaml:"title"`
	Description            string       `json:"description" yaml:"description"`
	Deposit                sdk.Coins    `json:"deposit" yaml:"deposit"`
	CredentialDefinitionID string       `json:"credential_definition_id" yaml:"credential_definition_id"`
}

// ProposePublicCredentialRESTHandler returns propose public credential handler
func ProposePublicCredentialRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "propose-public-credential",
		Handler:  newProposePublicCredential(clientCtx),
	}
}

// newProposePublicCredential creates a handler for new public credential
func newProposePublicCredential(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ProposePublicCredentialID

		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		content := credential.NewProposePublicCredentialID(req.Title, req.Description, req.CredentialDefinitionID)
		msg, err := govtypes.NewMsgSubmitProposal(content, req.Deposit, fromAddr)
		if rest.CheckBadRequestError(w, err) {
			return
		}
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type ProposeRemovePublicCredentialID struct {
	BaseReq                rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title                  string       `json:"title" yaml:"title"`
	Description            string       `json:"description" yaml:"description"`
	Deposit                sdk.Coins    `json:"deposit" yaml:"deposit"`
	CredentialDefinitionID string       `json:"credential_definition_id" yaml:"credential_definition_id"`
}

// ProposeRemovePublicCredentialRESTHandler returns proposal handler to remove a public credential
func ProposeRemovePublicCredentialRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "remove-public-credential",
		Handler:  newRemoveProposePublicCredential(clientCtx),
	}
}

// newRemoveProposePublicCredential creates a handler for removing public credential id from allow list
func newRemoveProposePublicCredential(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ProposeRemovePublicCredentialID

		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		content := credential.NewProposeRemovePublicCredentialID(req.Title, req.Description, req.CredentialDefinitionID)
		msg, err := govtypes.NewMsgSubmitProposal(content, req.Deposit, fromAddr)
		if rest.CheckBadRequestError(w, err) {
			return
		}
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
