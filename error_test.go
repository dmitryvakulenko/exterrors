package exterrors

import (
	"errors"
	"fmt"
	"testing"
)

func TestUnwrapThis(t *testing.T) {
	e1 := New("test", nil)
	e2 := fmt.Errorf("test2: %w", e1)

	unwrapped := errors.Unwrap(e2)
	if unwrapped != e1 {
		t.Error("Unwrap this error")
	}
}

func TestWrapUnwrap(t *testing.T) {
	e1 := fmt.Errorf("test")
	e2 := Wrap(e1, "test2", nil)

	unwrapped := errors.Unwrap(e2)
	if unwrapped != e1 {
		t.Error("Unwrap error")
	}
}

func TestIsAs(t *testing.T) {
	e1 := New("test", nil)

	if !errors.Is(e1, e1) {
		t.Error("Is error")
	}

	e2 := New("test2", nil)
	if !errors.As(e1, &e2) {
		t.Error("As error")
	}
}
