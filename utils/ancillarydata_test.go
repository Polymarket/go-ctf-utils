package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAncillaryData(t *testing.T) {
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

	actualBytes, err := CreateAncillaryData(title, desc, creator, outcomes)
	assert.Nil(t, err)

	actual = string(actualBytes)
	expected = `q: title: ETH greater than 10000?, description: Will the price of ETH on the ETH/USDC Uniswap V3 5bps pool be greater than 10000 USDC by December 31st 2024? res_data: p1: 0, p2: 1, p3: 0.5. Where p1 corresponds to No, p2 to Yes, p3 to unknown/50-50. Updates made by the question creator via the bulletin board at 0x6A9D222616C90FcA5754cd1333cFD9b7fb6a4F74 as described by https://polygonscan.com/tx/0xa14f01b115c4913624fc3f508f960f4dea252758e73c28f5f07f8e19d7bca066 should be considered.,initializer:6d8c4e9adf5748af82dabe2c6225207770d6b4fa`
	assert.Equal(t, expected, actual)

	// Fail cases
	badCreator := `` //invalid address
	_, err = CreateAncillaryData(title, desc, badCreator, outcomes)
	assert.NotNil(t, err)

	badOutcomes := []string{} //empty outcomes
	_, err = CreateAncillaryData(title, desc, creator, badOutcomes)
	assert.NotNil(t, err)
}
