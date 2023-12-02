package handrange

import (
	"github.com/chehsunliu/poker"
)

type Range struct {
	matrix *[13][13]int8
}

func NewRange(r *[13][13]int8) *Range {
	brange := &Range{}
	// rangeMatrix := [13][13]int8{
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// }
	brange.matrix = r
	return brange
}

type Hand struct {
	Fcard      poker.Card
	Scard      poker.Card
	Rank       int32
	RankString string
}

func NewHand(f poker.Card, s poker.Card) Hand {
	hand := Hand{}
	hand.Fcard = f
	hand.Scard = s
	return hand
}

func (h *Hand) String() string {
	return h.Fcard.String() + " " + h.Scard.String()
}

func (r *Range) GetComboHands() []*Hand {
	var combos []*Hand
	strRanks := map[int]string{
		0:  "A",
		1:  "K",
		2:  "Q",
		3:  "J",
		4:  "T",
		5:  "9",
		6:  "8",
		7:  "7",
		8:  "6",
		9:  "5",
		10: "4",
		11: "3",
		12: "2",
	}
	for i := 0; i < 13; i++ {
		for j := 0; j < 13; j++ {
			if r.matrix[i][j] == 1 {
				// Off suits
				if i > j {
					combos = append(combos, r.GetOffSuitHands(strRanks[i], strRanks[j])...)
				}
				// Pair
				if i == j {
					combos = append(combos, r.GetPairHands(strRanks[i])...)
				}
				// Suits
				if i < j {
					combos = append(combos, r.GetSuitHands(strRanks[i], strRanks[j])...)
				}
			}
		}
	}
	return combos
}

func (r *Range) GetOffSuitHands(s1 string, s2 string) []*Hand {
	var returnHand []*Hand
	suits := map[int]string{0: "s", 1: "h", 2: "d", 3: "c"}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i < j {
				shHand := NewHand(poker.NewCard(s1+suits[i]), poker.NewCard(s2+suits[j]))
				shHand2 := NewHand(poker.NewCard(s1+suits[j]), poker.NewCard(s2+suits[i]))
				hands := []*Hand{&shHand, &shHand2}
				returnHand = append(returnHand, hands...)
			}
		}
	}
	return returnHand
}

func (r *Range) GetSuitHands(s1 string, s2 string) []*Hand {
	var returnHand []*Hand
	suits := map[int]string{0: "s", 1: "h", 2: "d", 3: "c"}

	for i := 0; i < 4; i++ {
		shHand := NewHand(poker.NewCard(s1+suits[i]), poker.NewCard(s2+suits[i]))
		hands := []*Hand{&shHand}
		returnHand = append(returnHand, hands...)
	}

	return returnHand
}

func (r *Range) GetPairHands(s1 string) []*Hand {
	var returnHand []*Hand
	suits := map[int]string{0: "s", 1: "h", 2: "d", 3: "c"}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i < j {
				shHand := NewHand(poker.NewCard(s1+suits[i]), poker.NewCard(s1+suits[j]))
				hands := []*Hand{&shHand}
				returnHand = append(returnHand, hands...)
			}
		}
	}

	return returnHand
}
