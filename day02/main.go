package main

import (
	"aoc23/util"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func getGame(row string) (int, string) {
	r := strings.Split(row, ":")
	id := strings.Split(r[0], " ")

	nbr, _ := strconv.ParseInt(id[1], 10, 32)
	return int(nbr), r[1]

}

func checkGame(game string) int {
	_, g := getGame(game)
	sets := getSets(g)

	red := int64(0)
	green := int64(0)
	blue := int64(0)
	for _, set := range sets {
		throw := strings.Split(strings.TrimSpace(set), ", ")
		for _, t := range throw {

			th := strings.Split(t, " ")
			amount, _ := strconv.ParseInt(th[0], 10, 32)
			if th[1] == "red" && (amount > red) {
				red = amount
			} else if th[1] == "green" && amount > green {
				green = amount
			} else if th[1] == "blue" && amount > blue {
				blue = amount
			}

		}

	}

	return int(red * green * blue)

}

func getSets(game string) []string {
	sets := strings.Split(game, ";")

	return sets
}

var rules = map[string]int{"red": 12, "green": 13, "blue": 14}

func isSetPoissible(set string) bool {

	for _, throws := range strings.Split(set, ",") {

		throw := strings.Split(strings.TrimSpace(throws), " ")

		amount, _ := strconv.ParseInt(throw[0], 10, 32)

		if int(amount) > rules[throw[1]] {

			return false
		}
	}

	return true
}

func part1(input []string) {
	sum := 0
	for _, row := range input {
		possible := true
		id, game := getGame(row)
		sets := getSets(game)
		for _, set := range sets {
			if !isSetPoissible(set) {
				possible = false
				break
			}
		}
		if possible {
			sum = sum + id
		}
	}
	fmt.Println("Part 1: ", sum)

}

func part2(input []string) {
	sum := 0
	for _, row := range input {
		g := checkGame(row)

		sum = sum + g
	}
	fmt.Println("Part 2: ", sum)
}

func main() {

	input, err := util.GetInputLines()
	if err != nil {
		log.Fatal(err.Error())
	}
	part1(input)

	part2(input)
}
