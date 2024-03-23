package dice

type DiceResult struct {
	Sum    int
	Result string
}

type Dice interface {
	Roll() DiceResult
}
