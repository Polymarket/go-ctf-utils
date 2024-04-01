package utils

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

const (
	ANC_DATA_FMT = `q: title: %s, description: %s res_data: %s,initializer:%s`
	RES_DATA_FMT = `p1: 0, p2: 1, p3: 0.5. Where p1 corresponds to %s, p2 to %s, p3 to unknown/50-50. Updates made by the question creator via the bulletin board at 0x6A9D222616C90FcA5754cd1333cFD9b7fb6a4F74 as described by https://polygonscan.com/tx/0xa14f01b115c4913624fc3f508f960f4dea252758e73c28f5f07f8e19d7bca066 should be considered.`
)

func CreateAncillaryData(title, desc, creator string, outcomes []string) ([]byte, error) {
	b := make([]byte, 0)
	init, err := createInitializerHex(creator)
	if err != nil {
		return b, err
	}

	res, err := createResDataString(outcomes)
	if err != nil {
		return b, err
	}

	ancDataStr := fmt.Sprintf(ANC_DATA_FMT, title, desc, res, init)
	return []byte(ancDataStr), nil
}

func createInitializerHex(creator string) (string, error) {
	if len(creator) == 0 && !common.IsHexAddress(creator) {
		return "", fmt.Errorf("invalid creator address")
	}
	// Lowercase and remove the 0x
	return strings.ToLower(creator[2:]), nil
}

func createResDataString(outcomes []string) (string, error) {
	if len(outcomes) != 2 {
		return "", fmt.Errorf("invalid number of outcomes")
	}
	return fmt.Sprintf(RES_DATA_FMT, outcomes[1], outcomes[0]), nil
}
