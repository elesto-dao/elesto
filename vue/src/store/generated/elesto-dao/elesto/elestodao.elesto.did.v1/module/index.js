// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
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
export const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    let client;
    if (addr) {
        client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    }
    else {
        client = await SigningStargateClient.offline(wallet, { registry });
    }
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgAddVerification: (data) => ({ typeUrl: "/elestodao.elesto.did.v1.MsgAddVerification", value: MsgAddVerification.fromPartial(data) }),
        msgAddController: (data) => ({ typeUrl: "/elestodao.elesto.did.v1.MsgAddController", value: MsgAddController.fromPartial(data) }),
        msgDeleteService: (data) => ({ typeUrl: "/elestodao.elesto.did.v1.MsgDeleteService", value: MsgDeleteService.fromPartial(data) }),
        msgCreateDidDocument: (data) => ({ typeUrl: "/elestodao.elesto.did.v1.MsgCreateDidDocument", value: MsgCreateDidDocument.fromPartial(data) }),
        msgRevokeVerification: (data) => ({ typeUrl: "/elestodao.elesto.did.v1.MsgRevokeVerification", value: MsgRevokeVerification.fromPartial(data) }),
        msgAddService: (data) => ({ typeUrl: "/elestodao.elesto.did.v1.MsgAddService", value: MsgAddService.fromPartial(data) }),
        msgSetVerificationRelationships: (data) => ({ typeUrl: "/elestodao.elesto.did.v1.MsgSetVerificationRelationships", value: MsgSetVerificationRelationships.fromPartial(data) }),
        msgDeleteController: (data) => ({ typeUrl: "/elestodao.elesto.did.v1.MsgDeleteController", value: MsgDeleteController.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
