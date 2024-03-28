package userinput

var PRE_IS_DICE_ROLL_MAX_LEN = 20

func validDiceChar(c rune) bool {
	return isNumber(c) || c == 'd' || c == 'a' || c == '+' || c == '-'
}

func isNumber(c rune) bool {
	return c >= '0' && c <= '9'
}

func IsDiceRoll(s string) bool {
	if s == "a" {
		return true
	}
	diceType := 'n'
	diceFacesFound := false
	plusFound := false
	for i, c := range s {
		switch {
		case i > PRE_IS_DICE_ROLL_MAX_LEN:
			return false
		case !validDiceChar(c):
			return false
		case diceType == 'n':
			switch c {
			case 'a', 'd':
				diceType = c
			}
		case diceType != 'n':
			switch c {
			case '+', '-':
				if plusFound {
					return false
				} else if i == len(s)-1 {
					return false
				} else if !diceFacesFound && diceType == 'd' {
					return false
				}
				plusFound = true
			case 'a', 'd':
				return false
			default:
				if diceType == 'd' {
					diceFacesFound = true
				}
			}
		}
	}
	return diceType != 'n' && (diceFacesFound || diceType == 'a')
}
