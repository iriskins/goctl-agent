package core

import (
	"testing"
)

func TestNewApplication(t *testing.T) {
	App := NewApplication()
	output, err := App.Connection.SystemStateContext(App.Context)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if output == nil {
		t.Errorf("Oops!")
	}
}
