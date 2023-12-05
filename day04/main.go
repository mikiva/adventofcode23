package main

import (
	"fmt"
	"math"
	"regexp"

	"aoc23/util"
)

type Card struct {
	id      string
	winning []string
	mine    []string
}

func intersect(l, m []string) []string {
	intersected := make([]string, 0)
	fmt.Println(l, m)
	for _, ll := range l {
		exists := false
		for _, mm := range m {
			if mm == ll {
				exists = true
				break
			}

		}
		if exists {
			intersected = append(intersected, ll)
		}
	}
	return intersected
}

func parseInput(input []string) []Card {
	var cards []Card = []Card{}

	re := regexp.MustCompile("[0-9]+")

	for _, row := range input {

		cardId, numbers := util.SplitInTwo(row, ":")
		winning, mine := util.SplitInTwo(numbers, "|")
		card := Card{id: cardId}

		w := re.FindAllString(winning, -1)
		fmt.Println("winning", w)
		card.winning = w

		m := re.FindAllString(mine, -1)
		card.mine = m
		//card.winning = w
		//for _, nbr := range strings.Split(winning, " ") {
		//	n, _ := strconv.ParseInt(nbr, 10, 32)
		//	if n == 0 {
		//		continue
		//	}
		//	card.winning = append(card.winning, int(n))
		//}

		//for _, nbr := range strings.Split(mine, " ") {
		//	n, _ := strconv.ParseInt(nbr, 10, 32)
		//	if n == 0 {
		//		continue
		//	}
		//	card.mine = append(card.mine, int(n))
		//}

		cards = append(cards, card)
	}
	return cards
}

func main() {
	input, _ := util.GetInputLines()

	games := parseInput(input)
	sum := 0
	for _, game := range games {

		intersected := intersect(game.winning, game.mine)
		power := math.Pow(2, float64(len(intersected))) / 2

		sum = sum + int(power)
	}
	fmt.Println("Part 1: ", sum)
}
