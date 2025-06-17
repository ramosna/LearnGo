package main

func Sum(number []int) int {
	sum := 0
	for _, num := range number {
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {

	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func Tail(number []int) []int {
	return number[1:]
}

func SumAllTails(numbersToSum ...[]int) []int {

	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(Tail(numbers)))
		}
	}

	return sums
}
