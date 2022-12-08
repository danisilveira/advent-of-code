package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	min int
	max int
}

func (p Pair) String() string {
	return fmt.Sprintf("%d-%d", p.min, p.max)
}

func NewPairFromString(pair string) Pair {
	splitted := strings.Split(pair, "-")
	min, _ := strconv.Atoi(splitted[0])
	max, _ := strconv.Atoi(splitted[1])

	return Pair{
		min: min,
		max: max,
	}
}

func SwapPairsIfNeeded(firstPair, secondPair Pair) (Pair, Pair) {
	if firstPair.min == secondPair.min {
		if firstPair.max > secondPair.max {
			return firstPair, secondPair
		}

		return secondPair, firstPair
	}

	if firstPair.min < secondPair.min {
		return firstPair, secondPair
	}

	return secondPair, firstPair
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("failed to open the file input")
	}
	defer input.Close()

	count := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), ",")

		firstPair := NewPairFromString(pairs[0])
		secondPair := NewPairFromString(pairs[1])
		fmt.Printf("firstPair : %v\n", firstPair)
		fmt.Printf("secondPair: %v\n", secondPair)

		firstPair, secondPair = SwapPairsIfNeeded(firstPair, secondPair)
		fmt.Printf("firstPair : %v\n", firstPair)
		fmt.Printf("secondPair: %v\n", secondPair)

		if firstPair.min <= secondPair.min && firstPair.max >= secondPair.max {
			fmt.Printf("count + 1\n")
			count++
		}

		fmt.Printf("\n")
	}

	fmt.Println(count)
}
