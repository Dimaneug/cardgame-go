package player

import "cardgame/pkg/deck"

type Player interface {
	IsBlackjack() bool
	MakeTurn(i int) bool
	Hit()
	// Stand()
	IsBust() bool
	GetScore() int
	GetCards() []*deck.Card
	SetDeck(deck deck.Deck)
}
