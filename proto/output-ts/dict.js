/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

$root.pb_dict = (function() {

    /**
     * Namespace pb_dict.
     * @exports pb_dict
     * @namespace
     */
    var pb_dict = {};

    /**
     * CMD enum.
     * @name pb_dict.CMD
     * @enum {number}
     * @property {number} NONE=0 NONE value
     * @property {number} pb_game_C2S_EstablishPlayer=1000 pb_game_C2S_EstablishPlayer value
     * @property {number} pb_game_C2S_GameState=1001 pb_game_C2S_GameState value
     * @property {number} pb_game_C2S_Ping=1002 pb_game_C2S_Ping value
     * @property {number} pb_game_S2C_EstablishPlayer=1100 pb_game_S2C_EstablishPlayer value
     * @property {number} pb_game_S2C_GameState=1101 pb_game_S2C_GameState value
     * @property {number} pb_game_S2C_Ping=1102 pb_game_S2C_Ping value
     * @property {number} pb_game_S2C_ServerMessage=1050 pb_game_S2C_ServerMessage value
     * @property {number} pb_round_C2S_DiscardCards=2000 pb_round_C2S_DiscardCards value
     * @property {number} pb_round_S2C_RoundStart=2100 pb_round_S2C_RoundStart value
     * @property {number} pb_round_S2C_DiscardPhaseEnded=2101 pb_round_S2C_DiscardPhaseEnded value
     * @property {number} pb_round_S2C_DiscardCards=2102 pb_round_S2C_DiscardCards value
     * @property {number} pb_round_S2C_EndGame=2103 pb_round_S2C_EndGame value
     * @property {number} pb_battle_C2S_Produce=3000 pb_battle_C2S_Produce value
     * @property {number} pb_battle_C2S_MoveTroops=3001 pb_battle_C2S_MoveTroops value
     * @property {number} pb_battle_C2S_Attack=3002 pb_battle_C2S_Attack value
     * @property {number} pb_battle_C2S_ToggleTankGuardMode=3003 pb_battle_C2S_ToggleTankGuardMode value
     * @property {number} pb_battle_S2C_Produce=3100 pb_battle_S2C_Produce value
     * @property {number} pb_battle_S2C_MoveTroops=3101 pb_battle_S2C_MoveTroops value
     * @property {number} pb_battle_S2C_Attack=3102 pb_battle_S2C_Attack value
     * @property {number} pb_battle_S2C_ToggleTankGuardMode=3103 pb_battle_S2C_ToggleTankGuardMode value
     * @property {number} pb_battle_S2C_MergeTroops=3104 pb_battle_S2C_MergeTroops value
     * @property {number} pb_battle_S2C_MoveTroopsPlanningPath=3105 pb_battle_S2C_MoveTroopsPlanningPath value
     */
    pb_dict.CMD = (function() {
        var valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "NONE"] = 0;
        values[valuesById[1000] = "pb_game_C2S_EstablishPlayer"] = 1000;
        values[valuesById[1001] = "pb_game_C2S_GameState"] = 1001;
        values[valuesById[1002] = "pb_game_C2S_Ping"] = 1002;
        values[valuesById[1100] = "pb_game_S2C_EstablishPlayer"] = 1100;
        values[valuesById[1101] = "pb_game_S2C_GameState"] = 1101;
        values[valuesById[1102] = "pb_game_S2C_Ping"] = 1102;
        values[valuesById[1050] = "pb_game_S2C_ServerMessage"] = 1050;
        values[valuesById[2000] = "pb_round_C2S_DiscardCards"] = 2000;
        values[valuesById[2100] = "pb_round_S2C_RoundStart"] = 2100;
        values[valuesById[2101] = "pb_round_S2C_DiscardPhaseEnded"] = 2101;
        values[valuesById[2102] = "pb_round_S2C_DiscardCards"] = 2102;
        values[valuesById[2103] = "pb_round_S2C_EndGame"] = 2103;
        values[valuesById[3000] = "pb_battle_C2S_Produce"] = 3000;
        values[valuesById[3001] = "pb_battle_C2S_MoveTroops"] = 3001;
        values[valuesById[3002] = "pb_battle_C2S_Attack"] = 3002;
        values[valuesById[3003] = "pb_battle_C2S_ToggleTankGuardMode"] = 3003;
        values[valuesById[3100] = "pb_battle_S2C_Produce"] = 3100;
        values[valuesById[3101] = "pb_battle_S2C_MoveTroops"] = 3101;
        values[valuesById[3102] = "pb_battle_S2C_Attack"] = 3102;
        values[valuesById[3103] = "pb_battle_S2C_ToggleTankGuardMode"] = 3103;
        values[valuesById[3104] = "pb_battle_S2C_MergeTroops"] = 3104;
        values[valuesById[3105] = "pb_battle_S2C_MoveTroopsPlanningPath"] = 3105;
        return values;
    })();

    return pb_dict;
})();

module.exports = $root;
