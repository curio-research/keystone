import * as $protobuf from "protobufjs";
import Long = require("long");
/** Namespace pb_game. */
export namespace pb_game {

    /** Properties of a C2S_EstablishPlayer. */
    interface IC2S_EstablishPlayer {

        /** C2S_EstablishPlayer PlayerId */
        PlayerId?: (number|Long|null);
    }

    /** Represents a C2S_EstablishPlayer. */
    class C2S_EstablishPlayer implements IC2S_EstablishPlayer {

        /**
         * Constructs a new C2S_EstablishPlayer.
         * @param [properties] Properties to set
         */
        constructor(properties?: pb_game.IC2S_EstablishPlayer);

        /** C2S_EstablishPlayer PlayerId. */
        public PlayerId: (number|Long);

        /**
         * Creates a new C2S_EstablishPlayer instance using the specified properties.
         * @param [properties] Properties to set
         * @returns C2S_EstablishPlayer instance
         */
        public static create(properties?: pb_game.IC2S_EstablishPlayer): pb_game.C2S_EstablishPlayer;

        /**
         * Encodes the specified C2S_EstablishPlayer message. Does not implicitly {@link pb_game.C2S_EstablishPlayer.verify|verify} messages.
         * @param message C2S_EstablishPlayer message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: pb_game.IC2S_EstablishPlayer, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified C2S_EstablishPlayer message, length delimited. Does not implicitly {@link pb_game.C2S_EstablishPlayer.verify|verify} messages.
         * @param message C2S_EstablishPlayer message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: pb_game.IC2S_EstablishPlayer, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a C2S_EstablishPlayer message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns C2S_EstablishPlayer
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): pb_game.C2S_EstablishPlayer;

        /**
         * Decodes a C2S_EstablishPlayer message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns C2S_EstablishPlayer
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): pb_game.C2S_EstablishPlayer;

        /**
         * Verifies a C2S_EstablishPlayer message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a C2S_EstablishPlayer message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns C2S_EstablishPlayer
         */
        public static fromObject(object: { [k: string]: any }): pb_game.C2S_EstablishPlayer;

        /**
         * Creates a plain object from a C2S_EstablishPlayer message. Also converts values to other types if specified.
         * @param message C2S_EstablishPlayer
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: pb_game.C2S_EstablishPlayer, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this C2S_EstablishPlayer to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for C2S_EstablishPlayer
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a C2S_GameState. */
    interface IC2S_GameState {

        /** C2S_GameState IdentityPayload */
        IdentityPayload?: (pb_base.IIdentityPayload|null);
    }

    /** Represents a C2S_GameState. */
    class C2S_GameState implements IC2S_GameState {

        /**
         * Constructs a new C2S_GameState.
         * @param [properties] Properties to set
         */
        constructor(properties?: pb_game.IC2S_GameState);

        /** C2S_GameState IdentityPayload. */
        public IdentityPayload?: (pb_base.IIdentityPayload|null);

        /**
         * Creates a new C2S_GameState instance using the specified properties.
         * @param [properties] Properties to set
         * @returns C2S_GameState instance
         */
        public static create(properties?: pb_game.IC2S_GameState): pb_game.C2S_GameState;

        /**
         * Encodes the specified C2S_GameState message. Does not implicitly {@link pb_game.C2S_GameState.verify|verify} messages.
         * @param message C2S_GameState message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: pb_game.IC2S_GameState, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified C2S_GameState message, length delimited. Does not implicitly {@link pb_game.C2S_GameState.verify|verify} messages.
         * @param message C2S_GameState message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: pb_game.IC2S_GameState, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a C2S_GameState message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns C2S_GameState
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): pb_game.C2S_GameState;

        /**
         * Decodes a C2S_GameState message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns C2S_GameState
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): pb_game.C2S_GameState;

        /**
         * Verifies a C2S_GameState message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a C2S_GameState message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns C2S_GameState
         */
        public static fromObject(object: { [k: string]: any }): pb_game.C2S_GameState;

        /**
         * Creates a plain object from a C2S_GameState message. Also converts values to other types if specified.
         * @param message C2S_GameState
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: pb_game.C2S_GameState, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this C2S_GameState to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for C2S_GameState
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a C2S_Ping. */
    interface IC2S_Ping {

        /** C2S_Ping timeStamp */
        timeStamp?: (number|Long|null);

        /** C2S_Ping IdentityPayload */
        IdentityPayload?: (pb_base.IIdentityPayload|null);
    }

    /** Represents a C2S_Ping. */
    class C2S_Ping implements IC2S_Ping {

        /**
         * Constructs a new C2S_Ping.
         * @param [properties] Properties to set
         */
        constructor(properties?: pb_game.IC2S_Ping);

        /** C2S_Ping timeStamp. */
        public timeStamp: (number|Long);

        /** C2S_Ping IdentityPayload. */
        public IdentityPayload?: (pb_base.IIdentityPayload|null);

        /**
         * Creates a new C2S_Ping instance using the specified properties.
         * @param [properties] Properties to set
         * @returns C2S_Ping instance
         */
        public static create(properties?: pb_game.IC2S_Ping): pb_game.C2S_Ping;

        /**
         * Encodes the specified C2S_Ping message. Does not implicitly {@link pb_game.C2S_Ping.verify|verify} messages.
         * @param message C2S_Ping message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: pb_game.IC2S_Ping, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified C2S_Ping message, length delimited. Does not implicitly {@link pb_game.C2S_Ping.verify|verify} messages.
         * @param message C2S_Ping message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: pb_game.IC2S_Ping, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a C2S_Ping message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns C2S_Ping
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): pb_game.C2S_Ping;

        /**
         * Decodes a C2S_Ping message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns C2S_Ping
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): pb_game.C2S_Ping;

        /**
         * Verifies a C2S_Ping message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a C2S_Ping message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns C2S_Ping
         */
        public static fromObject(object: { [k: string]: any }): pb_game.C2S_Ping;

        /**
         * Creates a plain object from a C2S_Ping message. Also converts values to other types if specified.
         * @param message C2S_Ping
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: pb_game.C2S_Ping, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this C2S_Ping to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for C2S_Ping
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a S2C_EstablishPlayer. */
    interface IS2C_EstablishPlayer {
    }

    /** Represents a S2C_EstablishPlayer. */
    class S2C_EstablishPlayer implements IS2C_EstablishPlayer {

        /**
         * Constructs a new S2C_EstablishPlayer.
         * @param [properties] Properties to set
         */
        constructor(properties?: pb_game.IS2C_EstablishPlayer);

        /**
         * Creates a new S2C_EstablishPlayer instance using the specified properties.
         * @param [properties] Properties to set
         * @returns S2C_EstablishPlayer instance
         */
        public static create(properties?: pb_game.IS2C_EstablishPlayer): pb_game.S2C_EstablishPlayer;

        /**
         * Encodes the specified S2C_EstablishPlayer message. Does not implicitly {@link pb_game.S2C_EstablishPlayer.verify|verify} messages.
         * @param message S2C_EstablishPlayer message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: pb_game.IS2C_EstablishPlayer, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified S2C_EstablishPlayer message, length delimited. Does not implicitly {@link pb_game.S2C_EstablishPlayer.verify|verify} messages.
         * @param message S2C_EstablishPlayer message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: pb_game.IS2C_EstablishPlayer, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a S2C_EstablishPlayer message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns S2C_EstablishPlayer
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): pb_game.S2C_EstablishPlayer;

        /**
         * Decodes a S2C_EstablishPlayer message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns S2C_EstablishPlayer
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): pb_game.S2C_EstablishPlayer;

        /**
         * Verifies a S2C_EstablishPlayer message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a S2C_EstablishPlayer message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns S2C_EstablishPlayer
         */
        public static fromObject(object: { [k: string]: any }): pb_game.S2C_EstablishPlayer;

        /**
         * Creates a plain object from a S2C_EstablishPlayer message. Also converts values to other types if specified.
         * @param message S2C_EstablishPlayer
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: pb_game.S2C_EstablishPlayer, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this S2C_EstablishPlayer to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for S2C_EstablishPlayer
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a S2C_GameState. */
    interface IS2C_GameState {

        /** S2C_GameState Round */
        Round?: (pb_base.IRound|null);

        /** S2C_GameState PlayerList */
        PlayerList?: (pb_base.IPlayer[]|null);

        /** S2C_GameState TileList */
        TileList?: (pb_base.ITile[]|null);

        /** S2C_GameState TroopStackList */
        TroopStackList?: (pb_base.ITroopStack[]|null);

        /** S2C_GameState CapitalList */
        CapitalList?: (pb_base.ICapital[]|null);
    }

    /** Represents a S2C_GameState. */
    class S2C_GameState implements IS2C_GameState {

        /**
         * Constructs a new S2C_GameState.
         * @param [properties] Properties to set
         */
        constructor(properties?: pb_game.IS2C_GameState);

        /** S2C_GameState Round. */
        public Round?: (pb_base.IRound|null);

        /** S2C_GameState PlayerList. */
        public PlayerList: pb_base.IPlayer[];

        /** S2C_GameState TileList. */
        public TileList: pb_base.ITile[];

        /** S2C_GameState TroopStackList. */
        public TroopStackList: pb_base.ITroopStack[];

        /** S2C_GameState CapitalList. */
        public CapitalList: pb_base.ICapital[];

        /**
         * Creates a new S2C_GameState instance using the specified properties.
         * @param [properties] Properties to set
         * @returns S2C_GameState instance
         */
        public static create(properties?: pb_game.IS2C_GameState): pb_game.S2C_GameState;

        /**
         * Encodes the specified S2C_GameState message. Does not implicitly {@link pb_game.S2C_GameState.verify|verify} messages.
         * @param message S2C_GameState message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: pb_game.IS2C_GameState, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified S2C_GameState message, length delimited. Does not implicitly {@link pb_game.S2C_GameState.verify|verify} messages.
         * @param message S2C_GameState message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: pb_game.IS2C_GameState, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a S2C_GameState message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns S2C_GameState
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): pb_game.S2C_GameState;

        /**
         * Decodes a S2C_GameState message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns S2C_GameState
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): pb_game.S2C_GameState;

        /**
         * Verifies a S2C_GameState message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a S2C_GameState message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns S2C_GameState
         */
        public static fromObject(object: { [k: string]: any }): pb_game.S2C_GameState;

        /**
         * Creates a plain object from a S2C_GameState message. Also converts values to other types if specified.
         * @param message S2C_GameState
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: pb_game.S2C_GameState, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this S2C_GameState to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for S2C_GameState
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a S2C_Ping. */
    interface IS2C_Ping {

        /** S2C_Ping TimeStamp */
        TimeStamp?: (number|Long|null);
    }

    /** Represents a S2C_Ping. */
    class S2C_Ping implements IS2C_Ping {

        /**
         * Constructs a new S2C_Ping.
         * @param [properties] Properties to set
         */
        constructor(properties?: pb_game.IS2C_Ping);

        /** S2C_Ping TimeStamp. */
        public TimeStamp: (number|Long);

        /**
         * Creates a new S2C_Ping instance using the specified properties.
         * @param [properties] Properties to set
         * @returns S2C_Ping instance
         */
        public static create(properties?: pb_game.IS2C_Ping): pb_game.S2C_Ping;

        /**
         * Encodes the specified S2C_Ping message. Does not implicitly {@link pb_game.S2C_Ping.verify|verify} messages.
         * @param message S2C_Ping message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: pb_game.IS2C_Ping, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified S2C_Ping message, length delimited. Does not implicitly {@link pb_game.S2C_Ping.verify|verify} messages.
         * @param message S2C_Ping message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: pb_game.IS2C_Ping, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a S2C_Ping message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns S2C_Ping
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): pb_game.S2C_Ping;

        /**
         * Decodes a S2C_Ping message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns S2C_Ping
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): pb_game.S2C_Ping;

        /**
         * Verifies a S2C_Ping message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a S2C_Ping message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns S2C_Ping
         */
        public static fromObject(object: { [k: string]: any }): pb_game.S2C_Ping;

        /**
         * Creates a plain object from a S2C_Ping message. Also converts values to other types if specified.
         * @param message S2C_Ping
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: pb_game.S2C_Ping, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this S2C_Ping to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for S2C_Ping
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a S2C_ServerMessage. */
    interface IS2C_ServerMessage {

        /** S2C_ServerMessage Content */
        Content?: (string|null);
    }

    /** Represents a S2C_ServerMessage. */
    class S2C_ServerMessage implements IS2C_ServerMessage {

        /**
         * Constructs a new S2C_ServerMessage.
         * @param [properties] Properties to set
         */
        constructor(properties?: pb_game.IS2C_ServerMessage);

        /** S2C_ServerMessage Content. */
        public Content: string;

        /**
         * Creates a new S2C_ServerMessage instance using the specified properties.
         * @param [properties] Properties to set
         * @returns S2C_ServerMessage instance
         */
        public static create(properties?: pb_game.IS2C_ServerMessage): pb_game.S2C_ServerMessage;

        /**
         * Encodes the specified S2C_ServerMessage message. Does not implicitly {@link pb_game.S2C_ServerMessage.verify|verify} messages.
         * @param message S2C_ServerMessage message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: pb_game.IS2C_ServerMessage, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified S2C_ServerMessage message, length delimited. Does not implicitly {@link pb_game.S2C_ServerMessage.verify|verify} messages.
         * @param message S2C_ServerMessage message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: pb_game.IS2C_ServerMessage, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a S2C_ServerMessage message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns S2C_ServerMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): pb_game.S2C_ServerMessage;

        /**
         * Decodes a S2C_ServerMessage message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns S2C_ServerMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): pb_game.S2C_ServerMessage;

        /**
         * Verifies a S2C_ServerMessage message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a S2C_ServerMessage message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns S2C_ServerMessage
         */
        public static fromObject(object: { [k: string]: any }): pb_game.S2C_ServerMessage;

        /**
         * Creates a plain object from a S2C_ServerMessage message. Also converts values to other types if specified.
         * @param message S2C_ServerMessage
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: pb_game.S2C_ServerMessage, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this S2C_ServerMessage to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for S2C_ServerMessage
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }
}

/** Namespace pb_base. */
export namespace pb_base {

    /** TroopStackType enum. */
    enum TroopStackType {
        Infantry = 0,
        Tank = 1,
        Fort = 2,
        Plane = 3
    }

    /** Properties of a Position. */
    interface IPosition {

        /** Position X */
        X?: (number|null);

        /** Position Y */
        Y?: (number|null);
    }

    /** Represents a Position. */
    class Position implements IPosition {

        /**
         * Constructs a new Position.
         * @param [properties] Properties to set
         */
        constructor(properties?: pb_base.IPosition);

        /** Position X. */
        public X: number;

        /** Position Y. */
        public Y: number;

        /**
         * Creates a new Position instance using the specified properties.
         * @param [properties] Properties to set
         * @returns Position instance
         */
        public static create(properties?: pb_base.IPosition): pb_base.Position;

        /**
         * Encodes the specified Position message. Does not implicitly {@link pb_base.Position.verify|verify} messages.
         * @param message Position message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: pb_base.IPosition, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified Position message, length delimited. Does not implicitly {@link pb_base.Position.verify|verify} messages.
         * @param message Position message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: pb_base.IPosition, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Position message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Position
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): pb_base.Position;

        /**
         * Decodes a Position message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns Position
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): pb_base.Position;

        /**
         * Verifies a Position message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a Position message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Position
         */
        public static fromObject(object: { [k: string]: any }): pb_base.Position;

        /**
         * Creates a plain object from a Position message. Also converts values to other types if specified.
         * @param message Position
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: pb_base.Position, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Position to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for Position
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a Round. */
    interface IRound {

        /** Round Count */
        Count?: (number|null);
    }

    /** Represents a Round. */
    class Round implements IRound {

        /**
         * Constructs a new Round.
         * @param [properties] Properties to set
         */
        constructor(properties?: pb_base.IRound);

        /** Round Count. */
        public Count: number;

        /**
         * Creates a new Round instance using the specified properties.
         * @param [properties] Properties to set
         * @returns Round instance
         */
        public static create(properties?: pb_base.IRound): pb_base.Round;

        /**
         * Encodes the specified Round message. Does not implicitly {@link pb_base.Round.verify|verify} messages.
         * @param message Round message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: pb_base.IRound, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified Round message, length delimited. Does not implicitly {@link pb_base.Round.verify|verify} messages.
         * @param message Round message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: pb_base.IRound, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Round message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Round
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): pb_base.Round;

        /**
         * Decodes a Round message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns Round
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): pb_base.Round;

        /**
         * Verifies a Round message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a Round message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Round
         */
        public static fromObject(object: { [k: string]: any }): pb_base.Round;

        /**
         * Creates a plain object from a Round message. Also converts values to other types if specified.
         * @param message Round
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: pb_base.Round, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Round to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for Round
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a Player. */
    interface IPlayer {

        /** Player Id */
        Id?: (number|Long|null);

        /** Player Name */
        Name?: (string|null);

        /** Player Role */
        Role?: (string|null);
    }

    /** Represents a Player. */
    class Player implements IPlayer {

        /**
         * Constructs a new Player.
         * @param [properties] Properties to set
         */
        constructor(properties?: pb_base.IPlayer);

        /** Player Id. */
        public Id: (number|Long);

        /** Player Name. */
        public Name: string;

        /** Player Role. */
        public Role: string;

        /**
         * Creates a new Player instance using the specified properties.
         * @param [properties] Properties to set
         * @returns Player instance
         */
        public static create(properties?: pb_base.IPlayer): pb_base.Player;

        /**
         * Encodes the specified Player message. Does not implicitly {@link pb_base.Player.verify|verify} messages.
         * @param message Player message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: pb_base.IPlayer, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified Player message, length delimited. Does not implicitly {@link pb_base.Player.verify|verify} messages.
         * @param message Player message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: pb_base.IPlayer, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Player message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Player
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): pb_base.Player;

        /**
         * Decodes a Player message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns Player
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): pb_base.Player;

        /**
         * Verifies a Player message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a Player message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Player
         */
        public static fromObject(object: { [k: string]: any }): pb_base.Player;

        /**
         * Creates a plain object from a Player message. Also converts values to other types if specified.
         * @param message Player
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: pb_base.Player, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Player to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for Player
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a Tile. */
    interface ITile {

        /** Tile Id */
        Id?: (number|Long|null);

        /** Tile OwnerId */
        OwnerId?: (number|Long|null);

        /** Tile Level */
        Level?: (number|null);

        /** Tile Position */
        Position?: (pb_base.IPosition|null);
    }

    /** Represents a Tile. */
    class Tile implements ITile {

        /**
         * Constructs a new Tile.
         * @param [properties] Properties to set
         */
        constructor(properties?: pb_base.ITile);

        /** Tile Id. */
        public Id: (number|Long);

        /** Tile OwnerId. */
        public OwnerId: (number|Long);

        /** Tile Level. */
        public Level: number;

        /** Tile Position. */
        public Position?: (pb_base.IPosition|null);

        /**
         * Creates a new Tile instance using the specified properties.
         * @param [properties] Properties to set
         * @returns Tile instance
         */
        public static create(properties?: pb_base.ITile): pb_base.Tile;

        /**
         * Encodes the specified Tile message. Does not implicitly {@link pb_base.Tile.verify|verify} messages.
         * @param message Tile message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: pb_base.ITile, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified Tile message, length delimited. Does not implicitly {@link pb_base.Tile.verify|verify} messages.
         * @param message Tile message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: pb_base.ITile, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Tile message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Tile
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): pb_base.Tile;

        /**
         * Decodes a Tile message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns Tile
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): pb_base.Tile;

        /**
         * Verifies a Tile message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a Tile message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Tile
         */
        public static fromObject(object: { [k: string]: any }): pb_base.Tile;

        /**
         * Creates a plain object from a Tile message. Also converts values to other types if specified.
         * @param message Tile
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: pb_base.Tile, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Tile to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for Tile
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a Capital. */
    interface ICapital {

        /** Capital Id */
        Id?: (number|Long|null);

        /** Capital OwnerId */
        OwnerId?: (number|Long|null);

        /** Capital Level */
        Level?: (number|null);

        /** Capital Position */
        Position?: (pb_base.IPosition|null);
    }

    /** Represents a Capital. */
    class Capital implements ICapital {

        /**
         * Constructs a new Capital.
         * @param [properties] Properties to set
         */
        constructor(properties?: pb_base.ICapital);

        /** Capital Id. */
        public Id: (number|Long);

        /** Capital OwnerId. */
        public OwnerId: (number|Long);

        /** Capital Level. */
        public Level: number;

        /** Capital Position. */
        public Position?: (pb_base.IPosition|null);

        /**
         * Creates a new Capital instance using the specified properties.
         * @param [properties] Properties to set
         * @returns Capital instance
         */
        public static create(properties?: pb_base.ICapital): pb_base.Capital;

        /**
         * Encodes the specified Capital message. Does not implicitly {@link pb_base.Capital.verify|verify} messages.
         * @param message Capital message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: pb_base.ICapital, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified Capital message, length delimited. Does not implicitly {@link pb_base.Capital.verify|verify} messages.
         * @param message Capital message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: pb_base.ICapital, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Capital message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Capital
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): pb_base.Capital;

        /**
         * Decodes a Capital message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns Capital
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): pb_base.Capital;

        /**
         * Verifies a Capital message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a Capital message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Capital
         */
        public static fromObject(object: { [k: string]: any }): pb_base.Capital;

        /**
         * Creates a plain object from a Capital message. Also converts values to other types if specified.
         * @param message Capital
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: pb_base.Capital, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Capital to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for Capital
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a TroopStack. */
    interface ITroopStack {

        /** TroopStack Id */
        Id?: (number|Long|null);

        /** TroopStack OwnerId */
        OwnerId?: (number|Long|null);

        /** TroopStack Type */
        Type?: (pb_base.TroopStackType|null);

        /** TroopStack Amount */
        Amount?: (number|null);

        /** TroopStack Position */
        Position?: (pb_base.IPosition|null);

        /** TroopStack MovementStamina */
        MovementStamina?: (number|null);
    }

    /** Represents a TroopStack. */
    class TroopStack implements ITroopStack {

        /**
         * Constructs a new TroopStack.
         * @param [properties] Properties to set
         */
        constructor(properties?: pb_base.ITroopStack);

        /** TroopStack Id. */
        public Id: (number|Long);

        /** TroopStack OwnerId. */
        public OwnerId: (number|Long);

        /** TroopStack Type. */
        public Type: pb_base.TroopStackType;

        /** TroopStack Amount. */
        public Amount: number;

        /** TroopStack Position. */
        public Position?: (pb_base.IPosition|null);

        /** TroopStack MovementStamina. */
        public MovementStamina: number;

        /**
         * Creates a new TroopStack instance using the specified properties.
         * @param [properties] Properties to set
         * @returns TroopStack instance
         */
        public static create(properties?: pb_base.ITroopStack): pb_base.TroopStack;

        /**
         * Encodes the specified TroopStack message. Does not implicitly {@link pb_base.TroopStack.verify|verify} messages.
         * @param message TroopStack message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: pb_base.ITroopStack, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified TroopStack message, length delimited. Does not implicitly {@link pb_base.TroopStack.verify|verify} messages.
         * @param message TroopStack message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: pb_base.ITroopStack, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a TroopStack message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns TroopStack
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): pb_base.TroopStack;

        /**
         * Decodes a TroopStack message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns TroopStack
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): pb_base.TroopStack;

        /**
         * Verifies a TroopStack message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a TroopStack message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns TroopStack
         */
        public static fromObject(object: { [k: string]: any }): pb_base.TroopStack;

        /**
         * Creates a plain object from a TroopStack message. Also converts values to other types if specified.
         * @param message TroopStack
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: pb_base.TroopStack, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this TroopStack to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for TroopStack
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of an IdentityPayload. */
    interface IIdentityPayload {

        /** IdentityPayload JwtToken */
        JwtToken?: (string|null);

        /** IdentityPayload PlayerId */
        PlayerId?: (number|Long|null);
    }

    /** Represents an IdentityPayload. */
    class IdentityPayload implements IIdentityPayload {

        /**
         * Constructs a new IdentityPayload.
         * @param [properties] Properties to set
         */
        constructor(properties?: pb_base.IIdentityPayload);

        /** IdentityPayload JwtToken. */
        public JwtToken: string;

        /** IdentityPayload PlayerId. */
        public PlayerId: (number|Long);

        /**
         * Creates a new IdentityPayload instance using the specified properties.
         * @param [properties] Properties to set
         * @returns IdentityPayload instance
         */
        public static create(properties?: pb_base.IIdentityPayload): pb_base.IdentityPayload;

        /**
         * Encodes the specified IdentityPayload message. Does not implicitly {@link pb_base.IdentityPayload.verify|verify} messages.
         * @param message IdentityPayload message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: pb_base.IIdentityPayload, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified IdentityPayload message, length delimited. Does not implicitly {@link pb_base.IdentityPayload.verify|verify} messages.
         * @param message IdentityPayload message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: pb_base.IIdentityPayload, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an IdentityPayload message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns IdentityPayload
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): pb_base.IdentityPayload;

        /**
         * Decodes an IdentityPayload message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns IdentityPayload
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): pb_base.IdentityPayload;

        /**
         * Verifies an IdentityPayload message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates an IdentityPayload message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns IdentityPayload
         */
        public static fromObject(object: { [k: string]: any }): pb_base.IdentityPayload;

        /**
         * Creates a plain object from an IdentityPayload message. Also converts values to other types if specified.
         * @param message IdentityPayload
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: pb_base.IdentityPayload, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this IdentityPayload to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for IdentityPayload
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }
}
