package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/tendermint/starport/starport/pkg/cosmoscmd"
	"google.golang.org/grpc"

	"github.com/elesto-dao/elesto/v2/x/did"
)

// Test data for the sees to run
const (
	GrpcConnectionTimeoutSeconds = 10
	// TestKey is defined in the 00_start_chain.sh seed using the add-genesis-account cmd
	// this key is used to create the did in this seed
	TestKey              = "0c7636a266c5b92a3940ba15c136a87f03e29a5fa111034b31bbbfc75d606739"
	TestKey2             = "d154a53e883271a6444edb933993594338e01b02244974861215d4eb77f020c9"
	TestKey3             = "b2210c3c17118f55f8fc70a8650cae033708149b814d4d89757655889a64cdf3"
	AccountAddressPrefix = "elesto"
	ChainID              = "elesto"
	SampleDIDID          = "bob"
)

// Various prefixes for accounts and public keys
var (
	AccountPubKeyPrefix    = AccountAddressPrefix + "pub"
	ValidatorAddressPrefix = AccountAddressPrefix + "valoper"
	ValidatorPubKeyPrefix  = AccountAddressPrefix + "valoperpub"
	ConsNodeAddressPrefix  = AccountAddressPrefix + "valcons"
	ConsNodePubKeyPrefix   = AccountAddressPrefix + "valconspub"
)

// SetConfig initialize the configuration instance for the sdk
func SetConfig() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, AccountPubKeyPrefix)
	config.SetBech32PrefixForValidator(ValidatorAddressPrefix, ValidatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(ConsNodeAddressPrefix, ConsNodePubKeyPrefix)
	config.Seal()
}

// CreateGRPCConnection createa a grpc connection to a given url
func CreateGRPCConnection(addr string) *grpc.ClientConn {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(GrpcConnectionTimeoutSeconds)*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)

	if err != nil {
		panic(err)
	}

	return conn
}

// SignMsg signs an sdk.Message with a given private key, account number and account sequence
// this returns a signed byte array for broadcasting to the network
func SignMsg(
	encCfg cosmoscmd.EncodingConfig,
	msg sdk.Msg,
	priv1 *secp256k1.PrivKey,
	accNum uint64,
	accSeq uint64,
) []byte {
	// create a new transaction builder and attach the msg
	txBuilder := encCfg.TxConfig.NewTxBuilder()
	err := txBuilder.SetMsgs(msg)
	if err != nil {
		panic(err)
	}

	txBuilder.SetGasLimit(200000000)
	txBuilder.SetFeeAmount(sdk.Coins{sdk.NewInt64Coin("stake", 0)})

	// set the signatories of the tx
	sigV2 := signing.SignatureV2{
		PubKey: priv1.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  encCfg.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: accSeq,
	}
	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		panic(err)
	}

	signerData := xauthsigning.SignerData{
		ChainID:       ChainID,
		AccountNumber: accNum,
		Sequence:      accSeq,
	}
	// sign the data with the private key
	sigV2, err = tx.SignWithPrivKey(
		encCfg.TxConfig.SignModeHandler().DefaultMode(), signerData,
		txBuilder, priv1, encCfg.TxConfig, accSeq)
	if err != nil {
		panic(err)
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		panic(err)
	}

	// Generated Protobuf-encoded bytes
	txBytes, err := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		panic(err)
	}

	return txBytes
}

func main() {
	SetConfig()

	addr := "localhost:9090"
	conn := CreateGRPCConnection(addr)
	moduleBasics := module.NewBasicManager()
	encCfg := cosmoscmd.MakeEncodingConfig(moduleBasics)
	txClient := txtypes.NewServiceClient(conn)
	qc := did.NewQueryClient(conn)

	// decode the test private key
	bz, _ := hex.DecodeString(TestKey)
	priv1 := &secp256k1.PrivKey{Key: bz}

	accNum := uint64(15)
	accSeq := uint64(0)

	// create a did for the user
	didID := did.NewChainDID(ChainID, SampleDIDID)
	pubKey := priv1.PubKey()
	accAddress := sdk.AccAddress(priv1.PubKey().Address())
	vmID := didID.NewVerificationMethodID(accAddress.String())
	vmType := did.EcdsaSecp256k1VerificationKey2019
	auth := did.NewVerification(
		did.NewVerificationMethod(
			vmID,
			didID,
			did.NewPublicKeyMultibase(pubKey.Bytes()),
			vmType,
		),
		[]string{did.Authentication},
		nil,
	)

	msg := did.NewMsgCreateDidDocument(
		didID.String(),
		did.Verifications{auth},
		did.Services{},
		accAddress.String(),
	)

	// Sign the create did message
	txBytes := SignMsg(encCfg, msg, priv1, accNum, accSeq)

	resp, err := txClient.BroadcastTx(
		context.Background(),
		&txtypes.BroadcastTxRequest{
			Mode:    txtypes.BroadcastMode_BROADCAST_MODE_BLOCK,
			TxBytes: txBytes,
		},
	)
	fmt.Println(resp)
	if err != nil {
		panic(err)
	}

	// query the did document to check it was created
	didDoc, err := qc.DidDocument(context.Background(), &did.QueryDidDocumentRequest{Id: didID.String()})
	if err != nil {
		panic(err)
	}
	fmt.Println(didDoc)

	// build a vm for the update
	vmID = didID.NewVerificationMethodID(accAddress.String())
	auth = did.NewVerification(
		did.NewVerificationMethod(
			vmID,
			didID,
			did.NewPublicKeyMultibase(pubKey.Bytes()),
			vmType,
		),
		[]string{did.Authentication},
		nil,
	)

	// build a vm for the update
	bz, _ = hex.DecodeString(TestKey2)
	priv2 := &secp256k1.PrivKey{Key: bz}
	accAddress2 := sdk.AccAddress(priv2.PubKey().Address())
	vmID2 := didID.NewVerificationMethodID(accAddress2.String())
	vmType2 := did.CosmosAccountAddress
	auth2 := did.NewVerification(
		did.NewVerificationMethod(
			vmID2,
			didID,
			did.NewBlockchainAccountID(ChainID, accAddress2.String()),
			vmType2,
		),
		[]string{did.KeyAgreement, did.AssertionMethod, did.Authentication},
		nil,
	)

	// build a vm for the update
	bz, _ = hex.DecodeString(TestKey3)
	priv3 := &secp256k1.PrivKey{Key: bz}
	accAddress3 := sdk.AccAddress(priv3.PubKey().Address())
	pubKey3 := priv3.PubKey()
	vmID3 := didID.NewVerificationMethodID(accAddress3.String())
	vmType3 := did.Bls12381G1Key2020
	auth3 := did.NewVerification(
		did.NewVerificationMethod(
			vmID3,
			didID,
			did.NewPublicKeyMultibase(pubKey3.Bytes()),
			vmType3,
		),
		[]string{did.CapabilityDelegation, did.AssertionMethod},
		nil,
	)

	// add multiple services to the did document
	service1 := did.NewService("service:emti-agent1", "DIDComm", "https://agents.elesto.app.beta.starport.cloud/emti")
	service2 := did.NewService("service:emti-agent2", "DIDComm", "https://agents.elesto.app.beta.starport.cloud/emti")
	service3 := did.NewService("service:emti-agent3", "DIDComm", "https://agents.elesto.app.beta.starport.cloud/emti")
	service4 := did.NewService("service:emti-agent4", "DIDComm", "https://agents.elesto.app.beta.starport.cloud/emti")
	service5 := did.NewService("service:emti-agent5", "DIDComm", "https://agents.elesto.app.beta.starport.cloud/emti")
	service6 := did.NewService("service:emti-agent6", "DIDComm", "https://agents.elesto.app.beta.starport.cloud/emti")

	newDidDoc, err := did.NewDidDocument(
		didID.String(),
		did.WithVerifications(auth, auth2, auth3),
		did.WithServices(service1, service2, service3, service4, service5, service6),
	)
	if err != nil {
		panic(err)
	}

	updateMsg := did.NewMsgUpdateDidDocument(
		&newDidDoc,
		accAddress.String(),
	)

	// Sign the update did message
	txBytes = SignMsg(encCfg, updateMsg, priv1, accNum, accSeq+1)
	if err != nil {
		panic(err)
	}

	_, err = txClient.BroadcastTx(
		context.Background(),
		&txtypes.BroadcastTxRequest{
			Mode:    txtypes.BroadcastMode_BROADCAST_MODE_BLOCK,
			TxBytes: txBytes,
		},
	)
	if err != nil {
		panic(err)
	}

	// query the did document to check it was updated
	didDoc, err = qc.DidDocument(context.Background(), &did.QueryDidDocumentRequest{Id: didID.String()})
	if err != nil {
		panic(err)
	}
	fmt.Println(didDoc)
}
