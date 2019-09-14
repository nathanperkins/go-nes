package nes

import (
	log "github.com/sirupsen/logrus"
)

// Joypad represents that IO operations needed to interact with controllers.
//
// http://wiki.nesdev.com/w/index.php/Input_devices
type Joypad struct{}

// NewJoypad returns an initialized Joypad.
func NewJoypad() *Joypad {
	return new(Joypad)
}

// Read performs a read operation at addr.
func (j Joypad) Read(addr uint16) byte {
	log.Warnf("Joypad not implemented: ($%x).", addr)
	return 0
}

// Write performs a write operation at addr.
func (j *Joypad) Write(addr uint16, val byte) {
	log.Warnf("Joypad not implemented: (val %d @ $%x).", addr, val)
}

// Name returns the device name.
func (j Joypad) Name() string {
	return "Joypad"
}
