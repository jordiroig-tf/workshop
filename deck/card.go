package deck

import (
	"errors"
	"fmt"
)

var Suits = [4]string{"spades", "diamonds", "clubs", "hearts"}
var Ranks = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

type Card struct {
	rank string
	suit string
}

func (c Card) GetRank() string {
	return c.rank
}

func (c Card) GetSuit() string {
	return c.suit
}

func NewCard(rank string, suit string) (Card, error) {
	if !isValidRank(rank) {
		return Card{}, errors.New(fmt.Sprintf("Card rank is not valid: %s", rank))
	}

	if !isValidSuit(suit) {
		return Card{}, errors.New(fmt.Sprintf("Card suit is not valid: %s", suit))
	}

	return Card{rank: rank, suit: suit}, nil
}

func isValidRank(rank string) bool {
	for _, item := range Ranks {
		if item == rank {
			return true
		}
	}

	return false
}

func isValidSuit(suit string) bool {
	for _, item := range Suits {
		if item == suit {
			return true
		}
	}

	return false
}
