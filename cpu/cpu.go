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
		c.SetC(c.mmu.Read(c.HL.U16()))
		return 8

	case 0x4f: // ld c, a
		c.SetC(c.A())
		return 4

	case 0x50: // ld d, b
		c.SetD(c.B())
		return 4

	case 0x51: // ld d, c
		c.SetD(c.C())
		return 4

	case 0x52: // ld d, d
		c.SetD(c.D())
		return 4

	case 0x53: // ld d, e
		c.SetD(c.E())
		return 4

	case 0x54: // ld d, h
		c.SetD(c.H())
		return 4

	case 0x55: // ld d, l
		c.SetD(c.L())
		return 4

	case 0x56: // ld d, [HL]
		c.SetD(c.mmu.Read(c.HL.U16()))
		return 8

	case 0x57: // ld d, a
		c.SetD(c.A())
		return 4

	case 0x58: // ld e, b
		c.SetE(c.B())
		return 4

	case 0x59: // ld e, c
		c.SetE(c.C())
		return 4

	case 0x5a: // ld e, d
		c.SetE(c.D())
		return 4

	case 0x5b: // ld e, e
		c.SetE(c.E())
		return 4

	case 0x5c: // ld e, h
		c.SetE(c.H())
		return 4

	case 0x5d: // ld e, l
		c.SetE(c.L())
		return 4

	case 0x5e: // ld e, [HL]
		c.SetE(c.mmu.Read(c.HL.U16()))
		return 8

	case 0x5f: // ld e, a
		c.SetE(c.A())
		return 4

	case 0x60: // ld h, b
		c.SetH(c.B())
		return 4

	case 0x61: // ld h, c
		c.SetH(c.C())
		return 4

	case 0x62: // ld h, d
		c.SetH(c.D())
		return 4

	case 0x63: // ld h, e
		c.SetH(c.E())
		return 4

	case 0x64: // ld h, h
		c.SetH(c.H())
		return 4

	case 0x65: // ld h, l
		c.SetH(c.L())
		return 4

	case 0x66: // ld h, [HL]
		c.SetH(c.mmu.Read(c.HL.U16()))
		return 8

	case 0x67: // ld h, a
		c.SetH(c.A())
		return 4

	case 0x68: // ld l, b
		c.SetL(c.B())
		return 4

	case 0x69: // ld l, c
		c.SetL(c.C())
		return 4

	case 0x6a: // ld l, d
		c.SetL(c.D())
		return 4

	case 0x6b: // ld l, e
		c.SetL(c.E())
		return 4

	case 0x6c: // ld l, h
		c.SetL(c.H())
		return 4

	case 0x6d: // ld l, l
		c.SetL(c.L())
		return 4

	case 0x6e: // ld L, [HL]
		c.SetL(c.mmu.Read(c.HL.U16()))
		return 8

	case 0x6f: // ld L, a
		c.SetL(c.A())
		return 4

	case 0x70: // ld (hl), b
		c.mmu.Write(c.HL.U16(), c.B())
		return 8

	case 0x71: // ld (hl), c
		c.mmu.Write(c.HL.U16(), c.C())
		return 8

	case 0x72: // ld (hl), d
		c.mmu.Write(c.HL.U16(), c.D())
		return 8

	case 0x73: // ld (hl), e
		c.mmu.Write(c.HL.U16(), c.E())
		return 8

	case 0x74: // ld (hl), h
		c.mmu.Write(c.HL.U16(), c.H())
		return 8

	case 0x75: // ld (hl), l
		c.mmu.Write(c.HL.U16(), c.L())
		return 8

	case 0x77: // ld (hl), a
		c.mmu.Write(c.HL.U16(), c.A())
		return 8

	case 0x78: // ld a, b
		c.SetA(c.B())
		return 4
	case 0x79: // ld a, c
		c.SetA(c.C())
		return 4
	case 0x7a: // ld a, d
		c.SetA(c.D())
		return 4
	case 0x7b: // ld a, e
		c.SetA(c.E())
		return 4

	case 0x7c: // ld a, h
		c.SetA(c.H())
		return 4
	case 0x7d: // ld a, l
		c.SetA(c.L())
		return 4
	case 0x7e: // ld a, (hl)
		c.SetA(c.mmu.Read(c.HL.U16()))
		return 8

	case 0x7f: // ld a, a
		c.SetA(c.A())
		return 4

	case 0x01: // ld BC, uint16

		lo := c.Fetch()
		hi := c.Fetch()

		c.SetC(lo)
		c.SetB(hi)
		return 12

	case 0x11: // ld DE, uint16

		lo := c.Fetch()
		hi := c.Fetch()

		c.SetE(lo)
		c.SetD(hi)
		return 12

	case 0x21: // ld HL, uint16

		lo := c.Fetch()
		hi := c.Fetch()
		c.SetL(lo)
		c.SetH(hi)
		return 12
	case 0x31: // ld SP, uint16
		lo := c.Fetch()
		hi := c.Fetch()

		c.SP = uint16(hi)<<8 | uint16(lo)

		return 12

	case 0x06: // ld B, uint8
		c.SetB(c.Fetch())
		return 8
	case 0x16: // ld D, uint8
		c.SetD(c.Fetch())
		return 8
	case 0x26: // ld H, uint8
		c.SetH(c.Fetch())
		return 8
	case 0x36: // ld (HL), uint8
		c.mmu.Write(c.HL.U16(), c.Fetch())
		return 12

	case 0x0E: // ld C, uint8
		c.SetC(c.Fetch())
		return 8
	case 0x1E: // ld E, uint8
		c.SetE(c.Fetch())
		return 8
	case 0x2E: // ld L, uint8
		c.SetL(c.Fetch())
		return 8
	case 0x3E: // ld A, uint8
		c.SetA(c.Fetch())
		return 8

	case 0x02: // ld (bc), a
		c.mmu.Write(c.BC.U16(), c.A())
		return 8

	case 0x12: // ld (DE), a
		c.mmu.Write(c.DE.U16(), c.A())
		return 8

	case 0x0A: // ld A, (BC)
		c.SetA(c.mmu.Read(c.BC.U16()))
		return 8

	case 0x1A: // ld A, (DE)
		c.SetA(c.mmu.Read(c.DE.U16()))
		return 8

	case 0xEA: // ld (uint16), A

		lo := c.Fetch()
		hi := c.Fetch()

		addr := uint16(hi)<<8 | uint16(lo)

		c.mmu.Write(addr, c.A())

		return 16

	case 0xFA: // ld A , (uint16)

		lo := c.Fetch()
		hi := c.Fetch()

		addr := uint16(hi)<<8 | uint16(lo)

		c.SetA(c.mmu.Read(addr))

		return 16

	case 0x22: // ld [HL+], A

		c.mmu.Write(c.HL.U16(), c.A())
		c.HL.SetU16(c.HL.U16() + 1)

		return 8

	case 0x32: // ld [HL-], A
		c.mmu.Write(c.HL.U16(), c.A())

		c.HL.SetU16(c.HL.U16() - 1)

		return 8

	case 0x2A: //ld A, [HL+]
		c.SetA(c.mmu.Read(c.HL.U16()))
		c.HL.SetU16(c.HL.U16() + 1)
		return 8

	case 0x3A: // ld A, [HL-]

		c.SetA(c.mmu.Read(c.HL.U16()))
		c.HL.SetU16(c.HL.U16() - 1)

		return 8

	case 0xE0: // ld (FF00+u8), A
		c.mmu.Write(uint16(0xff00)+uint16(c.Fetch()), c.A())
		return 12

	case 0xF0: // ld A, (FF00+u8)
		c.SetA(c.mmu.Read(uint16(0xff00) + uint16(c.Fetch())))
		return 12

	case 0xE2: // ld (FF00+C), A
		c.mmu.Write(uint16(0xff00)+uint16(c.C()), c.A())
		return 8

	case 0xF2: // ld A, (FF00+c)
		c.SetA(c.mmu.Read(uint16(0xff00) + uint16(c.C())))
		return 8

	case 0xF9: // ld SP, HL
		c.SP = c.HL.U16()
		return 8

	case 0xF8: // ld HL, SP + i8
		e := int8(c.Fetch())

		offset := uint16(int16(c.SP) + int16(e))

		c.HL.SetU16(offset)

		spLo := c.SP & 0xFF
		eLo := uint16(uint8(e))

		c.SetHFlag((c.SP&0x0F)+(eLo&0x0F) > 0x0F)
		c.SetCFlag(spLo+eLo > 0xFF)
		c.SetZFlag(false)
		c.SetNFlag(false)

		return 12

	case 0x08: // ld (u16), SP
		lo := c.Fetch()
		hi := c.Fetch()

		addr := uint16(hi)<<8 | uint16(lo)
		c.mmu.Write(addr, uint8(c.SP))
		c.mmu.Write(addr+1, uint8(c.SP>>8))

		return 20

	default:
		panic(fmt.Sprintf("Unknown opcode 0x%02x at 0x%04x", opcode, c.PC-1))

	}
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
