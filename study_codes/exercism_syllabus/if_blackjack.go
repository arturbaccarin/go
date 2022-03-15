package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var valueCard int
	switch card {
	case "ace":
		valueCard = 11
	case "two":
		valueCard = 2
	case "three":
		valueCard = 3
	case "four":
		valueCard = 4
	case "five":
		valueCard = 5
	case "six":
		valueCard = 6
	case "seven":
		valueCard = 7
	case "eight":
		valueCard = 8
	case "nine":
		valueCard = 9
	case "ten":
		valueCard = 10
	case "jack":
		valueCard = 10
	case "queen":
		valueCard = 10
	case "king":
		valueCard = 10
	default:
		valueCard = 0
	}
	return valueCard
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	var valueCard1 int = ParseCard(card1)
	var valueCard2 int = ParseCard(card2)

	if (valueCard1 + valueCard2) == 21 {
		return true
	} else {
		return false
	}
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	if isBlackjack {
		if dealerScore < 10 {
			return "W"
		}
		return "S"
	}
	return "P"
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	if handScore >= 17 {
		return "S"
	} else if handScore <= 11 {
		return "H"
	} else {
		if dealerScore >= 7 {
			return "S"
		} else {
			return "H"
		}
	}
}

/*
// FirstTrun ...
func FirstTurn(card1 string, card2 string, dealer string) string {
    var valueCard1 int = ParseCard(card1)
    var valueCard2 int = ParseCard(card2)
    var dealerCard int = ParseCard(dealer)

    var handScore int = valueCard1 + valueCard2

    var smallHand string = SmallHand(handScore, dealerCard)

	var messageOut string = smallHand + ", want S"

    return messageOut
}
*/
