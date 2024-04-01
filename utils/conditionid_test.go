package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConditionID(t *testing.T) {
	var title string
	var desc string
	var outcomes []string
	var creator string

	var actual string
	var expected string

	title = `ETH greater than 10000?`
	desc = `Will the price of ETH on the ETH/USDC Uniswap V3 5bps pool be greater than 10000 USDC by December 31st 2024?`
	outcomes = []string{"Yes", "No"}
	creator = `0x6d8c4e9adf5748af82dabe2c6225207770d6b4fa`

	ancillaryData, err := CreateAncillaryData(title, desc, creator, outcomes)
	assert.Nil(t, err)

	questionIDHash := CalculateQuestionIDHash(ancillaryData)

	actual = GetConditionIDUsingDefaults(questionIDHash).Hex()
	expected = `0x491b47c68ed1de5b01c359fd5d14a285b68af60b14ec7939acfee2afbfbb8ec8`
	assert.Equal(t, expected, actual)
}
