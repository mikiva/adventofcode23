package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLines(lines []string) int {
	sum := 0

	for _, s := range lines {
		var nums []int
		for _, l := range s {
			if i, err := strconv.ParseInt(string(l), 10, 32); err == nil {
				nums = append(nums, int(i))
			}

		}
		values, _ := strconv.ParseInt(fmt.Sprintf("%d%d", nums[0], nums[len(nums)-1]), 10, 32)
		sum = sum + int(values)
	}
	return sum
}

func parsePart2(lines []string) []string {

	var newLines []string
	for _, line := range lines {
		newLine := line
		for i, num := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
			newLine = strings.Replace(newLine, num, fmt.Sprintf("%s%d%s", num, i+1, num), -1)
		}
		newLines = append(newLines, newLine)
	}
	return newLines
}

func main() {
	var filename string
	flag.StringVar(&filename, "file", "test", "Input filename")
	flag.Parse()
	filename = fmt.Sprintf("%s.txt", filename)
	content, _ := os.ReadFile(filename)
	str := string(content)
	lines := strings.Split(str, "\n")
	part1 := parseLines(lines)
	fmt.Println(part1)
	parsedLines := parsePart2(lines)
	part2 := parseLines(parsedLines)
	fmt.Println(part2)

}
