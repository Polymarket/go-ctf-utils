package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/polymarket/go-ctf-utils/utils"
)

func main() {
	// Ensure there are exactly two arguments (excluding the program name itself).
	if len(os.Args) != 4 {
		fmt.Println("Usage: <program> <conditionId> <collateral> <outcomeIndex>")
		os.Exit(1)
	}

	conditionIdStr, collateralStr, outcomeIndexStr := os.Args[1], os.Args[2], os.Args[3]

	conditionId := common.HexToHash(conditionIdStr)
	collateral := common.HexToAddress(collateralStr)
	// outcomeIndex, ok := new(big.Int).SetString(outcomeIndexStr, 10)
	// if !ok {
	// 	fmt.Printf("Invalid outcome index: %s\n", outcomeIndexStr)
	// 	os.Exit(1)
	// }

	// Call the utility function
	result := utils.CalculatePositionIds(collateral, conditionId)

	// Print the result
	fmt.Printf("Result of %s AND %s = %s\n", outcomeIndexStr, result)
}
