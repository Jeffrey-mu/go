package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4, 5}
	fmt.Print(arr[2:4])
}
func Array(arr []int) int {
	var i int
	for _, v := range arr {
		i += v
	}
	return i
}
func ArrayAll(numbersToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Array(numbers)
	}
	return
}
