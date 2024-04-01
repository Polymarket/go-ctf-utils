package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateQuestionIDHash(t *testing.T) {
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
	actual = questionIDHash.Hex()
	expected = `0x01741d802f72305df80da4d6e8ecd3a50287f09ec62edb3bd95ac7c395b2f5ef`
	assert.Equal(t, expected, actual)
}
