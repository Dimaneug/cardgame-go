package deck

import (
	"log"
	"strconv"
)

type RegularDeck struct {
	cards        []*Card
	playersLimit int
}

func (rd *RegularDeck) GetCardValue(card Card) int {
	rank := card.Rank
	switch rank {
	case "A":
		return 11
	case "K", "Q", "J":
		return 10
	default:
		num, err := strconv.Atoi(rank)
		if err != nil {
			log.Println("Error in rank convertion:", err)
		}
		return num
	}
}

func (rd *RegularDeck) GetScore(cards []*Card) int {
	score := 0
	aceCount := 0

	for _, card := range cards {
		if card.Rank == "A" {
			aceCount++
		}
		score += rd.GetCardValue(*card)
	}

	for i := 0; i < aceCount; i++ {
		if score > 21 {
			score -= 10
		}
	}

	return score
}

func (rd *RegularDeck) DeleteCard(i int) {
	copy(rd.cards[i:], rd.cards[i+1:])
	rd.cards[len(rd.cards)-1] = nil
	rd.cards = rd.cards[:len(rd.cards)-1]
}

func (rd *RegularDeck) GetPlayersLimit() int {
	return rd.playersLimit
}

func (rd *RegularDeck) GetCards() []*Card {
	return rd.cards
}

type RegularDeckFactory struct{}

func (rdf *RegularDeckFactory) MakeDeck() Deck {
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	suits := []string{"♥️", "♠️", "♣️", "♦️"}
	cards := []*Card{}
	for _, rank := range ranks {
		for _, suit := range suits {
			card := &Card{Rank: rank, Suit: suit}
			cards = append(cards, card)
		}
	}
	return &RegularDeck{cards: cards, playersLimit: 3}
}
