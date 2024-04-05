// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std/Test.sol";
import {CTHelpers} from "./CTHelpers.sol";

contract PositionsTest is Test {
    function test_Increment(address _collateral, bytes32 _conditionId, uint256 _outcomeIndex) public {
        string[] memory inputs = new string[](4);

        uint256 outcomeIndex = _outcomeIndex % 2;
        inputs[0] = "./app";
        inputs[1] = vm.toString(_collateral);
        inputs[2] = vm.toString(_conditionId);
        inputs[3] = vm.toString(outcomeIndex);

        bytes memory res = vm.ffi(inputs);

        console.logBytes(res);
        uint256 output = abi.decode(res, (uint256));

        bytes32 collectionId = CTHelpers.getCollectionId(bytes32(0), _conditionId, 1 << outcomeIndex);
        uint256 positionId = CTHelpers.getPositionId(_collateral, collectionId);

        assertEq(output, positionId);
    }
}
