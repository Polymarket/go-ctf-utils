package utils

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ONE  = big.NewInt(1)
	ZERO = big.NewInt(0)
	P, _ = new(big.Int).SetString("21888242871839275222246405745257275088696311157297823662689037894645226208583", 10)
	B    = big.NewInt(3)
)

func eq(a, b *big.Int) bool {
	return a.Cmp(b) == 0
}

func addModP(a, b *big.Int) *big.Int {
	z := new(big.Int)
	z.Add(a, b)
	z.Mod(z, P)
	return z
}

func mulModP(a, b *big.Int) *big.Int {
	z := new(big.Int)
	z.Mul(a, b)
	z.Mod(z, P)
	return z
}

func bitAnd(a, b *big.Int) *big.Int {
	z := new(big.Int)
	z.And(a, b)
	return z
}

func powModP(a, b *big.Int) *big.Int {
	res := big.NewInt(1)

	at, _ := new(big.Int).SetString(a.String(), 10)
	bt, _ := new(big.Int).SetString(b.String(), 10)

	for {
		if !eq(bitAnd(bt, ONE), ZERO) {
			res = mulModP(res, at)
		}
		at = mulModP(at, at)
		bt.Rsh(bt, 1)
		if eq(bt, ZERO) {
			break
		}
	}
	return res
}

// a^((P-1)/2) mod P
// powModP(a, P.minus(BigInt.fromI32(1)).rightShift(1));
func legendreSymbol(a *big.Int) *big.Int {
	z := new(big.Int)
	z.Sub(P, ONE)
	z.Rsh(z, 1)
	res := powModP(a, z)
	return res
}

func computeCollectionIdHash(conditionId common.Hash, outcomeIndex *big.Int) common.Hash {
	var hashPayload []byte
	// first 32 bytes is conditionId
	hashPayload = append(hashPayload, conditionId.Bytes()...)

	// second 32 bytes is index set
	indexSet := big.NewInt(0)
	indexSet = indexSet.Lsh(ONE, uint(outcomeIndex.Int64()))
	hashPayload = append(hashPayload, common.LeftPadBytes(indexSet.Bytes(), 32)...)

	h := crypto.Keccak256Hash(hashPayload)
	hashRes := h.Bytes()

	hashResBigInt := new(big.Int).SetBytes(hashRes[:])

	// check if the msb is set
	ii := big.NewInt(0)
	ii = ii.Rsh(hashResBigInt, 255)
	odd := ii.Cmp(ZERO) == 1 // (ii.Cmp(ZERO) == 0) == false

	x1 := hashResBigInt
	var yy *big.Int

	// At this point, all the inputs are correct

	// increment x1 until we find a point on the curve
	// i.e. if there exists y1 so y1^2 = x1^3 + 3 (mod P)
	// if the legendreSymbol is not 1, then the number is not a quadratic residue
	// i.e., its not a square mod P, and its not on the curve
	for {
		x1 = addModP(x1, ONE)
		yy = addModP(mulModP(x1, mulModP(x1, x1)), B)
		if eq(legendreSymbol(yy), ONE) {
			break
		}
	}

	ii = new(big.Int)
	oddToggle := ii.Lsh(ONE, 254)
	if odd {
		if bitAnd(x1, oddToggle).Cmp(ZERO) == 0 {
			x1 = x1.Add(x1, oddToggle)
		} else {
			x1 = x1.Sub(x1, oddToggle)
		}
	}

	x1Hash := common.BigToHash(x1)
	return x1Hash
}

func computePositionIdHash(collateral common.Address, collectionId common.Hash) common.Hash {
	var hashPayload []byte
	hashPayload = append(hashPayload, collateral.Bytes()...)
	hashPayload = append(hashPayload, collectionId.Bytes()...)
	return crypto.Keccak256Hash(hashPayload)
}

func CalculatePositionId(collateral common.Address, conditionId common.Hash, outcomeIndex *big.Int) *big.Int {
	collectionId := computeCollectionIdHash(conditionId, outcomeIndex)
	positionIdHash := computePositionIdHash(collateral, collectionId)

	positionId := new(big.Int).SetBytes(positionIdHash.Bytes())
	return positionId
}

func CalculatePositionIds(collateral common.Address, conditionId common.Hash) []string {
	token0 := CalculatePositionId(collateral, conditionId, big.NewInt(0))
	token1 := CalculatePositionId(collateral, conditionId, big.NewInt(1))

	return []string{token0.String(), token1.String()}
}
