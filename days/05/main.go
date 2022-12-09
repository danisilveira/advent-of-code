package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"

	"github.com/danisilveira/algorithms-and-data-structures/datastructures/stack"
)

type Action struct {
	Quantity int
	From     stack.Stack[rune]
	To       stack.Stack[rune]
}

func NewActionFromString(sentence string, stacks []stack.Stack[rune]) Action {
	// move 5 from 8 to 2
	re := regexp.MustCompile("[0-9]+")
	numbers := re.FindAllString(sentence, -1)

	quantity, _ := strconv.Atoi(numbers[0])
	fromIndex, _ := strconv.Atoi(numbers[1])
	toIndex, _ := strconv.Atoi(numbers[2])

	return Action{
		Quantity: quantity,
		From:     stacks[fromIndex-1],
		To:       stacks[toIndex-1],
	}
}

func (a Action) Do() {
	items := []rune{}
	for i := 0; i < a.Quantity; i++ {
		item, _ := a.From.Pop()
		items = append(items, item)
	}

	for i := len(items) - 1; i >= 0; i-- {
		item := items[i]
		a.To.Push(item)
	}
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("failed to open the input file")
	}
	defer input.Close()

	inputs := []string{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}

		inputs = append(inputs, text)
	}

	stacks := build(inputs)
	fmt.Print("\n")

	debugStacks(stacks)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)

		action := NewActionFromString(text, stacks)
		action.Do()
		debugStacks(stacks)
	}

	for _, stack := range stacks {
		r, _ := stack.Pop()
		fmt.Print(string(r))
	}
	fmt.Printf("\n")
}

func build(inputs []string) []stack.Stack[rune] {
	stacks := make([]stack.Stack[rune], 9)
	for i := 0; i < 9; i++ {
		stacks[i] = stack.New[rune](1000)
	}

	fmt.Println(inputs[len(inputs)-1])
	for i := (len(inputs) - 2); i >= 0; i-- {
		input := inputs[i]
		fmt.Println(input)

		for j, r := range input {
			if unicode.IsSpace(r) {
				continue
			}

			if r == '[' || r == ']' {
				continue
			}

			stackIndex, _ := strconv.Atoi(string(inputs[len(inputs)-1][j]))

			stack := stacks[stackIndex-1]
			stack.Push(r)
		}
	}

	return stacks
}

func debugStacks(stacks []stack.Stack[rune]) {
	for i, stack := range stacks {
		fmt.Printf("[Stack %d] ", i)
		items := []rune{}
		for !stack.Empty() {
			r, _ := stack.Pop()
			items = append(items, r)
		}

		for i := len(items) - 1; i >= 0; i-- {
			item := items[i]
			fmt.Printf("%s ", string(item))
			stack.Push(item)
		}

		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
