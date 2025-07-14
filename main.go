package main

import (
	"fmt"

	"github.com/medunes/go-dsa/sort"
)

func main() {
	array := []int{-1, -10000, -12345, -2032, -23, 0, 0, 0, 0, 10, 10000, 1024, 1024354, 155, 174, 1955, 2, 255, 3, 322, 4741, 96524}
	sort.Selection(array)
	fmt.Printf("%v\n", sort.Bubble(array))
}
