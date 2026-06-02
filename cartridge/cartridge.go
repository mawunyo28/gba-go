package cartridge

import "errors"

// should be about 80 bytes
type Cartridge struct {
	Title         [16]byte
	CartridgeType byte
	RomSize       byte
	RamSize       byte
	Checksum      byte
}

func (c *Cartridge) ValidateChecksum(rom []byte) bool {
	var computed byte

	for address := 0x0134; address <= 0x014c; address++ {
		computed = computed - rom[address] - 1
	}

	return computed == c.Checksum

}

func New(rom []byte) (Cartridge, error) {
	var err error

	c := Cartridge{
		CartridgeType: rom[0x0147],
		RomSize:       rom[0x0148],
		RamSize:       rom[0x0149],
		Checksum:      rom[0x014d],
	}

	copy(c.Title[:], rom[0x0134:0x0143])

	if !c.ValidateChecksum(rom) {
		err = errors.New("Unable to validate checksum.\nFile possibly corrupted")
	}

	return c, err
}
