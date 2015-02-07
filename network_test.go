package dali

import (
	"testing"
)

func TestSenddata(t *testing.T) {
	var buffer string = "A test string"

	dlconn, err := NewDLCP(":16000", "test")
	defer FreeDLCP(dlconn)
	if err != nil {
		t.Error("unable to create new DLCP")
		return
	}

	_, err2 := dlconn.Connect()
	if err2 != nil {
		t.Error("unable to connect to local server")
	}

	_, err3 := dlconn.Senddata(([]byte)(buffer), (uint)(len(buffer)))
	if err3 != nil {
		t.Error("unable to send data")
		return
	}

	dlconn.Terminate()
}

func TestRecvdata(t *testing.T) {
	var buffer []byte

	dlconn, err := NewDLCP(":16000", "test")
	defer FreeDLCP(dlconn)
	if err != nil {
		t.Error("unable to create new DLCP")
		return
	}

	_, err2 := dlconn.Connect()
	if err2 != nil {
		t.Error("unable to connect to local server")
	}

	_, err3 := dlconn.Recvdata(buffer, (uint)(len(buffer)), true)
	if err3 != nil {
		t.Error("unable to recv data")
		return
	}

	dlconn.Terminate()
}
