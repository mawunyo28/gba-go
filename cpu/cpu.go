package cpu

import (
	"fmt"

	"github.com/mawunyo28/gba-go/mmu"
)

type Reg16 struct {
	hi, lo uint8
}

func (r *Reg16) SetU16(v uint16) {
	r.hi = uint8(v >> 8)
	r.lo = uint8(v & 0xFF)
}

func (r *Reg16) U16() uint16 {
	return uint16(r.hi)<<8 | uint16(r.lo)
}

type CPU struct {
	AF     Reg16
	BC     Reg16
	DE     Reg16
	HL     Reg16
	SP     uint16
	PC     uint16
	mmu    *mmu.MMU
	halted bool
}

func NewCPU(m *mmu.MMU) *CPU {
	return &CPU{mmu: m}
}

func (c *CPU) Fetch() byte {
	v := c.mmu.Read(c.PC) // fetch byte from memory
	c.PC += 1
	return v
}

func (c *CPU) Step() int {
	opcode := c.Fetch()

	switch opcode {

	case 0x00:
		return 4

	case 0x76:
		c.halted = true
		return 4

	case 0x40: // ld b, b
		c.SetB(c.B())
		return 4

	case 0x41: // ld b, c
		c.SetB(c.C())
		return 4

	case 0x42: // ld b, d
		c.SetB(c.D())
		return 4

	case 0x43: // ld b, e
		c.SetB(c.E())
		return 4

	case 0x44: // ld b, h
		c.SetB(c.H())
		return 4

	case 0x45: // ld b, l
		c.SetB(c.L())
		return 4

	case 0x46: // ld b, [HL]
		c.SetB(c.mmu.Read(c.HL.U16()))
		return 8

	case 0x47: // ld b, a
		c.SetB(c.A())
		return 4

	case 0x48: // ld c, b
		c.SetC(c.B())
		return 4

	case 0x49: // ld c, c
		c.SetC(c.C())
		return 4

	case 0x4a: // ld c, d
		c.SetC(c.D())
		return 4

	case 0x4b: // ld c, e
		c.SetC(c.E())
		return 4

	case 0x4c: // ld c, h
		c.SetC(c.H())
		return 4

	case 0x4d: // ld c, l
		c.SetC(c.L())
		return 4

	case 0x4e: // ld c, [HL]
		c.SetB(c.mmu.Read(c.HL.U16()))
		return 8

	case 0x4f: // ld b, a
		c.SetB(c.A())
		return 4

	default:
		fmt.Sprintf("Unknow opcode %x", opcode)
		return 1

	}

	return 0
}

func (c *CPU) A() uint8 { return c.AF.hi }
func (c *CPU) F() uint8 { return c.AF.lo }
func (c *CPU) B() uint8 { return c.BC.hi }
func (c *CPU) C() uint8 { return c.BC.lo }
func (c *CPU) H() uint8 { return c.HL.hi }
func (c *CPU) L() uint8 { return c.HL.lo }
func (c *CPU) D() uint8 { return c.DE.hi }
func (c *CPU) E() uint8 { return c.DE.lo }

func (c *CPU) SetA(v uint8) { c.AF.hi = v }
func (c *CPU) SetF(v uint8) { c.AF.lo = v }
func (c *CPU) SetB(v uint8) { c.BC.hi = v }
func (c *CPU) SetC(v uint8) { c.BC.lo = v }
func (c *CPU) SetH(v uint8) { c.HL.hi = v }
func (c *CPU) SetL(v uint8) { c.HL.lo = v }
func (c *CPU) SetD(v uint8) { c.DE.hi = v }
func (c *CPU) SetE(v uint8) { c.DE.lo = v }

func (c *CPU) ZFlag() bool { return c.AF.lo&0x80 != 0 }
func (c *CPU) SetZFlag(v bool) {

	if v {

		c.SetF(c.AF.lo | 0x80)

	} else {
		c.SetF(c.AF.lo & ^uint8(0x80))
	}

}

func (c *CPU) HFlag() bool { return c.AF.lo&0x20 != 0 }
func (c *CPU) SetHFlag(v bool) {

	if v {

		c.SetF(c.AF.lo | 0x20)

	} else {
		c.SetF(c.AF.lo & ^uint8(0x20))
	}

}

func (c *CPU) NFlag() bool { return c.AF.lo&0x40 != 0 }
func (c *CPU) SetNFlag(v bool) {

	if v {

		c.SetF(c.AF.lo | 0x40)

	} else {
		c.SetF(c.AF.lo & ^uint8(0x40))
	}

}
func (c *CPU) CFlag() bool { return c.AF.lo&0x10 != 0 }
func (c *CPU) SetCFlag(v bool) {

	if v {

		c.SetF(c.AF.lo | 0x10)

	} else {
		c.SetF(c.AF.lo & ^uint8(0x10))
	}

}
