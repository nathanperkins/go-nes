package nes

type execFunc func(c *CPU, addr uint16, b byte)

// execSEI (Set Interrupt Disable Flag) sets the Interrupt Flag in the Processor Status Register by setting the 2nd bit 1.
// http://www.thealmightyguru.com/Games/Hacking/Wiki/index.php?title=SEI
func execSEI(c *CPU, addr uint16, b byte) {
	c.p |= cpuStatusInterruptDisable
}

// execCLD (Clear Decimal Flag) clears the Decimal Flag in the Processor Status Register by setting the 3rd bit 0.
// http://www.thealmightyguru.com/Games/Hacking/Wiki/index.php?title=CLD
func execCLD(c *CPU, addr uint16, b byte) {
	c.p &^= cpuStatusDecimalMode
}

// execLDA (Load Accumulator With Memory) loads the accumulator with specified memory.
// http://www.thealmightyguru.com/Games/Hacking/Wiki/index.php?title=LDA
func execLDA(c *CPU, addr uint16, b byte) {
	if b == 0 {
		c.p |= cpuStatusZero
	} else {
		c.p &^= cpuStatusZero
	}
	if b > 0x7F {
		c.p |= cpuStatusNegative
	} else {
		c.p &^= cpuStatusNegative
	}
	c.a = b
}

// LDX (Load X Index With Memory) loads the X Index Register with the specified memory.
// http://www.thealmightyguru.com/Games/Hacking/Wiki/index.php?title=LDX
func execLDX(c *CPU, addr uint16, b byte) {
	if b == 0 {
		c.p |= cpuStatusZero
	} else {
		c.p &^= cpuStatusZero
	}
	if b > 0x7F {
		c.p |= cpuStatusNegative
	} else {
		c.p &^= cpuStatusNegative
	}
	c.x = b
}

// LDY (Load Y Index With Memory) loads the Y Index Register with the specified memory.
// http://www.thealmightyguru.com/Games/Hacking/Wiki/index.php?title=LDY
func execLDY(c *CPU, addr uint16, b byte) {
	if b == 0 {
		c.p |= cpuStatusZero
	} else {
		c.p &^= cpuStatusZero
	}
	if b > 0x7F {
		c.p |= cpuStatusNegative
	} else {
		c.p &^= cpuStatusNegative
	}
	c.y = b
}

// execSTA (Store Accumulator In Memory) stores the accumulator into a specified memory address.
// http://www.thealmightyguru.com/Games/Hacking/Wiki/index.php?title=STA
func execSTA(c *CPU, addr uint16, b byte) {
	c.bus.Write(addr, c.a)
}
