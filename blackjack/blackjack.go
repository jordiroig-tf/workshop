package blackjack

import (
	"errors"

	"github.com/jordiroig-tf/workshop/deck"
)

const dealer = "Dealer"

var values = map[string]int{
	deck.Ranks[0]:  1, //A
	deck.Ranks[1]:  2,
	deck.Ranks[2]:  3,
	deck.Ranks[3]:  4,
	deck.Ranks[4]:  5,
	deck.Ranks[5]:  6,
	deck.Ranks[6]:  7,
	deck.Ranks[7]:  8,
	deck.Ranks[8]:  9,
	deck.Ranks[9]:  10,
	deck.Ranks[10]: 10, //J
	deck.Ranks[11]: 10, //Q
	deck.Ranks[12]: 10, //K
}

type BlackJack struct {
	players     []Player
	deck        deck.Deck
	playerDecks map[Player]deck.Deck
}

type Player string

func New(players []Player) (BlackJack, error) {
	dealer := Player(dealer)
	players = append(players, dealer)

	newDeck, err := deck.New()

	if err != nil {
		return BlackJack{}, errors.New("error creating the deck")
	}

	newDeck.Shuffle()

	playerDecks := map[Player]deck.Deck{}
	for _, player := range players {
		playerDecks[player] = deck.Deck{}
	}

	return BlackJack{players: players, deck: newDeck, playerDecks: playerDecks}, nil
}
