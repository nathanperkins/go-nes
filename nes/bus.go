package nes

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Bus handles the mapping of memory addresses to devices.
type Bus struct {
	mappings []BusMapping
}

// BusMapping represents one mapping of an address space to a device.
// low and high represent the min and max address range, inclusive.
// http://wiki.nesdev.com/w/index.php/CPU_memory_map
type BusMapping struct {
	device    Device
	low, high uint16
}

// NewBus creates a new Bus.
func NewBus() *Bus {
	return new(Bus)
}

// Read polls the proper device to respond with a byte at the proper address.
func (b Bus) Read(addr uint16) byte {
	for _, m := range b.mappings {
		if m.low <= addr && addr <= m.high {
			return m.device.Read(addr - m.low)
		}
	}
	log.Errorf("Bus Read() error: no devices on addr %x", addr)
	return 0
}

// Write writes a byte on the proper device at the proper address.
func (b *Bus) Write(addr uint16, val byte) {
	for _, m := range b.mappings {
		if m.low <= addr && addr <= m.high {
			m.device.Write(addr-m.low, val)
			return
		}
	}
	log.Errorf("Bus Write() err: no devices on addr %x", addr)
}

// AddDevice creates a mapping for the device and given range.
func (b *Bus) AddDevice(d Device, low, high uint16) error {
	for _, m := range b.mappings {
		if m.low <= low && low <= m.high || m.low <= high && high <= m.high {
			return fmt.Errorf("cannot map device: %x,%x overlaps with device %s", low, high, m.device.Name())
		}
	}
	log.Infof("Adding %s at address range $%x to $%x.", d.Name(), low, high)
	b.mappings = append(b.mappings, BusMapping{
		device: d,
		low:    low,
		high:   high,
	})
	return nil
}
