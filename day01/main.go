package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(lines []string) int {
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

func minmax(vals []int) (int, int) {
	min := vals[0]
	max := vals[0]

	for _, num := range vals {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}

func part2(lines []string) []string {

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
	sum := part1(lines)
	fmt.Println(sum)
	newLines := part2(lines)
	newSum := part1(newLines)
	fmt.Println(newSum)

}
