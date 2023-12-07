package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

type Hand struct {
	cards    string
	bid      int
	handType HandType
}

// Hand type
type HandType int

const (
	HIGH_CARD HandType = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

func getHandType(cards string) HandType {
	cardsCount := map[string]int{
		"2": 0,
		"3": 0,
		"4": 0,
		"5": 0,
		"6": 0,
		"7": 0,
		"8": 0,
		"9": 0,
		"T": 0,
		"J": 0,
		"Q": 0,
		"K": 0,
		"A": 0,
	}
	for _, card := range cards {
		cardsCount[string(card)]++
	}
	for key, count := range cardsCount {
		if count == 5 {
			return FIVE_OF_A_KIND
		}
		if count == 4 {
			return FOUR_OF_A_KIND
		}
		if count == 3 {
			for _, count2 := range cardsCount {
				if count2 == 2 {
					return FULL_HOUSE
				}
			}
			return THREE_OF_A_KIND
		}
		if count == 2 {
			for key2, count2 := range cardsCount {
				if count2 == 3 {
					return FULL_HOUSE
				}
				if count2 == 2 && key != key2 {
					return TWO_PAIR
				}
			}
			return ONE_PAIR
		}
	}
	return HIGH_CARD
}

func cardScore(card string) int {
	switch card {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 11
	case "T":
		return 10
	case "9":
		return 9
	case "8":
		return 8
	case "7":
		return 7
	case "6":
		return 6
	case "5":
		return 5
	case "4":
		return 4
	case "3":
		return 3
	case "2":
		return 2
	default:
		return 0
	}
}

func compareHands(hand1 Hand, hand2 Hand) int {
	hand1Type := getHandType(hand1.cards)
	hand2Type := getHandType(hand2.cards)
	if hand1Type > hand2Type {
		return 1
	}
	if hand1Type < hand2Type {
		return -1
	}
	if hand1Type == hand2Type {
		return compareGeneric(hand1, hand2)
	}
	return 0
}

func compareGeneric(hand1 Hand, hand2 Hand) int {
	hand1Score := cardScore(string(hand1.cards[0]))
	hand2Score := cardScore(string(hand2.cards[0]))

	for i := 1; hand1Score == hand2Score; i++ {
		hand1Score = cardScore(string(hand1.cards[i]))
		hand2Score = cardScore(string(hand2.cards[i]))
	}

	if hand1Score > hand2Score {
		return 1
	}
	if hand1Score < hand2Score {
		return -1
	}
	return 0
}

func parseData(path string) []Hand {
	file, err := os.Open(path)
	check(err)
	result := []Hand{}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		newData := Hand{}
		var split = strings.Fields(line)
		bid, err := strconv.Atoi(split[1])
		check(err)
		newData.bid = bid
		newData.cards = split[0]
		newData.handType = getHandType(newData.cards)

		result = append(result, newData)
	}
	file.Close()
	return result
}
func sortHands(hands []Hand) []Hand {
	for i := 0; i < len(hands); i++ {
		for j := i + 1; j < len(hands); j++ {
			if compareHands(hands[i], hands[j]) == 1 {
				hands[i], hands[j] = hands[j], hands[i]
			}
		}
	}
	return hands
}

func main() {
	result := []Hand{}
	result = parseData("../data.txt")
	// result = parseData("../sample.txt")
	sortedHandsByType := make([][]Hand, 7)
	for _, hand := range result {
		sortedHandsByType[hand.handType] = append(sortedHandsByType[hand.handType], hand)
	}
	for _, handsByType := range sortedHandsByType {
		handsByType = sortHands(handsByType)
	}

	rank := 1
	score := 0
	for _, handsByType := range sortedHandsByType {
		for _, hand := range handsByType {
			score += rank * hand.bid
			rank++
		}
	}

	fmt.Println("Score:", score)
}
