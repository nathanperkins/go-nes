package nes

import (
	log "github.com/sirupsen/logrus"
)

// PPU handles graphic processing for the NES.
//
// http://wiki.nesdev.com/w/index.php/PPU
type PPU struct{}

// NewPPU returns an initialized PPU.
func NewPPU() *PPU {
	return new(PPU)
}

// Read performs a read operation at addr.
func (p PPU) Read(addr uint16) byte {
	log.Warnf("PPU Read not implemented: ($%x).", addr)
	return 0
}

// Write performs a write operation at addr.
func (p *PPU) Write(addr uint16, val byte) {
	log.Warnf("PPU Write not implemented: (val %d @ $%x).", addr, val)
}

// Name returns the name of the device.
func (p PPU) Name() string {
	return "PPU"
}
