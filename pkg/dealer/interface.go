package dealer

import "cardgame/pkg/deck"

type Dealer interface {
	IsBlackjack() bool
	MakeTurn() int
	Hit()
	IsBust() bool
	GetScore(hidden bool) int
	GetCards() []*deck.Card
	SetDeck(deck deck.Deck)
	MakeAllTurns() bool
}
