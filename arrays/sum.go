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
	var sumArray []int
	for _, numbers := range numbersToSum {
		sumArray = append(sumArray, Sum(numbers))
	}
	return sumArray
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sumAllTails []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sumAllTails = append(sumAllTails, 0)
		} else {
			sumAllTails = append(sumAllTails, Sum(numbers[1:]))
		}
	}
	return sumAllTails
}
