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

func NewPairFromString(pair string) Pair {
	splitted := strings.Split(pair, "-")
	min, _ := strconv.Atoi(splitted[0])
	max, _ := strconv.Atoi(splitted[1])

	return Pair{
		min: min,
		max: max,
	}
}

func (p Pair) IsFullyContained(otherPair Pair) bool {
	return p.min <= otherPair.min && p.max >= otherPair.max
}

func (p Pair) HasOverlap(otherPair Pair) bool {
	return p.max >= otherPair.min && p.max <= otherPair.max
}

func (p Pair) String() string {
	return fmt.Sprintf("%d-%d", p.min, p.max)
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

		firstPair, secondPair = SwapPairsIfNeeded(firstPair, secondPair)

		if firstPair.IsFullyContained(secondPair) || firstPair.HasOverlap(secondPair) {
			count++
		}
	}

	fmt.Println(count)
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
