package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Revenue: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Expenses: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Tax Rate: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%f.1\n", ebt)
	fmt.Printf("%f.1\n", profit)
	fmt.Printf("%f.3\n", ratio)
	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %3.f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value must be a positive number")
	}

	return userInput, nil
}
