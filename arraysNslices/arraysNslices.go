package arraysnslices

// array has a fixed size
// slice has a dynamic size
func Sum(arr []int) int {

	sum := 0

	for _, num := range arr {
		sum += num
	}

	return sum
}

func SumAll(arrays ...[]int) []int {

	sums := []int{}

	for _, numbers := range arrays {
		sums = append(sums, Sum(numbers))
	}

	return sums

}

func SumAllTails(arrays ...[]int) []int {

	sums := []int{}

	for _, numbers := range arrays {
		if len(numbers) > 0 {
			sums = append(sums, Sum(numbers)-numbers[0])
		} else {
			sums = append(sums, 0)
		}
	}

	return sums

}
