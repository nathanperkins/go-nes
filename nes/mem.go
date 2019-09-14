package nes

import (
	log "github.com/sirupsen/logrus"
)

// Mem represents an arbitrary block of memory on the BUS.
type Mem struct {
	data     []byte
	name     string
	readOnly bool
}

// NewRAM returns a Mem that is marked as read/write.
func NewRAM(name string, size int) *Mem {
	return &Mem{
		data:     make([]byte, size),
		name:     name,
		readOnly: false,
	}
}

// NewROM returns a Mem that is marked as read only.
func NewROM(name string, size int) *Mem {
	return &Mem{
		data:     make([]byte, size),
		name:     name,
		readOnly: true,
	}
}

// Read performs a read operation at addr.
func (r Mem) Read(addr uint16) byte {
	addr %= uint16(len(r.data))
	return r.data[addr]
}

// Write performs a write operation at addr.
func (r *Mem) Write(addr uint16, val byte) {
	if r.readOnly {
		log.Fatalf("Trying to write to ")
	}
	addr %= uint16(len(r.data))
	r.data[addr] = val
}

// Name returns the name of the device.
func (r Mem) Name() string {
	return r.name
}
