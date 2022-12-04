package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("failed to open the file input")
	}

	scanner := bufio.NewScanner(input)

	maxCalories := 0
	calories := 0

	for scanner.Scan() {
		calorieAsString := scanner.Text()

		if calorieAsString == "" {
			if calories > maxCalories {
				maxCalories = calories
			}

			calories = 0
		}

		calorie, _ := strconv.Atoi(calorieAsString)

		calories += calorie
	}

	log.Println(maxCalories)
}
