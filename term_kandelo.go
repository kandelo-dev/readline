package readline

import "syscall"

type Termios = syscall.Termios

func getTermios(fd int) (*Termios, error) {
	return syscall.GetTermios(fd)
}

func setTermios(fd int, termios *Termios) error {
	return syscall.SetTermios(fd, termios)
}

func GetSize(fd int) (int, int, error) {
	winsize, err := syscall.GetWinsize(fd)
	if err != nil {
		return 0, 0, err
	}
	return int(winsize.Col), int(winsize.Row), nil
}
