package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("failed to open the file input")
	}
	defer input.Close()

	duplicated := []rune{}
	sum := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		rucksack := scanner.Text()

		firstCompartment := rucksack[:(len(rucksack) / 2)]
		secondCompartment := rucksack[(len(rucksack) / 2):]

		firstCompartmentMap := make(map[rune]int, len(firstCompartment))
		sharedItemType := ' '

		for _, r := range firstCompartment {
			firstCompartmentMap[r] = 1
		}

		for _, r := range secondCompartment {
			if _, ok := firstCompartmentMap[r]; !ok {
				continue
			}

			sharedItemType = r
		}

		if !unicode.IsSpace(sharedItemType) {
			duplicated = append(duplicated, sharedItemType)
		}
	}

	for _, r := range duplicated {
		subtract := 96
		if unicode.IsUpper(r) {
			subtract = 38
		}

		sum += (int(r) - subtract)
	}

	fmt.Println(sum)
}
