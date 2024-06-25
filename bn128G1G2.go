package main

import (
	"bn128/bn128"
	"fmt"
	"math/big"
	//bn128 "github.com/fentec-project/bn256"
)

func main() {

	fmt.Println(new(bn128.G1).ScalarBaseMult(big.NewInt(1)))
	fmt.Println(new(bn128.G2).ScalarBaseMult(big.NewInt(1)))

}
