import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "elestodao.elesto.did.v1";
/** DidDocumentCreatedEvent is an event triggered on a DID document creation */
export interface DidDocumentCreatedEvent {
    /** the did being created */
    did: string;
    /** the signer account creating the did */
    signer: string;
}
/** DidDocumentUpdatedEvent is an event triggered on a DID document update */
export interface DidDocumentUpdatedEvent {
    /** the did being updated */
    did: string;
    /** the signer account of the change */
    signer: string;
}
export declare const DidDocumentCreatedEvent: {
    encode(message: DidDocumentCreatedEvent, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): DidDocumentCreatedEvent;
    fromJSON(object: any): DidDocumentCreatedEvent;
    toJSON(message: DidDocumentCreatedEvent): unknown;
    fromPartial(object: DeepPartial<DidDocumentCreatedEvent>): DidDocumentCreatedEvent;
};
export declare const DidDocumentUpdatedEvent: {
    encode(message: DidDocumentUpdatedEvent, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): DidDocumentUpdatedEvent;
    fromJSON(object: any): DidDocumentUpdatedEvent;
    toJSON(message: DidDocumentUpdatedEvent): unknown;
    fromPartial(object: DeepPartial<DidDocumentUpdatedEvent>): DidDocumentUpdatedEvent;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
