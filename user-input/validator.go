package userinput

var PRE_IS_DICE_ROLL_MAX_LEN = 100

func validDiceChar(c rune) bool {
	return c >= '0' && c <= '9' || c == 'd'
}

func IsDiceRoll(s string) bool {
	dFound := false
	diceTypeFound := false
	for i, c := range s {
		switch {
		case i > PRE_IS_DICE_ROLL_MAX_LEN:
			return false
		case !validDiceChar(c):
			return false
		case !dFound:
			switch {
			case c == 'd':
				dFound = true
			}
		case dFound:
			if c == 'd' {
				return false
			}
			diceTypeFound = true
		}
	}
	return dFound && diceTypeFound
}
