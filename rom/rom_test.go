package rom_test

import (
	"testing"

	"github.com/nathanperkins/go-nes/rom"
)

func TestMustStartWithNES(t *testing.T) {
	_, err := rom.Open("testdata/bad.nes")
	if err == nil {
		t.Errorf("Open: should give error.")
	}
	_, err = rom.Open("testdata/nestest.nes")
	if err != nil {
		t.Errorf("Open: should not give error: %v", err)
	}
}
