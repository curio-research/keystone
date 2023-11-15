import * as $protobuf from "protobufjs";
import Long = require("long");
/** Namespace pb_round. */
export namespace pb_round {

    /** Properties of a C2S_DiscardCards. */
    interface IC2S_DiscardCards {

        /** C2S_DiscardCards PokerCards */
        PokerCards?: ((number|Long)[]|null);

        /** C2S_DiscardCards IdentityPayload */
        IdentityPayload?: (pb_base.IIdentityPayload|null);
    }

    /** Represents a C2S_DiscardCards. */
    class C2S_DiscardCards implements IC2S_DiscardCards {

        /**
         * Constructs a new C2S_DiscardCards.
         * @param [properties] Properties to set
         */
        constructor(properties?: pb_round.IC2S_DiscardCards);

        /** C2S_DiscardCards PokerCards. */
        public PokerCards: (number|Long)[];

        /** C2S_DiscardCards IdentityPayload. */
        public IdentityPayload?: (pb_base.IIdentityPayload|null);

        /**
         * Creates a new C2S_DiscardCards instance using the specified properties.
         * @param [properties] Properties to set
         * @returns C2S_DiscardCards instance
         */
        public static create(properties?: pb_round.IC2S_DiscardCards): pb_round.C2S_DiscardCards;

        /**
         * Encodes the specified C2S_DiscardCards message. Does not implicitly {@link pb_round.C2S_DiscardCards.verify|verify} messages.
         * @param message C2S_DiscardCards message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: pb_round.IC2S_DiscardCards, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified C2S_DiscardCards message, length delimited. Does not implicitly {@link pb_round.C2S_DiscardCards.verify|verify} messages.
         * @param message C2S_DiscardCards message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: pb_round.IC2S_DiscardCards, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a C2S_DiscardCards message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns C2S_DiscardCards
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): pb_round.C2S_DiscardCards;

        /**
         * Decodes a C2S_DiscardCards message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns C2S_DiscardCards
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): pb_round.C2S_DiscardCards;

        /**
         * Verifies a C2S_DiscardCards message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a C2S_DiscardCards message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns C2S_DiscardCards
         */
        public static fromObject(object: { [k: string]: any }): pb_round.C2S_DiscardCards;

        /**
         * Creates a plain object from a C2S_DiscardCards message. Also converts values to other types if specified.
         * @param message C2S_DiscardCards
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: pb_round.C2S_DiscardCards, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this C2S_DiscardCards to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for C2S_DiscardCards
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a S2C_RoundStart. */
    interface IS2C_RoundStart {

        /** S2C_RoundStart Round */
        Round?: (number|Long|null);
    }

    /** Represents a S2C_RoundStart. */
    class S2C_RoundStart implements IS2C_RoundStart {

        /**
         * Constructs a new S2C_RoundStart.
         * @param [properties] Properties to set
         */
        constructor(properties?: pb_round.IS2C_RoundStart);

        /** S2C_RoundStart Round. */
        public Round: (number|Long);

        /**
         * Creates a new S2C_RoundStart instance using the specified properties.
         * @param [properties] Properties to set
         * @returns S2C_RoundStart instance
         */
        public static create(properties?: pb_round.IS2C_RoundStart): pb_round.S2C_RoundStart;

        /**
         * Encodes the specified S2C_RoundStart message. Does not implicitly {@link pb_round.S2C_RoundStart.verify|verify} messages.
         * @param message S2C_RoundStart message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: pb_round.IS2C_RoundStart, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified S2C_RoundStart message, length delimited. Does not implicitly {@link pb_round.S2C_RoundStart.verify|verify} messages.
         * @param message S2C_RoundStart message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: pb_round.IS2C_RoundStart, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a S2C_RoundStart message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns S2C_RoundStart
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): pb_round.S2C_RoundStart;

        /**
         * Decodes a S2C_RoundStart message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns S2C_RoundStart
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): pb_round.S2C_RoundStart;

        /**
         * Verifies a S2C_RoundStart message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a S2C_RoundStart message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns S2C_RoundStart
         */
        public static fromObject(object: { [k: string]: any }): pb_round.S2C_RoundStart;

        /**
         * Creates a plain object from a S2C_RoundStart message. Also converts values to other types if specified.
         * @param message S2C_RoundStart
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: pb_round.S2C_RoundStart, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this S2C_RoundStart to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for S2C_RoundStart
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a S2C_DiscardPhaseEnded. */
    interface IS2C_DiscardPhaseEnded {
    }

    /** Represents a S2C_DiscardPhaseEnded. */
    class S2C_DiscardPhaseEnded implements IS2C_DiscardPhaseEnded {

        /**
         * Constructs a new S2C_DiscardPhaseEnded.
         * @param [properties] Properties to set
         */
        constructor(properties?: pb_round.IS2C_DiscardPhaseEnded);

        /**
         * Creates a new S2C_DiscardPhaseEnded instance using the specified properties.
         * @param [properties] Properties to set
         * @returns S2C_DiscardPhaseEnded instance
         */
        public static create(properties?: pb_round.IS2C_DiscardPhaseEnded): pb_round.S2C_DiscardPhaseEnded;

        /**
         * Encodes the specified S2C_DiscardPhaseEnded message. Does not implicitly {@link pb_round.S2C_DiscardPhaseEnded.verify|verify} messages.
         * @param message S2C_DiscardPhaseEnded message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: pb_round.IS2C_DiscardPhaseEnded, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified S2C_DiscardPhaseEnded message, length delimited. Does not implicitly {@link pb_round.S2C_DiscardPhaseEnded.verify|verify} messages.
         * @param message S2C_DiscardPhaseEnded message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: pb_round.IS2C_DiscardPhaseEnded, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a S2C_DiscardPhaseEnded message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns S2C_DiscardPhaseEnded
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): pb_round.S2C_DiscardPhaseEnded;

        /**
         * Decodes a S2C_DiscardPhaseEnded message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns S2C_DiscardPhaseEnded
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): pb_round.S2C_DiscardPhaseEnded;

        /**
         * Verifies a S2C_DiscardPhaseEnded message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a S2C_DiscardPhaseEnded message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns S2C_DiscardPhaseEnded
         */
        public static fromObject(object: { [k: string]: any }): pb_round.S2C_DiscardPhaseEnded;

        /**
         * Creates a plain object from a S2C_DiscardPhaseEnded message. Also converts values to other types if specified.
         * @param message S2C_DiscardPhaseEnded
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: pb_round.S2C_DiscardPhaseEnded, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this S2C_DiscardPhaseEnded to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for S2C_DiscardPhaseEnded
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a S2C_DiscardCards. */
    interface IS2C_DiscardCards {

        /** S2C_DiscardCards PokerCardIds */
        PokerCardIds?: ((number|Long)[]|null);
    }

    /** Represents a S2C_DiscardCards. */
    class S2C_DiscardCards implements IS2C_DiscardCards {

        /**
         * Constructs a new S2C_DiscardCards.
         * @param [properties] Properties to set
         */
        constructor(properties?: pb_round.IS2C_DiscardCards);

        /** S2C_DiscardCards PokerCardIds. */
        public PokerCardIds: (number|Long)[];

        /**
         * Creates a new S2C_DiscardCards instance using the specified properties.
         * @param [properties] Properties to set
         * @returns S2C_DiscardCards instance
         */
        public static create(properties?: pb_round.IS2C_DiscardCards): pb_round.S2C_DiscardCards;

        /**
         * Encodes the specified S2C_DiscardCards message. Does not implicitly {@link pb_round.S2C_DiscardCards.verify|verify} messages.
         * @param message S2C_DiscardCards message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: pb_round.IS2C_DiscardCards, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified S2C_DiscardCards message, length delimited. Does not implicitly {@link pb_round.S2C_DiscardCards.verify|verify} messages.
         * @param message S2C_DiscardCards message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: pb_round.IS2C_DiscardCards, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a S2C_DiscardCards message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns S2C_DiscardCards
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): pb_round.S2C_DiscardCards;

        /**
         * Decodes a S2C_DiscardCards message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns S2C_DiscardCards
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): pb_round.S2C_DiscardCards;

        /**
         * Verifies a S2C_DiscardCards message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a S2C_DiscardCards message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns S2C_DiscardCards
         */
        public static fromObject(object: { [k: string]: any }): pb_round.S2C_DiscardCards;

        /**
         * Creates a plain object from a S2C_DiscardCards message. Also converts values to other types if specified.
         * @param message S2C_DiscardCards
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: pb_round.S2C_DiscardCards, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this S2C_DiscardCards to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for S2C_DiscardCards
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a S2C_EndGame. */
    interface IS2C_EndGame {

        /** S2C_EndGame RoomName */
        RoomName?: (string|null);

        /** S2C_EndGame Results */
        Results?: (pb_round.IResult[]|null);
    }

    /** Represents a S2C_EndGame. */
    class S2C_EndGame implements IS2C_EndGame {

        /**
         * Constructs a new S2C_EndGame.
         * @param [properties] Properties to set
         */
        constructor(properties?: pb_round.IS2C_EndGame);

        /** S2C_EndGame RoomName. */
        public RoomName: string;

        /** S2C_EndGame Results. */
        public Results: pb_round.IResult[];

        /**
         * Creates a new S2C_EndGame instance using the specified properties.
         * @param [properties] Properties to set
         * @returns S2C_EndGame instance
         */
        public static create(properties?: pb_round.IS2C_EndGame): pb_round.S2C_EndGame;

        /**
         * Encodes the specified S2C_EndGame message. Does not implicitly {@link pb_round.S2C_EndGame.verify|verify} messages.
         * @param message S2C_EndGame message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: pb_round.IS2C_EndGame, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified S2C_EndGame message, length delimited. Does not implicitly {@link pb_round.S2C_EndGame.verify|verify} messages.
         * @param message S2C_EndGame message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: pb_round.IS2C_EndGame, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a S2C_EndGame message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns S2C_EndGame
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): pb_round.S2C_EndGame;

        /**
         * Decodes a S2C_EndGame message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns S2C_EndGame
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): pb_round.S2C_EndGame;

        /**
         * Verifies a S2C_EndGame message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a S2C_EndGame message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns S2C_EndGame
         */
        public static fromObject(object: { [k: string]: any }): pb_round.S2C_EndGame;

        /**
         * Creates a plain object from a S2C_EndGame message. Also converts values to other types if specified.
         * @param message S2C_EndGame
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: pb_round.S2C_EndGame, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this S2C_EndGame to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for S2C_EndGame
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a Result. */
    interface IResult {

        /** Result Username */
        Username?: (string|null);

        /** Result Elo */
        Elo?: (number|Long|null);
    }

    /** Represents a Result. */
    class Result implements IResult {

        /**
         * Constructs a new Result.
         * @param [properties] Properties to set
         */
        constructor(properties?: pb_round.IResult);

        /** Result Username. */
        public Username: string;

        /** Result Elo. */
        public Elo: (number|Long);

        /**
         * Creates a new Result instance using the specified properties.
         * @param [properties] Properties to set
         * @returns Result instance
         */
        public static create(properties?: pb_round.IResult): pb_round.Result;

        /**
         * Encodes the specified Result message. Does not implicitly {@link pb_round.Result.verify|verify} messages.
         * @param message Result message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: pb_round.IResult, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified Result message, length delimited. Does not implicitly {@link pb_round.Result.verify|verify} messages.
         * @param message Result message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: pb_round.IResult, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Result message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Result
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): pb_round.Result;

        /**
         * Decodes a Result message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns Result
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): pb_round.Result;

        /**
         * Verifies a Result message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a Result message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Result
         */
        public static fromObject(object: { [k: string]: any }): pb_round.Result;

        /**
         * Creates a plain object from a Result message. Also converts values to other types if specified.
         * @param message Result
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: pb_round.Result, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Result to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for Result
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
