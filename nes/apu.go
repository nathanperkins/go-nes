package nes

import (
	log "github.com/sirupsen/logrus"
)

// APU is the audio processing unit in the NES console which generates sound for games.
//
// http://wiki.nesdev.com/w/index.php/APU
type APU struct{}

// NewAPU returns an initialized APU.
func NewAPU() *APU {
	return new(APU)
}

// Read performs a read operation at addr.
func (a APU) Read(addr uint16) byte {
	log.Warnf("APU Read not implemented: ($%x).", addr)
	return 0
}

// Write performs a write operation at addr.
func (a *APU) Write(addr uint16, val byte) {
	log.Warnf("APU Write not implemented: (val %d @ $%x).", addr, val)
}

// Name returns the name of the device.
func (a APU) Name() string {
	return "APU"
}
