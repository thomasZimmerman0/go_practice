// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	const inflationRate = 2.5
// 	var investmentAmount, years, expectedReturnRate float64

// 	outputText("Investment Amount: ")
// 	fmt.Scan(&investmentAmount)

// 	fmt.Print("Years: ")
// 	fmt.Scan(&years)

// 	outputText("Return Rate: ")
// 	fmt.Scan(&expectedReturnRate)

// 	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, inflationRate, years)

// 	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
// 	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f", futureRealValue)

// 	fmt.Print(formattedFV, formattedRFV);
// }

// func outputText(text string){
// 	fmt.Print(text)
// }

// func calculateFutureValue(investmentAmount, expectedReturnRate, inflationRate, years float64) (float64, float64) {
// 	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	rfv := fv/ math.Pow(1+inflationRate/100, years)
// 	return fv, rfv
// }