package main

import "fmt"

func main() {
	revenue := print("Revenue: ")
	expenses := print("Expenses:")
	taxRate := print("Tax Rate: ")

	ebt, profit, ratio := calculate(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func print(text string) float64 {
	var value float64
	fmt.Print(text)
	fmt.Scan(&value)
	return value
}

func calculate(revenue float64, expenses float64, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}
