package blackjack
//Concepts - Switch statemetn in golang

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var ans int
    switch card {
        case "ace" :
    		ans = 11
        case "two" :
    		ans = 2
        case "three" :
    		ans = 3
        case "four" :
   			ans = 4
        case "five" :
    		ans = 5
        case "six" :
   			ans = 6
        case "seven" :
   			ans = 7
        case "eight" :
    		ans = 8
        case "nine" :
   			ans = 9
        case "ten","jack","queen","king" :
   			ans = 10
        default :
    		ans = 0
    }
	return ans
}


// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	return (21 == ParseCard(card1)+ParseCard(card2))
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	if isBlackjack {
    	if dealerScore < 10 {
            return "W"
        }else{
        	return "S"
        }
    }else{
    	return "P"
    }
}        

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	switch  {
		case handScore >=17 :
    		return "S"
        case handScore <=11 :
    		return "H"
        case handScore >=12 && handScore <= 17 && dealerScore <=6 :
    		return "S"
        case handScore >=12 && handScore <= 17 && dealerScore >=6 :
    		return "H"
    }
	return " "
}