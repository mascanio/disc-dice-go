package roller

import (
	"fmt"
)

type Ability struct {
	Name   string
	Base   int
	Roller Roller
}

type abilityRoll struct {
	Ability
	Roll
}

func (d *abilityRoll) RollPrefix() string {
	return fmt.Sprintf("Rolling %v (%v): ", d.Name, d.Base)
}

func (d *abilityRoll) Message() string {
	return fmt.Sprintf("%v%v\nTotal: %v\n", d.RollPrefix(), d.RollStr(), d.RollSum())
}

func (d *abilityRoll) RollSum() int {
	return d.Base + d.Roll.RollSum()
}

func (d *abilityRoll) RollStr() string {
	return fmt.Sprintf("%v + %v", d.Base, d.Roll.RollStr())
}

func (d Ability) Roll() Roll {
	baseRoll := d.Roller.Roll()
	roll := &abilityRoll{Ability: d, Roll: baseRoll}
	return roll
}
