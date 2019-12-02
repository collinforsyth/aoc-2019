package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// GetFuelRequirements returns the fuel required for a given mass
func CalculateBaseFuel(mass int) int {
	fuel := (mass / 3) - 2
	if fuel < 0 {
		return 0
	}
	return fuel
}

func CalculateActualFuel(mass int) int {
	actual := 0
	fuel := CalculateBaseFuel(mass)
	for {
		actual += fuel
		fuel = CalculateBaseFuel(fuel)
		if fuel == 0 {
			break
		}
	}
	return actual
}

func GetTotalFuel(fileName string, accountForFuel bool) int {
	result := 0
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		if !accountForFuel {
			result = result + CalculateBaseFuel(num)
		} else {
			result = result + CalculateActualFuel(num)
		}
	}

	return result
}

func Driver() {
	initialFuel := GetTotalFuel("day01/input.txt", false)
	actualFuel := GetTotalFuel("day01/input.txt", true)
	fmt.Printf("Initial Fuel Requirements=%d, Actual Fuel Requirements=%d\n", initialFuel, actualFuel)
}
