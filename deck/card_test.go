package deck_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jordiroig-tf/workshop/deck"
)

func TestNewCard(t *testing.T) {
	value := deck.Min
	suit := deck.Suits[0]

	card, err := deck.NewCard(value, suit)

	assert.IsType(t, deck.Card{}, card)
	assert.Nil(t, err)

	assert.Equal(t, value, card.GetValue())
	assert.Equal(t, suit, card.GetSuit())
}

func TestNewCardWrongSuit(t *testing.T) {
	value := deck.Min
	suit := "wrong"

	card, err := deck.NewCard(value, suit)

	assert.Error(t, err)
	assert.Empty(t, card)
	assert.EqualError(t, err, fmt.Sprintf("Card suit is not valid: %s", suit))
}

func TestNewCardWrongValue(t *testing.T) {
	value := -1
	suit := deck.Suits[0]

	card, err := deck.NewCard(value, suit)

	assert.Error(t, err)
	assert.Empty(t, card)
	assert.EqualError(t, err, fmt.Sprintf("Card value is not valid: %d", value))
}
