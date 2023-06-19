package deck

type DeckFactory interface {
	MakeDeck() Deck
}

type Deck interface {
	GetCardValue(card Card) int
	GetScore(cards []*Card) int
	DeleteCard(i int)
	GetPlayersLimit() int
	GetCards() []*Card
}

type Card struct {
	Rank string
	Suit string
}
