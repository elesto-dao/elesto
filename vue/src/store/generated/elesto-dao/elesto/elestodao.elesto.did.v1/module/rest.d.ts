export interface ProtobufAny {
    "@type"?: string;
}
export interface RpcStatus {
    /** @format int32 */
    code?: number;
    message?: string;
    details?: ProtobufAny[];
}
/**
 * DidDocument represents a dencentralised identifer.
 */
export interface V1DidDocument {
    /** @context is spec for did document. */
    context?: string[];
    /** id represents the id for the did document. */
    id?: string;
    controller?: string[];
    verificationMethod?: V1VerificationMethod[];
    service?: V1Service[];
    authentication?: string[];
    assertionMethod?: string[];
    keyAgreement?: string[];
    capabilityInvocation?: string[];
    capabilityDelegation?: string[];
}
export interface V1DidMetadata {
    versionId?: string;
    /** @format date-time */
    created?: string;
    /** @format date-time */
    updated?: string;
    deactivated?: boolean;
}
export declare type V1MsgAddControllerResponse = object;
export declare type V1MsgAddServiceResponse = object;
export declare type V1MsgAddVerificationResponse = object;
export declare type V1MsgCreateDidDocumentResponse = object;
export declare type V1MsgDeleteControllerResponse = object;
export declare type V1MsgDeleteServiceResponse = object;
export declare type V1MsgRevokeVerificationResponse = object;
export declare type V1MsgSetVerificationRelationshipsResponse = object;
export interface V1PublicKeyJwk {
    Kid?: string;
    Crv?: string;
    X?: string;
    Y?: string;
    Kty?: string;
}
export interface V1QueryDidDocumentResponse {
    /** DidDocument represents a dencentralised identifer. */
    didDocument?: V1DidDocument;
    didMetadata?: V1DidMetadata;
}
export interface V1Service {
    id?: string;
    type?: string;
    serviceEndpoint?: string;
}
export interface V1Verification {
    relationships?: string[];
    method?: V1VerificationMethod;
    context?: string[];
}
export interface V1VerificationMethod {
    id?: string;
    type?: string;
    controller?: string;
    blockchainAccountID?: string;
    publicKeyHex?: string;
    publicKeyMultibase?: string;
    PublicKeyJwk?: V1PublicKeyJwk;
}
export declare type QueryParamsType = Record<string | number, any>;
export declare type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;
export interface FullRequestParams extends Omit<RequestInit, "body"> {
    /** set parameter to `true` for call `securityWorker` for this request */
    secure?: boolean;
    /** request path */
    path: string;
    /** content type of request body */
    type?: ContentType;
    /** query params */
    query?: QueryParamsType;
    /** format of response (i.e. response.json() -> format: "json") */
    format?: keyof Omit<Body, "body" | "bodyUsed">;
    /** request body */
    body?: unknown;
    /** base url */
    baseUrl?: string;
    /** request cancellation token */
    cancelToken?: CancelToken;
}
export declare type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;
export interface ApiConfig<SecurityDataType = unknown> {
    baseUrl?: string;
    baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
    securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}
export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
    data: D;
    error: E;
}
declare type CancelToken = Symbol | string | number;
export declare enum ContentType {
    Json = "application/json",
    FormData = "multipart/form-data",
    UrlEncoded = "application/x-www-form-urlencoded"
}
export declare class HttpClient<SecurityDataType = unknown> {
    baseUrl: string;
    private securityData;
    private securityWorker;
    private abortControllers;
    private baseApiParams;
    constructor(apiConfig?: ApiConfig<SecurityDataType>);
    setSecurityData: (data: SecurityDataType) => void;
    private addQueryParam;
    protected toQueryString(rawQuery?: QueryParamsType): string;
    protected addQueryParams(rawQuery?: QueryParamsType): string;
    private contentFormatters;
    private mergeRequestParams;
    private createAbortSignal;
    abortRequest: (cancelToken: CancelToken) => void;
    request: <T = any, E = any>({ body, secure, path, type, query, format, baseUrl, cancelToken, ...params }: FullRequestParams) => Promise<HttpResponse<T, E>>;
}
/**
 * @title did/v1/did.proto
 * @version version not set
 */
export declare class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryDidDocument
     * @summary DidDocument queries a did documents with an id.
     * @request GET:/elesto/did/dids/{id}
     */
    queryDidDocument: (id: string, params?: RequestParams) => Promise<HttpResponse<V1QueryDidDocumentResponse, RpcStatus>>;
}
export {};
