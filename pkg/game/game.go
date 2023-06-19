package game

import (
	"fmt"

	"cardgame/pkg/dealer"
	"cardgame/pkg/deck"
	"cardgame/pkg/player"
)

type Game struct {
	deck    deck.Deck
	players []player.Player
	dealer  dealer.Dealer
}

func (g *Game) AddPlayer(player player.Player) {
	if g.deck.GetPlayersLimit() > len(g.players) {
		g.players = append(g.players, player)
	}
}

func (g *Game) SetDeck(deck deck.Deck) {
	g.deck = deck
}

func (g *Game) SetDealer(dealer dealer.Dealer) {
	g.dealer = dealer
}

func (g *Game) Start() {
	g.dealer.SetDeck(g.deck)
	for i := 0; i < 2; i++ {
		g.dealer.Hit()
	}

	fmt.Printf("\nКарты раздающего: {Скрытая карта} %v\n", *g.dealer.GetCards()[1])
	fmt.Println("Счёт:", g.dealer.GetScore(true))
	fmt.Println()

	for _, player := range g.players {
		player.SetDeck(g.deck)
		for i := 0; i < 2; i++ {
			player.Hit()
		}
	}

	if g.dealer.IsBlackjack() {
		fmt.Println("У раздающего блэкджек.")
		isBlackjack := false
		for i, player := range g.players {
			if player.IsBlackjack() {
				fmt.Printf("Игрок %d, ", i+1)
				isBlackjack = true
			}
		}
		if isBlackjack {
			fmt.Println(" - ничья. Остальные проиграли")
		} else {
			fmt.Println("Все проиграли")
		}
		return
	}

	playersBust := 0
	for i, player := range g.players {
		fmt.Printf("Карты игрока %d:", i+1)
		for _, card := range player.GetCards() {
			fmt.Printf(" %v", *card)
		}
		fmt.Printf("\nСчёт игрока %d: %d\n\n", i+1, player.GetScore())

		if player.IsBlackjack() {
			fmt.Printf("У игрока %d блэкджек.\n", i+1)
			continue
		}

		for {
			if player.MakeTurn(i + 1) {
				fmt.Printf("У игрока %d счёт %d\n\n", i+1, player.GetScore())
				break
			}
			fmt.Printf("\nКарты игрока %d:", i+1)
			for _, card := range player.GetCards() {
				fmt.Printf(" %v", *card)
			}
			fmt.Printf("\nСчёт игрока %d: %d\n\n", i+1, player.GetScore())
			if player.IsBust() {
				fmt.Printf("У игрока %d перебор.\n", i+1)
				playersBust++
				break
			}
		}
	}

	if playersBust == len(g.players) {
		fmt.Println("Все проиграли")
		return
	}

	if g.dealer.MakeAllTurns() {
		return
	}
	dealerScore := g.dealer.GetScore(false)

	fmt.Println("--------------------")
	fmt.Println("\nИтоговый счёт")
	fmt.Printf("Очки раздающего: %d", dealerScore)
	for _, card := range g.dealer.GetCards() {
		fmt.Printf(" %v", *card)
	}
	fmt.Println()
	for i, player := range g.players {
		score := player.GetScore()
		fmt.Printf("Игрок %d: %d", i+1, score)
		if score > dealerScore && score < 22 {
			fmt.Println(" - победил")
		} else {
			fmt.Println(" - проиграл")
		}
	}

}
