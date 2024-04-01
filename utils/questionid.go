package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func CalculateQuestionIDHash(ancillaryData []byte) common.Hash {
	return crypto.Keccak256Hash(ancillaryData)
}
