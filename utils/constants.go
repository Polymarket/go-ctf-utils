package utils

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	ORACLE             = common.HexToAddress("0x6A9D222616C90FcA5754cd1333cFD9b7fb6a4F74")
	OUTCOME_SLOT_COUNT = big.NewInt(2)
)
