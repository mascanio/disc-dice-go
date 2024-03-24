package userinput

var PRE_IS_DICE_ROLL_MAX_LEN = 20

func validDiceChar(c rune) bool {
	return isNumber(c) || c == 'd' || c == 'a' || c == '+'
}

func isNumber(c rune) bool {
	return c >= '0' && c <= '9'
}

func IsDiceRoll(s string) bool {
	if s == "a" {
		return true
	}
	diceTypeFound := false
	diceFacesFound := false
	plusFound := false
	for i, c := range s {
		switch {
		case i > PRE_IS_DICE_ROLL_MAX_LEN:
			return false
		case !validDiceChar(c):
			return false
		case !diceTypeFound:
			switch c {
			case 'a':
			case 'd':
				diceTypeFound = true
			}
		case diceTypeFound:
			switch c {
			case '+':
				if plusFound {
					return false
				} else if i == len(s)-1 {
					return false
				} else if !diceFacesFound {
					return false
				}
				plusFound = true
			case 'a':
			case 'd':
				return false
			default:
				diceFacesFound = true
			}
		}
	}
	return diceTypeFound && diceFacesFound
}
