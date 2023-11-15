/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

$root.pb_game = (function() {

    /**
     * Namespace pb_game.
     * @exports pb_game
     * @namespace
     */
    var pb_game = {};

    pb_game.C2S_EstablishPlayer = (function() {

        /**
         * Properties of a C2S_EstablishPlayer.
         * @memberof pb_game
         * @interface IC2S_EstablishPlayer
         * @property {number|Long|null} [PlayerId] C2S_EstablishPlayer PlayerId
         */

        /**
         * Constructs a new C2S_EstablishPlayer.
         * @memberof pb_game
         * @classdesc Represents a C2S_EstablishPlayer.
         * @implements IC2S_EstablishPlayer
         * @constructor
         * @param {pb_game.IC2S_EstablishPlayer=} [properties] Properties to set
         */
        function C2S_EstablishPlayer(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * C2S_EstablishPlayer PlayerId.
         * @member {number|Long} PlayerId
         * @memberof pb_game.C2S_EstablishPlayer
         * @instance
         */
        C2S_EstablishPlayer.prototype.PlayerId = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Creates a new C2S_EstablishPlayer instance using the specified properties.
         * @function create
         * @memberof pb_game.C2S_EstablishPlayer
         * @static
         * @param {pb_game.IC2S_EstablishPlayer=} [properties] Properties to set
         * @returns {pb_game.C2S_EstablishPlayer} C2S_EstablishPlayer instance
         */
        C2S_EstablishPlayer.create = function create(properties) {
            return new C2S_EstablishPlayer(properties);
        };

        /**
         * Encodes the specified C2S_EstablishPlayer message. Does not implicitly {@link pb_game.C2S_EstablishPlayer.verify|verify} messages.
         * @function encode
         * @memberof pb_game.C2S_EstablishPlayer
         * @static
         * @param {pb_game.IC2S_EstablishPlayer} message C2S_EstablishPlayer message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        C2S_EstablishPlayer.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.PlayerId != null && Object.hasOwnProperty.call(message, "PlayerId"))
                writer.uint32(/* id 1, wireType 0 =*/8).int64(message.PlayerId);
            return writer;
        };

        /**
         * Encodes the specified C2S_EstablishPlayer message, length delimited. Does not implicitly {@link pb_game.C2S_EstablishPlayer.verify|verify} messages.
         * @function encodeDelimited
         * @memberof pb_game.C2S_EstablishPlayer
         * @static
         * @param {pb_game.IC2S_EstablishPlayer} message C2S_EstablishPlayer message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        C2S_EstablishPlayer.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a C2S_EstablishPlayer message from the specified reader or buffer.
         * @function decode
         * @memberof pb_game.C2S_EstablishPlayer
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {pb_game.C2S_EstablishPlayer} C2S_EstablishPlayer
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        C2S_EstablishPlayer.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.pb_game.C2S_EstablishPlayer();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.PlayerId = reader.int64();
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a C2S_EstablishPlayer message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof pb_game.C2S_EstablishPlayer
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {pb_game.C2S_EstablishPlayer} C2S_EstablishPlayer
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        C2S_EstablishPlayer.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a C2S_EstablishPlayer message.
         * @function verify
         * @memberof pb_game.C2S_EstablishPlayer
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        C2S_EstablishPlayer.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.PlayerId != null && message.hasOwnProperty("PlayerId"))
                if (!$util.isInteger(message.PlayerId) && !(message.PlayerId && $util.isInteger(message.PlayerId.low) && $util.isInteger(message.PlayerId.high)))
                    return "PlayerId: integer|Long expected";
            return null;
        };

        /**
         * Creates a C2S_EstablishPlayer message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof pb_game.C2S_EstablishPlayer
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {pb_game.C2S_EstablishPlayer} C2S_EstablishPlayer
         */
        C2S_EstablishPlayer.fromObject = function fromObject(object) {
            if (object instanceof $root.pb_game.C2S_EstablishPlayer)
                return object;
            var message = new $root.pb_game.C2S_EstablishPlayer();
            if (object.PlayerId != null)
                if ($util.Long)
                    (message.PlayerId = $util.Long.fromValue(object.PlayerId)).unsigned = false;
                else if (typeof object.PlayerId === "string")
                    message.PlayerId = parseInt(object.PlayerId, 10);
                else if (typeof object.PlayerId === "number")
                    message.PlayerId = object.PlayerId;
                else if (typeof object.PlayerId === "object")
                    message.PlayerId = new $util.LongBits(object.PlayerId.low >>> 0, object.PlayerId.high >>> 0).toNumber();
            return message;
        };

        /**
         * Creates a plain object from a C2S_EstablishPlayer message. Also converts values to other types if specified.
         * @function toObject
         * @memberof pb_game.C2S_EstablishPlayer
         * @static
         * @param {pb_game.C2S_EstablishPlayer} message C2S_EstablishPlayer
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        C2S_EstablishPlayer.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.PlayerId = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.PlayerId = options.longs === String ? "0" : 0;
            if (message.PlayerId != null && message.hasOwnProperty("PlayerId"))
                if (typeof message.PlayerId === "number")
                    object.PlayerId = options.longs === String ? String(message.PlayerId) : message.PlayerId;
                else
                    object.PlayerId = options.longs === String ? $util.Long.prototype.toString.call(message.PlayerId) : options.longs === Number ? new $util.LongBits(message.PlayerId.low >>> 0, message.PlayerId.high >>> 0).toNumber() : message.PlayerId;
            return object;
        };

        /**
         * Converts this C2S_EstablishPlayer to JSON.
         * @function toJSON
         * @memberof pb_game.C2S_EstablishPlayer
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        C2S_EstablishPlayer.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for C2S_EstablishPlayer
         * @function getTypeUrl
         * @memberof pb_game.C2S_EstablishPlayer
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        C2S_EstablishPlayer.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/pb_game.C2S_EstablishPlayer";
        };

        return C2S_EstablishPlayer;
    })();

    pb_game.C2S_GameState = (function() {

        /**
         * Properties of a C2S_GameState.
         * @memberof pb_game
         * @interface IC2S_GameState
         * @property {pb_base.IIdentityPayload|null} [IdentityPayload] C2S_GameState IdentityPayload
         */

        /**
         * Constructs a new C2S_GameState.
         * @memberof pb_game
         * @classdesc Represents a C2S_GameState.
         * @implements IC2S_GameState
         * @constructor
         * @param {pb_game.IC2S_GameState=} [properties] Properties to set
         */
        function C2S_GameState(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * C2S_GameState IdentityPayload.
         * @member {pb_base.IIdentityPayload|null|undefined} IdentityPayload
         * @memberof pb_game.C2S_GameState
         * @instance
         */
        C2S_GameState.prototype.IdentityPayload = null;

        /**
         * Creates a new C2S_GameState instance using the specified properties.
         * @function create
         * @memberof pb_game.C2S_GameState
         * @static
         * @param {pb_game.IC2S_GameState=} [properties] Properties to set
         * @returns {pb_game.C2S_GameState} C2S_GameState instance
         */
        C2S_GameState.create = function create(properties) {
            return new C2S_GameState(properties);
        };

        /**
         * Encodes the specified C2S_GameState message. Does not implicitly {@link pb_game.C2S_GameState.verify|verify} messages.
         * @function encode
         * @memberof pb_game.C2S_GameState
         * @static
         * @param {pb_game.IC2S_GameState} message C2S_GameState message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        C2S_GameState.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.IdentityPayload != null && Object.hasOwnProperty.call(message, "IdentityPayload"))
                $root.pb_base.IdentityPayload.encode(message.IdentityPayload, writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified C2S_GameState message, length delimited. Does not implicitly {@link pb_game.C2S_GameState.verify|verify} messages.
         * @function encodeDelimited
         * @memberof pb_game.C2S_GameState
         * @static
         * @param {pb_game.IC2S_GameState} message C2S_GameState message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        C2S_GameState.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a C2S_GameState message from the specified reader or buffer.
         * @function decode
         * @memberof pb_game.C2S_GameState
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {pb_game.C2S_GameState} C2S_GameState
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        C2S_GameState.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.pb_game.C2S_GameState();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.IdentityPayload = $root.pb_base.IdentityPayload.decode(reader, reader.uint32());
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a C2S_GameState message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof pb_game.C2S_GameState
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {pb_game.C2S_GameState} C2S_GameState
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        C2S_GameState.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a C2S_GameState message.
         * @function verify
         * @memberof pb_game.C2S_GameState
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        C2S_GameState.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.IdentityPayload != null && message.hasOwnProperty("IdentityPayload")) {
                var error = $root.pb_base.IdentityPayload.verify(message.IdentityPayload);
                if (error)
                    return "IdentityPayload." + error;
            }
            return null;
        };

        /**
         * Creates a C2S_GameState message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof pb_game.C2S_GameState
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {pb_game.C2S_GameState} C2S_GameState
         */
        C2S_GameState.fromObject = function fromObject(object) {
            if (object instanceof $root.pb_game.C2S_GameState)
                return object;
            var message = new $root.pb_game.C2S_GameState();
            if (object.IdentityPayload != null) {
                if (typeof object.IdentityPayload !== "object")
                    throw TypeError(".pb_game.C2S_GameState.IdentityPayload: object expected");
                message.IdentityPayload = $root.pb_base.IdentityPayload.fromObject(object.IdentityPayload);
            }
            return message;
        };

        /**
         * Creates a plain object from a C2S_GameState message. Also converts values to other types if specified.
         * @function toObject
         * @memberof pb_game.C2S_GameState
         * @static
         * @param {pb_game.C2S_GameState} message C2S_GameState
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        C2S_GameState.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.IdentityPayload = null;
            if (message.IdentityPayload != null && message.hasOwnProperty("IdentityPayload"))
                object.IdentityPayload = $root.pb_base.IdentityPayload.toObject(message.IdentityPayload, options);
            return object;
        };

        /**
         * Converts this C2S_GameState to JSON.
         * @function toJSON
         * @memberof pb_game.C2S_GameState
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        C2S_GameState.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for C2S_GameState
         * @function getTypeUrl
         * @memberof pb_game.C2S_GameState
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        C2S_GameState.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/pb_game.C2S_GameState";
        };

        return C2S_GameState;
    })();

    pb_game.C2S_Ping = (function() {

        /**
         * Properties of a C2S_Ping.
         * @memberof pb_game
         * @interface IC2S_Ping
         * @property {number|Long|null} [timeStamp] C2S_Ping timeStamp
         * @property {pb_base.IIdentityPayload|null} [IdentityPayload] C2S_Ping IdentityPayload
         */

        /**
         * Constructs a new C2S_Ping.
         * @memberof pb_game
         * @classdesc Represents a C2S_Ping.
         * @implements IC2S_Ping
         * @constructor
         * @param {pb_game.IC2S_Ping=} [properties] Properties to set
         */
        function C2S_Ping(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * C2S_Ping timeStamp.
         * @member {number|Long} timeStamp
         * @memberof pb_game.C2S_Ping
         * @instance
         */
        C2S_Ping.prototype.timeStamp = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * C2S_Ping IdentityPayload.
         * @member {pb_base.IIdentityPayload|null|undefined} IdentityPayload
         * @memberof pb_game.C2S_Ping
         * @instance
         */
        C2S_Ping.prototype.IdentityPayload = null;

        /**
         * Creates a new C2S_Ping instance using the specified properties.
         * @function create
         * @memberof pb_game.C2S_Ping
         * @static
         * @param {pb_game.IC2S_Ping=} [properties] Properties to set
         * @returns {pb_game.C2S_Ping} C2S_Ping instance
         */
        C2S_Ping.create = function create(properties) {
            return new C2S_Ping(properties);
        };

        /**
         * Encodes the specified C2S_Ping message. Does not implicitly {@link pb_game.C2S_Ping.verify|verify} messages.
         * @function encode
         * @memberof pb_game.C2S_Ping
         * @static
         * @param {pb_game.IC2S_Ping} message C2S_Ping message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        C2S_Ping.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.timeStamp != null && Object.hasOwnProperty.call(message, "timeStamp"))
                writer.uint32(/* id 1, wireType 0 =*/8).int64(message.timeStamp);
            if (message.IdentityPayload != null && Object.hasOwnProperty.call(message, "IdentityPayload"))
                $root.pb_base.IdentityPayload.encode(message.IdentityPayload, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified C2S_Ping message, length delimited. Does not implicitly {@link pb_game.C2S_Ping.verify|verify} messages.
         * @function encodeDelimited
         * @memberof pb_game.C2S_Ping
         * @static
         * @param {pb_game.IC2S_Ping} message C2S_Ping message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        C2S_Ping.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a C2S_Ping message from the specified reader or buffer.
         * @function decode
         * @memberof pb_game.C2S_Ping
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {pb_game.C2S_Ping} C2S_Ping
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        C2S_Ping.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.pb_game.C2S_Ping();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.timeStamp = reader.int64();
                        break;
                    }
                case 2: {
                        message.IdentityPayload = $root.pb_base.IdentityPayload.decode(reader, reader.uint32());
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a C2S_Ping message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof pb_game.C2S_Ping
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {pb_game.C2S_Ping} C2S_Ping
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        C2S_Ping.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a C2S_Ping message.
         * @function verify
         * @memberof pb_game.C2S_Ping
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        C2S_Ping.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.timeStamp != null && message.hasOwnProperty("timeStamp"))
                if (!$util.isInteger(message.timeStamp) && !(message.timeStamp && $util.isInteger(message.timeStamp.low) && $util.isInteger(message.timeStamp.high)))
                    return "timeStamp: integer|Long expected";
            if (message.IdentityPayload != null && message.hasOwnProperty("IdentityPayload")) {
                var error = $root.pb_base.IdentityPayload.verify(message.IdentityPayload);
                if (error)
                    return "IdentityPayload." + error;
            }
            return null;
        };

        /**
         * Creates a C2S_Ping message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof pb_game.C2S_Ping
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {pb_game.C2S_Ping} C2S_Ping
         */
        C2S_Ping.fromObject = function fromObject(object) {
            if (object instanceof $root.pb_game.C2S_Ping)
                return object;
            var message = new $root.pb_game.C2S_Ping();
            if (object.timeStamp != null)
                if ($util.Long)
                    (message.timeStamp = $util.Long.fromValue(object.timeStamp)).unsigned = false;
                else if (typeof object.timeStamp === "string")
                    message.timeStamp = parseInt(object.timeStamp, 10);
                else if (typeof object.timeStamp === "number")
                    message.timeStamp = object.timeStamp;
                else if (typeof object.timeStamp === "object")
                    message.timeStamp = new $util.LongBits(object.timeStamp.low >>> 0, object.timeStamp.high >>> 0).toNumber();
            if (object.IdentityPayload != null) {
                if (typeof object.IdentityPayload !== "object")
                    throw TypeError(".pb_game.C2S_Ping.IdentityPayload: object expected");
                message.IdentityPayload = $root.pb_base.IdentityPayload.fromObject(object.IdentityPayload);
            }
            return message;
        };

        /**
         * Creates a plain object from a C2S_Ping message. Also converts values to other types if specified.
         * @function toObject
         * @memberof pb_game.C2S_Ping
         * @static
         * @param {pb_game.C2S_Ping} message C2S_Ping
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        C2S_Ping.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.timeStamp = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.timeStamp = options.longs === String ? "0" : 0;
                object.IdentityPayload = null;
            }
            if (message.timeStamp != null && message.hasOwnProperty("timeStamp"))
                if (typeof message.timeStamp === "number")
                    object.timeStamp = options.longs === String ? String(message.timeStamp) : message.timeStamp;
                else
                    object.timeStamp = options.longs === String ? $util.Long.prototype.toString.call(message.timeStamp) : options.longs === Number ? new $util.LongBits(message.timeStamp.low >>> 0, message.timeStamp.high >>> 0).toNumber() : message.timeStamp;
            if (message.IdentityPayload != null && message.hasOwnProperty("IdentityPayload"))
                object.IdentityPayload = $root.pb_base.IdentityPayload.toObject(message.IdentityPayload, options);
            return object;
        };

        /**
         * Converts this C2S_Ping to JSON.
         * @function toJSON
         * @memberof pb_game.C2S_Ping
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        C2S_Ping.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for C2S_Ping
         * @function getTypeUrl
         * @memberof pb_game.C2S_Ping
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        C2S_Ping.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/pb_game.C2S_Ping";
        };

        return C2S_Ping;
    })();

    pb_game.S2C_EstablishPlayer = (function() {

        /**
         * Properties of a S2C_EstablishPlayer.
         * @memberof pb_game
         * @interface IS2C_EstablishPlayer
         */

        /**
         * Constructs a new S2C_EstablishPlayer.
         * @memberof pb_game
         * @classdesc Represents a S2C_EstablishPlayer.
         * @implements IS2C_EstablishPlayer
         * @constructor
         * @param {pb_game.IS2C_EstablishPlayer=} [properties] Properties to set
         */
        function S2C_EstablishPlayer(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Creates a new S2C_EstablishPlayer instance using the specified properties.
         * @function create
         * @memberof pb_game.S2C_EstablishPlayer
         * @static
         * @param {pb_game.IS2C_EstablishPlayer=} [properties] Properties to set
         * @returns {pb_game.S2C_EstablishPlayer} S2C_EstablishPlayer instance
         */
        S2C_EstablishPlayer.create = function create(properties) {
            return new S2C_EstablishPlayer(properties);
        };

        /**
         * Encodes the specified S2C_EstablishPlayer message. Does not implicitly {@link pb_game.S2C_EstablishPlayer.verify|verify} messages.
         * @function encode
         * @memberof pb_game.S2C_EstablishPlayer
         * @static
         * @param {pb_game.IS2C_EstablishPlayer} message S2C_EstablishPlayer message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        S2C_EstablishPlayer.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            return writer;
        };

        /**
         * Encodes the specified S2C_EstablishPlayer message, length delimited. Does not implicitly {@link pb_game.S2C_EstablishPlayer.verify|verify} messages.
         * @function encodeDelimited
         * @memberof pb_game.S2C_EstablishPlayer
         * @static
         * @param {pb_game.IS2C_EstablishPlayer} message S2C_EstablishPlayer message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        S2C_EstablishPlayer.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a S2C_EstablishPlayer message from the specified reader or buffer.
         * @function decode
         * @memberof pb_game.S2C_EstablishPlayer
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {pb_game.S2C_EstablishPlayer} S2C_EstablishPlayer
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        S2C_EstablishPlayer.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.pb_game.S2C_EstablishPlayer();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a S2C_EstablishPlayer message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof pb_game.S2C_EstablishPlayer
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {pb_game.S2C_EstablishPlayer} S2C_EstablishPlayer
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        S2C_EstablishPlayer.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a S2C_EstablishPlayer message.
         * @function verify
         * @memberof pb_game.S2C_EstablishPlayer
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        S2C_EstablishPlayer.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            return null;
        };

        /**
         * Creates a S2C_EstablishPlayer message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof pb_game.S2C_EstablishPlayer
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {pb_game.S2C_EstablishPlayer} S2C_EstablishPlayer
         */
        S2C_EstablishPlayer.fromObject = function fromObject(object) {
            if (object instanceof $root.pb_game.S2C_EstablishPlayer)
                return object;
            return new $root.pb_game.S2C_EstablishPlayer();
        };

        /**
         * Creates a plain object from a S2C_EstablishPlayer message. Also converts values to other types if specified.
         * @function toObject
         * @memberof pb_game.S2C_EstablishPlayer
         * @static
         * @param {pb_game.S2C_EstablishPlayer} message S2C_EstablishPlayer
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        S2C_EstablishPlayer.toObject = function toObject() {
            return {};
        };

        /**
         * Converts this S2C_EstablishPlayer to JSON.
         * @function toJSON
         * @memberof pb_game.S2C_EstablishPlayer
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        S2C_EstablishPlayer.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for S2C_EstablishPlayer
         * @function getTypeUrl
         * @memberof pb_game.S2C_EstablishPlayer
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        S2C_EstablishPlayer.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/pb_game.S2C_EstablishPlayer";
        };

        return S2C_EstablishPlayer;
    })();

    pb_game.S2C_GameState = (function() {

        /**
         * Properties of a S2C_GameState.
         * @memberof pb_game
         * @interface IS2C_GameState
         * @property {pb_base.IRound|null} [Round] S2C_GameState Round
         * @property {Array.<pb_base.IPlayer>|null} [PlayerList] S2C_GameState PlayerList
         * @property {Array.<pb_base.ITile>|null} [TileList] S2C_GameState TileList
         * @property {Array.<pb_base.ITroopStack>|null} [TroopStackList] S2C_GameState TroopStackList
         * @property {Array.<pb_base.ICapital>|null} [CapitalList] S2C_GameState CapitalList
         */

        /**
         * Constructs a new S2C_GameState.
         * @memberof pb_game
         * @classdesc Represents a S2C_GameState.
         * @implements IS2C_GameState
         * @constructor
         * @param {pb_game.IS2C_GameState=} [properties] Properties to set
         */
        function S2C_GameState(properties) {
            this.PlayerList = [];
            this.TileList = [];
            this.TroopStackList = [];
            this.CapitalList = [];
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * S2C_GameState Round.
         * @member {pb_base.IRound|null|undefined} Round
         * @memberof pb_game.S2C_GameState
         * @instance
         */
        S2C_GameState.prototype.Round = null;

        /**
         * S2C_GameState PlayerList.
         * @member {Array.<pb_base.IPlayer>} PlayerList
         * @memberof pb_game.S2C_GameState
         * @instance
         */
        S2C_GameState.prototype.PlayerList = $util.emptyArray;

        /**
         * S2C_GameState TileList.
         * @member {Array.<pb_base.ITile>} TileList
         * @memberof pb_game.S2C_GameState
         * @instance
         */
        S2C_GameState.prototype.TileList = $util.emptyArray;

        /**
         * S2C_GameState TroopStackList.
         * @member {Array.<pb_base.ITroopStack>} TroopStackList
         * @memberof pb_game.S2C_GameState
         * @instance
         */
        S2C_GameState.prototype.TroopStackList = $util.emptyArray;

        /**
         * S2C_GameState CapitalList.
         * @member {Array.<pb_base.ICapital>} CapitalList
         * @memberof pb_game.S2C_GameState
         * @instance
         */
        S2C_GameState.prototype.CapitalList = $util.emptyArray;

        /**
         * Creates a new S2C_GameState instance using the specified properties.
         * @function create
         * @memberof pb_game.S2C_GameState
         * @static
         * @param {pb_game.IS2C_GameState=} [properties] Properties to set
         * @returns {pb_game.S2C_GameState} S2C_GameState instance
         */
        S2C_GameState.create = function create(properties) {
            return new S2C_GameState(properties);
        };

        /**
         * Encodes the specified S2C_GameState message. Does not implicitly {@link pb_game.S2C_GameState.verify|verify} messages.
         * @function encode
         * @memberof pb_game.S2C_GameState
         * @static
         * @param {pb_game.IS2C_GameState} message S2C_GameState message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        S2C_GameState.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.Round != null && Object.hasOwnProperty.call(message, "Round"))
                $root.pb_base.Round.encode(message.Round, writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            if (message.PlayerList != null && message.PlayerList.length)
                for (var i = 0; i < message.PlayerList.length; ++i)
                    $root.pb_base.Player.encode(message.PlayerList[i], writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
            if (message.TileList != null && message.TileList.length)
                for (var i = 0; i < message.TileList.length; ++i)
                    $root.pb_base.Tile.encode(message.TileList[i], writer.uint32(/* id 3, wireType 2 =*/26).fork()).ldelim();
            if (message.TroopStackList != null && message.TroopStackList.length)
                for (var i = 0; i < message.TroopStackList.length; ++i)
                    $root.pb_base.TroopStack.encode(message.TroopStackList[i], writer.uint32(/* id 4, wireType 2 =*/34).fork()).ldelim();
            if (message.CapitalList != null && message.CapitalList.length)
                for (var i = 0; i < message.CapitalList.length; ++i)
                    $root.pb_base.Capital.encode(message.CapitalList[i], writer.uint32(/* id 5, wireType 2 =*/42).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified S2C_GameState message, length delimited. Does not implicitly {@link pb_game.S2C_GameState.verify|verify} messages.
         * @function encodeDelimited
         * @memberof pb_game.S2C_GameState
         * @static
         * @param {pb_game.IS2C_GameState} message S2C_GameState message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        S2C_GameState.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a S2C_GameState message from the specified reader or buffer.
         * @function decode
         * @memberof pb_game.S2C_GameState
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {pb_game.S2C_GameState} S2C_GameState
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        S2C_GameState.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.pb_game.S2C_GameState();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.Round = $root.pb_base.Round.decode(reader, reader.uint32());
                        break;
                    }
                case 2: {
                        if (!(message.PlayerList && message.PlayerList.length))
                            message.PlayerList = [];
                        message.PlayerList.push($root.pb_base.Player.decode(reader, reader.uint32()));
                        break;
                    }
                case 3: {
                        if (!(message.TileList && message.TileList.length))
                            message.TileList = [];
                        message.TileList.push($root.pb_base.Tile.decode(reader, reader.uint32()));
                        break;
                    }
                case 4: {
                        if (!(message.TroopStackList && message.TroopStackList.length))
                            message.TroopStackList = [];
                        message.TroopStackList.push($root.pb_base.TroopStack.decode(reader, reader.uint32()));
                        break;
                    }
                case 5: {
                        if (!(message.CapitalList && message.CapitalList.length))
                            message.CapitalList = [];
                        message.CapitalList.push($root.pb_base.Capital.decode(reader, reader.uint32()));
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a S2C_GameState message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof pb_game.S2C_GameState
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {pb_game.S2C_GameState} S2C_GameState
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        S2C_GameState.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a S2C_GameState message.
         * @function verify
         * @memberof pb_game.S2C_GameState
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        S2C_GameState.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.Round != null && message.hasOwnProperty("Round")) {
                var error = $root.pb_base.Round.verify(message.Round);
                if (error)
                    return "Round." + error;
            }
            if (message.PlayerList != null && message.hasOwnProperty("PlayerList")) {
                if (!Array.isArray(message.PlayerList))
                    return "PlayerList: array expected";
                for (var i = 0; i < message.PlayerList.length; ++i) {
                    var error = $root.pb_base.Player.verify(message.PlayerList[i]);
                    if (error)
                        return "PlayerList." + error;
                }
            }
            if (message.TileList != null && message.hasOwnProperty("TileList")) {
                if (!Array.isArray(message.TileList))
                    return "TileList: array expected";
                for (var i = 0; i < message.TileList.length; ++i) {
                    var error = $root.pb_base.Tile.verify(message.TileList[i]);
                    if (error)
                        return "TileList." + error;
                }
            }
            if (message.TroopStackList != null && message.hasOwnProperty("TroopStackList")) {
                if (!Array.isArray(message.TroopStackList))
                    return "TroopStackList: array expected";
                for (var i = 0; i < message.TroopStackList.length; ++i) {
                    var error = $root.pb_base.TroopStack.verify(message.TroopStackList[i]);
                    if (error)
                        return "TroopStackList." + error;
                }
            }
            if (message.CapitalList != null && message.hasOwnProperty("CapitalList")) {
                if (!Array.isArray(message.CapitalList))
                    return "CapitalList: array expected";
                for (var i = 0; i < message.CapitalList.length; ++i) {
                    var error = $root.pb_base.Capital.verify(message.CapitalList[i]);
                    if (error)
                        return "CapitalList." + error;
                }
            }
            return null;
        };

        /**
         * Creates a S2C_GameState message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof pb_game.S2C_GameState
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {pb_game.S2C_GameState} S2C_GameState
         */
        S2C_GameState.fromObject = function fromObject(object) {
            if (object instanceof $root.pb_game.S2C_GameState)
                return object;
            var message = new $root.pb_game.S2C_GameState();
            if (object.Round != null) {
                if (typeof object.Round !== "object")
                    throw TypeError(".pb_game.S2C_GameState.Round: object expected");
                message.Round = $root.pb_base.Round.fromObject(object.Round);
            }
            if (object.PlayerList) {
                if (!Array.isArray(object.PlayerList))
                    throw TypeError(".pb_game.S2C_GameState.PlayerList: array expected");
                message.PlayerList = [];
                for (var i = 0; i < object.PlayerList.length; ++i) {
                    if (typeof object.PlayerList[i] !== "object")
                        throw TypeError(".pb_game.S2C_GameState.PlayerList: object expected");
                    message.PlayerList[i] = $root.pb_base.Player.fromObject(object.PlayerList[i]);
                }
            }
            if (object.TileList) {
                if (!Array.isArray(object.TileList))
                    throw TypeError(".pb_game.S2C_GameState.TileList: array expected");
                message.TileList = [];
                for (var i = 0; i < object.TileList.length; ++i) {
                    if (typeof object.TileList[i] !== "object")
                        throw TypeError(".pb_game.S2C_GameState.TileList: object expected");
                    message.TileList[i] = $root.pb_base.Tile.fromObject(object.TileList[i]);
                }
            }
            if (object.TroopStackList) {
                if (!Array.isArray(object.TroopStackList))
                    throw TypeError(".pb_game.S2C_GameState.TroopStackList: array expected");
                message.TroopStackList = [];
                for (var i = 0; i < object.TroopStackList.length; ++i) {
                    if (typeof object.TroopStackList[i] !== "object")
                        throw TypeError(".pb_game.S2C_GameState.TroopStackList: object expected");
                    message.TroopStackList[i] = $root.pb_base.TroopStack.fromObject(object.TroopStackList[i]);
                }
            }
            if (object.CapitalList) {
                if (!Array.isArray(object.CapitalList))
                    throw TypeError(".pb_game.S2C_GameState.CapitalList: array expected");
                message.CapitalList = [];
                for (var i = 0; i < object.CapitalList.length; ++i) {
                    if (typeof object.CapitalList[i] !== "object")
                        throw TypeError(".pb_game.S2C_GameState.CapitalList: object expected");
                    message.CapitalList[i] = $root.pb_base.Capital.fromObject(object.CapitalList[i]);
                }
            }
            return message;
        };

        /**
         * Creates a plain object from a S2C_GameState message. Also converts values to other types if specified.
         * @function toObject
         * @memberof pb_game.S2C_GameState
         * @static
         * @param {pb_game.S2C_GameState} message S2C_GameState
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        S2C_GameState.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.arrays || options.defaults) {
                object.PlayerList = [];
                object.TileList = [];
                object.TroopStackList = [];
                object.CapitalList = [];
            }
            if (options.defaults)
                object.Round = null;
            if (message.Round != null && message.hasOwnProperty("Round"))
                object.Round = $root.pb_base.Round.toObject(message.Round, options);
            if (message.PlayerList && message.PlayerList.length) {
                object.PlayerList = [];
                for (var j = 0; j < message.PlayerList.length; ++j)
                    object.PlayerList[j] = $root.pb_base.Player.toObject(message.PlayerList[j], options);
            }
            if (message.TileList && message.TileList.length) {
                object.TileList = [];
                for (var j = 0; j < message.TileList.length; ++j)
                    object.TileList[j] = $root.pb_base.Tile.toObject(message.TileList[j], options);
            }
            if (message.TroopStackList && message.TroopStackList.length) {
                object.TroopStackList = [];
                for (var j = 0; j < message.TroopStackList.length; ++j)
                    object.TroopStackList[j] = $root.pb_base.TroopStack.toObject(message.TroopStackList[j], options);
            }
            if (message.CapitalList && message.CapitalList.length) {
                object.CapitalList = [];
                for (var j = 0; j < message.CapitalList.length; ++j)
                    object.CapitalList[j] = $root.pb_base.Capital.toObject(message.CapitalList[j], options);
            }
            return object;
        };

        /**
         * Converts this S2C_GameState to JSON.
         * @function toJSON
         * @memberof pb_game.S2C_GameState
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        S2C_GameState.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for S2C_GameState
         * @function getTypeUrl
         * @memberof pb_game.S2C_GameState
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        S2C_GameState.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/pb_game.S2C_GameState";
        };

        return S2C_GameState;
    })();

    pb_game.S2C_Ping = (function() {

        /**
         * Properties of a S2C_Ping.
         * @memberof pb_game
         * @interface IS2C_Ping
         * @property {number|Long|null} [TimeStamp] S2C_Ping TimeStamp
         */

        /**
         * Constructs a new S2C_Ping.
         * @memberof pb_game
         * @classdesc Represents a S2C_Ping.
         * @implements IS2C_Ping
         * @constructor
         * @param {pb_game.IS2C_Ping=} [properties] Properties to set
         */
        function S2C_Ping(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * S2C_Ping TimeStamp.
         * @member {number|Long} TimeStamp
         * @memberof pb_game.S2C_Ping
         * @instance
         */
        S2C_Ping.prototype.TimeStamp = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Creates a new S2C_Ping instance using the specified properties.
         * @function create
         * @memberof pb_game.S2C_Ping
         * @static
         * @param {pb_game.IS2C_Ping=} [properties] Properties to set
         * @returns {pb_game.S2C_Ping} S2C_Ping instance
         */
        S2C_Ping.create = function create(properties) {
            return new S2C_Ping(properties);
        };

        /**
         * Encodes the specified S2C_Ping message. Does not implicitly {@link pb_game.S2C_Ping.verify|verify} messages.
         * @function encode
         * @memberof pb_game.S2C_Ping
         * @static
         * @param {pb_game.IS2C_Ping} message S2C_Ping message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        S2C_Ping.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.TimeStamp != null && Object.hasOwnProperty.call(message, "TimeStamp"))
                writer.uint32(/* id 1, wireType 0 =*/8).int64(message.TimeStamp);
            return writer;
        };

        /**
         * Encodes the specified S2C_Ping message, length delimited. Does not implicitly {@link pb_game.S2C_Ping.verify|verify} messages.
         * @function encodeDelimited
         * @memberof pb_game.S2C_Ping
         * @static
         * @param {pb_game.IS2C_Ping} message S2C_Ping message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        S2C_Ping.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a S2C_Ping message from the specified reader or buffer.
         * @function decode
         * @memberof pb_game.S2C_Ping
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {pb_game.S2C_Ping} S2C_Ping
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        S2C_Ping.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.pb_game.S2C_Ping();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.TimeStamp = reader.int64();
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a S2C_Ping message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof pb_game.S2C_Ping
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {pb_game.S2C_Ping} S2C_Ping
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        S2C_Ping.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a S2C_Ping message.
         * @function verify
         * @memberof pb_game.S2C_Ping
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        S2C_Ping.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.TimeStamp != null && message.hasOwnProperty("TimeStamp"))
                if (!$util.isInteger(message.TimeStamp) && !(message.TimeStamp && $util.isInteger(message.TimeStamp.low) && $util.isInteger(message.TimeStamp.high)))
                    return "TimeStamp: integer|Long expected";
            return null;
        };

        /**
         * Creates a S2C_Ping message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof pb_game.S2C_Ping
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {pb_game.S2C_Ping} S2C_Ping
         */
        S2C_Ping.fromObject = function fromObject(object) {
            if (object instanceof $root.pb_game.S2C_Ping)
                return object;
            var message = new $root.pb_game.S2C_Ping();
            if (object.TimeStamp != null)
                if ($util.Long)
                    (message.TimeStamp = $util.Long.fromValue(object.TimeStamp)).unsigned = false;
                else if (typeof object.TimeStamp === "string")
                    message.TimeStamp = parseInt(object.TimeStamp, 10);
                else if (typeof object.TimeStamp === "number")
                    message.TimeStamp = object.TimeStamp;
                else if (typeof object.TimeStamp === "object")
                    message.TimeStamp = new $util.LongBits(object.TimeStamp.low >>> 0, object.TimeStamp.high >>> 0).toNumber();
            return message;
        };

        /**
         * Creates a plain object from a S2C_Ping message. Also converts values to other types if specified.
         * @function toObject
         * @memberof pb_game.S2C_Ping
         * @static
         * @param {pb_game.S2C_Ping} message S2C_Ping
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        S2C_Ping.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.TimeStamp = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.TimeStamp = options.longs === String ? "0" : 0;
            if (message.TimeStamp != null && message.hasOwnProperty("TimeStamp"))
                if (typeof message.TimeStamp === "number")
                    object.TimeStamp = options.longs === String ? String(message.TimeStamp) : message.TimeStamp;
                else
                    object.TimeStamp = options.longs === String ? $util.Long.prototype.toString.call(message.TimeStamp) : options.longs === Number ? new $util.LongBits(message.TimeStamp.low >>> 0, message.TimeStamp.high >>> 0).toNumber() : message.TimeStamp;
            return object;
        };

        /**
         * Converts this S2C_Ping to JSON.
         * @function toJSON
         * @memberof pb_game.S2C_Ping
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        S2C_Ping.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for S2C_Ping
         * @function getTypeUrl
         * @memberof pb_game.S2C_Ping
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        S2C_Ping.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/pb_game.S2C_Ping";
        };

        return S2C_Ping;
    })();

    pb_game.S2C_ServerMessage = (function() {

        /**
         * Properties of a S2C_ServerMessage.
         * @memberof pb_game
         * @interface IS2C_ServerMessage
         * @property {string|null} [Content] S2C_ServerMessage Content
         */

        /**
         * Constructs a new S2C_ServerMessage.
         * @memberof pb_game
         * @classdesc Represents a S2C_ServerMessage.
         * @implements IS2C_ServerMessage
         * @constructor
         * @param {pb_game.IS2C_ServerMessage=} [properties] Properties to set
         */
        function S2C_ServerMessage(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * S2C_ServerMessage Content.
         * @member {string} Content
         * @memberof pb_game.S2C_ServerMessage
         * @instance
         */
        S2C_ServerMessage.prototype.Content = "";

        /**
         * Creates a new S2C_ServerMessage instance using the specified properties.
         * @function create
         * @memberof pb_game.S2C_ServerMessage
         * @static
         * @param {pb_game.IS2C_ServerMessage=} [properties] Properties to set
         * @returns {pb_game.S2C_ServerMessage} S2C_ServerMessage instance
         */
        S2C_ServerMessage.create = function create(properties) {
            return new S2C_ServerMessage(properties);
        };

        /**
         * Encodes the specified S2C_ServerMessage message. Does not implicitly {@link pb_game.S2C_ServerMessage.verify|verify} messages.
         * @function encode
         * @memberof pb_game.S2C_ServerMessage
         * @static
         * @param {pb_game.IS2C_ServerMessage} message S2C_ServerMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        S2C_ServerMessage.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.Content != null && Object.hasOwnProperty.call(message, "Content"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.Content);
            return writer;
        };

        /**
         * Encodes the specified S2C_ServerMessage message, length delimited. Does not implicitly {@link pb_game.S2C_ServerMessage.verify|verify} messages.
         * @function encodeDelimited
         * @memberof pb_game.S2C_ServerMessage
         * @static
         * @param {pb_game.IS2C_ServerMessage} message S2C_ServerMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        S2C_ServerMessage.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a S2C_ServerMessage message from the specified reader or buffer.
         * @function decode
         * @memberof pb_game.S2C_ServerMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {pb_game.S2C_ServerMessage} S2C_ServerMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        S2C_ServerMessage.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.pb_game.S2C_ServerMessage();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.Content = reader.string();
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a S2C_ServerMessage message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof pb_game.S2C_ServerMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {pb_game.S2C_ServerMessage} S2C_ServerMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        S2C_ServerMessage.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a S2C_ServerMessage message.
         * @function verify
         * @memberof pb_game.S2C_ServerMessage
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        S2C_ServerMessage.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.Content != null && message.hasOwnProperty("Content"))
                if (!$util.isString(message.Content))
                    return "Content: string expected";
            return null;
        };

        /**
         * Creates a S2C_ServerMessage message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof pb_game.S2C_ServerMessage
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {pb_game.S2C_ServerMessage} S2C_ServerMessage
         */
        S2C_ServerMessage.fromObject = function fromObject(object) {
            if (object instanceof $root.pb_game.S2C_ServerMessage)
                return object;
            var message = new $root.pb_game.S2C_ServerMessage();
            if (object.Content != null)
                message.Content = String(object.Content);
            return message;
        };

        /**
         * Creates a plain object from a S2C_ServerMessage message. Also converts values to other types if specified.
         * @function toObject
         * @memberof pb_game.S2C_ServerMessage
         * @static
         * @param {pb_game.S2C_ServerMessage} message S2C_ServerMessage
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        S2C_ServerMessage.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.Content = "";
            if (message.Content != null && message.hasOwnProperty("Content"))
                object.Content = message.Content;
            return object;
        };

        /**
         * Converts this S2C_ServerMessage to JSON.
         * @function toJSON
         * @memberof pb_game.S2C_ServerMessage
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        S2C_ServerMessage.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for S2C_ServerMessage
         * @function getTypeUrl
         * @memberof pb_game.S2C_ServerMessage
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        S2C_ServerMessage.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/pb_game.S2C_ServerMessage";
        };

        return S2C_ServerMessage;
    })();

    return pb_game;
})();

$root.pb_base = (function() {

    /**
     * Namespace pb_base.
     * @exports pb_base
     * @namespace
     */
    var pb_base = {};

    /**
     * TroopStackType enum.
     * @name pb_base.TroopStackType
     * @enum {number}
     * @property {number} Infantry=0 Infantry value
     * @property {number} Tank=1 Tank value
     * @property {number} Fort=2 Fort value
     * @property {number} Plane=3 Plane value
     */
    pb_base.TroopStackType = (function() {
        var valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "Infantry"] = 0;
        values[valuesById[1] = "Tank"] = 1;
        values[valuesById[2] = "Fort"] = 2;
        values[valuesById[3] = "Plane"] = 3;
        return values;
    })();

    pb_base.Position = (function() {

        /**
         * Properties of a Position.
         * @memberof pb_base
         * @interface IPosition
         * @property {number|null} [X] Position X
         * @property {number|null} [Y] Position Y
         */

        /**
         * Constructs a new Position.
         * @memberof pb_base
         * @classdesc Represents a Position.
         * @implements IPosition
         * @constructor
         * @param {pb_base.IPosition=} [properties] Properties to set
         */
        function Position(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Position X.
         * @member {number} X
         * @memberof pb_base.Position
         * @instance
         */
        Position.prototype.X = 0;

        /**
         * Position Y.
         * @member {number} Y
         * @memberof pb_base.Position
         * @instance
         */
        Position.prototype.Y = 0;

        /**
         * Creates a new Position instance using the specified properties.
         * @function create
         * @memberof pb_base.Position
         * @static
         * @param {pb_base.IPosition=} [properties] Properties to set
         * @returns {pb_base.Position} Position instance
         */
        Position.create = function create(properties) {
            return new Position(properties);
        };

        /**
         * Encodes the specified Position message. Does not implicitly {@link pb_base.Position.verify|verify} messages.
         * @function encode
         * @memberof pb_base.Position
         * @static
         * @param {pb_base.IPosition} message Position message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Position.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.X != null && Object.hasOwnProperty.call(message, "X"))
                writer.uint32(/* id 1, wireType 0 =*/8).int32(message.X);
            if (message.Y != null && Object.hasOwnProperty.call(message, "Y"))
                writer.uint32(/* id 2, wireType 0 =*/16).int32(message.Y);
            return writer;
        };

        /**
         * Encodes the specified Position message, length delimited. Does not implicitly {@link pb_base.Position.verify|verify} messages.
         * @function encodeDelimited
         * @memberof pb_base.Position
         * @static
         * @param {pb_base.IPosition} message Position message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Position.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Position message from the specified reader or buffer.
         * @function decode
         * @memberof pb_base.Position
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {pb_base.Position} Position
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Position.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.pb_base.Position();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.X = reader.int32();
                        break;
                    }
                case 2: {
                        message.Y = reader.int32();
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Position message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof pb_base.Position
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {pb_base.Position} Position
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Position.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Position message.
         * @function verify
         * @memberof pb_base.Position
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Position.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.X != null && message.hasOwnProperty("X"))
                if (!$util.isInteger(message.X))
                    return "X: integer expected";
            if (message.Y != null && message.hasOwnProperty("Y"))
                if (!$util.isInteger(message.Y))
                    return "Y: integer expected";
            return null;
        };

        /**
         * Creates a Position message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof pb_base.Position
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {pb_base.Position} Position
         */
        Position.fromObject = function fromObject(object) {
            if (object instanceof $root.pb_base.Position)
                return object;
            var message = new $root.pb_base.Position();
            if (object.X != null)
                message.X = object.X | 0;
            if (object.Y != null)
                message.Y = object.Y | 0;
            return message;
        };

        /**
         * Creates a plain object from a Position message. Also converts values to other types if specified.
         * @function toObject
         * @memberof pb_base.Position
         * @static
         * @param {pb_base.Position} message Position
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Position.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.X = 0;
                object.Y = 0;
            }
            if (message.X != null && message.hasOwnProperty("X"))
                object.X = message.X;
            if (message.Y != null && message.hasOwnProperty("Y"))
                object.Y = message.Y;
            return object;
        };

        /**
         * Converts this Position to JSON.
         * @function toJSON
         * @memberof pb_base.Position
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Position.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for Position
         * @function getTypeUrl
         * @memberof pb_base.Position
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        Position.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/pb_base.Position";
        };

        return Position;
    })();

    pb_base.Round = (function() {

        /**
         * Properties of a Round.
         * @memberof pb_base
         * @interface IRound
         * @property {number|null} [Count] Round Count
         */

        /**
         * Constructs a new Round.
         * @memberof pb_base
         * @classdesc Represents a Round.
         * @implements IRound
         * @constructor
         * @param {pb_base.IRound=} [properties] Properties to set
         */
        function Round(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Round Count.
         * @member {number} Count
         * @memberof pb_base.Round
         * @instance
         */
        Round.prototype.Count = 0;

        /**
         * Creates a new Round instance using the specified properties.
         * @function create
         * @memberof pb_base.Round
         * @static
         * @param {pb_base.IRound=} [properties] Properties to set
         * @returns {pb_base.Round} Round instance
         */
        Round.create = function create(properties) {
            return new Round(properties);
        };

        /**
         * Encodes the specified Round message. Does not implicitly {@link pb_base.Round.verify|verify} messages.
         * @function encode
         * @memberof pb_base.Round
         * @static
         * @param {pb_base.IRound} message Round message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Round.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.Count != null && Object.hasOwnProperty.call(message, "Count"))
                writer.uint32(/* id 1, wireType 0 =*/8).int32(message.Count);
            return writer;
        };

        /**
         * Encodes the specified Round message, length delimited. Does not implicitly {@link pb_base.Round.verify|verify} messages.
         * @function encodeDelimited
         * @memberof pb_base.Round
         * @static
         * @param {pb_base.IRound} message Round message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Round.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Round message from the specified reader or buffer.
         * @function decode
         * @memberof pb_base.Round
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {pb_base.Round} Round
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Round.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.pb_base.Round();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.Count = reader.int32();
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Round message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof pb_base.Round
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {pb_base.Round} Round
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Round.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Round message.
         * @function verify
         * @memberof pb_base.Round
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Round.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.Count != null && message.hasOwnProperty("Count"))
                if (!$util.isInteger(message.Count))
                    return "Count: integer expected";
            return null;
        };

        /**
         * Creates a Round message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof pb_base.Round
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {pb_base.Round} Round
         */
        Round.fromObject = function fromObject(object) {
            if (object instanceof $root.pb_base.Round)
                return object;
            var message = new $root.pb_base.Round();
            if (object.Count != null)
                message.Count = object.Count | 0;
            return message;
        };

        /**
         * Creates a plain object from a Round message. Also converts values to other types if specified.
         * @function toObject
         * @memberof pb_base.Round
         * @static
         * @param {pb_base.Round} message Round
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Round.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.Count = 0;
            if (message.Count != null && message.hasOwnProperty("Count"))
                object.Count = message.Count;
            return object;
        };

        /**
         * Converts this Round to JSON.
         * @function toJSON
         * @memberof pb_base.Round
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Round.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for Round
         * @function getTypeUrl
         * @memberof pb_base.Round
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        Round.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/pb_base.Round";
        };

        return Round;
    })();

    pb_base.Player = (function() {

        /**
         * Properties of a Player.
         * @memberof pb_base
         * @interface IPlayer
         * @property {number|Long|null} [Id] Player Id
         * @property {string|null} [Name] Player Name
         * @property {string|null} [Role] Player Role
         */

        /**
         * Constructs a new Player.
         * @memberof pb_base
         * @classdesc Represents a Player.
         * @implements IPlayer
         * @constructor
         * @param {pb_base.IPlayer=} [properties] Properties to set
         */
        function Player(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Player Id.
         * @member {number|Long} Id
         * @memberof pb_base.Player
         * @instance
         */
        Player.prototype.Id = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Player Name.
         * @member {string} Name
         * @memberof pb_base.Player
         * @instance
         */
        Player.prototype.Name = "";

        /**
         * Player Role.
         * @member {string} Role
         * @memberof pb_base.Player
         * @instance
         */
        Player.prototype.Role = "";

        /**
         * Creates a new Player instance using the specified properties.
         * @function create
         * @memberof pb_base.Player
         * @static
         * @param {pb_base.IPlayer=} [properties] Properties to set
         * @returns {pb_base.Player} Player instance
         */
        Player.create = function create(properties) {
            return new Player(properties);
        };

        /**
         * Encodes the specified Player message. Does not implicitly {@link pb_base.Player.verify|verify} messages.
         * @function encode
         * @memberof pb_base.Player
         * @static
         * @param {pb_base.IPlayer} message Player message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Player.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.Id != null && Object.hasOwnProperty.call(message, "Id"))
                writer.uint32(/* id 1, wireType 0 =*/8).int64(message.Id);
            if (message.Name != null && Object.hasOwnProperty.call(message, "Name"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.Name);
            if (message.Role != null && Object.hasOwnProperty.call(message, "Role"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.Role);
            return writer;
        };

        /**
         * Encodes the specified Player message, length delimited. Does not implicitly {@link pb_base.Player.verify|verify} messages.
         * @function encodeDelimited
         * @memberof pb_base.Player
         * @static
         * @param {pb_base.IPlayer} message Player message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Player.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Player message from the specified reader or buffer.
         * @function decode
         * @memberof pb_base.Player
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {pb_base.Player} Player
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Player.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.pb_base.Player();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.Id = reader.int64();
                        break;
                    }
                case 2: {
                        message.Name = reader.string();
                        break;
                    }
                case 3: {
                        message.Role = reader.string();
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Player message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof pb_base.Player
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {pb_base.Player} Player
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Player.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Player message.
         * @function verify
         * @memberof pb_base.Player
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Player.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.Id != null && message.hasOwnProperty("Id"))
                if (!$util.isInteger(message.Id) && !(message.Id && $util.isInteger(message.Id.low) && $util.isInteger(message.Id.high)))
                    return "Id: integer|Long expected";
            if (message.Name != null && message.hasOwnProperty("Name"))
                if (!$util.isString(message.Name))
                    return "Name: string expected";
            if (message.Role != null && message.hasOwnProperty("Role"))
                if (!$util.isString(message.Role))
                    return "Role: string expected";
            return null;
        };

        /**
         * Creates a Player message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof pb_base.Player
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {pb_base.Player} Player
         */
        Player.fromObject = function fromObject(object) {
            if (object instanceof $root.pb_base.Player)
                return object;
            var message = new $root.pb_base.Player();
            if (object.Id != null)
                if ($util.Long)
                    (message.Id = $util.Long.fromValue(object.Id)).unsigned = false;
                else if (typeof object.Id === "string")
                    message.Id = parseInt(object.Id, 10);
                else if (typeof object.Id === "number")
                    message.Id = object.Id;
                else if (typeof object.Id === "object")
                    message.Id = new $util.LongBits(object.Id.low >>> 0, object.Id.high >>> 0).toNumber();
            if (object.Name != null)
                message.Name = String(object.Name);
            if (object.Role != null)
                message.Role = String(object.Role);
            return message;
        };

        /**
         * Creates a plain object from a Player message. Also converts values to other types if specified.
         * @function toObject
         * @memberof pb_base.Player
         * @static
         * @param {pb_base.Player} message Player
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Player.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.Id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.Id = options.longs === String ? "0" : 0;
                object.Name = "";
                object.Role = "";
            }
            if (message.Id != null && message.hasOwnProperty("Id"))
                if (typeof message.Id === "number")
                    object.Id = options.longs === String ? String(message.Id) : message.Id;
                else
                    object.Id = options.longs === String ? $util.Long.prototype.toString.call(message.Id) : options.longs === Number ? new $util.LongBits(message.Id.low >>> 0, message.Id.high >>> 0).toNumber() : message.Id;
            if (message.Name != null && message.hasOwnProperty("Name"))
                object.Name = message.Name;
            if (message.Role != null && message.hasOwnProperty("Role"))
                object.Role = message.Role;
            return object;
        };

        /**
         * Converts this Player to JSON.
         * @function toJSON
         * @memberof pb_base.Player
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Player.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for Player
         * @function getTypeUrl
         * @memberof pb_base.Player
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        Player.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/pb_base.Player";
        };

        return Player;
    })();

    pb_base.Tile = (function() {

        /**
         * Properties of a Tile.
         * @memberof pb_base
         * @interface ITile
         * @property {number|Long|null} [Id] Tile Id
         * @property {number|Long|null} [OwnerId] Tile OwnerId
         * @property {number|null} [Level] Tile Level
         * @property {pb_base.IPosition|null} [Position] Tile Position
         */

        /**
         * Constructs a new Tile.
         * @memberof pb_base
         * @classdesc Represents a Tile.
         * @implements ITile
         * @constructor
         * @param {pb_base.ITile=} [properties] Properties to set
         */
        function Tile(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Tile Id.
         * @member {number|Long} Id
         * @memberof pb_base.Tile
         * @instance
         */
        Tile.prototype.Id = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Tile OwnerId.
         * @member {number|Long} OwnerId
         * @memberof pb_base.Tile
         * @instance
         */
        Tile.prototype.OwnerId = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Tile Level.
         * @member {number} Level
         * @memberof pb_base.Tile
         * @instance
         */
        Tile.prototype.Level = 0;

        /**
         * Tile Position.
         * @member {pb_base.IPosition|null|undefined} Position
         * @memberof pb_base.Tile
         * @instance
         */
        Tile.prototype.Position = null;

        /**
         * Creates a new Tile instance using the specified properties.
         * @function create
         * @memberof pb_base.Tile
         * @static
         * @param {pb_base.ITile=} [properties] Properties to set
         * @returns {pb_base.Tile} Tile instance
         */
        Tile.create = function create(properties) {
            return new Tile(properties);
        };

        /**
         * Encodes the specified Tile message. Does not implicitly {@link pb_base.Tile.verify|verify} messages.
         * @function encode
         * @memberof pb_base.Tile
         * @static
         * @param {pb_base.ITile} message Tile message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Tile.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.Id != null && Object.hasOwnProperty.call(message, "Id"))
                writer.uint32(/* id 1, wireType 0 =*/8).int64(message.Id);
            if (message.OwnerId != null && Object.hasOwnProperty.call(message, "OwnerId"))
                writer.uint32(/* id 2, wireType 0 =*/16).int64(message.OwnerId);
            if (message.Level != null && Object.hasOwnProperty.call(message, "Level"))
                writer.uint32(/* id 3, wireType 0 =*/24).int32(message.Level);
            if (message.Position != null && Object.hasOwnProperty.call(message, "Position"))
                $root.pb_base.Position.encode(message.Position, writer.uint32(/* id 4, wireType 2 =*/34).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified Tile message, length delimited. Does not implicitly {@link pb_base.Tile.verify|verify} messages.
         * @function encodeDelimited
         * @memberof pb_base.Tile
         * @static
         * @param {pb_base.ITile} message Tile message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Tile.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Tile message from the specified reader or buffer.
         * @function decode
         * @memberof pb_base.Tile
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {pb_base.Tile} Tile
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Tile.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.pb_base.Tile();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.Id = reader.int64();
                        break;
                    }
                case 2: {
                        message.OwnerId = reader.int64();
                        break;
                    }
                case 3: {
                        message.Level = reader.int32();
                        break;
                    }
                case 4: {
                        message.Position = $root.pb_base.Position.decode(reader, reader.uint32());
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Tile message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof pb_base.Tile
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {pb_base.Tile} Tile
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Tile.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Tile message.
         * @function verify
         * @memberof pb_base.Tile
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Tile.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.Id != null && message.hasOwnProperty("Id"))
                if (!$util.isInteger(message.Id) && !(message.Id && $util.isInteger(message.Id.low) && $util.isInteger(message.Id.high)))
                    return "Id: integer|Long expected";
            if (message.OwnerId != null && message.hasOwnProperty("OwnerId"))
                if (!$util.isInteger(message.OwnerId) && !(message.OwnerId && $util.isInteger(message.OwnerId.low) && $util.isInteger(message.OwnerId.high)))
                    return "OwnerId: integer|Long expected";
            if (message.Level != null && message.hasOwnProperty("Level"))
                if (!$util.isInteger(message.Level))
                    return "Level: integer expected";
            if (message.Position != null && message.hasOwnProperty("Position")) {
                var error = $root.pb_base.Position.verify(message.Position);
                if (error)
                    return "Position." + error;
            }
            return null;
        };

        /**
         * Creates a Tile message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof pb_base.Tile
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {pb_base.Tile} Tile
         */
        Tile.fromObject = function fromObject(object) {
            if (object instanceof $root.pb_base.Tile)
                return object;
            var message = new $root.pb_base.Tile();
            if (object.Id != null)
                if ($util.Long)
                    (message.Id = $util.Long.fromValue(object.Id)).unsigned = false;
                else if (typeof object.Id === "string")
                    message.Id = parseInt(object.Id, 10);
                else if (typeof object.Id === "number")
                    message.Id = object.Id;
                else if (typeof object.Id === "object")
                    message.Id = new $util.LongBits(object.Id.low >>> 0, object.Id.high >>> 0).toNumber();
            if (object.OwnerId != null)
                if ($util.Long)
                    (message.OwnerId = $util.Long.fromValue(object.OwnerId)).unsigned = false;
                else if (typeof object.OwnerId === "string")
                    message.OwnerId = parseInt(object.OwnerId, 10);
                else if (typeof object.OwnerId === "number")
                    message.OwnerId = object.OwnerId;
                else if (typeof object.OwnerId === "object")
                    message.OwnerId = new $util.LongBits(object.OwnerId.low >>> 0, object.OwnerId.high >>> 0).toNumber();
            if (object.Level != null)
                message.Level = object.Level | 0;
            if (object.Position != null) {
                if (typeof object.Position !== "object")
                    throw TypeError(".pb_base.Tile.Position: object expected");
                message.Position = $root.pb_base.Position.fromObject(object.Position);
            }
            return message;
        };

        /**
         * Creates a plain object from a Tile message. Also converts values to other types if specified.
         * @function toObject
         * @memberof pb_base.Tile
         * @static
         * @param {pb_base.Tile} message Tile
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Tile.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.Id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.Id = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.OwnerId = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.OwnerId = options.longs === String ? "0" : 0;
                object.Level = 0;
                object.Position = null;
            }
            if (message.Id != null && message.hasOwnProperty("Id"))
                if (typeof message.Id === "number")
                    object.Id = options.longs === String ? String(message.Id) : message.Id;
                else
                    object.Id = options.longs === String ? $util.Long.prototype.toString.call(message.Id) : options.longs === Number ? new $util.LongBits(message.Id.low >>> 0, message.Id.high >>> 0).toNumber() : message.Id;
            if (message.OwnerId != null && message.hasOwnProperty("OwnerId"))
                if (typeof message.OwnerId === "number")
                    object.OwnerId = options.longs === String ? String(message.OwnerId) : message.OwnerId;
                else
                    object.OwnerId = options.longs === String ? $util.Long.prototype.toString.call(message.OwnerId) : options.longs === Number ? new $util.LongBits(message.OwnerId.low >>> 0, message.OwnerId.high >>> 0).toNumber() : message.OwnerId;
            if (message.Level != null && message.hasOwnProperty("Level"))
                object.Level = message.Level;
            if (message.Position != null && message.hasOwnProperty("Position"))
                object.Position = $root.pb_base.Position.toObject(message.Position, options);
            return object;
        };

        /**
         * Converts this Tile to JSON.
         * @function toJSON
         * @memberof pb_base.Tile
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Tile.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for Tile
         * @function getTypeUrl
         * @memberof pb_base.Tile
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        Tile.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/pb_base.Tile";
        };

        return Tile;
    })();

    pb_base.Capital = (function() {

        /**
         * Properties of a Capital.
         * @memberof pb_base
         * @interface ICapital
         * @property {number|Long|null} [Id] Capital Id
         * @property {number|Long|null} [OwnerId] Capital OwnerId
         * @property {number|null} [Level] Capital Level
         * @property {pb_base.IPosition|null} [Position] Capital Position
         */

        /**
         * Constructs a new Capital.
         * @memberof pb_base
         * @classdesc Represents a Capital.
         * @implements ICapital
         * @constructor
         * @param {pb_base.ICapital=} [properties] Properties to set
         */
        function Capital(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Capital Id.
         * @member {number|Long} Id
         * @memberof pb_base.Capital
         * @instance
         */
        Capital.prototype.Id = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Capital OwnerId.
         * @member {number|Long} OwnerId
         * @memberof pb_base.Capital
         * @instance
         */
        Capital.prototype.OwnerId = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Capital Level.
         * @member {number} Level
         * @memberof pb_base.Capital
         * @instance
         */
        Capital.prototype.Level = 0;

        /**
         * Capital Position.
         * @member {pb_base.IPosition|null|undefined} Position
         * @memberof pb_base.Capital
         * @instance
         */
        Capital.prototype.Position = null;

        /**
         * Creates a new Capital instance using the specified properties.
         * @function create
         * @memberof pb_base.Capital
         * @static
         * @param {pb_base.ICapital=} [properties] Properties to set
         * @returns {pb_base.Capital} Capital instance
         */
        Capital.create = function create(properties) {
            return new Capital(properties);
        };

        /**
         * Encodes the specified Capital message. Does not implicitly {@link pb_base.Capital.verify|verify} messages.
         * @function encode
         * @memberof pb_base.Capital
         * @static
         * @param {pb_base.ICapital} message Capital message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Capital.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.Id != null && Object.hasOwnProperty.call(message, "Id"))
                writer.uint32(/* id 1, wireType 0 =*/8).int64(message.Id);
            if (message.OwnerId != null && Object.hasOwnProperty.call(message, "OwnerId"))
                writer.uint32(/* id 2, wireType 0 =*/16).int64(message.OwnerId);
            if (message.Level != null && Object.hasOwnProperty.call(message, "Level"))
                writer.uint32(/* id 3, wireType 0 =*/24).int32(message.Level);
            if (message.Position != null && Object.hasOwnProperty.call(message, "Position"))
                $root.pb_base.Position.encode(message.Position, writer.uint32(/* id 4, wireType 2 =*/34).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified Capital message, length delimited. Does not implicitly {@link pb_base.Capital.verify|verify} messages.
         * @function encodeDelimited
         * @memberof pb_base.Capital
         * @static
         * @param {pb_base.ICapital} message Capital message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Capital.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Capital message from the specified reader or buffer.
         * @function decode
         * @memberof pb_base.Capital
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {pb_base.Capital} Capital
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Capital.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.pb_base.Capital();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.Id = reader.int64();
                        break;
                    }
                case 2: {
                        message.OwnerId = reader.int64();
                        break;
                    }
                case 3: {
                        message.Level = reader.int32();
                        break;
                    }
                case 4: {
                        message.Position = $root.pb_base.Position.decode(reader, reader.uint32());
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Capital message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof pb_base.Capital
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {pb_base.Capital} Capital
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Capital.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Capital message.
         * @function verify
         * @memberof pb_base.Capital
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Capital.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.Id != null && message.hasOwnProperty("Id"))
                if (!$util.isInteger(message.Id) && !(message.Id && $util.isInteger(message.Id.low) && $util.isInteger(message.Id.high)))
                    return "Id: integer|Long expected";
            if (message.OwnerId != null && message.hasOwnProperty("OwnerId"))
                if (!$util.isInteger(message.OwnerId) && !(message.OwnerId && $util.isInteger(message.OwnerId.low) && $util.isInteger(message.OwnerId.high)))
                    return "OwnerId: integer|Long expected";
            if (message.Level != null && message.hasOwnProperty("Level"))
                if (!$util.isInteger(message.Level))
                    return "Level: integer expected";
            if (message.Position != null && message.hasOwnProperty("Position")) {
                var error = $root.pb_base.Position.verify(message.Position);
                if (error)
                    return "Position." + error;
            }
            return null;
        };

        /**
         * Creates a Capital message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof pb_base.Capital
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {pb_base.Capital} Capital
         */
        Capital.fromObject = function fromObject(object) {
            if (object instanceof $root.pb_base.Capital)
                return object;
            var message = new $root.pb_base.Capital();
            if (object.Id != null)
                if ($util.Long)
                    (message.Id = $util.Long.fromValue(object.Id)).unsigned = false;
                else if (typeof object.Id === "string")
                    message.Id = parseInt(object.Id, 10);
                else if (typeof object.Id === "number")
                    message.Id = object.Id;
                else if (typeof object.Id === "object")
                    message.Id = new $util.LongBits(object.Id.low >>> 0, object.Id.high >>> 0).toNumber();
            if (object.OwnerId != null)
                if ($util.Long)
                    (message.OwnerId = $util.Long.fromValue(object.OwnerId)).unsigned = false;
                else if (typeof object.OwnerId === "string")
                    message.OwnerId = parseInt(object.OwnerId, 10);
                else if (typeof object.OwnerId === "number")
                    message.OwnerId = object.OwnerId;
                else if (typeof object.OwnerId === "object")
                    message.OwnerId = new $util.LongBits(object.OwnerId.low >>> 0, object.OwnerId.high >>> 0).toNumber();
            if (object.Level != null)
                message.Level = object.Level | 0;
            if (object.Position != null) {
                if (typeof object.Position !== "object")
                    throw TypeError(".pb_base.Capital.Position: object expected");
                message.Position = $root.pb_base.Position.fromObject(object.Position);
            }
            return message;
        };

        /**
         * Creates a plain object from a Capital message. Also converts values to other types if specified.
         * @function toObject
         * @memberof pb_base.Capital
         * @static
         * @param {pb_base.Capital} message Capital
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Capital.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.Id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.Id = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.OwnerId = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.OwnerId = options.longs === String ? "0" : 0;
                object.Level = 0;
                object.Position = null;
            }
            if (message.Id != null && message.hasOwnProperty("Id"))
                if (typeof message.Id === "number")
                    object.Id = options.longs === String ? String(message.Id) : message.Id;
                else
                    object.Id = options.longs === String ? $util.Long.prototype.toString.call(message.Id) : options.longs === Number ? new $util.LongBits(message.Id.low >>> 0, message.Id.high >>> 0).toNumber() : message.Id;
            if (message.OwnerId != null && message.hasOwnProperty("OwnerId"))
                if (typeof message.OwnerId === "number")
                    object.OwnerId = options.longs === String ? String(message.OwnerId) : message.OwnerId;
                else
                    object.OwnerId = options.longs === String ? $util.Long.prototype.toString.call(message.OwnerId) : options.longs === Number ? new $util.LongBits(message.OwnerId.low >>> 0, message.OwnerId.high >>> 0).toNumber() : message.OwnerId;
            if (message.Level != null && message.hasOwnProperty("Level"))
                object.Level = message.Level;
            if (message.Position != null && message.hasOwnProperty("Position"))
                object.Position = $root.pb_base.Position.toObject(message.Position, options);
            return object;
        };

        /**
         * Converts this Capital to JSON.
         * @function toJSON
         * @memberof pb_base.Capital
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Capital.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for Capital
         * @function getTypeUrl
         * @memberof pb_base.Capital
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        Capital.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/pb_base.Capital";
        };

        return Capital;
    })();

    pb_base.TroopStack = (function() {

        /**
         * Properties of a TroopStack.
         * @memberof pb_base
         * @interface ITroopStack
         * @property {number|Long|null} [Id] TroopStack Id
         * @property {number|Long|null} [OwnerId] TroopStack OwnerId
         * @property {pb_base.TroopStackType|null} [Type] TroopStack Type
         * @property {number|null} [Amount] TroopStack Amount
         * @property {pb_base.IPosition|null} [Position] TroopStack Position
         * @property {number|null} [MovementStamina] TroopStack MovementStamina
         */

        /**
         * Constructs a new TroopStack.
         * @memberof pb_base
         * @classdesc Represents a TroopStack.
         * @implements ITroopStack
         * @constructor
         * @param {pb_base.ITroopStack=} [properties] Properties to set
         */
        function TroopStack(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * TroopStack Id.
         * @member {number|Long} Id
         * @memberof pb_base.TroopStack
         * @instance
         */
        TroopStack.prototype.Id = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * TroopStack OwnerId.
         * @member {number|Long} OwnerId
         * @memberof pb_base.TroopStack
         * @instance
         */
        TroopStack.prototype.OwnerId = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * TroopStack Type.
         * @member {pb_base.TroopStackType} Type
         * @memberof pb_base.TroopStack
         * @instance
         */
        TroopStack.prototype.Type = 0;

        /**
         * TroopStack Amount.
         * @member {number} Amount
         * @memberof pb_base.TroopStack
         * @instance
         */
        TroopStack.prototype.Amount = 0;

        /**
         * TroopStack Position.
         * @member {pb_base.IPosition|null|undefined} Position
         * @memberof pb_base.TroopStack
         * @instance
         */
        TroopStack.prototype.Position = null;

        /**
         * TroopStack MovementStamina.
         * @member {number} MovementStamina
         * @memberof pb_base.TroopStack
         * @instance
         */
        TroopStack.prototype.MovementStamina = 0;

        /**
         * Creates a new TroopStack instance using the specified properties.
         * @function create
         * @memberof pb_base.TroopStack
         * @static
         * @param {pb_base.ITroopStack=} [properties] Properties to set
         * @returns {pb_base.TroopStack} TroopStack instance
         */
        TroopStack.create = function create(properties) {
            return new TroopStack(properties);
        };

        /**
         * Encodes the specified TroopStack message. Does not implicitly {@link pb_base.TroopStack.verify|verify} messages.
         * @function encode
         * @memberof pb_base.TroopStack
         * @static
         * @param {pb_base.ITroopStack} message TroopStack message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        TroopStack.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.Id != null && Object.hasOwnProperty.call(message, "Id"))
                writer.uint32(/* id 1, wireType 0 =*/8).int64(message.Id);
            if (message.OwnerId != null && Object.hasOwnProperty.call(message, "OwnerId"))
                writer.uint32(/* id 2, wireType 0 =*/16).int64(message.OwnerId);
            if (message.Type != null && Object.hasOwnProperty.call(message, "Type"))
                writer.uint32(/* id 3, wireType 0 =*/24).int32(message.Type);
            if (message.Amount != null && Object.hasOwnProperty.call(message, "Amount"))
                writer.uint32(/* id 4, wireType 0 =*/32).int32(message.Amount);
            if (message.Position != null && Object.hasOwnProperty.call(message, "Position"))
                $root.pb_base.Position.encode(message.Position, writer.uint32(/* id 5, wireType 2 =*/42).fork()).ldelim();
            if (message.MovementStamina != null && Object.hasOwnProperty.call(message, "MovementStamina"))
                writer.uint32(/* id 6, wireType 0 =*/48).int32(message.MovementStamina);
            return writer;
        };

        /**
         * Encodes the specified TroopStack message, length delimited. Does not implicitly {@link pb_base.TroopStack.verify|verify} messages.
         * @function encodeDelimited
         * @memberof pb_base.TroopStack
         * @static
         * @param {pb_base.ITroopStack} message TroopStack message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        TroopStack.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a TroopStack message from the specified reader or buffer.
         * @function decode
         * @memberof pb_base.TroopStack
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {pb_base.TroopStack} TroopStack
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        TroopStack.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.pb_base.TroopStack();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.Id = reader.int64();
                        break;
                    }
                case 2: {
                        message.OwnerId = reader.int64();
                        break;
                    }
                case 3: {
                        message.Type = reader.int32();
                        break;
                    }
                case 4: {
                        message.Amount = reader.int32();
                        break;
                    }
                case 5: {
                        message.Position = $root.pb_base.Position.decode(reader, reader.uint32());
                        break;
                    }
                case 6: {
                        message.MovementStamina = reader.int32();
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a TroopStack message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof pb_base.TroopStack
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {pb_base.TroopStack} TroopStack
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        TroopStack.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a TroopStack message.
         * @function verify
         * @memberof pb_base.TroopStack
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        TroopStack.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.Id != null && message.hasOwnProperty("Id"))
                if (!$util.isInteger(message.Id) && !(message.Id && $util.isInteger(message.Id.low) && $util.isInteger(message.Id.high)))
                    return "Id: integer|Long expected";
            if (message.OwnerId != null && message.hasOwnProperty("OwnerId"))
                if (!$util.isInteger(message.OwnerId) && !(message.OwnerId && $util.isInteger(message.OwnerId.low) && $util.isInteger(message.OwnerId.high)))
                    return "OwnerId: integer|Long expected";
            if (message.Type != null && message.hasOwnProperty("Type"))
                switch (message.Type) {
                default:
                    return "Type: enum value expected";
                case 0:
                case 1:
                case 2:
                case 3:
                    break;
                }
            if (message.Amount != null && message.hasOwnProperty("Amount"))
                if (!$util.isInteger(message.Amount))
                    return "Amount: integer expected";
            if (message.Position != null && message.hasOwnProperty("Position")) {
                var error = $root.pb_base.Position.verify(message.Position);
                if (error)
                    return "Position." + error;
            }
            if (message.MovementStamina != null && message.hasOwnProperty("MovementStamina"))
                if (!$util.isInteger(message.MovementStamina))
                    return "MovementStamina: integer expected";
            return null;
        };

        /**
         * Creates a TroopStack message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof pb_base.TroopStack
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {pb_base.TroopStack} TroopStack
         */
        TroopStack.fromObject = function fromObject(object) {
            if (object instanceof $root.pb_base.TroopStack)
                return object;
            var message = new $root.pb_base.TroopStack();
            if (object.Id != null)
                if ($util.Long)
                    (message.Id = $util.Long.fromValue(object.Id)).unsigned = false;
                else if (typeof object.Id === "string")
                    message.Id = parseInt(object.Id, 10);
                else if (typeof object.Id === "number")
                    message.Id = object.Id;
                else if (typeof object.Id === "object")
                    message.Id = new $util.LongBits(object.Id.low >>> 0, object.Id.high >>> 0).toNumber();
            if (object.OwnerId != null)
                if ($util.Long)
                    (message.OwnerId = $util.Long.fromValue(object.OwnerId)).unsigned = false;
                else if (typeof object.OwnerId === "string")
                    message.OwnerId = parseInt(object.OwnerId, 10);
                else if (typeof object.OwnerId === "number")
                    message.OwnerId = object.OwnerId;
                else if (typeof object.OwnerId === "object")
                    message.OwnerId = new $util.LongBits(object.OwnerId.low >>> 0, object.OwnerId.high >>> 0).toNumber();
            switch (object.Type) {
            default:
                if (typeof object.Type === "number") {
                    message.Type = object.Type;
                    break;
                }
                break;
            case "Infantry":
            case 0:
                message.Type = 0;
                break;
            case "Tank":
            case 1:
                message.Type = 1;
                break;
            case "Fort":
            case 2:
                message.Type = 2;
                break;
            case "Plane":
            case 3:
                message.Type = 3;
                break;
            }
            if (object.Amount != null)
                message.Amount = object.Amount | 0;
            if (object.Position != null) {
                if (typeof object.Position !== "object")
                    throw TypeError(".pb_base.TroopStack.Position: object expected");
                message.Position = $root.pb_base.Position.fromObject(object.Position);
            }
            if (object.MovementStamina != null)
                message.MovementStamina = object.MovementStamina | 0;
            return message;
        };

        /**
         * Creates a plain object from a TroopStack message. Also converts values to other types if specified.
         * @function toObject
         * @memberof pb_base.TroopStack
         * @static
         * @param {pb_base.TroopStack} message TroopStack
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        TroopStack.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.Id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.Id = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.OwnerId = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.OwnerId = options.longs === String ? "0" : 0;
                object.Type = options.enums === String ? "Infantry" : 0;
                object.Amount = 0;
                object.Position = null;
                object.MovementStamina = 0;
            }
            if (message.Id != null && message.hasOwnProperty("Id"))
                if (typeof message.Id === "number")
                    object.Id = options.longs === String ? String(message.Id) : message.Id;
                else
                    object.Id = options.longs === String ? $util.Long.prototype.toString.call(message.Id) : options.longs === Number ? new $util.LongBits(message.Id.low >>> 0, message.Id.high >>> 0).toNumber() : message.Id;
            if (message.OwnerId != null && message.hasOwnProperty("OwnerId"))
                if (typeof message.OwnerId === "number")
                    object.OwnerId = options.longs === String ? String(message.OwnerId) : message.OwnerId;
                else
                    object.OwnerId = options.longs === String ? $util.Long.prototype.toString.call(message.OwnerId) : options.longs === Number ? new $util.LongBits(message.OwnerId.low >>> 0, message.OwnerId.high >>> 0).toNumber() : message.OwnerId;
            if (message.Type != null && message.hasOwnProperty("Type"))
                object.Type = options.enums === String ? $root.pb_base.TroopStackType[message.Type] === undefined ? message.Type : $root.pb_base.TroopStackType[message.Type] : message.Type;
            if (message.Amount != null && message.hasOwnProperty("Amount"))
                object.Amount = message.Amount;
            if (message.Position != null && message.hasOwnProperty("Position"))
                object.Position = $root.pb_base.Position.toObject(message.Position, options);
            if (message.MovementStamina != null && message.hasOwnProperty("MovementStamina"))
                object.MovementStamina = message.MovementStamina;
            return object;
        };

        /**
         * Converts this TroopStack to JSON.
         * @function toJSON
         * @memberof pb_base.TroopStack
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        TroopStack.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for TroopStack
         * @function getTypeUrl
         * @memberof pb_base.TroopStack
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        TroopStack.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/pb_base.TroopStack";
        };

        return TroopStack;
    })();

    pb_base.IdentityPayload = (function() {

        /**
         * Properties of an IdentityPayload.
         * @memberof pb_base
         * @interface IIdentityPayload
         * @property {string|null} [JwtToken] IdentityPayload JwtToken
         * @property {number|Long|null} [PlayerId] IdentityPayload PlayerId
         */

        /**
         * Constructs a new IdentityPayload.
         * @memberof pb_base
         * @classdesc Represents an IdentityPayload.
         * @implements IIdentityPayload
         * @constructor
         * @param {pb_base.IIdentityPayload=} [properties] Properties to set
         */
        function IdentityPayload(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * IdentityPayload JwtToken.
         * @member {string} JwtToken
         * @memberof pb_base.IdentityPayload
         * @instance
         */
        IdentityPayload.prototype.JwtToken = "";

        /**
         * IdentityPayload PlayerId.
         * @member {number|Long} PlayerId
         * @memberof pb_base.IdentityPayload
         * @instance
         */
        IdentityPayload.prototype.PlayerId = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Creates a new IdentityPayload instance using the specified properties.
         * @function create
         * @memberof pb_base.IdentityPayload
         * @static
         * @param {pb_base.IIdentityPayload=} [properties] Properties to set
         * @returns {pb_base.IdentityPayload} IdentityPayload instance
         */
        IdentityPayload.create = function create(properties) {
            return new IdentityPayload(properties);
        };

        /**
         * Encodes the specified IdentityPayload message. Does not implicitly {@link pb_base.IdentityPayload.verify|verify} messages.
         * @function encode
         * @memberof pb_base.IdentityPayload
         * @static
         * @param {pb_base.IIdentityPayload} message IdentityPayload message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        IdentityPayload.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.JwtToken != null && Object.hasOwnProperty.call(message, "JwtToken"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.JwtToken);
            if (message.PlayerId != null && Object.hasOwnProperty.call(message, "PlayerId"))
                writer.uint32(/* id 2, wireType 0 =*/16).int64(message.PlayerId);
            return writer;
        };

        /**
         * Encodes the specified IdentityPayload message, length delimited. Does not implicitly {@link pb_base.IdentityPayload.verify|verify} messages.
         * @function encodeDelimited
         * @memberof pb_base.IdentityPayload
         * @static
         * @param {pb_base.IIdentityPayload} message IdentityPayload message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        IdentityPayload.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an IdentityPayload message from the specified reader or buffer.
         * @function decode
         * @memberof pb_base.IdentityPayload
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {pb_base.IdentityPayload} IdentityPayload
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        IdentityPayload.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.pb_base.IdentityPayload();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.JwtToken = reader.string();
                        break;
                    }
                case 2: {
                        message.PlayerId = reader.int64();
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an IdentityPayload message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof pb_base.IdentityPayload
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {pb_base.IdentityPayload} IdentityPayload
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        IdentityPayload.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an IdentityPayload message.
         * @function verify
         * @memberof pb_base.IdentityPayload
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        IdentityPayload.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.JwtToken != null && message.hasOwnProperty("JwtToken"))
                if (!$util.isString(message.JwtToken))
                    return "JwtToken: string expected";
            if (message.PlayerId != null && message.hasOwnProperty("PlayerId"))
                if (!$util.isInteger(message.PlayerId) && !(message.PlayerId && $util.isInteger(message.PlayerId.low) && $util.isInteger(message.PlayerId.high)))
                    return "PlayerId: integer|Long expected";
            return null;
        };

        /**
         * Creates an IdentityPayload message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof pb_base.IdentityPayload
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {pb_base.IdentityPayload} IdentityPayload
         */
        IdentityPayload.fromObject = function fromObject(object) {
            if (object instanceof $root.pb_base.IdentityPayload)
                return object;
            var message = new $root.pb_base.IdentityPayload();
            if (object.JwtToken != null)
                message.JwtToken = String(object.JwtToken);
            if (object.PlayerId != null)
                if ($util.Long)
                    (message.PlayerId = $util.Long.fromValue(object.PlayerId)).unsigned = false;
                else if (typeof object.PlayerId === "string")
                    message.PlayerId = parseInt(object.PlayerId, 10);
                else if (typeof object.PlayerId === "number")
                    message.PlayerId = object.PlayerId;
                else if (typeof object.PlayerId === "object")
                    message.PlayerId = new $util.LongBits(object.PlayerId.low >>> 0, object.PlayerId.high >>> 0).toNumber();
            return message;
        };

        /**
         * Creates a plain object from an IdentityPayload message. Also converts values to other types if specified.
         * @function toObject
         * @memberof pb_base.IdentityPayload
         * @static
         * @param {pb_base.IdentityPayload} message IdentityPayload
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        IdentityPayload.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.JwtToken = "";
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.PlayerId = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.PlayerId = options.longs === String ? "0" : 0;
            }
            if (message.JwtToken != null && message.hasOwnProperty("JwtToken"))
                object.JwtToken = message.JwtToken;
            if (message.PlayerId != null && message.hasOwnProperty("PlayerId"))
                if (typeof message.PlayerId === "number")
                    object.PlayerId = options.longs === String ? String(message.PlayerId) : message.PlayerId;
                else
                    object.PlayerId = options.longs === String ? $util.Long.prototype.toString.call(message.PlayerId) : options.longs === Number ? new $util.LongBits(message.PlayerId.low >>> 0, message.PlayerId.high >>> 0).toNumber() : message.PlayerId;
            return object;
        };

        /**
         * Converts this IdentityPayload to JSON.
         * @function toJSON
         * @memberof pb_base.IdentityPayload
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        IdentityPayload.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for IdentityPayload
         * @function getTypeUrl
         * @memberof pb_base.IdentityPayload
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        IdentityPayload.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/pb_base.IdentityPayload";
        };

        return IdentityPayload;
    })();

    return pb_base;
})();

module.exports = $root;
