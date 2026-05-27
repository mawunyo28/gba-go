package mmu

import (
	"testing"
)

func TestWriteWRAM(t *testing.T) {
	addr := uint16(0xc0af)
	var value uint8 = 0xe4

	memory := NewMMU()

	memory.Write(addr, value)

	// read

	var value1 uint8 = memory.Read(addr)

	if value != value1 {

		t.Errorf("wrote %x but read %x", value, value1)
	}

}

func TestWriteROM(t *testing.T) {
	var addr uint16 = 0x0000

	memory := NewMMU()

	memory.LoadRom([]byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9})

	value := memory.Read(addr)

	memory.Write(addr, 0xe2)

	value1 := memory.Read(addr)

	if value1 != value {
		t.Errorf("rom was mutated")
	}
}

func TestReadFromUnstable(t *testing.T) {
	var addr uint16 = 0xfea0

	memory := NewMMU()

	if memory.Read(addr) != 0xff {
		t.Errorf("unwanted return")
	}
}

func TestWriteToEchoReadFromWRAM(t *testing.T) {
	var echoAddr uint16 = 0xfdff
	var readAddr uint16 = 0xdfff
	memory := NewMMU()

	if memory.Read(echoAddr) != memory.Read(readAddr) {

		t.Errorf("echo mapping is wrong")

	}

}
