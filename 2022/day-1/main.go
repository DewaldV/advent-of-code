package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileContent, err := os.ReadFile("input")
	if err != nil {
		fmt.Printf("error reading input file, %s\n", err)
		return
	}

	calories, err := mostCalories(string(fileContent))
	if err != nil {
		fmt.Printf("error converting calorie list, %s\n", err)
		return
	}

	fmt.Printf("Most Calories carried: %d\n", calories)
}

func mostCalories(input string) (int, error) {
	cleanInput := strings.TrimSpace(input)
	elves := strings.Split(cleanInput, "\n\n")

	mostCalories := 0

	for _, elf := range elves {
		items := strings.Split(elf, "\n")
		totalCalories := 0
		for _, calories := range items {
			i, err := strconv.Atoi(calories)
			if err != nil {
				return 0, fmt.Errorf("unable to convert elf calories, %w", err)
			}

			totalCalories += i
		}

		if totalCalories > mostCalories {
			mostCalories = totalCalories
		}
	}

	return mostCalories, nil
}
