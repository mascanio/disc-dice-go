package parser

import (
	"fmt"
	"strconv"

	"github.com/mascanio/disc-dice-go/internal/roller"
	nre "github.com/mascanio/regexp-named"
)

var re_dice = nre.MustCompile(`(?P<n>\d+)?d(?P<d>\d+)`)

var max_n_dices = 100
var max_faces_dice = 1000

func getNDicesAndFaces(s string) (int, int, error) {
	_, mm := re_dice.FindStringNamed(s)
	if mm == nil {
		return 0, 0, fmt.Errorf("no dice found")
	}
	nDices, err := strconv.Atoi(mm["n"])
	if err != nil {
		nDices = 1
	}
	nFaces, err := strconv.Atoi(mm["d"])
	if err != nil {
		return 0, 0, fmt.Errorf("error converting d to int")
	}
	if nDices > max_n_dices || nFaces > max_faces_dice {
		return 0, 0, fmt.Errorf("too many dices or incorrect dice faces")
	}
	if nFaces < 2 {
		return 0, 0, fmt.Errorf("dice faces must be at least 2")
	}
	return nDices, nFaces, nil
}

func ParseDice(s string) roller.Roller {
	nDices, nFaces, err := getNDicesAndFaces(s)
	if err != nil {
		return nil
	}
	switch nDices {
	case 1:
		return roller.GenericDice{Faces: nFaces}
	default:
		return roller.MultiDice{Faces: nFaces, Dices: nDices}
	}
}
