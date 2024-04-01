package utils

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestComputeCollectionID(t *testing.T) {

	// questionID: 0xa23eaa48f1d77a8ed0660f8198198b08ad33cf059ef05d5d87f254325c62b18a
	// conditionID: 0x41771a29f1fa3b5ac743ddcf224017f802bd69152c1a65230ec666abfc22b708
	// clob_token_ids: ["87848146419241057657677458104196204655830537958664996373577118668208015365957", "11831001752042525186810643219442170205387399550590230096735412415464402686073"]

	conditionId := common.HexToHash("0x41771a29f1fa3b5ac743ddcf224017f802bd69152c1a65230ec666abfc22b708")
	outcomeIndex0 := big.NewInt(0)
	outcomeIndex1 := big.NewInt(1)
	expectedCollectionId0 := "0x63ecd1f555d88721e1d063640ec7719a904925b55bf9a8c3e9d1dddfa86b1a6b"
	expectedCollectionId1 := "0x610f95f837bfcb4a52eedebd551bbdd7f273edeb47ad140cfba8287e6fed1929"

	collectionIdHash0 := computeCollectionIdHash(conditionId, outcomeIndex0)
	collectionIdHex0 := strings.ToLower(collectionIdHash0.Hex())
	assert.Equal(t, expectedCollectionId0, collectionIdHex0)

	collectionIdHash1 := computeCollectionIdHash(conditionId, outcomeIndex1)
	collectionIdHex1 := strings.ToLower(collectionIdHash1.Hex())
	assert.Equal(t, expectedCollectionId1, collectionIdHex1)
}

func TestCalculatePositionId(t *testing.T) {
	conditionId := common.HexToHash("0x41771a29f1fa3b5ac743ddcf224017f802bd69152c1a65230ec666abfc22b708")
	collateral := common.HexToAddress("0x2791bca1f2de4661ed88a30c99a7a9449aa84174")
	outcomeIndex0 := big.NewInt(0)
	outcomeIndex1 := big.NewInt(1)

	token0 := "87848146419241057657677458104196204655830537958664996373577118668208015365957"
	token1 := "11831001752042525186810643219442170205387399550590230096735412415464402686073"

	posId0 := CalculatePositionId(collateral, conditionId, outcomeIndex0)
	assert.Equal(t, token0, posId0.String())

	posId1 := CalculatePositionId(collateral, conditionId, outcomeIndex1)
	assert.Equal(t, token1, posId1.String())
}
