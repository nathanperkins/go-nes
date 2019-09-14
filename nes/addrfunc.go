package nes

// AddrFunc is used to handle fetching the operand for execution.
type AddrFunc func(*CPU) (uint16, byte)

// AddrImplied occurs when there is no operand.
// http://www.thealmightyguru.com/Games/Hacking/Wiki/index.php?title=Implied_Addressing
func AddrImplied(c *CPU) (addr uint16, b byte) {
	return
}

// AddrImmediate is used when the operand's value is given in the instruction itself.
// http://www.thealmightyguru.com/Games/Hacking/Wiki/index.php?title=Immediate_Addressing
func AddrImmediate(c *CPU) (addr uint16, b byte) {
	b = c.bus.Read(c.pc)
	c.pc++
	return
}

// AddrAbs is used when the operand is a 2-byte memory address to be read directly.
// http://www.thealmightyguru.com/Games/Hacking/Wiki/index.php?title=Absolute_Addressing
func AddrAbs(c *CPU) (addr uint16, b byte) {
	addr = uint16(c.bus.Read(c.pc))
	c.pc++
	addr |= uint16(c.bus.Read(c.pc)) << 8
	c.pc++
	b = c.bus.Read(addr)
	return
}
