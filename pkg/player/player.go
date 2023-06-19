package player

import (
	"fmt"
	"math/rand"

	"cardgame/pkg/deck"
)

type RegularPlayer struct {
	playerDeck deck.Deck
	cards      []*deck.Card
}

func (rp *RegularPlayer) IsBlackjack() bool {
	if len(rp.cards) != 2 {
		return false
	}
	return rp.playerDeck.GetScore(rp.cards) == 21
}

func (rp *RegularPlayer) MakeTurn(i int) bool {

	for {
		fmt.Printf("Игрок %d, выберите действие:\n", i)
		fmt.Println("1 - Взять")
		fmt.Println("2 - Пас")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			rp.Hit()
			return false
		case 2:
			return true
		default:
			fmt.Println("Неверный номер действия.")
			continue
		}
	}

}

func (rp *RegularPlayer) Hit() {
	randIndex := rand.Intn(len(rp.playerDeck.GetCards()))
	rp.cards = append(rp.cards, rp.playerDeck.GetCards()[randIndex])
	rp.playerDeck.DeleteCard(randIndex)
}

func (rp *RegularPlayer) IsBust() bool {
	return rp.playerDeck.GetScore(rp.cards) > 21
}

func (rp *RegularPlayer) GetScore() int {
	return rp.playerDeck.GetScore(rp.cards)
}

func (rp *RegularPlayer) GetCards() []*deck.Card {
	return rp.cards
}

func (rp *RegularPlayer) SetDeck(deck deck.Deck) {
	rp.playerDeck = deck
}
