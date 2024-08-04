package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 21000000

	const d = 3e20 / n
	fmt.Println(d)

	//fmt.Println(int64(d))
	fmt.Println(float64(d))

	fmt.Println(math.Sin(n))
}
