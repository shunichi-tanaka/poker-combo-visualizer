package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/chehsunliu/poker"
	handrange "github.com/shunichi-tanaka/poker-combo-visualizer/pkg"
)

func main() {
	fmt.Println("abc")
	deck := poker.NewDeck()
	board := deck.Draw(3)
	//enemyHand := deck.Draw(2)

	//fmt.Println(hand)
	//fmt.Println(enemyHand)

	//evHand := append(hand, board...)
	//fmt.Println(board)
	//evEnemyHand := append(enemyHand, board...)
	//fmt.Println(poker.Evaluate(evHand))
	//fmt.Println(poker.Evaluate(evEnemyHand))

	//fmt.Println(enemyHand[0].String())

	// Rangeの用意
	rangeMatrix := [13][13]int8{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	myRange := handrange.NewRange(&rangeMatrix)
	hands := myRange.GetComboHands()
	fmt.Println("--- Hands ---")
	fmt.Println(board)

	for _, v := range hands {
		strength := append(board, v.Fcard, v.Scard)
		v.Rank = poker.Evaluate(strength)
		v.RankString = poker.RankString(v.Rank)
		//fmt.Print(v.String() + "  ")
		//fmt.Println(poker.Evaluate(strength))
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Rank < hands[j].Rank
	})

	rankHighCard := 0
	rankPair := 0
	rankTwoPair := 0
	rankThreeOfAKind := 0
	rankStraight := 0
	rankFlush := 0
	rankFullHouse := 0
	rankForOfAKind := 0

	for _, v := range hands {
		if v.RankString == "High Card" {
			rankHighCard++
		}
		if v.RankString == "Pair" {
			rankPair++
		}
		if v.RankString == "Two Pair" {
			rankTwoPair++
		}
		if v.RankString == "Three of a Kind" {
			rankThreeOfAKind++
		}
		if v.RankString == "Straight" {
			rankStraight++
		}
		if v.RankString == "Flush" {
			rankFlush++
		}
		if v.RankString == "Full House" {
			rankFullHouse++
		}
		if v.RankString == "Four of a Kind" {
			rankForOfAKind++
		}
		//fmt.Print(v.String() + "  " + v.RankString + "  ")
		//fmt.Println(v.rank)
	}

	fmt.Print("Straight" + "  ")
	fmt.Print("Three of a Kind" + "  ")
	fmt.Print(math.Round(float64(rankThreeOfAKind)*100/float64(len(hands))) / 100)
	fmt.Println("  " + strconv.Itoa(rankThreeOfAKind))
	fmt.Print("Two Pair" + "  ")
	fmt.Print(math.Round(float64(rankTwoPair)*100/float64(len(hands))) / 100)
	fmt.Println("  " + strconv.Itoa(rankTwoPair))
	fmt.Print("Pair" + "  ")
	fmt.Print(math.Round(float64(rankPair)*100/float64(len(hands))) / 100)
	fmt.Println("  " + strconv.Itoa(rankPair))
	fmt.Print("High Card" + "  ")
	fmt.Print(math.Round(float64(rankHighCard)*100/float64(len(hands))) / 100)
	fmt.Println("  " + strconv.Itoa(rankHighCard))
}
