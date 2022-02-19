import { Reader, Writer } from "protobufjs/minimal";
import { DidDocument, DidMetadata } from "../../did/v1/did";
export declare const protobufPackage = "elestodao.elesto.did.v1";
/** QueryDidDocumentsRequest is request type for Query/DidDocuments RPC method. */
export interface QueryDidDocumentRequest {
    /** did id used for getting a did document by did id */
    id: string;
}
/** QueryDidDocumentsResponse is response type for the Query/DidDocuments RPC method */
export interface QueryDidDocumentResponse {
    /** Returns a did document */
    didDocument: DidDocument | undefined;
    /** Returns a did documents metadata */
    didMetadata: DidMetadata | undefined;
}
export declare const QueryDidDocumentRequest: {
    encode(message: QueryDidDocumentRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryDidDocumentRequest;
    fromJSON(object: any): QueryDidDocumentRequest;
    toJSON(message: QueryDidDocumentRequest): unknown;
    fromPartial(object: DeepPartial<QueryDidDocumentRequest>): QueryDidDocumentRequest;
};
export declare const QueryDidDocumentResponse: {
    encode(message: QueryDidDocumentResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryDidDocumentResponse;
    fromJSON(object: any): QueryDidDocumentResponse;
    toJSON(message: QueryDidDocumentResponse): unknown;
    fromPartial(object: DeepPartial<QueryDidDocumentResponse>): QueryDidDocumentResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** DidDocument queries a did documents with an id. */
    DidDocument(request: QueryDidDocumentRequest): Promise<QueryDidDocumentResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    DidDocument(request: QueryDidDocumentRequest): Promise<QueryDidDocumentResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
