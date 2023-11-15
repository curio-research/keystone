/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

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
