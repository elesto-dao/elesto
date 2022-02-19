import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgAddVerification } from "./types/did/v1/tx";
import { MsgAddController } from "./types/did/v1/tx";
import { MsgDeleteService } from "./types/did/v1/tx";
import { MsgCreateDidDocument } from "./types/did/v1/tx";
import { MsgRevokeVerification } from "./types/did/v1/tx";
import { MsgAddService } from "./types/did/v1/tx";
import { MsgSetVerificationRelationships } from "./types/did/v1/tx";
import { MsgDeleteController } from "./types/did/v1/tx";
export declare const MissingWalletError: Error;
export declare const registry: Registry;
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }?: SignAndBroadcastOptions) => any;
    msgAddVerification: (data: MsgAddVerification) => EncodeObject;
    msgAddController: (data: MsgAddController) => EncodeObject;
    msgDeleteService: (data: MsgDeleteService) => EncodeObject;
    msgCreateDidDocument: (data: MsgCreateDidDocument) => EncodeObject;
    msgRevokeVerification: (data: MsgRevokeVerification) => EncodeObject;
    msgAddService: (data: MsgAddService) => EncodeObject;
    msgSetVerificationRelationships: (data: MsgSetVerificationRelationships) => EncodeObject;
    msgDeleteController: (data: MsgDeleteController) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
