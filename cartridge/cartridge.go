package cartridge

import "errors"

// should be about 80 bytes
type Cartridge struct {
	title         [16]byte
	cartridgeType byte
	romSize       byte
	ramSize       byte
	checksum      byte
}

func (c *Cartridge) ValidateChecksum(rom []byte) bool {
	var computed byte

	for address := 0x0134; address <= 0x014c; address++ {
		computed = computed - rom[address] - 1
	}

	return computed == c.checksum

}

func New(rom []byte) (Cartridge, error) {
	var err error

	c := Cartridge{
		cartridgeType: rom[0x0147],
		romSize:       rom[0x0148],
		ramSize:       rom[0x0149],
		checksum:      rom[0x014d],
	}

	copy(c.title[:], rom[0x0134:0x0143])

	if !c.ValidateChecksum(rom) {
		err = errors.New("Unable to validate checksum.\nFile possibly corrupted")
	}

	return c, err
}
