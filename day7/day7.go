package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var CARDSTR = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

var CARDSTR2 = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
}

var HANDTYPE = map[int]string{
	25: "Five of a kind",
	17: "Four of a kind",
	13: "Full house",
	11: "Three of a kind",
	9:  "Two pair",
	7:  "One pair",
	5:  "High card",
}

type card struct {
	hand string
	bid  int
}

func removeEmptyElement(arr []string) []string {
	ptr1 := 0
	for ptr2, val := range arr {
		if val != "" {
			arr[ptr1] = arr[ptr2]
			ptr1++
		}
	}

	return arr[:ptr1]
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func checkPower(hand string) int {
	charMap := make(map[string]int)
	handPower := 0

	for _, char := range hand {
		charMap[string(char)]++
	}

	for _, char := range hand {
		if value, found := charMap[string(char)]; found {
			switch value {
			case 1:
				handPower += 1
			case 2:
				handPower += 2
			case 3:
				handPower += 3
			case 4:
				handPower += 4
			case 5:
				handPower += 5
			}
		}
	}

	return handPower
}

func checkPower2(hand string) int {
	charMap := make(map[string]int)
	handPower := 0
	jokerFound := 0
	mostDupes := ""

	for _, char := range hand {
		if string(char) == "J" {
			jokerFound++
			continue
		}
		charMap[string(char)]++
	}

	if jokerFound > 0 {
		temp := 0
		for _, char := range hand {
			if charMap[string(char)] > temp {
				temp = charMap[string(char)]
				mostDupes = string(char)
			}
		}
		charMap[mostDupes] += jokerFound
		charMap["J"] = charMap[mostDupes]
	}

	for _, char := range hand {
		if value, found := charMap[string(char)]; found {
			switch value {
			case 1:
				handPower += 1
			case 2:
				handPower += 2
			case 3:
				handPower += 3
			case 4:
				handPower += 4
			case 5:
				handPower += 5
			}
		}
	}

	return handPower
}

func sortCards(cards []card) []card {
	powercmp := func(a, b card) int {
		aStr := checkPower(a.hand)
		bStr := checkPower(b.hand)

		if aStr == bStr {
			for i := range a.hand {
				cardCmp := cmp.Compare(CARDSTR[string(a.hand[i])], CARDSTR[string(b.hand[i])])
				if cardCmp != 0 {
					return cardCmp
				}
			}
		}

		return cmp.Compare(aStr, bStr)
	}

	slices.SortFunc(cards, powercmp)

	return cards
}

func sortCards2(cards []card) []card {
	powercmp := func(a, b card) int {
		aStr := checkPower2(a.hand)
		bStr := checkPower2(b.hand)

		if aStr == bStr {
			for i := range a.hand {
				cardCmp := cmp.Compare(CARDSTR2[string(a.hand[i])], CARDSTR2[string(b.hand[i])])
				if cardCmp != 0 {
					return cardCmp
				}
			}
		}

		return cmp.Compare(aStr, bStr)
	}

	slices.SortFunc(cards, powercmp)

	return cards
}

func part1(input []string) {
	var cards []card
	sum := 0

	for _, line := range input {
		str := strings.TrimSpace(line)
		strArr := strings.Split(str, " ")
		bidNum, err := strconv.Atoi(strArr[1])
		checkError(err)
		currentCard := card{strArr[0], bidNum}
		cards = append(cards, currentCard)
	}

	sortedCards := sortCards(cards)

	for i, hand := range sortedCards {
		sum += hand.bid * (i + 1)
	}

	fmt.Println(sum)
}

func part2(input []string) {
	var cards []card
	sum := 0

	for _, line := range input {
		str := strings.TrimSpace(line)
		strArr := strings.Split(str, " ")
		bidNum, err := strconv.Atoi(strArr[1])
		checkError(err)
		currentCard := card{strArr[0], bidNum}
		cards = append(cards, currentCard)
	}

	sortedCards := sortCards2(cards)

	for i, hand := range sortedCards {
		sum += hand.bid * (i + 1)
	}

	fmt.Println(sum)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	part1(input)
	part2(input)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
