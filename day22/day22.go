package day22

import (
	"strings"

	"github.com/anderslundholm/advent_of_code_2020/pkg/util"
)

const deckLen = 25

// Deck of cards.
type Deck []int

// Decks of cards.
type Decks []Deck

func (d Deck) drawCard() (int, Deck) {
	return d[0], d[1:]
}

func (d Deck) checkIfSubGame(d2 Deck, t1, t2 int) bool {
	return len(d) >= t1 && len(d2) >= t2
}

func parseInput(input []string) Decks {
	var deck1, deck2 Deck
	var decks Decks
	list := &deck1
	for _, line := range input {
		if line == "" {
			list = &deck2
			continue
		}
		if strings.HasPrefix(line, "Player") {
			continue
		}
		*list = append(*list, util.MustAtoi(line))
	}
	decks = append(decks, deck1, deck2)
	return decks
}

func calculateScore(input []string, part2 bool) int {
	var result int
	var deck []int
	decks := parseInput(input)
	if part2 {
		deck, _ = recursiveCombat(decks[0], decks[1])
	} else {
		deck = combat(decks[0], decks[1])
	}
	for i, num := range deck {
		result += (len(deck) - i) * num
	}
	return result
}

func combat(deck1, deck2 Deck) Deck {
	for len(deck1) != 0 && len(deck2) != 0 {
		var top1, top2 int
		top1, deck1 = deck1.drawCard()
		top2, deck2 = deck2.drawCard()

		if top1 > top2 {
			deck1 = append(deck1, top1, top2)
		} else {
			deck2 = append(deck2, top2, top1)
		}
	}
	if len(deck1) != 0 {
		return deck1
	}
	return deck2
}

func checkInfiniteLoop(deck1, deck2 []int, previosDecks map[[2][deckLen]int]bool) bool {
	var deckArrayCopy1 [deckLen]int
	var deckArrayCopy2 [deckLen]int
	copy(deckArrayCopy1[:], deck1)
	copy(deckArrayCopy2[:], deck2)

	if _, ok := previosDecks[[2][deckLen]int{deckArrayCopy1, deckArrayCopy2}]; ok {
		return true
	}
	previosDecks[[2][deckLen]int{deckArrayCopy1, deckArrayCopy2}] = true
	return false
}

func recursiveCombat(deck1, deck2 Deck) (Deck, int) {
	previosDecks := make(map[[2][deckLen]int]bool)
	for len(deck1) != 0 && len(deck2) != 0 {
		if checkInfiniteLoop(deck1, deck2, previosDecks) {
			return deck1, 1
		}
		var top1, top2 int
		top1, deck1 = deck1.drawCard()
		top2, deck2 = deck2.drawCard()

		if deck1.checkIfSubGame(deck2, top1, top2) {
			deckCopy1 := make([]int, top1)
			deckCopy2 := make([]int, top2)
			copy(deckCopy1, deck1[:top1])
			copy(deckCopy2, deck2[:top2])
			_, winner := recursiveCombat(deckCopy1, deckCopy2)
			if winner == 1 {
				deck1 = append(deck1, top1, top2)
			} else {
				deck2 = append(deck2, top2, top1)
			}
		} else {
			if top1 > top2 {
				deck1 = append(deck1, top1, top2)
			} else {
				deck2 = append(deck2, top2, top1)
			}
		}
	}
	if len(deck1) != 0 {
		return deck1, 1
	}
	return deck2, 2
}
