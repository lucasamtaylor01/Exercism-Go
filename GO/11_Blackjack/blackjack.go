package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch {
	case card == "ace":
		return 11
	case card == "two":
		return 2
	case card == "three":
		return 3
	case card == "four":
		return 4
	case card == "five":
		return 5
	case card == "six":
		return 6
	case card == "seven":
		return 7
	case card == "eight":
		return 8
	case card == "nine":
		return 9
	case card == "ten" || card == "jack" || card == "king" || card == "queen":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	//Hands classification
	hands := ParseCard(card1) + ParseCard(card2)
	dealer := ParseCard(dealerCard)
	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"

	case hands == 21 && dealerCard != "jack" && dealerCard != "queen" && dealerCard != "king" && dealerCard != "ace":
		return "W"

	case hands <= 11:
		return "H"

	case hands >= 12 && hands <= 16 && dealer >= 7:
		return "H"

	case hands == 21 && dealer == 10 || dealer == 11:
		return "S"

	case hands >= 17 && hands <= 20:
		return "S"

	case hands >= 12 && hands <= 16 && dealer < 7:
		return "S"

	default:
		return "H"

	}
}
