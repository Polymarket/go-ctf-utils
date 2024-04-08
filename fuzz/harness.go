package main

import (
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/polymarket/go-ctf-utils/utils"
)

func main() {
	// Ensure there are exactly two arguments (excluding the program name itself).
	if len(os.Args) != 4 {
		fmt.Println("Usage: <program> <collateral> <conditionId> <outcomeIndex>")
		os.Exit(1)
	}

	collateralStr, conditionIdStr, outcomeIndexStr := os.Args[1], os.Args[2], os.Args[3]

	collateral := common.HexToAddress(collateralStr)
	conditionId := common.HexToHash(conditionIdStr)
	outcomeIndex, ok := new(big.Int).SetString(outcomeIndexStr, 10)
	if !ok {
		fmt.Printf("Invalid outcome index: %s\n", outcomeIndexStr)
		os.Exit(1)
	}

	// Call the utility function
	result := utils.CalculatePositionId(collateral, conditionId, outcomeIndex)

	hexString := result.Text(16)
	// Pad the hex string to 64 characters (32 bytes) with leading zeros
	paddedHexString := fmt.Sprintf("%064s", hexString)
	// Print the padded hex string
	fmt.Println(paddedHexString)
}
