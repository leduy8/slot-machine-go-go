package main

import (
	"fmt"
	"math/rand"
)

func GenerateSymbolArray(symbols map[string]uint) []string {
	var symbolArr []string
	for symbol, count := range symbols {
		for i := 0; i < int(count); i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}

	return symbolArr
}

func getRandomNumber(min int, max int) int {
	return rand.Intn(max-min) + min + 1
}

func GetSpin(reel []string, rows int, cols int) [][]string {
	var result [][]string

	for i := 0; i < rows; i++ {
		result = append(result, []string{})
	}

	for col := 0; col < cols; col++ {
		selected := map[int]bool{}
		for row := 0; row < rows; row++ {
			for {
				randomIndex := getRandomNumber(0, len(reel)-1)
				_, exists := selected[randomIndex]
				if !exists {
					selected[randomIndex] = true
					result[row] = append(result[row], reel[randomIndex])
					break
				}
			}
		}
	}

	return result
}

func PrintSpin(spin [][]string) {
	for _, row := range spin {
		for j, symbol := range row {
			fmt.Printf(symbol)
			if j != len(row)-1 {
				fmt.Printf(" | ")
			}
		}
		fmt.Println()
	}
}
