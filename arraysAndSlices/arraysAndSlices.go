package arraysandslices

func Sum(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}

	return sum
}

func SumAll(slices ...[]int) []int {
	var sums []int

	for _, val := range slices {
		sums = append(sums, Sum(val))
	}

	return sums
}

func SumAllTails(slices ...[]int) []int {
	var tailSums []int
	for _, val := range slices {
		if len(val) == 0 {
			tailSums = append(tailSums, 0)
		} else {

			tail := val[1:]
			tailSums = append(tailSums, Sum(tail))
		}
	}

	return tailSums
}
