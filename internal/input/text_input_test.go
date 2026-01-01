package input

import (
	"testing"

	"github.com/nsf/termbox-go"
)

// TestChopLast tests the chopLast function which removes the last rune from a string
func TestChopLast(t *testing.T) {
	// Table-driven tests are a Go best practice
	// Each test case has a name, input, and expected output
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "removes last character from simple string",
			input:    "hello",
			expected: "hell",
		},
		{
			name:     "removes last character from single character",
			input:    "a",
			expected: "",
		},
		{
			name:     "empty string returns empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "handles spaces",
			input:    "hello world",
			expected: "hello worl",
		},
		{
			name:     "handles unicode emoji",
			input:    "hello👋",
			expected: "hello",
		},
		{
			name:     "handles unicode characters",
			input:    "café",
			expected: "caf",
		},
		{
			name:     "handles multiple unicode characters",
			input:    "日本語",
			expected: "日本",
		},
	}

	// Run each test case
	for _, tt := range tests {
		// t.Run creates a subtest - makes output clearer
		t.Run(tt.name, func(t *testing.T) {
			result := chopLast(tt.input)
			if result != tt.expected {
				t.Errorf("chopLast(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// TestCursorDefault tests that the default cursor is set
func TestCursorDefault(t *testing.T) {
	// Testing package-level variables
	expected := "█"
	if Cursor != expected {
		t.Errorf("Cursor = %q; want %q", Cursor, expected)
	}
}

// TestGetTextInput_SimpleInput tests basic text input
func TestGetTextInput_SimpleInput(t *testing.T) {
	// Setup mock dependencies
	events := &mockEventSource{
		events: []termbox.Event{
			{Ch: 'h'},
			{Ch: 'e'},
			{Ch: 'l'},
			{Ch: 'l'},
			{Ch: 'o'},
			{Key: termbox.KeyEnter},
		},
	}
	renderer := &mockRenderer{}
	termInfo := &mockTerminalInfo{width: 80}
	spinner := &mockSpinnerRunner{}

	// Execute
	result := GetTextInput(10, 5, "Name: ", events, renderer, termInfo, spinner)

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
		events: []termbox.Event{
			{Ch: 'h'},
			{Ch: 'i'},
			{Key: termbox.KeySpace},
			{Ch: 't'},
			{Ch: 'h'},
			{Ch: 'e'},
			{Ch: 'r'},
			{Ch: 'e'},
			{Key: termbox.KeyEnter},
		},
	}
	renderer := &mockRenderer{}
	termInfo := &mockTerminalInfo{width: 80}
	spinner := &mockSpinnerRunner{}

	result := GetTextInput(0, 0, "", events, renderer, termInfo, spinner)

	expected := "hi there"
	if result != expected {
		t.Errorf("GetTextInput() = %q; want %q", result, expected)
	}
}

// TestGetTextInput_WithBackspace tests backspace functionality
func TestGetTextInput_WithBackspace(t *testing.T) {
	events := &mockEventSource{
		events: []termbox.Event{
			{Ch: 'h'},
			{Ch: 'e'},
			{Ch: 'l'},
			{Ch: 'l'},
			{Ch: 'l'}, // Extra 'l'
			{Ch: 'o'},
			{Key: termbox.KeyBackspace}, // Remove 'o'
			{Key: termbox.KeyBackspace}, // Remove 'l'
			{Ch: 'o'},                   // Add 'o' back
			{Key: termbox.KeyEnter},
		},
	}
	renderer := &mockRenderer{}
	termInfo := &mockTerminalInfo{width: 80}
	spinner := &mockSpinnerRunner{}

	result := GetTextInput(0, 0, "", events, renderer, termInfo, spinner)

	expected := "hello"
	if result != expected {
		t.Errorf("GetTextInput() = %q; want %q", result, expected)
	}
}

// TestGetTextInput_BackspaceOnEmpty tests backspace on empty input
func TestGetTextInput_BackspaceOnEmpty(t *testing.T) {
	events := &mockEventSource{
		events: []termbox.Event{
			{Key: termbox.KeyBackspace}, // Backspace on empty
			{Key: termbox.KeyBackspace}, // Another one
			{Ch: 'a'},
			{Key: termbox.KeyEnter},
		},
	}
	renderer := &mockRenderer{}
	termInfo := &mockTerminalInfo{width: 80}
	spinner := &mockSpinnerRunner{}

	result := GetTextInput(0, 0, "", events, renderer, termInfo, spinner)

	expected := "a"
	if result != expected {
		t.Errorf("GetTextInput() = %q; want %q", result, expected)
	}
}

// TestGetTextInput_EmptyInput tests pressing Enter immediately
func TestGetTextInput_EmptyInput(t *testing.T) {
	events := &mockEventSource{
		events: []termbox.Event{
			{Key: termbox.KeyEnter},
		},
	}
	renderer := &mockRenderer{}
	termInfo := &mockTerminalInfo{width: 80}
	spinner := &mockSpinnerRunner{}

	result := GetTextInput(0, 0, "Prompt: ", events, renderer, termInfo, spinner)

	expected := ""
	if result != expected {
		t.Errorf("GetTextInput() = %q; want %q", result, expected)
	}
}

// TestGetTextInput_UnicodeInput tests unicode character input
func TestGetTextInput_UnicodeInput(t *testing.T) {
	events := &mockEventSource{
		events: []termbox.Event{
			{Ch: '日'},
			{Ch: '本'},
			{Ch: '語'},
			{Key: termbox.KeyEnter},
		},
	}
	renderer := &mockRenderer{}
	termInfo := &mockTerminalInfo{width: 80}
	spinner := &mockSpinnerRunner{}

	result := GetTextInput(0, 0, "", events, renderer, termInfo, spinner)

	expected := "日本語"
	if result != expected {
		t.Errorf("GetTextInput() = %q; want %q", result, expected)
	}
}

// TestGetTextInput_PositionParameters tests that position is passed correctly
func TestGetTextInput_PositionParameters(t *testing.T) {
	events := &mockEventSource{
		events: []termbox.Event{
			{Key: termbox.KeyEnter},
		},
	}
	renderer := &mockRenderer{}
	termInfo := &mockTerminalInfo{width: 80}
	spinner := &mockSpinnerRunner{}

	x, y := 15, 20
	prompt := "Test: "

	GetTextInput(x, y, prompt, events, renderer, termInfo, spinner)

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
