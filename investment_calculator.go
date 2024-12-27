package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Investment Expected Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Investment Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formatedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formatedFRV := fmt.Sprintf("Future value( adjust for Inflation): %2f ", futureRealValue)

	// fmt.Println(futureValue)
	// fmt.Println(futureRealValue)

	// fmt.Printf("future value: %2f\nFuture value( adjust for Inflation): %2f ", futureValue, futureRealValue)

	fmt.Print(formatedFV, formatedFRV)

}
