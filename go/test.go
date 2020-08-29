package main

import (
	"github.com/CMiksche/benford"
)

func main() {
	chiSquared := benford.CalcBenfords([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).ChiSquared
}
