package cli

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/noandrea/rl2020"
	"github.com/spf13/cobra"

	"github.com/elesto-dao/elesto/x/credential"
	"github.com/elesto-dao/elesto/x/did"
)

func NewQueryCredentialStatusCmd() *cobra.Command {
	var (
		command = "credential-status"
	)

	cmd := &cobra.Command{
		Use:     use(command, "credentialJSONFile"),
		Short:   "verify the credential status of a credential",
		Example: exQuery(command, "my-credential.json"),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			var (
				queryClient    = credential.NewQueryClient(clientCtx)
				credentialFile = args[0]
				wc             *credential.WrappedCredential
				rs             revocationStatus
			)

			// read the credential from file
			if wc, err = credential.NewWrappedPublicCredentialFromFile(credentialFile); err != nil {
				println("error building credential definition", err)
				return err
			}
			// check for revocation
			if rs, err = checkRevocation(queryClient, wc); err != nil {
				println("error processing credential revocation:", err)
				return err
			}
			// print results
			return clientCtx.PrintBytes(rs.GetBytes())
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func NewQueryPublicCredentialStatusCmd() *cobra.Command {
	var (
		command = "public-credential-status"
	)
	cmd := &cobra.Command{
		Use:     use(command, "credentialID"),
		Short:   "verify the credential status of a credential",
		Example: exQuery(command, "https://exmaple.resolver/credential/1234"),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			//
			var (
				queryClient  = credential.NewQueryClient(clientCtx)
				credentialID = args[0]
				wc           *credential.WrappedCredential
				rs           revocationStatus
			)
			// retrieve the public credential
			if wc, err = queryPublicCredential(queryClient, credentialID); err != nil {
				fmt.Println(err)
				return err
			}
			// check for revocation
			if rs, err = checkRevocation(queryClient, wc); err != nil {
				println("error processing credential revocation:", err)
				return err
			}
			// print results
			return clientCtx.PrintBytes(rs.GetBytes())
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func NewMakeCredentialFromSchemaCmd() *cobra.Command {
	var credentialFileOut string
	cmd := &cobra.Command{
		Use:     "prepare-credential definitionID",
		Short:   "interactively build a credential based on the definition",
		Example: "elestod query credential prepare-credential cd-1",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			var (
				qc            = credential.NewQueryClient(clientCtx)
				schema        CredentialSchema
				wc            *credential.WrappedCredential          // wrapped credential
				rlcs          []*credential.WrappedCredential        // the revocation list credentials
				vc            *credential.PublicVerifiableCredential // verifiable credential
				definitionDID = did.NewChainDID(clientCtx.ChainID, args[0])
			)

			result, err := qc.CredentialDefinition(
				context.Background(),
				&credential.QueryCredentialDefinitionRequest{
					Id: definitionDID.String(),
				},
			)
			if err != nil {
				return err
			}
			// unmarshall the schema
			if err = json.Unmarshal([]byte(result.Definition.Schema), &schema); err != nil {
				return err
			}

			// credential ID
			cID := askString("", "What is the credential ID?")
			var cIssuer did.DID
			// ask default data
			if cIssuer, err = askDid("Enter the issuer DID"); err != nil {
				println("not a valid did, but carrying on anyway")
			}

			vc = credential.NewPublicVerifiableCredential(cID,
				credential.WithIssuanceDate(time.Now()),
				credential.WithType(fmt.Sprint(schema.Title, "Credential")),
				credential.WithContext(fmt.Sprintf("https://resolver.cc/context/%s", definitionDID)),
				credential.WithIssuerDID(cIssuer),
			)

			if _, ok := schema.Properties["credentialStatus"]; ok {
				println("Setting up the revocation list")

				if rlcs, err = revocationListCredentials(qc, vc.Issuer); err != nil {
					return err
				}
				if len(rlcs) == 0 {
					fmt.Println("there are no revocation lists defined for this ")
					return fmt.Errorf("there are no revocation list credentials for this issuer: %s", vc.Issuer)
				}

				rlID := selectRevocationList("Select the revocation list to associate the credential to", rlcs)
				var (
					idx int
				)
				if idx, err = askPositiveNumber("enter the credential index"); err != nil {
					println("not a valid did, but carrying on anyway")
				}
				vc.CredentialStatus = credential.NewCredentialStatus(rlID, idx)
			}
			// wrap the credential
			if wc, err = credential.NewWrappedCredential(vc); err != nil {
				return err
			}
			// build the subject
			println("Setting up the credential subject")
			// now recursively build the credential subject
			var sbj = make(map[string]interface{})
			cs := schema.Properties["credentialSubject"].(map[string]interface{})["properties"].(map[string]interface{})
			for p, s := range cs {
				var desc string
				if v, ok := s.(map[string]interface{})["description"]; ok {
					desc = v.(string)
				}

				v := askString(desc, "enter the value for", p)
				if v != "" {
					sbj[p] = v
				}

			}
			// set the subject
			if err = wc.SetSubject(sbj); err != nil {
				return err
			}

			// save the file
			id, _ := wc.GetSubjectID()
			id = strings.TrimPrefix(id, "did:cosmos:elesto:")
			outFile := strings.ToLower(fmt.Sprintf("credential.%s.%s.json", schema.Title, id))

			wcB, err := wc.GetBytes()
			if err != nil {
				return err
			}

			if n := q("leave it empty not to save", "credential filename", outFile); !credential.IsEmpty(n) {
				return os.WriteFile(n, wcB, 0600)
			}
			return clientCtx.PrintBytes(wcB)
		},
	}

	cmd.Flags().StringVar(&credentialFileOut, "export", "", "export the unsigned credential to a json file")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

//
func queryPublicCredential(qc credential.QueryClient, credentialID string) (wc *credential.WrappedCredential, err error) {
	// query the public credential
	result, err := qc.PublicCredential(
		context.Background(),
		&credential.QueryPublicCredentialRequest{
			Id: credentialID,
		},
	)
	if err != nil {
		err = fmt.Errorf("public credential %s not found: %w", credentialID, err)
		return
	}
	wc, err = credential.NewWrappedCredential(result.Credential)
	if err != nil {
		err = fmt.Errorf("error processing credential %s: %w", credentialID, err)
	}
	return
}

type revocationStatus struct {
	Revoked    bool
	Credential *credential.WrappedCredential
}

func (rs revocationStatus) GetBytes() []byte {
	bs, err := json.Marshal(rs)
	if err != nil {
		panic(err)
	}
	return bs
}

// checkRevocation perform a revocation list check for a credential
func checkRevocation(qc credential.QueryClient, wc *credential.WrappedCredential) (rs revocationStatus, err error) {

	// is there a credential status to use?
	if wc.CredentialStatus == nil {
		err = fmt.Errorf("missing credentialStatus definition from the credential, revocation cannot be checked")
		return
	}
	// retrieve the revocation list
	res, err := qc.PublicCredential(
		context.Background(),
		&credential.QueryPublicCredentialRequest{
			Id: wc.CredentialStatus.RevocationListCredential,
		},
	)
	if err != nil {
		err = fmt.Errorf("revocation list credential %s not found: %w", wc.CredentialStatus.RevocationListCredential, err)
		return
	}
	// check issuer
	if res.Credential.Issuer != wc.Issuer {
		err = fmt.Errorf("credential issuer mismatch, expected: %v, got %v", wc.Issuer, res.Credential.Issuer)
		return
	}
	// load the revocation list
	rl, err := rl2020.NewRevocationListFromJSON(res.Credential.CredentialSubject)
	if err != nil {
		err = fmt.Errorf("error parsing the credential revocation list: %w", err)
		return
	}
	rs.Credential = wc
	rs.Revoked, err = rl.IsRevoked(*wc.CredentialStatus)
	return
}

type CredentialSchema struct {
	Schema   string `json:"$schema"`
	ID       string `json:"$id"`
	Metadata struct {
		Slug         string `json:"slug"`
		Version      string `json:"version"`
		Icon         string `json:"icon"`
		Discoverable bool   `json:"discoverable"`
		Uris         struct {
			JSONLdContextPlus string `json:"jsonLdContextPlus"`
			JSONLdContext     string `json:"jsonLdContext"`
			JSONSchema        string `json:"jsonSchema"`
		} `json:"uris"`
	} `json:"$metadata"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Type        string                 `json:"type"`
	Required    []string               `json:"required"`
	Properties  map[string]interface{} `json:"properties"`
}

func revocationListCredentials(qc credential.QueryClient, issuerDID string) (pwcs []*credential.WrappedCredential, err error) {
	var res *credential.QueryPublicCredentialsByIssuerResponse
	// query credentials
	if res, err = qc.PublicCredentialsByIssuer(
		context.Background(),
		&credential.QueryPublicCredentialsByIssuerRequest{
			Did: issuerDID,
		},
	); err != nil {
		return
	}
	// process credentials
	for _, pvc := range res.Credential {
		pwc, err := credential.NewWrappedCredential(pvc)
		if err != nil {
			fmt.Printf("warning, cannot process credential %v: %v, for further inspection run the 'public-credential' command with the '--native' flag", pvc.Id, err)
			continue
		}
		if pwc.HasType(rl2020.TypeRevocationList2020) {
			pwcs = append(pwcs, pwc)
		}
	}
	return
}

func selectRevocationList(msg string, rlcs []*credential.WrappedCredential) (v string) {
	var opts []string
	for _, rl := range rlcs {
		opts = append(opts, rl.Id)
	}
	prompt := &survey.Select{
		Message: msg,
		Options: opts,
		Default: opts[0],
	}
	if err := survey.AskOne(prompt, &v); err != nil {
		panic(err)
	}
	return
}

func q(help, msg, def string) (a string) {
	prompt := &survey.Input{
		Message: msg,
		Help:    help,
		Default: def,
	}
	_ = survey.AskOne(prompt, &a)
	return
}

func askString(help string, msg ...string) (v string) {
	prompt := &survey.Input{
		Message: fmt.Sprint(msg),
		Help:    help,
	}
	_ = survey.AskOne(prompt, &v)
	return
}

func askDid(what ...string) (d did.DID, err error) {
	ds := askString("", what...)
	if !did.IsValidDID(ds) {
		err = fmt.Errorf("this is not a valid did")
	}
	d = did.DID(ds)
	return
}

func askPositiveNumber(what string) (v int, err error) {
	var (
		result string
		v64    int64
	)
	result = askString("", what)
	if v64, err = strconv.ParseInt(result, 10, 32); err != nil {
		return
	}
	if v < 0 {
		err = errors.New("Must be a positive number")
	}
	v = int(v64)
	return
}

func NewAddressInfoCmd() *cobra.Command {
	var credentialFileOut string
	cmd := &cobra.Command{
		Use:     "address-info name",
		Short:   "print info about an address",
		Example: "elestod query credential address-info validator",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			var (
				info keyring.Info
				path string
			)

			if info, err = clientCtx.Keyring.Key(args[0]); err != nil {
				return err
			}

			if hdPath, err := info.GetPath(); err == nil {
				path = hdPath.String()
			}

			v := fmt.Sprint(
				fmt.Sprintf("Name       %v\n", info.GetName()),
				fmt.Sprintf("Address    %v\n", info.GetAddress().String()),
				fmt.Sprintf("PubKeyHex  %v\n", hex.EncodeToString(info.GetPubKey().Bytes())),
				fmt.Sprintf("PubKeyLen  %v\n", len(info.GetPubKey().Bytes())),
				fmt.Sprintf("PubKeyType %v\n", info.GetPubKey().Type()),
				fmt.Sprintf("Type       %v\n", info.GetType().String()),
				fmt.Sprintf("Path       %v\n", path),
				fmt.Sprintf("PubKeyB64  %v\n", base64.StdEncoding.EncodeToString(info.GetPubKey().Bytes())),
			)

			return clientCtx.PrintString(v)
		},
	}

	cmd.Flags().StringVar(&credentialFileOut, "export", "", "export the unsigned credential to a json file")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
