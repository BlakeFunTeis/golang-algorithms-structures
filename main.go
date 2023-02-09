package main

import (
	"fmt"
	"golang-algorithms-structures/BinarySearch"
)

func main() {
	BinarySearchData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result, err := BinarySearch.Run(BinarySearchData, 99)
	fmt.Println(result, err)
}
