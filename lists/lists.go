package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[1:1])
	
	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices, prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	fmt.Println(prices)
// 	fmt.Println(productNames)
// 	fmt.Println(prices[0])

// 	featuredPrices := prices[1:]
// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// }
