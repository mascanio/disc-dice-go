package roller

import "fmt"

type RollAdder struct {
	Base  Roller
	Adder int
}

type rollAdderRoll struct {
	RollAdder
	roll Roll
}

func (d *rollAdderRoll) RollPrefix() string {
	return ""
}

func (d *rollAdderRoll) Message() string {
	return fmt.Sprintf("%v %v + %v\nTotal: %v\n", d.roll.RollPrefix(), d.Adder, d.roll.RollStr(), d.RollSum())
}

func (d *rollAdderRoll) RollSum() int {
	return d.Adder + d.roll.RollSum()
}

func (d *rollAdderRoll) RollStr() string {
	return fmt.Sprintf("%v + %v", d.Adder, d.roll.RollStr())
}

func (d RollAdder) Roll() Roll {
	roll := d.Base.Roll()
	return &rollAdderRoll{RollAdder: d, roll: roll}
}
