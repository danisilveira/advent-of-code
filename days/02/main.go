package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/danisilveira/algorithms-and-data-structures/datastructures/hashtable"
)

type Result string

const (
	Win  Result = "Win"
	Lose Result = "Lose"
	Draw Result = "Draw"
)

func (r Result) Score() int {
	switch r {
	case Win:
		return 6
	case Draw:
		return 3
	case Lose:
		return 0
	}

	return 0
}

type Shape string

const (
	Rock     Shape = "Rock"
	Paper    Shape = "Paper"
	Scissors Shape = "Scissors"
)

func (m Shape) Score() int {
	switch m {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}

	return 0
}

func (m Shape) Compare(otherMove Shape) Result {
	if m == otherMove {
		return Draw
	}

	if m == Rock && otherMove == Paper {
		return Lose
	}

	if m == Paper && otherMove == Scissors {
		return Lose
	}

	if m == Scissors && otherMove == Rock {
		return Lose
	}

	return Win
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("failed to open the file input")
	}
	defer input.Close()

	score := 0

	opponentMoveToShape := hashtable.New[string, Shape](3)
	opponentMoveToShape.Set("A", Rock)
	opponentMoveToShape.Set("B", Paper)
	opponentMoveToShape.Set("C", Scissors)

	myMoveToShape := hashtable.New[string, Shape](3)
	myMoveToShape.Set("X", Rock)
	myMoveToShape.Set("Y", Paper)
	myMoveToShape.Set("Z", Scissors)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		moves := strings.Split(line, " ")

		opponentMove, _ := opponentMoveToShape.Get(moves[0])
		myMove, _ := myMoveToShape.Get(moves[1])

		result := myMove.Compare(opponentMove)
		score += (myMove.Score() + result.Score())
	}

	log.Println(score)
}
