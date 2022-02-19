// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgAddVerification } from "./types/did/v1/tx";
import { MsgAddController } from "./types/did/v1/tx";
import { MsgDeleteService } from "./types/did/v1/tx";
import { MsgCreateDidDocument } from "./types/did/v1/tx";
import { MsgRevokeVerification } from "./types/did/v1/tx";
import { MsgAddService } from "./types/did/v1/tx";
import { MsgSetVerificationRelationships } from "./types/did/v1/tx";
import { MsgDeleteController } from "./types/did/v1/tx";


const types = [
  ["/elestodao.elesto.did.v1.MsgAddVerification", MsgAddVerification],
  ["/elestodao.elesto.did.v1.MsgAddController", MsgAddController],
  ["/elestodao.elesto.did.v1.MsgDeleteService", MsgDeleteService],
  ["/elestodao.elesto.did.v1.MsgCreateDidDocument", MsgCreateDidDocument],
  ["/elestodao.elesto.did.v1.MsgRevokeVerification", MsgRevokeVerification],
  ["/elestodao.elesto.did.v1.MsgAddService", MsgAddService],
  ["/elestodao.elesto.did.v1.MsgSetVerificationRelationships", MsgSetVerificationRelationships],
  ["/elestodao.elesto.did.v1.MsgDeleteController", MsgDeleteController],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgAddVerification: (data: MsgAddVerification): EncodeObject => ({ typeUrl: "/elestodao.elesto.did.v1.MsgAddVerification", value: MsgAddVerification.fromPartial( data ) }),
    msgAddController: (data: MsgAddController): EncodeObject => ({ typeUrl: "/elestodao.elesto.did.v1.MsgAddController", value: MsgAddController.fromPartial( data ) }),
    msgDeleteService: (data: MsgDeleteService): EncodeObject => ({ typeUrl: "/elestodao.elesto.did.v1.MsgDeleteService", value: MsgDeleteService.fromPartial( data ) }),
    msgCreateDidDocument: (data: MsgCreateDidDocument): EncodeObject => ({ typeUrl: "/elestodao.elesto.did.v1.MsgCreateDidDocument", value: MsgCreateDidDocument.fromPartial( data ) }),
    msgRevokeVerification: (data: MsgRevokeVerification): EncodeObject => ({ typeUrl: "/elestodao.elesto.did.v1.MsgRevokeVerification", value: MsgRevokeVerification.fromPartial( data ) }),
    msgAddService: (data: MsgAddService): EncodeObject => ({ typeUrl: "/elestodao.elesto.did.v1.MsgAddService", value: MsgAddService.fromPartial( data ) }),
    msgSetVerificationRelationships: (data: MsgSetVerificationRelationships): EncodeObject => ({ typeUrl: "/elestodao.elesto.did.v1.MsgSetVerificationRelationships", value: MsgSetVerificationRelationships.fromPartial( data ) }),
    msgDeleteController: (data: MsgDeleteController): EncodeObject => ({ typeUrl: "/elestodao.elesto.did.v1.MsgDeleteController", value: MsgDeleteController.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
