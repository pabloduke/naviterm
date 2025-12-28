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

func Width() int {
	width, _ := termbox.Size()
	return width
}

func Height() int {
	_, height := termbox.Size()
	return height
}
