package utils

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetConditionIDUsingDefaults(questionID common.Hash) common.Hash {
	return getConditionID(ORACLE, questionID, OUTCOME_SLOT_COUNT)
}

func GetConditionID(oracle common.Address, questionID common.Hash, outcomeSlotCount *big.Int) common.Hash {
	return getConditionID(oracle, questionID, outcomeSlotCount)
}

// Gets the conditionID given the oracle, questionID and outcomeSlotCount
//
//	function getConditionId(address oracle, bytes32 questionId, uint outcomeSlotCount) internal pure returns (bytes32) {
//		return keccak256(abi.encodePacked(oracle, questionId, outcomeSlotCount));
//	}
func getConditionID(oracle common.Address, questionID common.Hash, outcomeSlotCount *big.Int) common.Hash {
	var encoded []byte
	encoded = append(encoded, oracle.Bytes()...)
	encoded = append(encoded, questionID.Bytes()...)
	encoded = append(encoded, common.LeftPadBytes(outcomeSlotCount.Bytes(), 32)...)
	return crypto.Keccak256Hash(encoded)
}
