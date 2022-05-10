package fun

import (
	"fmt"
	"testing"
)

// string转*big.Int
func TestStrToBigInt(t *testing.T) {
	bigInt, ok := StrToBigInt("0.123", 3)
	fmt.Println(bigInt.String(), ok)
}