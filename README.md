# go-ctf-utils

Calculate CTF ERC1155 position ids

## Usage

```go
collateral := common.HexToAddress("0x2791bca1f2de4661ed88a30c99a7a9449aa84174")
conditionId := common.HexToHash("0x41771a29f1fa3b5ac743ddcf224017f802bd69152c1a65230ec666abfc22b708")
outcomeIndex := big.NewInt(0)
positionid := ctfutils.CalculatePositionId(collateral, conditionId, outcomeIndex)
fmt.Printf("PositionID: %s\n", positionid.String())
```
