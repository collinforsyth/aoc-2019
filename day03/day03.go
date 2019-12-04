package day03

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

type coordinate struct {
	x     int
	y     int
	steps int
}

func (c coordinate) String() string {
	return fmt.Sprintf("x=%d, y=%d, steps=%d", c.x, c.y, c.steps)
}

type Line struct {
	start coordinate
	end   coordinate
}

func (l Line) String() string {
	return fmt.Sprintf("start={%s}, end={%s}", l.start, l.end)
}

func buildVectors(input string) ([]Line, []Line) {
	inputs := strings.Split(input, ",")
	origin := coordinate{0, 0, 0}
	vertical, horizontal := []Line{}, []Line{}
	// steps is the amount of steps that it took to reach the current start
	steps := 0
	for _, s := range inputs {
		var end coordinate
		length, err := strconv.Atoi(s[1:])
		if err != nil {
			log.Fatal(err)
		}
		steps += length
		switch string(s[0]) {
		case "U":
			end = coordinate{origin.x, origin.y + length, steps}
			vertical = append(vertical, Line{start: origin, end: end})
		case "R":
			end = coordinate{origin.x + length, origin.y, steps}
			horizontal = append(horizontal, Line{start: origin, end: end})
		case "D":
			end = coordinate{origin.x, origin.y - length, steps}
			// Write coordinates so that start < end
			vertical = append(vertical, Line{start: end, end: origin})
		case "L":
			end = coordinate{origin.x - length, origin.y, steps}
			// We write coordinates so that start < end
			horizontal = append(horizontal, Line{start: end, end: origin})
		default:
			log.Fatalf("Unknown direction %s", string(s[0]))
		}
		origin = end
	}
	sort.Slice(vertical, func(i, j int) bool { return vertical[i].start.x < vertical[j].start.x })
	sort.Slice(horizontal, func(i, j int) bool { return horizontal[i].start.y < horizontal[j].start.y })
	return vertical, horizontal
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// Get intersections should record where all intersections are.
// If we consider a line, (a, b) that lies the x axis, it can only
// intersect with a line on the y axis. Therefore, we must consider
// the candidate lines in the second set to have an x value, C, s.t.
// a <= C <= b
func (l Line) GetHorizontalIntersections(lines []Line) []coordinate {
	intersections := []coordinate{}
	a, b, c := l.start.y, l.end.y, l.start.x
	for _, line := range lines {
		if line.start.y < a {
			continue
		} else if line.start.y > b {
			// No longer in possible set of candidates
			break
		}
		// This will be all lines with a <= line.start.y <= b
		if (line.start.x <= c) && (c <= line.end.x) {
			var xTotal, yTotal int
			if l.start.steps < l.end.steps { // l is the vertical line
				yTotal = l.start.steps + Abs(line.start.y-l.start.y)
			} else {
				yTotal = l.end.steps + Abs(line.start.y-l.end.y)
			}
			if line.start.steps < line.end.steps {
				xTotal = line.start.steps + Abs(c-line.start.x)
			} else {
				xTotal = line.end.steps + Abs(c-line.end.x)
			}
			intersections = append(intersections, coordinate{c, line.start.y, yTotal + xTotal})
		}
	}
	return intersections
}

func (l Line) GetVerticalIntersections(lines []Line) []coordinate {
	intersections := []coordinate{}
	a, b, c := l.start.x, l.end.x, l.start.y
	for _, line := range lines {
		if line.start.x < a {
			continue
		} else if line.start.x > b {
			// No longer in possible set of candidates
			break
		}
		// This will be all lines with a <= line.start.y <= b
		if (line.start.y <= c) && (c <= line.end.y) {
			var xTotal, yTotal int
			if l.start.steps < l.end.steps { // l is the horizontal line
				xTotal = l.start.steps + Abs(line.start.x-l.start.x)
			} else {
				xTotal = l.end.steps + Abs(line.start.x-l.end.x)
			}
			if line.start.steps < line.end.steps {
				yTotal = line.start.steps + Abs(c-line.start.y)
			} else {
				yTotal = line.start.steps + Abs(c-line.end.y)
			}
			intersections = append(intersections, coordinate{line.start.x, c, yTotal + xTotal})
		}
	}
	return intersections

}

func ManhattanDistance(i coordinate, j coordinate) (int, int) {
	x := i.x - j.x
	y := i.y - j.y
	return x, y
}

func Distance(i coordinate, j coordinate) float64 {
	return math.Sqrt(math.Pow((math.Abs(float64(i.x))-float64(j.x)), 2) + math.Pow(math.Abs(float64(i.y)-float64(j.y)), 2))
}

func Driver() {
	in, err := ioutil.ReadFile("day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	split := strings.Split(string(in), "\n")
	var intersections []coordinate
	firstVertical, firstHorizontal := buildVectors(split[0])
	secondVertical, secondHorizontal := buildVectors(split[1])
	//fmt.Printf("%v\n%v\n%v\n%v\n", firstVertical, firstHorizontal, secondVertical, secondHorizontal)
	for _, l := range firstVertical {
		intersections = append(intersections, l.GetHorizontalIntersections(secondHorizontal)...)
	}
	for _, l := range firstHorizontal {
		intersections = append(intersections, l.GetVerticalIntersections(secondVertical)...)
	}
	minDistance := math.Inf(1)
	origin := coordinate{0, 0, 0}
	var closest coordinate
	minSteps := int(^uint(0) >> 1)
	var minStepsCoord coordinate
	for _, intersection := range intersections {
		dist := Distance(origin, intersection)
		if dist < minDistance {
			minDistance = dist
			closest = intersection
		}
		if intersection.steps < minSteps {
			minStepsCoord = intersection
			minSteps = intersection.steps
		}
	}
	fmt.Printf("Part1: Min distance is: %s, %d + %d = %d\n", closest, closest.x, closest.y, closest.x+closest.y)
	fmt.Printf("Part2: Min steps coordinate is %s, steps=%d", minStepsCoord, minStepsCoord.steps)

}
