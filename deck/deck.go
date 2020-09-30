package deck

import (
	"math/rand"
	"time"
)

type Deck []Card

func New() (Deck, error) {
	var deck []Card

	for _, suit := range Suits {
		for _, rank := range Ranks {

			c, err := NewCard(rank, suit)

			if err != nil {
				return Deck{}, err
			}

			deck = append(deck, c)
		}
	}

	return deck, nil
}

func (d Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
