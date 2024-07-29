package arraysandslices

func Sum(number []int) int {
	var sum int
	for _, number := range number {
		sum += number
	}
	return sum
}
