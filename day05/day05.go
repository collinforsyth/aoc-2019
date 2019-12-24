package day05

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type OpCode int

const (
	Add OpCode = iota + 1
	Multiply
	Input
	Output
	JumpIfTrue
	JumpIfFalse
	LessThan
	Equals
	Finish = 99
)

type Computer struct {
	pointer int
	inst    []int
}

// getArgument gets the positional argument of each
func (c *Computer) getArgument(position int) int {
	mode := (c.inst[c.pointer] / int(math.Pow10(2+position))) % 10
	if mode == 1 { // if we are in immediate mode
		return c.inst[c.pointer+position+1]
	}
	return c.inst[c.inst[c.pointer+position+1]]
}

// Execute executes an instruction
func ExecuteIntCode(in []int, input int) {
	c := Computer{0, in}
	for c.pointer < len(c.inst) {
		instruction := c.inst[c.pointer]
		switch OpCode(instruction % 100) {
		case Add:
			// Add takes 3 parameters, check each is positional or immediate mode
			first := c.getArgument(0)
			second := c.getArgument(1)
			result := first + second
			// parameters that an instruction writes to will never be in immediate mode
			c.inst[c.inst[c.pointer+3]] = result
			c.pointer += 4
		case Multiply:
			// Add takes 3 parameters, check each is positional or immediate mode
			first := c.getArgument(0)
			second := c.getArgument(1)
			result := first * second
			// parameters that an instruction writes to will never be in immediate mode
			c.inst[c.inst[c.pointer+3]] = result
			c.pointer += 4
		case Input:
			c.inst[c.inst[c.pointer+1]] = input
			c.pointer += 2
		case Output:
			fmt.Printf("%d\n", c.inst[c.inst[c.pointer+1]])
			c.pointer += 2
		case JumpIfTrue:
			first, second := c.getArgument(0), c.getArgument(1)
			if first != 0 {
				c.pointer = second
			} else {
				c.pointer += 3
			}
		case JumpIfFalse:
			first, second := c.getArgument(0), c.getArgument(1)
			if first == 0 {
				c.pointer = second
			} else {
				c.pointer += 3
			}
		case LessThan:
			first, second := c.getArgument(0), c.getArgument(1)
			if first < second {
				c.inst[c.inst[c.pointer+3]] = 1
			} else {
				c.inst[c.inst[c.pointer+3]] = 0
			}
			c.pointer += 4
		case Equals:
			first, second := c.getArgument(0), c.getArgument(1)
			if first == second {
				c.inst[c.inst[c.pointer+3]] = 1
			} else {
				c.inst[c.inst[c.pointer+3]] = 0
			}
			c.pointer += 4
		case Finish:
			return
		default:
			log.Fatalf("Invalid instruction: %d\n", instruction)
		}
	}
}

func Driver() {
	s, err := ioutil.ReadFile("day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Get rid of newline character
	words := strings.Split(strings.TrimRight(string(s), "\n"), ",")
	arr := make([]int, len(words))
	for i, num := range words {
		arr[i], err = strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
	}
	firstArray := make([]int, len(arr))
	secondArray := make([]int, len(arr))
	copy(firstArray, arr)
	copy(secondArray, arr)
	ExecuteIntCode(firstArray, 1)
	ExecuteIntCode(secondArray, 5)

}
