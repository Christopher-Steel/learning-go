package arrays

func Sum(numbers []int) (sum int) {
	for _, i := range numbers {
		sum += i
	}
	return
}

func SumAll(numbers ...[]int) (sums []int) {
	for _, nums := range numbers {
		sums = append(sums, Sum(nums))
	}
	return
}