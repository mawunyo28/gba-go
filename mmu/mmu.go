package mmu

type MMU struct {
	rom   []byte // read only
	vram  [8192]byte
	wram  [8192]byte
	exram []byte
	oam   [160]byte
	ioreg [128]byte
	hram  [127]byte
	ie    byte
}

func (m *MMU) Read(addr uint16) uint8 {

	switch {
	case addr >= 0xe000 && addr <= 0xfdff: // echo ram
		return m.wram[addr-0xe000]

	case addr >= 0xFEA0 && addr <= 0xFEFF: // prohibited
		return 0xff

	case addr == 0xffff: // Interrupt enable register
		return m.ie

	case addr <= 0x3fff:
		return m.rom[addr-0x0000] // rom bank 1

	case addr >= 0x4000 && addr <= 0x7fff: // rom bank nn
		return m.rom[addr]

	case addr >= 0x8000 && addr <= 0x9fff: // vram
		return m.vram[addr-0x8000]

	case addr >= 0xA000 && addr <= 0xBfff: // external ram
		if m.exram == nil {
			return 0xff
		}
		return m.exram[addr-0xA000]

	case addr >= 0xc000 && addr <= 0xdfff: // work ram
		return m.wram[addr-0xc000]

	case addr >= 0xfe00 && addr <= 0xfe9f: // object attribute memory
		return m.oam[addr-0xfe00]

	case addr >= 0xff00 && addr <= 0xff7f: // io registers
		return m.ioreg[addr-0xff00]

	case addr >= 0xff80 && addr <= 0xfffe: // high ram
		return m.hram[addr-0xff80]

	default:
		return 0xff

	}
}

func (m *MMU) Write(addr uint16, val uint8) {

	switch {
	case addr >= 0xe000 && addr <= 0xfdff: // echo ram

		// do nothing
		m.wram[addr-0xe000] = val

	case addr >= 0xFEA0 && addr <= 0xFEFF: // prohibited

		// do nothing
		// 0xff

	case addr == 0xffff: // Interrupt enable register
		m.ie = val

	case addr <= 0x3fff: // rom bank 1
		// do nothing

	case addr >= 0x4000 && addr <= 0x7fff: // rom bank nn
		// do nothing

	case addr >= 0x8000 && addr <= 0x9fff: // vram
		m.vram[addr-0x8000] = val

	case addr >= 0xA000 && addr <= 0xBfff: // external ram
		if m.exram != nil {
			m.exram[addr-0xA000] = val
		}

	case addr >= 0xc000 && addr <= 0xdfff: // work ram
		m.wram[addr-0xc000] = val

	case addr >= 0xfe00 && addr <= 0xfe9f: // object attribute memory
		m.oam[addr-0xfe00] = val

	case addr >= 0xff00 && addr <= 0xff7f: // io registers
		m.ioreg[addr-0xff00] = val

	case addr >= 0xff80 && addr <= 0xfffe: // high ram
		m.hram[addr-0xff80] = val

	}
}

func NewMMU() *MMU {
	return &MMU{}
}

func (m *MMU) LoadRom(rom []byte) {
	m.rom = rom
}
