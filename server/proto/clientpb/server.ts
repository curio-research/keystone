/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "serverpb";

export enum CMD {
  S2C_Error = 0,
  UNRECOGNIZED = -1,
}

export function cMDFromJSON(object: any): CMD {
  switch (object) {
    case 0:
    case "S2C_Error":
      return CMD.S2C_Error;
    case -1:
    case "UNRECOGNIZED":
    default:
      return CMD.UNRECOGNIZED;
  }
}

export function cMDToJSON(object: CMD): string {
  switch (object) {
    case CMD.S2C_Error:
      return "S2C_Error";
    case CMD.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface S2CErrorMessage {
  Content: string;
}

function createBaseS2CErrorMessage(): S2CErrorMessage {
  return { Content: "" };
}

export const S2CErrorMessage = {
  encode(message: S2CErrorMessage, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.Content !== "") {
      writer.uint32(10).string(message.Content);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): S2CErrorMessage {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseS2CErrorMessage();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.Content = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): S2CErrorMessage {
    return { Content: isSet(object.Content) ? globalThis.String(object.Content) : "" };
  },

  toJSON(message: S2CErrorMessage): unknown {
    const obj: any = {};
    if (message.Content !== "") {
      obj.Content = message.Content;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<S2CErrorMessage>, I>>(base?: I): S2CErrorMessage {
    return S2CErrorMessage.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<S2CErrorMessage>, I>>(object: I): S2CErrorMessage {
    const message = createBaseS2CErrorMessage();
    message.Content = object.Content ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
