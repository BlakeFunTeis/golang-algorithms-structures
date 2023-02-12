package main

import (
	"fmt"
	"golang-algorithms-structures/BinarySearch"
	"golang-algorithms-structures/SelectionSort"
)

func main() {
	binarySearchData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	binarySearchResult, err := BinarySearch.Run(binarySearchData, 99)
	fmt.Println(binarySearchResult, err)
	selectionSortData := []int{555, 447, 223, 11, 5, 6, 7, 8, 9, 10}
	selectionSortResult := SelectionSort.Run(selectionSortData)
	fmt.Println(selectionSortResult)
}
