package readline

import (
	"os"
	"testing"
)

func TestKandeloTerminal(t *testing.T) {
	terminal, err := os.OpenFile("/dev/ptmx", os.O_RDWR, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer terminal.Close()
	fd := int(terminal.Fd())
	if !IsTerminal(fd) {
		t.Fatal("PTY master is not a terminal")
	}
	state, err := MakeRaw(fd)
	if err != nil {
		t.Fatal(err)
	}
	if err := Restore(fd, state); err != nil {
		t.Fatal(err)
	}
	if _, _, err := GetSize(fd); err != nil {
		t.Fatal(err)
	}
}
