package nes

import (
	"fmt"
	"io"
	"os"

	"github.com/google/go-cmp/cmp"
	log "github.com/sirupsen/logrus"
)

// Cart is based on the iNES format spec from
// https://wiki.nesdev.com/w/index.php/INES.
type Cart struct {
	Header cartHeader
	// Trainer, if present (0 or 512 bytes)
	Trainer []byte
	// PRG ROM data (16384 * x bytes)
	PRG *Mem
	// CHR ROM data, if present (8192 * y bytes)
	CHR *Mem
}

// NewCart creates a new Cart based on the given file.
func NewCart(r io.Reader) (*Cart, error) {
	newCart := new(Cart)
	log.Infof("Reading header (%d bytes).", 16)
	newCart.Header = make([]byte, 16)
	if _, err := r.Read(newCart.Header); err != nil {
		return nil, fmt.Errorf("could not read header: %v", err)
	}

	// Read Trainer
	if newCart.Header.trainerPresent() {
		log.Infof("Reading trainer (%d bytes).", trainerSize)
		newCart.Trainer = make([]byte, trainerSize)
		if n, err := r.Read(newCart.Trainer); err != nil {
			return nil, fmt.Errorf("could not read trainer: %v", err)
		} else if n != trainerSize {
			return nil, fmt.Errorf("only read %d bytes of trainer, wanted %d", n, trainerSize)
		}
	}

	// Read PRG
	prgSize := newCart.Header.prgSize()
	log.Infof("Reading PRG ROM (%d bytes).", prgSize)
	newCart.PRG = NewROM("PRG ROM", prgSize)
	if n, err := r.Read(newCart.PRG.data); err != nil {
		return nil, fmt.Errorf("could not read PRG: %v", err)
	} else if n != prgSize {
		return nil, fmt.Errorf("only read %d bytes of PRG, wanted %d", n, prgSize)
	}

	// Read CHR
	chrSize := newCart.Header.chrSize()
	log.Infof("Reading CHR ROM (%d bytes).", chrSize)
	newCart.CHR = NewROM("CHR ROM", chrSize)
	if n, err := r.Read(newCart.CHR.data); err != nil {
		return nil, fmt.Errorf("could not read CHR: %v", err)
	} else if n != chrSize {
		return nil, fmt.Errorf("only read %d bytes of CHR, wanted %d", n, chrSize)
	}

	if err := newCart.Validate(); err != nil {
		return nil, err
	}
	return newCart, nil
}

// OpenCart opens the file and runs NewCart().
func OpenCart(filename string) (*Cart, error) {
	log.Infof("Opening %q.", filename)
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return NewCart(f)
}

// Validate checks that the cart is valid.
func (r Cart) Validate() error {
	if diff := cmp.Diff(r.Header[0:4], cartHeader{0x4E, 0x45, 0x53, 0x1A}); diff != "" {
		return fmt.Errorf("invalid ROM: does not start with NES: %v", diff)
	}
	return nil
}

const (
	trainerSize = 512
)

type cartHeaderFlag6 int

const (
	horizontalMirror cartHeaderFlag6 = 1 << iota
	verticalMirror
	containsBatteryPRGRAM
	trainerPresent
	ignoreMirror
)

// cartHeader is the 16 byte header at the beginning of a .nes file.
// 0-3: Constant $4E $45 $53 $1A ("NES" followed by MS-DOS end-of-file)
// 4: Size of PRG ROM in 16 KB units
// 5: Size of CHR ROM in 8 KB units (Value 0 means the board uses CHR RAM)
// 6: Flags 6 - Mapper, mirroring, battery, trainer
// 7: Flags 7 - Mapper, VS/Playchoice, NES 2.0
// 8: Flags 8 - PRG-RAM size (rarely used extension)
// 9: Flags 9 - TV system (rarely used extension)
// 10: Flags 10 - TV system, PRG-RAM presence (unofficial, rarely used extension)
// 11-15: Unused padding (should be filled with zero, but some rippers put their name across bytes 7-15)
type cartHeader []byte

func (h cartHeader) prgSize() int {
	return 16384 * int(h[4])
}

func (h cartHeader) chrSize() int {
	return 8192 * int(h[5])
}

func (h cartHeader) trainerPresent() bool {
	return h[6]&(1<<4) != 0
}
