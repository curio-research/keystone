import * as $protobuf from "protobufjs";
import Long = require("long");
/** Namespace pb_dict. */
export namespace pb_dict {

    /** CMD enum. */
    enum CMD {
        NONE = 0,
        pb_game_C2S_EstablishPlayer = 1000,
        pb_game_C2S_GameState = 1001,
        pb_game_C2S_Ping = 1002,
        pb_game_S2C_EstablishPlayer = 1100,
        pb_game_S2C_GameState = 1101,
        pb_game_S2C_Ping = 1102,
        pb_game_S2C_ServerMessage = 1050,
        pb_round_C2S_DiscardCards = 2000,
        pb_round_S2C_RoundStart = 2100,
        pb_round_S2C_DiscardPhaseEnded = 2101,
        pb_round_S2C_DiscardCards = 2102,
        pb_round_S2C_EndGame = 2103,
        pb_battle_C2S_Produce = 3000,
        pb_battle_C2S_MoveTroops = 3001,
        pb_battle_C2S_Attack = 3002,
        pb_battle_C2S_ToggleTankGuardMode = 3003,
        pb_battle_S2C_Produce = 3100,
        pb_battle_S2C_MoveTroops = 3101,
        pb_battle_S2C_Attack = 3102,
        pb_battle_S2C_ToggleTankGuardMode = 3103,
        pb_battle_S2C_MergeTroops = 3104,
        pb_battle_S2C_MoveTroopsPlanningPath = 3105
    }
}
