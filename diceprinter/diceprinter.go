package diceprinter

import (
	"fmt"
	"github.com/AndFran/dice"
)

func RollAndPrint(sides int) {
	fmt.Println(dice.Roll(sides))
}
