package terminal

import "github.com/nsf/termbox-go"

func Init() error {
	err := termbox.Init()
	if err != nil {
		return err
	}

	return nil
}

func Close() {
	termbox.Close()
}

func Flush() {
	termbox.Flush()
}
