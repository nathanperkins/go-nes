package nes

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// NES represents an abstraction of the NES hardware itself.
type NES struct {
	cpu  *CPU
	bus  *Bus
	cart *Cart
}

// NewNES creates a new NES.
func NewNES(filename string) (*NES, error) {
	nes := new(NES)
	nes.bus = NewBus()
	nes.cpu = NewCPU(nes.bus)

	// Read the ROM into a cart
	cart, err := OpenCart(filename)
	if err != nil {
		return nil, err
	}
	nes.cart = cart

	// Create Bus devices and add mappings.
	nes.bus.AddDevice(NewRAM("CPU RAM", 0x2000), 0x0000, 0x1FFF)
	nes.bus.AddDevice(NewPPU(), 0x2000, 0x3FFF)
	nes.bus.AddDevice(NewAPU(), 0x4000, 0x4015)
	nes.bus.AddDevice(NewJoypad(), 0x4016, 0x4017)
	nes.bus.AddDevice(NewDisabledDevice(), 0x4018, 0x401F)
	nes.bus.AddDevice(nes.cart.PRG, 0x8000, 0xFFFF)

	nes.cpu.PowerOn()
	log.Infof("%s", nes.cpu)
	return nes, nil
}

// Run starts the execution loop.
func (n *NES) Run() error {
	for {
		if err := n.cpu.Tick(); err != nil {
			return fmt.Errorf("NES run error: %v", err)
		}
	}
}
