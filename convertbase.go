package piscine

import "math/big"

func ConvertBase(nbr, baseFrom, baseTo string) string {
	num := new(big.Int)
	num.SetString(nbr, len(baseFrom))

	result := ""
	bigBaseTo := big.NewInt(int64(len(baseTo)))
	zero := big.NewInt(0)

	for num.Cmp(zero) > 0 {
		rem := new(big.Int)
		num.QuoRem(num, bigBaseTo, rem)
		result = string(baseTo[rem.Int64()]) + result
	}

	return result
}
