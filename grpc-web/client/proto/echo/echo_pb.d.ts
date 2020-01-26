import * as jspb from "google-protobuf"

export class EchoRequest extends jspb.Message {
  getBody(): string;
  setBody(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EchoRequest.AsObject;
  static toObject(includeInstance: boolean, msg: EchoRequest): EchoRequest.AsObject;
  static serializeBinaryToWriter(message: EchoRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EchoRequest;
  static deserializeBinaryFromReader(message: EchoRequest, reader: jspb.BinaryReader): EchoRequest;
}

export namespace EchoRequest {
  export type AsObject = {
    body: string,
  }
}

export class EchoResponse extends jspb.Message {
  getBody(): string;
  setBody(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EchoResponse.AsObject;
  static toObject(includeInstance: boolean, msg: EchoResponse): EchoResponse.AsObject;
  static serializeBinaryToWriter(message: EchoResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EchoResponse;
  static deserializeBinaryFromReader(message: EchoResponse, reader: jspb.BinaryReader): EchoResponse;
}

export namespace EchoResponse {
  export type AsObject = {
    body: string,
  }
}

export class EchoReverseRequest extends jspb.Message {
  getBody(): string;
  setBody(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EchoReverseRequest.AsObject;
  static toObject(includeInstance: boolean, msg: EchoReverseRequest): EchoReverseRequest.AsObject;
  static serializeBinaryToWriter(message: EchoReverseRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EchoReverseRequest;
  static deserializeBinaryFromReader(message: EchoReverseRequest, reader: jspb.BinaryReader): EchoReverseRequest;
}

export namespace EchoReverseRequest {
  export type AsObject = {
    body: string,
  }
}

export class EchoReverseResponse extends jspb.Message {
  getBody(): string;
  setBody(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EchoReverseResponse.AsObject;
  static toObject(includeInstance: boolean, msg: EchoReverseResponse): EchoReverseResponse.AsObject;
  static serializeBinaryToWriter(message: EchoReverseResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EchoReverseResponse;
  static deserializeBinaryFromReader(message: EchoReverseResponse, reader: jspb.BinaryReader): EchoReverseResponse;
}

export namespace EchoReverseResponse {
  export type AsObject = {
    body: string,
  }
}

