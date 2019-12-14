package applications

import (
	"fmt"

	"opcode"
)

var AdjustRelativeBase = &ARBApp{}

type ARBApp struct{}

func (a *ARBApp) Opcode() int {
	return 9
}

func (a *ARBApp) Exec(os *opcode.OS, c *opcode.OPCode, cursor int) (*int, error) {
	param := os.Memory.GetAt(cursor+1, c.Param1Mode)

	if os.Debug {
		fmt.Printf("%02d (arb): %d (crb) + %d\n", c.Code, os.Memory.RelativeBase, param)
	}

	os.Memory.RelativeBase += param

	if os.Debug {
		fmt.Printf("%02d (arb): relativeBase now %d\n", c.Code, os.Memory.RelativeBase)
	}

	return opcode.IntP(cursor + 2), nil
}
