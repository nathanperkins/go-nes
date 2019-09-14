package nes

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nathanperkins/go-nes/nes/fake"
)

func NewTestCPU() *CPU {
	return &CPU{
		bus: &Bus{
			mappings: []BusMapping{
				BusMapping{
					device: &fake.Device{},
					low:    0x0000,
					high:   0xFFFF,
				},
			},
		},
	}
}

func TestAddrImplied(t *testing.T) {
	tests := []struct {
		name string
		addr uint16
	}{
		{
			"Zero",
			0x0000,
		},
		{
			"Max",
			0xFFFF,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := NewTestCPU()
			cpu.pc = test.addr
			device := cpu.bus.mappings[0].device.(*fake.Device)
			device.ReadBytes = []byte{0xFF}
			gotAddr, gotByte := AddrImplied(cpu)
			if len(device.ReadAddrs) > 0 {
				t.Errorf("AddrImplied() should not read from device.")
			}
			if gotAddr != 0 {
				t.Errorf("AddrImplied(%#x) got addr %#x, want %#x", test.addr, gotAddr, 0)
			}
			if gotByte != 0 {
				t.Errorf("AddrImplied(%#x) got byte %#x, want %#x", test.addr, gotByte, 0)
			}
		})
	}
}

func TestAddrImmediate(t *testing.T) {
	tests := []struct {
		name string
		addr uint16
		want byte
	}{
		{
			"Zero",
			0x0000,
			0x0000,
		},
		{
			"Max",
			0xFFFF,
			0x00FF,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := NewTestCPU()
			cpu.pc = test.addr
			device := cpu.bus.mappings[0].device.(*fake.Device)
			device.ReadBytes = []byte{test.want}
			gotAddr, gotByte := AddrImmediate(cpu)
			if len(device.ReadAddrs) != 1 {
				t.Errorf("AddrImplied() should read once, got %v", device.ReadAddrs)
			}
			if gotAddr != 0 {
				t.Errorf("AddrImplied(%#x) got addr %#x, want %#x", test.addr, gotAddr, test.addr)
			}
			if gotByte != test.want {
				t.Errorf("AddrImplied(%#x) got byte %#x, want %#x", test.addr, gotByte, test.want)
			}
			if diff := cmp.Diff(device.ReadAddrs, []uint16{test.addr}); diff != "" {
				t.Errorf("Device addrs read diff (-got, + want):\n%s", diff)
			}
		})
	}
}
