package deck

import (
	"errors"
	"fmt"
)

const Min = 1
const Max = 13

var Suits = [4]string{"spades", "diamonds", "clubs", "hearts"}

type Card struct {
	value int
	suit string
}

func (c Card) GetValue() int {
	return c.value
}

func (c Card) GetSuit() string {
	return c.suit
}

func NewCard(value int, suit string) (Card, error) {
	if value < Min || value > Max {
		return Card{}, errors.New(fmt.Sprintf("Card value is not valid: %d", value))
	}

	if !isValidSuit(suit) {
		return Card{}, errors.New(fmt.Sprintf("Card suit is not valid: %s", suit))
	}

	return Card{value: value, suit: suit}, nil
}

func isValidSuit(suit string) bool {
	for _, item := range Suits {
		if item == suit {
			return true
		}
	}

	return false
}
