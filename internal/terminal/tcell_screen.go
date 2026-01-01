package terminal

import (
	"github.com/gdamore/tcell/v3"
)

// TcellScreen implements Screen using tcell v3
type TcellScreen struct {
	screen tcell.Screen
}

// NewTcellScreen creates a new tcell screen adapter
func NewTcellScreen() (*TcellScreen, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	return &TcellScreen{screen: screen}, nil
}

func (s *TcellScreen) Init() error {
	if err := s.screen.Init(); err != nil {
		return err
	}
	s.screen.Clear()
	return nil
}

func (s *TcellScreen) Close() {
	s.screen.Fini()
}

func (s *TcellScreen) Clear() {
	s.screen.Clear()
}

func (s *TcellScreen) Flush() {
	s.screen.Show()
}

func (s *TcellScreen) Size() (int, int) {
	return s.screen.Size()
}

func (s *TcellScreen) SetCell(x, y int, ch rune, fg, bg Attribute) {
	style := tcell.StyleDefault.
		Foreground(convertToTcellColor(fg)).
		Background(convertToTcellColor(bg))
	s.screen.SetContent(x, y, ch, nil, style)
}

func (s *TcellScreen) PollEvent() Event {
	// In tcell v3, events come from a channel
	tcellEvent := <-s.screen.EventQ()
	if tcellEvent == nil {
		return Event{Type: EventError}
	}
	return convertTcellEvent(tcellEvent)
}

// convertToTcellColor maps terminal.Attribute to tcell.Color
func convertToTcellColor(attr Attribute) tcell.Color {
	switch attr {
	case 0: // ColorDefault
		return tcell.ColorDefault
	case 1: // ColorBlack
		return tcell.ColorBlack
	case 2: // ColorRed
		return tcell.ColorRed
	case 3: // ColorGreen
		return tcell.ColorGreen
	case 4: // ColorYellow
		return tcell.ColorYellow
	case 5: // ColorBlue
		return tcell.ColorBlue
	case 6: // ColorMagenta
		return tcell.ColorMaroon
	case 7: // ColorCyan
		return tcell.ColorTeal
	case 8: // ColorWhite
		return tcell.ColorWhite
	case 14: // ColorLightCyan
		return tcell.ColorAqua
	default:
		return tcell.ColorDefault
	}
}

// convertTcellEvent converts tcell events to terminal.Event
func convertTcellEvent(ev tcell.Event) Event {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		// In tcell v3, Str() returns the string representation
		// Convert to rune (take first character if present)
		var ch rune
		str := ev.Str()
		if len(str) > 0 {
			ch = []rune(str)[0]
		}
		return Event{
			Type: EventKey,
			Key:  convertTcellKey(ev.Key()),
			Ch:   ch,
			Mod:  convertTcellMod(ev.Modifiers()),
		}
	case *tcell.EventResize:
		return Event{
			Type: EventResize,
		}
	default:
		return Event{
			Type: EventError,
		}
	}
}

// convertTcellKey converts tcell key constants to terminal.Key
func convertTcellKey(key tcell.Key) Key {
	switch key {
	case tcell.KeyEnter:
		return KeyEnter
	case tcell.KeyBackspace, tcell.KeyBackspace2:
		return KeyBackspace
	case tcell.KeyUp:
		return KeyArrowUp
	case tcell.KeyDown:
		return KeyArrowDown
	case tcell.KeyLeft:
		return KeyArrowLeft
	case tcell.KeyRight:
		return KeyArrowRight
	case tcell.KeyEscape:
		return KeyEsc
	case tcell.KeyTab:
		return KeyTab
	case tcell.KeyRune:
		// Character is in the Ch field, check for space
		return KeyNone
	default:
		return KeyNone
	}
}

// convertTcellMod converts tcell modifiers to terminal.Modifier
func convertTcellMod(mod tcell.ModMask) Modifier {
	var result Modifier
	if mod&tcell.ModAlt != 0 {
		result |= ModAlt
	}
	if mod&tcell.ModCtrl != 0 {
		result |= ModCtrl
	}
	if mod&tcell.ModShift != 0 {
		result |= ModShift
	}
	return result
}
