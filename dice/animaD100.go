package dice

import (
	"fmt"
	"math/rand"
	"strings"
)

type AnimaD100 struct {
	StdOpen               bool
	AditionalOpen         bool
	CriticalFailThreshold int
}

type animaD100Roll struct {
	AnimaD100
	result         int
	resultChainStr string
	acuOpen        int
}

func (d *animaD100Roll) Message() string {
	return fmt.Sprintf("Rolling anima 1d100: %v\nTotal: %v\n", d.resultChainStr, d.result)
}

func (d *animaD100Roll) RollSum() int {
	return d.result
}

func (d *animaD100Roll) RollStr() string {
	return d.resultChainStr
}

func (d *animaD100Roll) addResult(result int) {
	d.resultChainStr += fmt.Sprintf("%v", result)
	d.result += result
}

func rOpen(prevRoll, acuOpen int, aditionalOpen bool) (int, []string) {
	if aditionalOpen {
		switch prevRoll {
		case 11:
		case 22:
		case 33:
		case 44:
		case 55:
		case 66:
		case 77:
		case 88:
		case 99:
			secondRoll := rand.Intn(10) + 1
			if secondRoll == prevRoll%10 {
				rSum, rStr := rOpen(rand.Intn(100)+1, acuOpen+1, aditionalOpen)
				// - !!(66,6->100) 192 [100, 92]
				newSum := 100 + rSum
				newStr := fmt.Sprintf("!!(%v,%v->100) %v [%v]", prevRoll, secondRoll, newSum, strings.Join(append([]string{fmt.Sprintf("%v", 100)}, rStr...), ", "))
				return newSum, []string{newStr}
			}
		}
	}
	if prevRoll >= (90+acuOpen) || prevRoll == 100 {
		rSum, rStr := rOpen(rand.Intn(100)+1, acuOpen+1, aditionalOpen)
		newSum := prevRoll + rSum
		// - !192 [100, 92]
		newStr := fmt.Sprintf("!%v [%v]", newSum, strings.Join(append([]string{fmt.Sprintf("%v", prevRoll)}, rStr...), ", "))
		return prevRoll + rSum, []string{newStr}
	}

	return prevRoll, []string{fmt.Sprintf("%v", prevRoll)}
}

func (d *animaD100Roll) roll() {
	result := rand.Intn(100) + 1
	switch {
	case d.result == 0 && result <= d.CriticalFailThreshold:
		failRoll := rand.Intn(100) + 1
		d.result += result - failRoll
		d.resultChainStr += fmt.Sprintf("[F %v]-%v", result, failRoll)
	case d.StdOpen || d.AditionalOpen:
		rSum, rStr := rOpen(result, d.acuOpen, d.AditionalOpen)
		d.result += rSum
		d.resultChainStr += strings.Join(rStr, ", ")
	default:
		d.addResult(result)
	}

}

func (d AnimaD100) Roll() Roll {
	result := animaD100Roll{AnimaD100: d}
	result.roll()
	return &result
}
