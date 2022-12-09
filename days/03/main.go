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
		log.Fatalln("failed to open the input file")
	}
	defer input.Close()

	agroupCh := make(chan string)
	findBadgeCh := make(chan []string)
	badges := make(chan rune)

	go func() {
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			rucksack := scanner.Text()
			agroupCh <- rucksack
		}

		close(agroupCh)
	}()

	go func() {
		group := []string{}
		for rucksack := range agroupCh {
			group = append(group, rucksack)

			if len(group) == 3 {
				findBadgeCh <- group
				group = []string{}
			}
		}

		close(findBadgeCh)
	}()

	go func() {
		for rucksacks := range findBadgeCh {
			firstRucksackMap := map[rune]int{}
			for _, r := range rucksacks[0] {
				firstRucksackMap[r] = 1
			}

			secondRucksackMap := map[rune]int{}
			for _, r := range rucksacks[1] {
				secondRucksackMap[r] = 1
			}

			sharedItemType := ' '
			for _, r := range rucksacks[2] {
				if _, ok := firstRucksackMap[r]; !ok {
					continue
				}

				if _, ok := secondRucksackMap[r]; !ok {
					continue
				}

				sharedItemType = r
			}

			if !unicode.IsSpace(sharedItemType) {
				badges <- sharedItemType
			}
		}

		close(badges)
	}()

	sum := 0
	for badge := range badges {
		subtract := 96
		if unicode.IsUpper(badge) {
			subtract = 38
		}

		sum += (int(badge) - subtract)
	}

	fmt.Println(sum)
}
