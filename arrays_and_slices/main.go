package arraysandslices

func Sum(number [5]int) int {
	var sum int
	for _, number := range number {
		sum += number
	}
	return sum
}
