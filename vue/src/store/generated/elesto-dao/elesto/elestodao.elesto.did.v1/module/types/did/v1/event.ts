/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "elestodao.elesto.did.v1";

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

const baseDidDocumentCreatedEvent: object = { did: "", signer: "" };

export const DidDocumentCreatedEvent = {
  encode(
    message: DidDocumentCreatedEvent,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    if (message.signer !== "") {
      writer.uint32(18).string(message.signer);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): DidDocumentCreatedEvent {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseDidDocumentCreatedEvent,
    } as DidDocumentCreatedEvent;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 2:
          message.signer = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DidDocumentCreatedEvent {
    const message = {
      ...baseDidDocumentCreatedEvent,
    } as DidDocumentCreatedEvent;
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer);
    } else {
      message.signer = "";
    }
    return message;
  },

  toJSON(message: DidDocumentCreatedEvent): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    message.signer !== undefined && (obj.signer = message.signer);
    return obj;
  },

  fromPartial(
    object: DeepPartial<DidDocumentCreatedEvent>
  ): DidDocumentCreatedEvent {
    const message = {
      ...baseDidDocumentCreatedEvent,
    } as DidDocumentCreatedEvent;
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer;
    } else {
      message.signer = "";
    }
    return message;
  },
};

const baseDidDocumentUpdatedEvent: object = { did: "", signer: "" };

export const DidDocumentUpdatedEvent = {
  encode(
    message: DidDocumentUpdatedEvent,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    if (message.signer !== "") {
      writer.uint32(18).string(message.signer);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): DidDocumentUpdatedEvent {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseDidDocumentUpdatedEvent,
    } as DidDocumentUpdatedEvent;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 2:
          message.signer = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DidDocumentUpdatedEvent {
    const message = {
      ...baseDidDocumentUpdatedEvent,
    } as DidDocumentUpdatedEvent;
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer);
    } else {
      message.signer = "";
    }
    return message;
  },

  toJSON(message: DidDocumentUpdatedEvent): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    message.signer !== undefined && (obj.signer = message.signer);
    return obj;
  },

  fromPartial(
    object: DeepPartial<DidDocumentUpdatedEvent>
  ): DidDocumentUpdatedEvent {
    const message = {
      ...baseDidDocumentUpdatedEvent,
    } as DidDocumentUpdatedEvent;
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer;
    } else {
      message.signer = "";
    }
    return message;
  },
};

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
