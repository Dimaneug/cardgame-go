package main

import (
	"cardgame/pkg/dealer"
	"cardgame/pkg/deck"
	"cardgame/pkg/game"
	"cardgame/pkg/player"
)

func main() {
	deckFactory := deck.RegularDeckFactory{}
	newDeck := deckFactory.MakeDeck()

	game := game.Game{}
	game.SetDeck(newDeck)
	game.SetDealer(&dealer.RegularDealer{})
	game.AddPlayer(&player.RegularPlayer{})
	game.Start()
}
