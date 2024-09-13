//Example of using functions as first class citizents (passing functions to other functions and as return values)

// package main

// import "fmt"

// type transformFn func(int) int

// func main() {
// 	numbers := []int{1, 2, 3, 4}
// 	moreNumbers := []int{5, 1, 2}
// 	doubled := transformNumbers(&numbers, double)
// 	tripled := transformNumbers(&numbers, triple)
// 	fmt.Println(doubled)
// 	fmt.Println(tripled)

// 	transformFn1 := getTranformerFunction(&numbers)
// 	transformeFn2 := getTranformerFunction(&moreNumbers)

// 	transfomredNumbers := transformNumbers(&numbers, transformFn1)
// 	moreTransfomredNumbers := transformNumbers(&numbers, transformeFn2)

// 	fmt.Println(transfomredNumbers)
// 	fmt.Println(moreTransfomredNumbers)
// }

// func transformNumbers(numbers *[]int, transform transformFn) []int {
// 	dNumbers := []int{}
// 	for _, val := range *numbers {
// 		dNumbers = append(dNumbers, transform(val))
// 	}
// 	return dNumbers
// }

// func getTranformerFunction(numbers *[]int) transformFn {
// 	if(*numbers)[0] == 1 {
// 		return double
// 	} else {
// 		return triple
// 	}
// }

// func double(number int) int {
// 	return number * 2
// }

// func triple(number int) int {
// 	return number * 3
// }

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTransfomrer(2)
	triple := createTransfomrer(3)

	transformed := transformNumbers(&numbers, func (number int) int {
		return number * 2
	})

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func createTransfomrer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}
