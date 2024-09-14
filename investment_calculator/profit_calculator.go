package main

import (
	"fmt"
	"os"
)

func main() {
	const taxRate float64 = 20
	var revenue, expenses float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	if revenue < 1 || expenses < 1 {
		fmt.Println("Revenue and Expenses must not be less than 1")
		return
	}

	earnings, profit, ratio := getData(revenue, expenses, taxRate)

	fmt.Println(earnings)
	fmt.Println(profit)
	fmt.Println(ratio)
	calculatedResults := fmt.Sprint(earnings, profit, ratio)
	os.WriteFile("calculated-results.txt", []byte(calculatedResults), 0644)
}

func getData(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earnings := revenue - expenses
	profit := earnings * (1 - taxRate / 100)
	ratio := earnings / profit
	return earnings, profit, ratio
}
