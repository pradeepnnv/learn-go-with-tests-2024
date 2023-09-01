package arrays

func Sum(numbers []int) (sum int) {
	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }
	for _, v := range numbers {
		sum += v
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	sumArray := make([]int, len(numbersToSum))
	for i, n := range numbersToSum {
		sumArray[i] = Sum(n)
	}
	return sumArray
}
