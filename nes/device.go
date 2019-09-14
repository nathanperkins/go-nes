package nes

import log "github.com/sirupsen/logrus"

// Device is an I/O device on the bus that can read or write.
type Device interface {
	Read(uint16) byte
	Write(uint16, byte)
	Name() string
}

// DisabledDevice is used to disable ranges of bus mapping.
type DisabledDevice struct{}

// NewDisabledDevice creates a new DisabledDevice.
func NewDisabledDevice() *DisabledDevice {
	return new(DisabledDevice)
}

// Read logs a warning and returns 0.
func (DisabledDevice) Read(addr uint16) byte {
	log.Warnf("Unexpected read on disabled device at address $%x.", addr)
	return 0
}

// Write logs a warning.
func (DisabledDevice) Write(addr uint16, val byte) {
	log.Warnf("Unexpected write on disabled device at $%x: %d.", addr, val)
}

// Name returns the name of this device.
func (DisabledDevice) Name() string {
	return "DisabledDevice"
}
