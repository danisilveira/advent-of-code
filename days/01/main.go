package main

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/danisilveira/algorithms-and-data-structures/algorithms/searching/quickselect"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("failed to open the file input")
	}

	calories := []int{}
	caloriesCount := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		calorieAsString := scanner.Text()
		if calorieAsString == "" {
			calories = append(calories, caloriesCount)
			caloriesCount = 0
			continue
		}

		calorie, _ := strconv.Atoi(calorieAsString)
		caloriesCount += calorie
	}

	firstElf := quickselect.Select(calories, len(calories))
	secondElf := quickselect.Select(calories, len(calories)-1)
	thirdElf := quickselect.Select(calories, len(calories)-2)

	log.Println(firstElf + secondElf + thirdElf)
}
