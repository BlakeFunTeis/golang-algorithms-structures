package QuickSort

func Run(data []int) (result []int) {
	if len(data) < 2 {
		return data
	}
	pivot := data[0]
	var less []int
	var greater []int
	for _, v := range data[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	result = append(result, Run(less)...)
	result = append(result, pivot)
	result = append(result, Run(greater)...)
	return result
}
