package main

import (
	"aoc23/util"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

const (
	HIGH_CARD = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

var CARD_VALUES = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

type Card struct {
	Hand   string
	Values []int
	Bid    int
	Points int
}

func countOccurence(arr []int, val int) int {
	count := 0
	for _, v := range arr {
		if v == val {

			count = count + 1
		}
	}
	return count
}

func (card *Card) CalculateStrength() {

	m := make(map[string]int)
	cardValues := make([]int, 0)
	for _, c := range card.Hand {
		s := string(c)
		v, ok := m[s]
		if ok {
			m[s] = v + 1
		} else {
			m[s] = 1
		}
		cardValues = append(cardValues, CARD_VALUES[s])
	}
	v := make([]int, 0)
	points := 0
	for _, val := range m {
		v = append(v, val)
	}

	if slices.Contains(v, 5) {
		points = FIVE_OF_A_KIND
	} else if slices.Contains(v, 4) {
		points = FOUR_OF_A_KIND
	} else if slices.Contains(v, 3) && slices.Contains(v, 2) {
		points = FULL_HOUSE
	} else if slices.Contains(v, 3) && !slices.Contains(v, 2) {
		points = THREE_OF_A_KIND
	} else if slices.Contains(v, 2) {
		count := 0
		for _, two := range v {
			if two == 2 {
				count = count + 1
			}
		}
		if count == 2 {
			points = TWO_PAIR
		} else {
			points = ONE_PAIR
		}
	} else {
		points = HIGH_CARD
	}

	card.Points = points
	card.Values = cardValues
}
func (card *Card) CalculateStrength2() {

	m := make(map[string]int)
	cardValues := make([]int, 0)
	jokers := make([]string, 0)
	for _, c := range card.Hand {
		s := string(c)
		v, ok := m[s]
		if ok {
			m[s] = v + 1
		} else {
			m[s] = 1
		}

		if s != "J" {
			cardValues = append(cardValues, CARD_VALUES[s])
		} else {

			jokers = append(jokers, s)
		}
	}

	fmt.Println(jokers)

	v := make([]int, 0)
	points := 0
	for _, val := range m {
		v = append(v, val)
	}
	if slices.Contains(v, 5) {
		points = FIVE_OF_A_KIND
	} else if slices.Contains(v, 4) {
		points = FOUR_OF_A_KIND
	} else if slices.Contains(v, 3) && slices.Contains(v, 2) {
		points = FULL_HOUSE
	} else if slices.Contains(v, 3) && !slices.Contains(v, 2) {
		points = THREE_OF_A_KIND
	} else if slices.Contains(v, 2) {
		count := 0
		for _, two := range v {
			if two == 2 {
				count = count + 1
			}
		}
		if count == 2 {
			points = TWO_PAIR
		} else {
			points = ONE_PAIR
		}
	} else {
		points = HIGH_CARD
	}

	card.Points = points
	card.Values = cardValues
}

func createCards(raw []string) []Card {

	cards := make([]Card, 0)

	for _, r := range raw {
		vals := strings.Split(r, " ")
		bid, _ := strconv.Atoi(vals[1])
		card := Card{Hand: vals[0], Bid: bid, Points: 0}
		card.CalculateStrength()
		cards = append(cards, card)
	}
	return cards
}

func createCards2(raw []string) []Card {

	cards := make([]Card, 0)

	for _, r := range raw {
		vals := strings.Split(r, " ")
		bid, _ := strconv.Atoi(vals[1])
		card := Card{Hand: vals[0], Bid: bid, Points: 0}
		card.CalculateStrength2()
		cards = append(cards, card)
	}
	return cards
}

func sortCardsByPoints(cards []Card) {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Points < cards[j].Points
	})
}

func sortCardsByCards(cards []Card) []Card {
	sort.Slice(cards, func(i, j int) bool {
		swap := false
		for k := 0; k < len(cards[i].Values); k++ {
			if cards[i].Values[k] == cards[j].Values[k] {
				continue
			} else {
				swap = cards[i].Values[k] < cards[j].Values[k]
				break
			}
		}
		return swap
	})
	return cards
}

func part1(input []string) {

	cards := createCards(input)
	sortCardsByPoints(cards)
	samePoints := make(map[int][]Card)
	for _, card := range cards {
		samePoints[card.Points] = append(samePoints[card.Points], card)
	}
	keys := make([]int, 0)
	for k := range samePoints {
		keys = append(keys, k)

	}
	sort.Ints(keys)
	sortedCards := make([]Card, 0)
	for _, v := range keys {
		newCards := sortCardsByCards(samePoints[v])

		sortedCards = append(sortedCards, newCards...)
	}
	sum := 0

	for i, card := range sortedCards {

		sum = sum + card.Bid*(i+1)
	}
	fmt.Println("Part 1: ", sum)
}

func part2(input []string) {

	CARD_VALUES["J"] = 1
	cards := createCards2(input)
	sortCardsByPoints(cards)
	samePoints := make(map[int][]Card)
	for _, card := range cards {
		samePoints[card.Points] = append(samePoints[card.Points], card)
	}
	keys := make([]int, 0)
	for k := range samePoints {
		keys = append(keys, k)

	}
	sort.Ints(keys)
	sortedCards := make([]Card, 0)
	for _, v := range keys {
		newCards := sortCardsByCards(samePoints[v])

		sortedCards = append(sortedCards, newCards...)
	}
	sum := 0

	for i, card := range sortedCards {

		sum = sum + card.Bid*(i+1)
	}
	fmt.Println("Part 2: ", sum)
}

func main() {
	input, _ := util.GetInputLines()
	part1(input)

	part2(input)
}
