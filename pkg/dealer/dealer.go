package dealer

import (
	"cardgame/pkg/deck"
	"math/rand"
)

type RegularDealer struct {
	dealerDeck deck.Deck
	cards      []*deck.Card
}

func (rd *RegularDealer) IsBlackjack() bool {
	if len(rd.cards) != 2 {
		return false
	}
	return rd.dealerDeck.GetScore(rd.cards) == 21
}

func (rd *RegularDealer) MakeTurn() int {
	rd.Hit()
	return rd.dealerDeck.GetScore(rd.cards)
}

func (rd *RegularDealer) Hit() {
	randIndex := rand.Intn(len(rd.dealerDeck.GetCards()))
	rd.cards = append(rd.cards, rd.dealerDeck.GetCards()[randIndex])
	rd.dealerDeck.DeleteCard(randIndex)
}

func (rd *RegularDealer) IsBust() bool {
	return rd.dealerDeck.GetScore(rd.cards) > 21
}

func (rd *RegularDealer) GetScore(hidden bool) int {
	if hidden {
		return rd.dealerDeck.GetScore(rd.cards[1:])
	}
	return rd.dealerDeck.GetScore(rd.cards)
}

func (rd *RegularDealer) GetCards() []*deck.Card {
	return rd.cards
}

func (rd *RegularDealer) SetDeck(deck deck.Deck) {
	rd.dealerDeck = deck
}
