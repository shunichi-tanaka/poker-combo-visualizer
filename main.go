package main

import (
	"fmt"

	"github.com/chehsunliu/poker"
)

func main() {
	deck := poker.NewDeck()
	hand := deck.Draw(2)
	board := deck.Draw(3)
	enemyHand := deck.Draw(2)

	fmt.Println(hand)
	fmt.Println(enemyHand)

	evHand := append(hand, board...)
	fmt.Println(board)
	evEnemyHand := append(enemyHand, board...)
	fmt.Println(poker.Evaluate(evHand))
	fmt.Println(poker.Evaluate(evEnemyHand))
}
