// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0;

// to test precompiles, paste contracts in https://remix.ethereum.org/ to easily visualize

contract ECS {
    uint256 public number;
    string public str;
    Pos public pos;
    address public addr;

    struct QueryCondition {
        string component;
        uint256 action;
        bytes value;
    }

    struct Pos {
        uint256 X;
        uint256 Y;
    }

    struct GetComponentValueRequest {
        string component;
        uint256 entity;
    }

    address ECS_QUERY_PRECOMPILE = address(5);
    address GET_COMPONENT_VALUE_PRECOMPILE = address(6);

    // query types
    uint256 Has = 0;
    uint256 Not = 1;
    uint256 HasExact = 2;

    function query() public returns (uint256) {
        QueryCondition[] memory query = new QueryCondition[](1);
        query[0] = QueryCondition({
            component: "Tag",
            action: Has,
            value: abi.encode("Hello")
        });

        (bool success, bytes memory result) = ECS_QUERY_PRECOMPILE.call(
            abi.encode(query)
        );

        uint256[] memory ecsResults = abi.decode(result, (uint256[]));

        // result is an array of uint256
        number = ecsResults.length;
        return ecsResults.length;
    }

    // string example
    function getComponentValueString() public {
        GetComponentValueRequest memory request = GetComponentValueRequest({
            component: "Tag",
            entity: 1
        });

        (bool success, bytes memory result) = GET_COMPONENT_VALUE_PRECOMPILE
            .call(abi.encode(request));

        str = abi.decode(result, (string));
    }

    // address example
    function getCompnentValuePosition() public {
        GetComponentValueRequest memory request = GetComponentValueRequest({
            component: "Position",
            entity: 1
        });

        (bool success, bytes memory result) = GET_COMPONENT_VALUE_PRECOMPILE
            .call(abi.encode(request));

        Pos memory position = abi.decode(result, (Pos));
        pos = position;
    }

    // uint256 example
    function getComponentValueUint256() public {
        GetComponentValueRequest memory request = GetComponentValueRequest({
            component: "Gold",
            entity: 1
        });

        (bool success, bytes memory result) = GET_COMPONENT_VALUE_PRECOMPILE
            .call(abi.encode(request));

        uint256 res = abi.decode(result, (uint256));
        number = res;
    }

    // address example
    function getComponentValueAddress() public {
        GetComponentValueRequest memory request = GetComponentValueRequest({
            component: "Treaty",
            entity: 1
        });

        (bool success, bytes memory result) = GET_COMPONENT_VALUE_PRECOMPILE
            .call(abi.encode(request));

        address res = abi.decode(result, (address));
        addr = res;
    }
}
