package SelectionSort

func Run(data []int) []int {
	length := len(data)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		data[i], data[min] = data[min], data[i]
	}
	return data
}
