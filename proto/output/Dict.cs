// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: dict.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace PbDict {

  /// <summary>Holder for reflection information generated from dict.proto</summary>
  public static partial class DictReflection {

    #region Descriptor
    /// <summary>File descriptor for dict.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static DictReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CgpkaWN0LnByb3RvEgdwYl9kaWN0KqMICgNDTUQSCAoETk9ORRAAEh4KGnBi",
            "X3Rlc3RfQzJTX1Jlc2V0R2FtZVN0YXRlEGQSHwoacGJfdGVzdF9TMkNfUmVz",
            "ZXRHYW1lU3RhdGUQyAESIAobcGJfZ2FtZV9DMlNfRXN0YWJsaXNoUGxheWVy",
            "EOgHEhoKFXBiX2dhbWVfQzJTX0dhbWVTdGF0ZRDpBxIVChBwYl9nYW1lX0My",
            "U19QaW5nEOoHEiAKG3BiX2dhbWVfUzJDX0VzdGFibGlzaFBsYXllchDMCBIa",
            "ChVwYl9nYW1lX1MyQ19HYW1lU3RhdGUQzQgSFQoQcGJfZ2FtZV9TMkNfUGlu",
            "ZxDOCBIeChlwYl9nYW1lX1MyQ19TZXJ2ZXJNZXNzYWdlEJoIEh4KGXBiX3Jv",
            "dW5kX0MyU19EaXNjYXJkQ2FyZHMQ0A8SHAoXcGJfcm91bmRfUzJDX1JvdW5k",
            "U3RhcnQQtBASIwoecGJfcm91bmRfUzJDX0Rpc2NhcmRQaGFzZUVuZGVkELUQ",
            "Eh4KGXBiX3JvdW5kX1MyQ19EaXNjYXJkQ2FyZHMQthASGQoUcGJfcm91bmRf",
            "UzJDX0VuZEdhbWUQtxASGgoVcGJfYmF0dGxlX0MyU19Qcm9kdWNlELgXEh0K",
            "GHBiX2JhdHRsZV9DMlNfTW92ZVRyb29wcxC5FxIZChRwYl9iYXR0bGVfQzJT",
            "X0F0dGFjaxC6FxImCiFwYl9iYXR0bGVfQzJTX1RvZ2dsZVRhbmtHdWFyZE1v",
            "ZGUQuxcSIgodcGJfYmF0dGxlX0MyU19Qcm9kdWNlQnVpbGRpbmcQvBcSHAoX",
            "cGJfYmF0dGxlX0MyU19QbGFuZUxvYWQQvRcSHgoZcGJfYmF0dGxlX0MyU19Q",
            "bGFuZVVubG9hZBC+FxIaChVwYl9iYXR0bGVfUzJDX1Byb2R1Y2UQnBgSHwoa",
            "cGJfYmF0dGxlX1MyQ19Nb3ZlQ29tcGxldGUQnRgSGQoUcGJfYmF0dGxlX1My",
            "Q19BdHRhY2sQnhgSJgohcGJfYmF0dGxlX1MyQ19Ub2dnbGVUYW5rR3VhcmRN",
            "b2RlEJ8YEh4KGXBiX2JhdHRsZV9TMkNfTWVyZ2VUcm9vcHMQoBgSKQokcGJf",
            "YmF0dGxlX1MyQ19Nb3ZlVHJvb3BzUGxhbm5pbmdQYXRoEKEYEiEKHHBiX2Jh",
            "dHRsZV9TMkNfQ3Jvc3NMYXJnZVRpbGUQohgSIgodcGJfYmF0dGxlX1MyQ19M",
            "YXJnZVRpbGVDaGFuZ2UQoxgSIgodcGJfYmF0dGxlX1MyQ19Qcm9kdWNlQnVp",
            "bGRpbmcQpBgSHAoXcGJfYmF0dGxlX1MyQ19QbGFuZUxvYWQQpRgSHgoZcGJf",
            "YmF0dGxlX1MyQ19QbGFuZVVubG9hZBCmGBIVChBwYl90ZXN0X0MyU19UZXN0",
            "EKhGQglaB3BiLmRpY3RiBnByb3RvMw=="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::PbDict.CMD), }, null, null));
    }
    #endregion

  }
  #region Enums
  public enum CMD {
    [pbr::OriginalName("NONE")] None = 0,
    /// <summary>
    /// Test
    /// </summary>
    [pbr::OriginalName("pb_test_C2S_ResetGameState")] PbTestC2SResetGameState = 100,
    [pbr::OriginalName("pb_test_S2C_ResetGameState")] PbTestS2CResetGameState = 200,
    /// <summary>
    /// Game Common
    /// </summary>
    [pbr::OriginalName("pb_game_C2S_EstablishPlayer")] PbGameC2SEstablishPlayer = 1000,
    [pbr::OriginalName("pb_game_C2S_GameState")] PbGameC2SGameState = 1001,
    [pbr::OriginalName("pb_game_C2S_Ping")] PbGameC2SPing = 1002,
    [pbr::OriginalName("pb_game_S2C_EstablishPlayer")] PbGameS2CEstablishPlayer = 1100,
    [pbr::OriginalName("pb_game_S2C_GameState")] PbGameS2CGameState = 1101,
    [pbr::OriginalName("pb_game_S2C_Ping")] PbGameS2CPing = 1102,
    [pbr::OriginalName("pb_game_S2C_ServerMessage")] PbGameS2CServerMessage = 1050,
    /// <summary>
    /// Round
    /// </summary>
    [pbr::OriginalName("pb_round_C2S_DiscardCards")] PbRoundC2SDiscardCards = 2000,
    [pbr::OriginalName("pb_round_S2C_RoundStart")] PbRoundS2CRoundStart = 2100,
    [pbr::OriginalName("pb_round_S2C_DiscardPhaseEnded")] PbRoundS2CDiscardPhaseEnded = 2101,
    [pbr::OriginalName("pb_round_S2C_DiscardCards")] PbRoundS2CDiscardCards = 2102,
    [pbr::OriginalName("pb_round_S2C_EndGame")] PbRoundS2CEndGame = 2103,
    /// <summary>
    /// Battle
    /// </summary>
    [pbr::OriginalName("pb_battle_C2S_Produce")] PbBattleC2SProduce = 3000,
    [pbr::OriginalName("pb_battle_C2S_MoveTroops")] PbBattleC2SMoveTroops = 3001,
    [pbr::OriginalName("pb_battle_C2S_Attack")] PbBattleC2SAttack = 3002,
    [pbr::OriginalName("pb_battle_C2S_ToggleTankGuardMode")] PbBattleC2SToggleTankGuardMode = 3003,
    [pbr::OriginalName("pb_battle_C2S_ProduceBuilding")] PbBattleC2SProduceBuilding = 3004,
    [pbr::OriginalName("pb_battle_C2S_PlaneLoad")] PbBattleC2SPlaneLoad = 3005,
    [pbr::OriginalName("pb_battle_C2S_PlaneUnload")] PbBattleC2SPlaneUnload = 3006,
    [pbr::OriginalName("pb_battle_S2C_Produce")] PbBattleS2CProduce = 3100,
    [pbr::OriginalName("pb_battle_S2C_MoveComplete")] PbBattleS2CMoveComplete = 3101,
    [pbr::OriginalName("pb_battle_S2C_Attack")] PbBattleS2CAttack = 3102,
    [pbr::OriginalName("pb_battle_S2C_ToggleTankGuardMode")] PbBattleS2CToggleTankGuardMode = 3103,
    [pbr::OriginalName("pb_battle_S2C_MergeTroops")] PbBattleS2CMergeTroops = 3104,
    [pbr::OriginalName("pb_battle_S2C_MoveTroopsPlanningPath")] PbBattleS2CMoveTroopsPlanningPath = 3105,
    [pbr::OriginalName("pb_battle_S2C_CrossLargeTile")] PbBattleS2CCrossLargeTile = 3106,
    [pbr::OriginalName("pb_battle_S2C_LargeTileChange")] PbBattleS2CLargeTileChange = 3107,
    [pbr::OriginalName("pb_battle_S2C_ProduceBuilding")] PbBattleS2CProduceBuilding = 3108,
    [pbr::OriginalName("pb_battle_S2C_PlaneLoad")] PbBattleS2CPlaneLoad = 3109,
    [pbr::OriginalName("pb_battle_S2C_PlaneUnload")] PbBattleS2CPlaneUnload = 3110,
    /// <summary>
    /// Test
    /// </summary>
    [pbr::OriginalName("pb_test_C2S_Test")] PbTestC2STest = 9000,
  }

  #endregion

}

#endregion Designer generated code
