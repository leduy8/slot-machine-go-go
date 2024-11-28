package main

import "fmt"

func getName() string {
	name := ""

	fmt.Println("Welcome to Go Go Slot")
	fmt.Printf("Enter your name: ")
	_, err := fmt.Scanln(&name)

	if err != nil {
		return ""
	}

	fmt.Printf("Welcome %s, let's play!\n", name)
	return name
}

func getBet(balance uint) uint {
	var bet uint
	for {
		fmt.Printf("Enter your bet, or 0 to quit (balance = $%d): ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("Bet cannot be greater than your balance")
		} else {
			break
		}
	}

	return bet
}

func generateSymbolArray(symbols map[string]uint) []string {
	var symbolArr []string
	for symbol, count := range symbols {
		for i := 0; i < int(count); i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}

	return symbolArr
}

func main() {
	balance := uint(200)
	// Symbol map
	// symbol: odds
	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 12,
		"D": 20,
	}
	symbolArr := generateSymbolArray(symbols)
	fmt.Println(symbolArr)
	// Multiplier map
	// multiplier: x times winning for a line (Ex: Line of A = 20 times of your bet)
	multiplier := map[string]uint{
		"A": 20,
		"B": 10,
		"C": 5,
		"D": 2,
	}

	getName()

	for balance > 0 {
		bet := getBet(balance)
		if bet == 0 {
			break
		}

		balance -= bet
	}

	fmt.Printf("You left with %d\n", balance)
}
