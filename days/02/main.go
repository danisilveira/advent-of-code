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
	Undefined Shape = "Undefined"
	Rock      Shape = "Rock"
	Paper     Shape = "Paper"
	Scissors  Shape = "Scissors"
)

func NewShapeToWin(shape Shape) Shape {
	if shape == Rock {
		return Paper
	}

	if shape == Paper {
		return Scissors
	}

	if shape == Scissors {
		return Rock
	}

	return Undefined
}

func NewShapeToDraw(shape Shape) Shape {
	return shape
}

func NewShapeToLose(shape Shape) Shape {
	if shape == Rock {
		return Scissors
	}

	if shape == Paper {
		return Rock
	}

	if shape == Scissors {
		return Paper
	}

	return Undefined
}

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

	roundResultNeeded := hashtable.New[string, Result](3)
	roundResultNeeded.Set("X", Lose)
	roundResultNeeded.Set("Y", Draw)
	roundResultNeeded.Set("Z", Win)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplitted := strings.Split(line, " ")

		opponentMove, _ := opponentMoveToShape.Get(lineSplitted[0])
		resultNeeded, _ := roundResultNeeded.Get(lineSplitted[1])

		var myMove Shape

		switch resultNeeded {
		case Win:
			myMove = NewShapeToWin(opponentMove)
		case Draw:
			myMove = NewShapeToDraw(opponentMove)
		case Lose:
			myMove = NewShapeToLose(opponentMove)
		}

		score += (myMove.Score() + resultNeeded.Score())
	}

	log.Println(score)
}
