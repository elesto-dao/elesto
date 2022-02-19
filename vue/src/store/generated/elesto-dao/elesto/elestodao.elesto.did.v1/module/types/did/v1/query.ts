/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { DidDocument, DidMetadata } from "../../did/v1/did";

export const protobufPackage = "elestodao.elesto.did.v1";

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

const baseQueryDidDocumentRequest: object = { id: "" };

export const QueryDidDocumentRequest = {
  encode(
    message: QueryDidDocumentRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryDidDocumentRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryDidDocumentRequest,
    } as QueryDidDocumentRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDidDocumentRequest {
    const message = {
      ...baseQueryDidDocumentRequest,
    } as QueryDidDocumentRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    return message;
  },

  toJSON(message: QueryDidDocumentRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryDidDocumentRequest>
  ): QueryDidDocumentRequest {
    const message = {
      ...baseQueryDidDocumentRequest,
    } as QueryDidDocumentRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    return message;
  },
};

const baseQueryDidDocumentResponse: object = {};

export const QueryDidDocumentResponse = {
  encode(
    message: QueryDidDocumentResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.didDocument !== undefined) {
      DidDocument.encode(
        message.didDocument,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.didMetadata !== undefined) {
      DidMetadata.encode(
        message.didMetadata,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryDidDocumentResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryDidDocumentResponse,
    } as QueryDidDocumentResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.didDocument = DidDocument.decode(reader, reader.uint32());
          break;
        case 2:
          message.didMetadata = DidMetadata.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDidDocumentResponse {
    const message = {
      ...baseQueryDidDocumentResponse,
    } as QueryDidDocumentResponse;
    if (object.didDocument !== undefined && object.didDocument !== null) {
      message.didDocument = DidDocument.fromJSON(object.didDocument);
    } else {
      message.didDocument = undefined;
    }
    if (object.didMetadata !== undefined && object.didMetadata !== null) {
      message.didMetadata = DidMetadata.fromJSON(object.didMetadata);
    } else {
      message.didMetadata = undefined;
    }
    return message;
  },

  toJSON(message: QueryDidDocumentResponse): unknown {
    const obj: any = {};
    message.didDocument !== undefined &&
      (obj.didDocument = message.didDocument
        ? DidDocument.toJSON(message.didDocument)
        : undefined);
    message.didMetadata !== undefined &&
      (obj.didMetadata = message.didMetadata
        ? DidMetadata.toJSON(message.didMetadata)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryDidDocumentResponse>
  ): QueryDidDocumentResponse {
    const message = {
      ...baseQueryDidDocumentResponse,
    } as QueryDidDocumentResponse;
    if (object.didDocument !== undefined && object.didDocument !== null) {
      message.didDocument = DidDocument.fromPartial(object.didDocument);
    } else {
      message.didDocument = undefined;
    }
    if (object.didMetadata !== undefined && object.didMetadata !== null) {
      message.didMetadata = DidMetadata.fromPartial(object.didMetadata);
    } else {
      message.didMetadata = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** DidDocument queries a did documents with an id. */
  DidDocument(
    request: QueryDidDocumentRequest
  ): Promise<QueryDidDocumentResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  DidDocument(
    request: QueryDidDocumentRequest
  ): Promise<QueryDidDocumentResponse> {
    const data = QueryDidDocumentRequest.encode(request).finish();
    const promise = this.rpc.request(
      "elestodao.elesto.did.v1.Query",
      "DidDocument",
      data
    );
    return promise.then((data) =>
      QueryDidDocumentResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
