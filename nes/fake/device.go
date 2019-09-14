package fake

// Device is a device intended for testing.
type Device struct {
	// TODO: Change this device so that we can actually store bytes to be read at specific addresses.

	// A queue of bytes provided by the tester to be read in order.
	ReadBytes []byte

	// A list of bytes written to the device.
	WriteBytes []byte

	// A list of the addresses that have been read from the device.
	ReadAddrs []uint16

	// A list of addresses that have been written on the device.
	WriteAddrs []uint16
}

// Read returns the first available byte in ReadBytes and records the addr request.
func (d *Device) Read(addr uint16) byte {
	b := d.ReadBytes[0]
	d.ReadBytes = d.ReadBytes[1:]
	d.ReadAddrs = append(d.ReadAddrs, addr)
	return b
}

// Write records the byte and the addr.
func (d *Device) Write(addr uint16, b byte) {
	d.WriteBytes = append(d.WriteBytes, b)
	d.WriteAddrs = append(d.WriteAddrs, addr)
}

// Name returns the name of the device.
func (d Device) Name() string {
	return "FakeDevice"
}
