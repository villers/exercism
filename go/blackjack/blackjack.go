package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int
	switch card {
	case "ace":
		value = 11
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "ten", "jack", "queen", "king":
		value = 10
	default:
		value = 0
	}

	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	valueC1 := ParseCard(card1)
	valueC2 := ParseCard(card2)
	valueD := ParseCard(dealerCard)

	total := valueC1 + valueC2

	var result string
	switch {

	case valueC1 == 11 && valueC2 == 11:
		result = "P"

	case
		total >= 12 && total <= 16 && valueD >= 7, total <= 11:
		result = "H"

	case total == 21 && dealerCard != "ace" && dealerCard != "ten" && valueD <= 9:
		result = "W"

	default:
		result = "S"
	}

	return result
}
