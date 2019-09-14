package nes

import (
	"fmt"
	"log"
)

// CPUStatus stores the CPU status flags and provides their operations.
// http://wiki.nesdev.com/w/index.php/Status_flags
type cpuStatus byte

const (
	cpuStatusCarry cpuStatus = 1 << iota
	cpuStatusZero
	cpuStatusInterruptDisable
	cpuStatusDecimalMode
	cpuStatusB1
	cpuStatusB2
	cpuStatusOverflow
	cpuStatusNegative
)

// CPU holds processor state and processes instructions.
//
// Registers:
//   a: accumulator
//   x: x
//   y: y
//   pc: program counter
//   s: stack pointer
//   p: status
//
// http://wiki.nesdev.com/w/index.php/CPU
type CPU struct {
	// TODO: These should probably be exported so that we can see them in debuggers.
	a     byte
	x     byte
	y     byte
	pc    uint16
	s     byte
	p     cpuStatus
	bus   *Bus
	cycle int
}

// NewCPU creates a new CPU with the given bus.
func NewCPU(b *Bus) *CPU {
	return &CPU{bus: b}
}

func (c CPU) String() string {
	return fmt.Sprintf(`CPU Status:
c:  %d
a:  %#04x
x:  %#04x
y:  %#04x
pc: %#04x
s:  %#04x
p:  %#04x
`, c.cycle, c.a, c.x, c.y, c.pc, c.s, c.p)
}

// PowerOn sets the CPU state after turning on the power.
// http://wiki.nesdev.com/w/index.php/CPU_power_up_state#At_power-up
func (c *CPU) PowerOn() {
	c.a, c.x, c.y = 0, 0, 0
	c.p = 0x34 // (IRQ disabled)
	c.s = 0xFD
	c.bus.Write(0x4015, 0x00) // (all channels disabled)
	c.bus.Write(0x4017, 0x00) // (frame irq enabled)
	c.resetProgramCounter()
}

// Reset sets the CPU after hitting the reset button.
// http://wiki.nesdev.com/w/index.php/CPU_power_up_state#After_reset
func (c *CPU) Reset() {
	c.s -= 3
	c.p |= cpuStatusInterruptDisable
	c.bus.Write(0x4015, 0x00)
	c.resetProgramCounter()
}

func (c *CPU) resetProgramCounter() {
	c.pc = uint16(c.bus.Read(0xFFFD))<<8 | uint16(c.bus.Read(0xFFFC))
}

// Tick runs through one CPU cycle.
func (c *CPU) Tick() error {
	// TODO: Track cycles and do nothing if the CPU is not ready.
	code := c.bus.Read(c.pc)
	c.pc++
	if err := c.Do(code); err != nil {
		return fmt.Errorf("opcode %02x: %v", code, err)
	}
	log.Println(c)
	return nil
}

// Do executes one operation given by the code byte.
func (c *CPU) Do(code byte) error {
	op := opCodes[code]
	if op.name == "" {
		return fmt.Errorf("not implemented")
	}
	log.Printf("Running %s $%2x", op.name, code)
	addr, b := op.addrFunc(c)
	op.execFunc(c, addr, b)
	if op.cycles == 0 {
		return fmt.Errorf("no cycle count")
	}
	c.cycle += op.cycles
	return nil
}
