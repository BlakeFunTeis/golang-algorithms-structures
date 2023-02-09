package BinarySearch

import "errors"

func Run(data []int, search int) (int, error) {
	length := len(data)
	low := 0
	high := length - 1
	for low <= high {
		mid := (low + high) / 2
		guess := data[mid]
		if guess == search {
			return mid, nil
		}
		if guess > search {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0, errors.New("not found")
}
