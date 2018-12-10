package main

import (
	"fmt"
	"strconv"
)

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

func main() {
	var i float64
	var ii, jj float64 = 10, 3
	i = ii / jj
	fmt.Println(i)
	fmt.Println(Decimal(i))
}
