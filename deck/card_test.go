package deck_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jordiroig-tf/workshop/deck"
)

func TestNewCard(t *testing.T) {
	rank := deck.Ranks[0]
	suit := deck.Suits[0]

	card, err := deck.NewCard(rank, suit)

	assert.IsType(t, deck.Card{}, card)
	assert.Nil(t, err)

	assert.Equal(t, rank, card.GetRank())
	assert.Equal(t, suit, card.GetSuit())
}

func TestNewCardWrongSuit(t *testing.T) {
	rank := deck.Ranks[0]
	suit := "wrong"

	card, err := deck.NewCard(rank, suit)

	assert.Error(t, err)
	assert.Empty(t, card)
	assert.EqualError(t, err, fmt.Sprintf("Card suit is not valid: %s", suit))
}

func TestNewCardWrongValue(t *testing.T) {
	rank := "wrong"
	suit := deck.Suits[0]

	card, err := deck.NewCard(rank, suit)

	assert.Error(t, err)
	assert.Empty(t, card)
	assert.EqualError(t, err, fmt.Sprintf("Card rank is not valid: %s", rank))
}
