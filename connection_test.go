package dali

import (
	"testing"
)

func TestNewDLCP(t *testing.T) {
	dlconn, err := NewDLCP("localhost", "test")
	defer FreeDLCP(dlconn)
	if err != nil {
		t.Error("unable to create new DLCP")
	}
}

func TestDLTerminate(t *testing.T) {
	dlconn, err := NewDLCP("localhost", "test")
	defer FreeDLCP(dlconn)
	if err != nil {
		t.Error("unable to create new DLCP")
	}
	dlconn.Terminate()
}
