/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { DidDocument, DidMetadata } from "../../did/v1/did";
export const protobufPackage = "elestodao.elesto.did.v1";
const baseQueryDidDocumentRequest = { id: "" };
export const QueryDidDocumentRequest = {
    encode(message, writer = Writer.create()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryDidDocumentRequest,
        };
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
    fromJSON(object) {
        const message = {
            ...baseQueryDidDocumentRequest,
        };
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryDidDocumentRequest,
        };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        return message;
    },
};
const baseQueryDidDocumentResponse = {};
export const QueryDidDocumentResponse = {
    encode(message, writer = Writer.create()) {
        if (message.didDocument !== undefined) {
            DidDocument.encode(message.didDocument, writer.uint32(10).fork()).ldelim();
        }
        if (message.didMetadata !== undefined) {
            DidMetadata.encode(message.didMetadata, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryDidDocumentResponse,
        };
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
    fromJSON(object) {
        const message = {
            ...baseQueryDidDocumentResponse,
        };
        if (object.didDocument !== undefined && object.didDocument !== null) {
            message.didDocument = DidDocument.fromJSON(object.didDocument);
        }
        else {
            message.didDocument = undefined;
        }
        if (object.didMetadata !== undefined && object.didMetadata !== null) {
            message.didMetadata = DidMetadata.fromJSON(object.didMetadata);
        }
        else {
            message.didMetadata = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
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
    fromPartial(object) {
        const message = {
            ...baseQueryDidDocumentResponse,
        };
        if (object.didDocument !== undefined && object.didDocument !== null) {
            message.didDocument = DidDocument.fromPartial(object.didDocument);
        }
        else {
            message.didDocument = undefined;
        }
        if (object.didMetadata !== undefined && object.didMetadata !== null) {
            message.didMetadata = DidMetadata.fromPartial(object.didMetadata);
        }
        else {
            message.didMetadata = undefined;
        }
        return message;
    },
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    DidDocument(request) {
        const data = QueryDidDocumentRequest.encode(request).finish();
        const promise = this.rpc.request("elestodao.elesto.did.v1.Query", "DidDocument", data);
        return promise.then((data) => QueryDidDocumentResponse.decode(new Reader(data)));
    }
}
