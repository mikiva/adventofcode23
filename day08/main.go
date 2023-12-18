package main

import (
	"aoc23/util"
	"fmt"
	"regexp"
	"strings"
)

type Node struct {
	Key   string
	Value map[string]string
}
type fn func(a string) bool

func checkP1(val string) bool {
	return val == "ZZZ"
}
func checkP2(val string) bool {
	return strings.HasSuffix(val, "Z")
}

func solve(instructions []string, nodes map[string]map[string]string, start string, comparing fn) int {

	steps := 1
	next := start
End:
	for {
		for _, dir := range instructions {
			node := nodes[next]
			next = node[dir]
			if comparing(next) {
				break End
			}
			steps++
		}
	}
	return steps

}

func parseNetwork(input []string) map[string]map[string]string {
	reg := `[A-Z0-9]{3}`
	re := regexp.MustCompile(reg)
	docs := make(map[string]map[string]string)
	for _, v := range input {
		matches := re.FindAllStringSubmatch(v, -1)
		docs[matches[0][0]] = map[string]string{"L": matches[1][0], "R": matches[2][0]}

	}
	return docs
}
func part2(instructions []string, nodes map[string]map[string]string) int {
	starts := make([]string, 0)
	for k := range nodes {
		if strings.HasSuffix(k, "A") {
			starts = append(starts, k)
		}
	}
	allSteps := make([]int, len(starts))
	for i, start := range starts {
		steps := solve(instructions, nodes, start, checkP2)
		allSteps[i] = steps
	}
	return util.LCM(allSteps...)
}

func main() {
	input, _ := util.GetInputLines()

	instructions := strings.Split(input[0], "")
	network := parseNetwork(input[2:])
	p1 := solve(instructions, network, "AAA", checkP1)
	fmt.Println("Part 1:", p1)

	p2 := part2(instructions, network)
	fmt.Println("Part 2: ", p2)

}
