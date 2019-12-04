package day04

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

// GetNextNondecreasingPermutation finds the next nondecreasing
// permutation of the byte sequence
func GetNextNondecreasingPermutation(num int) int {
	b := []byte(strconv.Itoa(num))
	i := 0
	for i < len(b)-1 {
		if b[i] > b[i+1] {
			break
		}
		i++
	}
	// Pivot is equal to b[i]
	for j := i + 1; j < len(b); j++ {
		b[j] = b[i]
	}
	ret, err := strconv.Atoi(string(b))
	if err != nil {
		log.Fatal(err)
	}
	return ret
}

func isValidPasswordPart1(num int) bool {
	// Range and 6 digits is already satisfied by the loop
	// Need to check that we have adjacent numbers and int is nondecreasing
	adjacentNums := false
	for i := 1; i < 6; i++ {
		j := (num % int(math.Pow10(i))) / int(math.Pow10(i-1))
		k := (num % int(math.Pow10(i+1))) / int(math.Pow10(i))
		if j < k {
			return false
		} else if j == k {
			adjacentNums = true
		}
	}
	return true && adjacentNums
}

// Password rules
// 1. It is a 6 digit number
// 2. Two adjacent digits must be the same
// 3. The value is within the range in your puzzle input
// 4. Going from left to right, the digits never decrease
func NumberOfPasswordsPart1(a, b int) int {
	count := 0
	i := a

	for i <= b {
		if isValidPasswordPart1(i) {
			count++
			i++
		} else {
			next := GetNextNondecreasingPermutation(i)
			if i == next {
				i++
			} else {
				i = next
			}
		}
	}
	return count
}

func isValidPasswordPart2(num int) bool {
	// Range and 6 digits is already satisfied by the loop
	// Need to check that we have adjacent numbers and int is nondecreasing
	// For part 2, adjacentNumbers need to be exactly size 2.
	adjacentNums := false
	// isMatching is a flag that tracks if we are tracking a matching state
	for i := 1; i < 6; i++ {
		j := (num % int(math.Pow10(i))) / int(math.Pow10(i-1))
		k := (num % int(math.Pow10(i+1))) / int(math.Pow10(i))
		if j < k {
			return false
		} else if j == k {
			// j is element i, k is element i-1, we can check bounds to confirm
			if i == 1 {
				// Can't check on left
				l := (num % int(math.Pow10(i+2))) / int(math.Pow10(i+1))
				if l != k {
					adjacentNums = true
				}
			} else if i == 5 {
				// Can't check on right
				l := (num % int(math.Pow10(i-1))) / int(math.Pow10(i-2))
				if l != k {
					adjacentNums = true
				}
			} else {
				m := (num % int(math.Pow10(i+2))) / int(math.Pow10(i+1))
				l := (num % int(math.Pow10(i-1))) / int(math.Pow10(i-2))
				if m != j && l != j {
					adjacentNums = true
				}
			}
		}
	}
	return true && adjacentNums
}

func NumberOfPasswordsPart2(a, b int) int {
	count := 0
	i := a

	for i <= b {
		if isValidPasswordPart2(i) {
			count++
			i++
		} else {
			next := GetNextNondecreasingPermutation(i)
			if i == next {
				i++
			} else {
				i = next
			}
		}
	}
	return count
}

func Driver() {
	a, b := 145852, 616942
	resultOne := NumberOfPasswordsPart1(a, b)
	resultTwo := NumberOfPasswordsPart2(a, b)
	fmt.Printf("Part1: Number of passwords is: %d\n", resultOne)
	fmt.Printf("Part2: Number of passwords is: %d\n", resultTwo)
}
