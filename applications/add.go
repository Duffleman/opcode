package applications

import (
	"fmt"

	"opcode"
)

var Add = &AddApp{}

type AddApp struct{}

func (a *AddApp) Opcode() int {
	return 1
}

func (a *AddApp) Exec(os *opcode.OS, c *opcode.OPCode, cursor int) (*int, error) {
	p1 := os.Memory.GetAt(cursor+1, c.Param1Mode)
	p2 := os.Memory.GetAt(cursor+2, c.Param2Mode)
	ptr := os.Memory.GetIndex(cursor+3, c.Param3Mode)

	val := p1 + p2

	if os.Debug {
		fmt.Printf("%02d (add): %d + %d = %d -> %d\n", c.Code, p1, p2, val, ptr)
		fmt.Printf("\t%d was %d, now %d\n", ptr, os.Memory.GetIndex(ptr, opcode.PositionMode), val)
	}

	os.Memory.Set(ptr, val)

	return opcode.IntP(cursor + 4), nil
}
