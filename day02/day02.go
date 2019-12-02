package day02

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// An Intcode progam is a list of integers separated by commas. To run
// one, start by looking at the first integer (pos 0). Here, you will
// find an opcode - either 1,2, or 99. The opcode indicates what to do.
func IntCode(arr []int) []int {
	i := 0
	for i < len(arr) {
		operand := arr[i]
		switch operand {
		case 1:
			first, second, third := arr[i+1], arr[i+2], arr[i+3]
			arr[third] = arr[first] + arr[second]
		case 2:
			first, second, third := arr[i+1], arr[i+2], arr[i+3]
			arr[third] = arr[first] * arr[second]
		case 99:
			return arr
		default:
			log.Fatalf("Invalid operand %d found at index: %d", operand, i)
		}
		i += 4
	}
	return arr
}

func Driver() {
	s, err := ioutil.ReadFile("day02/input.txt")
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
	initialArray := make([]int, len(arr))
	copy(initialArray, arr)
	// As we need to restore the gravity assist program state, position 1
	// in the array must be set to 12, and position 2 must be set to 2
	arr[1], arr[2] = 12, 2
	res := IntCode(arr)
	fmt.Printf("Part 1: Position 0 element is %d\n", res[0])

	// The next part requires a bit of brute forcing. We need to determine
	// which noun and verb ends with the output
	// If you were more clever than me - you could do a binary search on the outputs.
	// Basically, start at i=50, j=50 and continue getting closer to the target
	// each iteration. This is because a map of the values produced will inclusively
	// be larger. I.e. for i, j in nouns, if i > j then res with i will be bigger.

	// allResults should give you an idea on how to implement a binary search on the
	// results if you want. I can't be bothered.
	allResults := make([][]int, 100)
	for i := range allResults {
		allResults[i] = make([]int, 100)
	}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			buffer := make([]int, len(initialArray))
			copy(buffer, initialArray)
			buffer[1], buffer[2] = i, j
			res := IntCode(buffer)
			if res[0] == 19690720 {
				fmt.Printf("Part 2: Res: %d Noun=%d, Verb=%d, 100 * Noun + Verb= %d\n", res[0], i, j, 100*i+j)
			}
			allResults[i][j] = res[0]
		}
	}

}
