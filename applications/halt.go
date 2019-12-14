package applications

import (
	"opcode"
)

var Halt = &HaltApp{}

type HaltApp struct{}

func (a *HaltApp) Opcode() int {
	return 99
}

func (a *HaltApp) Exec(os opcode.OS, _ *opcode.OPCode, cursor int) (*int, error) {
	return opcode.IntP(cursor + 1), opcode.ErrHalt
}
