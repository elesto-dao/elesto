import { Reader, Writer } from "protobufjs/minimal";
import { VerificationMethod, Service } from "../../did/v1/did";
export declare const protobufPackage = "elestodao.elesto.did.v1";
/**
 * Verification is a message that allows to assign a verification method
 * to one or more verification relationships
 */
export interface Verification {
    /**
     * verificationRelationships defines which relationships
     * are allowed to use the verification method
     */
    relationships: string[];
    /** public key associated with the did document. */
    method: VerificationMethod | undefined;
    /** additional contexts (json ld schemas) */
    context: string[];
}
/** MsgCreateDidDocument defines a SDK message for creating a new did. */
export interface MsgCreateDidDocument {
    /** the did */
    id: string;
    /** the list of controller DIDs */
    controllers: string[];
    /** the list of verification methods and relationships */
    verifications: Verification[];
    /** the list of services */
    services: Service[];
    /** address of the account signing the message */
    signer: string;
}
export interface MsgCreateDidDocumentResponse {
}
export interface MsgAddVerification {
    /** the did */
    id: string;
    /** the verification to add */
    verification: Verification | undefined;
    /** address of the account signing the message */
    signer: string;
}
export interface MsgAddVerificationResponse {
}
export interface MsgSetVerificationRelationships {
    /** the did */
    id: string;
    /** the verification method id */
    methodId: string;
    /** the list of relationships to set */
    relationships: string[];
    /** address of the account signing the message */
    signer: string;
}
export interface MsgSetVerificationRelationshipsResponse {
}
export interface MsgRevokeVerification {
    /** the did */
    id: string;
    /** the verification method id */
    methodId: string;
    /** address of the account signing the message */
    signer: string;
}
export interface MsgRevokeVerificationResponse {
}
export interface MsgAddService {
    /** the did */
    id: string;
    /** the service data to add */
    serviceData: Service | undefined;
    /** address of the account signing the message */
    signer: string;
}
export interface MsgAddServiceResponse {
}
export interface MsgDeleteService {
    /** the did */
    id: string;
    /** the service id */
    serviceId: string;
    /** address of the account signing the message */
    signer: string;
}
export interface MsgDeleteServiceResponse {
}
export interface MsgAddController {
    /** the did of the document */
    id: string;
    /** the did to add as a controller of the did document */
    controllerDid: string;
    /** address of the account signing the message */
    signer: string;
}
export interface MsgAddControllerResponse {
}
export interface MsgDeleteController {
    /** the did of the document */
    id: string;
    /** the did to remove from the list of controllers of the did document */
    controllerDid: string;
    /** address of the account signing the message */
    signer: string;
}
export interface MsgDeleteControllerResponse {
}
export declare const Verification: {
    encode(message: Verification, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Verification;
    fromJSON(object: any): Verification;
    toJSON(message: Verification): unknown;
    fromPartial(object: DeepPartial<Verification>): Verification;
};
export declare const MsgCreateDidDocument: {
    encode(message: MsgCreateDidDocument, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateDidDocument;
    fromJSON(object: any): MsgCreateDidDocument;
    toJSON(message: MsgCreateDidDocument): unknown;
    fromPartial(object: DeepPartial<MsgCreateDidDocument>): MsgCreateDidDocument;
};
export declare const MsgCreateDidDocumentResponse: {
    encode(_: MsgCreateDidDocumentResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateDidDocumentResponse;
    fromJSON(_: any): MsgCreateDidDocumentResponse;
    toJSON(_: MsgCreateDidDocumentResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateDidDocumentResponse>): MsgCreateDidDocumentResponse;
};
export declare const MsgAddVerification: {
    encode(message: MsgAddVerification, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddVerification;
    fromJSON(object: any): MsgAddVerification;
    toJSON(message: MsgAddVerification): unknown;
    fromPartial(object: DeepPartial<MsgAddVerification>): MsgAddVerification;
};
export declare const MsgAddVerificationResponse: {
    encode(_: MsgAddVerificationResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddVerificationResponse;
    fromJSON(_: any): MsgAddVerificationResponse;
    toJSON(_: MsgAddVerificationResponse): unknown;
    fromPartial(_: DeepPartial<MsgAddVerificationResponse>): MsgAddVerificationResponse;
};
export declare const MsgSetVerificationRelationships: {
    encode(message: MsgSetVerificationRelationships, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetVerificationRelationships;
    fromJSON(object: any): MsgSetVerificationRelationships;
    toJSON(message: MsgSetVerificationRelationships): unknown;
    fromPartial(object: DeepPartial<MsgSetVerificationRelationships>): MsgSetVerificationRelationships;
};
export declare const MsgSetVerificationRelationshipsResponse: {
    encode(_: MsgSetVerificationRelationshipsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetVerificationRelationshipsResponse;
    fromJSON(_: any): MsgSetVerificationRelationshipsResponse;
    toJSON(_: MsgSetVerificationRelationshipsResponse): unknown;
    fromPartial(_: DeepPartial<MsgSetVerificationRelationshipsResponse>): MsgSetVerificationRelationshipsResponse;
};
export declare const MsgRevokeVerification: {
    encode(message: MsgRevokeVerification, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRevokeVerification;
    fromJSON(object: any): MsgRevokeVerification;
    toJSON(message: MsgRevokeVerification): unknown;
    fromPartial(object: DeepPartial<MsgRevokeVerification>): MsgRevokeVerification;
};
export declare const MsgRevokeVerificationResponse: {
    encode(_: MsgRevokeVerificationResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRevokeVerificationResponse;
    fromJSON(_: any): MsgRevokeVerificationResponse;
    toJSON(_: MsgRevokeVerificationResponse): unknown;
    fromPartial(_: DeepPartial<MsgRevokeVerificationResponse>): MsgRevokeVerificationResponse;
};
export declare const MsgAddService: {
    encode(message: MsgAddService, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddService;
    fromJSON(object: any): MsgAddService;
    toJSON(message: MsgAddService): unknown;
    fromPartial(object: DeepPartial<MsgAddService>): MsgAddService;
};
export declare const MsgAddServiceResponse: {
    encode(_: MsgAddServiceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddServiceResponse;
    fromJSON(_: any): MsgAddServiceResponse;
    toJSON(_: MsgAddServiceResponse): unknown;
    fromPartial(_: DeepPartial<MsgAddServiceResponse>): MsgAddServiceResponse;
};
export declare const MsgDeleteService: {
    encode(message: MsgDeleteService, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteService;
    fromJSON(object: any): MsgDeleteService;
    toJSON(message: MsgDeleteService): unknown;
    fromPartial(object: DeepPartial<MsgDeleteService>): MsgDeleteService;
};
export declare const MsgDeleteServiceResponse: {
    encode(_: MsgDeleteServiceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteServiceResponse;
    fromJSON(_: any): MsgDeleteServiceResponse;
    toJSON(_: MsgDeleteServiceResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteServiceResponse>): MsgDeleteServiceResponse;
};
export declare const MsgAddController: {
    encode(message: MsgAddController, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddController;
    fromJSON(object: any): MsgAddController;
    toJSON(message: MsgAddController): unknown;
    fromPartial(object: DeepPartial<MsgAddController>): MsgAddController;
};
export declare const MsgAddControllerResponse: {
    encode(_: MsgAddControllerResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddControllerResponse;
    fromJSON(_: any): MsgAddControllerResponse;
    toJSON(_: MsgAddControllerResponse): unknown;
    fromPartial(_: DeepPartial<MsgAddControllerResponse>): MsgAddControllerResponse;
};
export declare const MsgDeleteController: {
    encode(message: MsgDeleteController, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteController;
    fromJSON(object: any): MsgDeleteController;
    toJSON(message: MsgDeleteController): unknown;
    fromPartial(object: DeepPartial<MsgDeleteController>): MsgDeleteController;
};
export declare const MsgDeleteControllerResponse: {
    encode(_: MsgDeleteControllerResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteControllerResponse;
    fromJSON(_: any): MsgDeleteControllerResponse;
    toJSON(_: MsgDeleteControllerResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteControllerResponse>): MsgDeleteControllerResponse;
};
/** Msg defines the identity Msg service. */
export interface Msg {
    /** CreateDidDocument defines a method for creating a new identity. */
    CreateDidDocument(request: MsgCreateDidDocument): Promise<MsgCreateDidDocumentResponse>;
    /** AddVerificationMethod adds a new verification method */
    AddVerification(request: MsgAddVerification): Promise<MsgAddVerificationResponse>;
    /** RevokeVerification remove the verification method and all associated verification Relations */
    RevokeVerification(request: MsgRevokeVerification): Promise<MsgRevokeVerificationResponse>;
    /** SetVerificationRelationships overwrite current verification relationships */
    SetVerificationRelationships(request: MsgSetVerificationRelationships): Promise<MsgSetVerificationRelationshipsResponse>;
    /** AddService add a new service */
    AddService(request: MsgAddService): Promise<MsgAddServiceResponse>;
    /** DeleteService delete an existing service */
    DeleteService(request: MsgDeleteService): Promise<MsgDeleteServiceResponse>;
    /** AddService add a new service */
    AddController(request: MsgAddController): Promise<MsgAddControllerResponse>;
    /** DeleteService delete an existing service */
    DeleteController(request: MsgDeleteController): Promise<MsgDeleteControllerResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateDidDocument(request: MsgCreateDidDocument): Promise<MsgCreateDidDocumentResponse>;
    AddVerification(request: MsgAddVerification): Promise<MsgAddVerificationResponse>;
    RevokeVerification(request: MsgRevokeVerification): Promise<MsgRevokeVerificationResponse>;
    SetVerificationRelationships(request: MsgSetVerificationRelationships): Promise<MsgSetVerificationRelationshipsResponse>;
    AddService(request: MsgAddService): Promise<MsgAddServiceResponse>;
    DeleteService(request: MsgDeleteService): Promise<MsgDeleteServiceResponse>;
    AddController(request: MsgAddController): Promise<MsgAddControllerResponse>;
    DeleteController(request: MsgDeleteController): Promise<MsgDeleteControllerResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};