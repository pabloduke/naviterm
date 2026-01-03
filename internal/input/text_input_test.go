package input_test

import (
	"testing"

	"github.com/pabloduke/naviterm/internal/input"
	"github.com/pabloduke/naviterm/internal/terminal"
)

// TestCursorDefault tests that the default cursor is set
func TestCursorDefault(t *testing.T) {
	// Testing the exported Cursor variable
	expected := "█"
	if input.Cursor != expected {
		t.Errorf("Cursor = %q; want %q", input.Cursor, expected)
	}
}

// TestGetTextInput_SimpleInput tests basic text input
func TestGetTextInput_SimpleInput(t *testing.T) {
	// Setup mock dependencies
	events := &mockEventSource{
		events: []terminal.Event{
			{Ch: 'h'},
			{Ch: 'e'},
			{Ch: 'l'},
			{Ch: 'l'},
			{Ch: 'o'},
			{Key: terminal.KeyEnter},
		},
	}
	renderer := &mockRenderer{}
	termInfo := &mockTerminalInfo{width: 80}
	spinner := &mockSpinnerRunner{}

	// Execute
	result := input.GetTextInput(10, 5, "Name: ", events, renderer, termInfo, spinner)

	// Assert
	if result != "hello" {
		t.Errorf("GetTextInput() = %q; want %q", result, "hello")
	}

	// Verify spinner was started
	if !spinner.started {
		t.Error("Expected spinner to be started")
	}

	// Verify spinner done channel was closed
	select {
	case <-spinner.doneChan:
		// Channel was closed, good!
	default:
		t.Error("Expected done channel to be closed")
	}

	// Verify flush was called
	if renderer.flushes < 1 {
		t.Errorf("Expected at least 1 flush, got %d", renderer.flushes)
	}
}

// TestGetTextInput_WithSpaces tests input containing spaces
func TestGetTextInput_WithSpaces(t *testing.T) {
	events := &mockEventSource{
		events: []terminal.Event{
			{Ch: 'h'},
			{Ch: 'i'},
			{Key: terminal.KeySpace},
			{Ch: 't'},
			{Ch: 'h'},
			{Ch: 'e'},
			{Ch: 'r'},
			{Ch: 'e'},
			{Key: terminal.KeyEnter},
		},
	}
	renderer := &mockRenderer{}
	termInfo := &mockTerminalInfo{width: 80}
	spinner := &mockSpinnerRunner{}

	result := input.GetTextInput(0, 0, "", events, renderer, termInfo, spinner)

	expected := "hi there"
	if result != expected {
		t.Errorf("GetTextInput() = %q; want %q", result, expected)
	}
}

// TestGetTextInput_WithBackspace tests backspace functionality
func TestGetTextInput_WithBackspace(t *testing.T) {
	events := &mockEventSource{
		events: []terminal.Event{
			{Ch: 'h'},
			{Ch: 'e'},
			{Ch: 'l'},
			{Ch: 'l'},
			{Ch: 'l'}, // Extra 'l'
			{Ch: 'o'},
			{Key: terminal.KeyBackspace}, // Remove 'o'
			{Key: terminal.KeyBackspace}, // Remove 'l'
			{Ch: 'o'},                    // Add 'o' back
			{Key: terminal.KeyEnter},
		},
	}
	renderer := &mockRenderer{}
	termInfo := &mockTerminalInfo{width: 80}
	spinner := &mockSpinnerRunner{}

	result := input.GetTextInput(0, 0, "", events, renderer, termInfo, spinner)

	expected := "hello"
	if result != expected {
		t.Errorf("GetTextInput() = %q; want %q", result, expected)
	}
}

// TestGetTextInput_BackspaceOnEmpty tests backspace on empty input
func TestGetTextInput_BackspaceOnEmpty(t *testing.T) {
	events := &mockEventSource{
		events: []terminal.Event{
			{Key: terminal.KeyBackspace}, // Backspace on empty
			{Key: terminal.KeyBackspace}, // Another one
			{Ch: 'a'},
			{Key: terminal.KeyEnter},
		},
	}
	renderer := &mockRenderer{}
	termInfo := &mockTerminalInfo{width: 80}
	spinner := &mockSpinnerRunner{}

	result := input.GetTextInput(0, 0, "", events, renderer, termInfo, spinner)

	expected := "a"
	if result != expected {
		t.Errorf("GetTextInput() = %q; want %q", result, expected)
	}
}

// TestGetTextInput_EmptyInput tests pressing Enter immediately
func TestGetTextInput_EmptyInput(t *testing.T) {
	events := &mockEventSource{
		events: []terminal.Event{
			{Key: terminal.KeyEnter},
		},
	}
	renderer := &mockRenderer{}
	termInfo := &mockTerminalInfo{width: 80}
	spinner := &mockSpinnerRunner{}

	result := input.GetTextInput(0, 0, "Prompt: ", events, renderer, termInfo, spinner)

	expected := ""
	if result != expected {
		t.Errorf("GetTextInput() = %q; want %q", result, expected)
	}
}

// TestGetTextInput_UnicodeInput tests unicode character input
func TestGetTextInput_UnicodeInput(t *testing.T) {
	events := &mockEventSource{
		events: []terminal.Event{
			{Ch: '日'},
			{Ch: '本'},
			{Ch: '語'},
			{Key: terminal.KeyEnter},
		},
	}
	renderer := &mockRenderer{}
	termInfo := &mockTerminalInfo{width: 80}
	spinner := &mockSpinnerRunner{}

	result := input.GetTextInput(0, 0, "", events, renderer, termInfo, spinner)

	expected := "日本語"
	if result != expected {
		t.Errorf("GetTextInput() = %q; want %q", result, expected)
	}
}

// TestGetTextInput_PositionParameters tests that position is passed correctly
func TestGetTextInput_PositionParameters(t *testing.T) {
	events := &mockEventSource{
		events: []terminal.Event{
			{Key: terminal.KeyEnter},
		},
	}
	renderer := &mockRenderer{}
	termInfo := &mockTerminalInfo{width: 80}
	spinner := &mockSpinnerRunner{}

	x, y := 15, 20
	prompt := "Test: "

	input.GetTextInput(x, y, prompt, events, renderer, termInfo, spinner)

	// Verify spinner position (should be x-1, y)
	if spinner.lastX != x-1 || spinner.lastY != y {
		t.Errorf("Spinner position = (%d, %d); want (%d, %d)",
			spinner.lastX, spinner.lastY, x-1, y)
	}

	// Verify initial draw call includes prompt
	if len(renderer.drawCalls) == 0 {
		t.Fatal("Expected at least one draw call")
	}
	firstCall := renderer.drawCalls[0]
	if firstCall.x != x || firstCall.y != y {
		t.Errorf("First draw at (%d, %d); want (%d, %d)",
			firstCall.x, firstCall.y, x, y)
	}
}
