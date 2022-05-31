package helper

import (
	cryptorand "crypto/rand"
	"fmt"
	"math/big"
)

// GenerateOrderNo helper
func GenerateOrderNo() string {
	bignum, _ := cryptorand.Int(cryptorand.Reader, big.NewInt(9999999999))
	return fmt.Sprintf("%d", bignum.Int64())
}
