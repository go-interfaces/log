package log

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"
)

func TestNoLog(t *testing.T) {
	nl := NoLog{}

	stdout := os.Stdout
	stderr := os.Stderr
	rOut, wOut, e := os.Pipe()
	if e != nil {
		t.Errorf("could not pipe stream: %v", e)
	}
	rErr, wErr, e := os.Pipe()
	if e != nil {
		t.Errorf("could not pipe stream: %v", e)
	}
	os.Stdout = wOut
	os.Stderr = wErr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()

	nl.Error(
		errors.New("test error"),
		"message for test error",
		KV{"a", "1"}, KV{"b", "2"},
	)
	wOut.Close()
	wErr.Close()

	var bufOut bytes.Buffer
	io.Copy(&bufOut, rOut)
	if bufOut.Len() != 0 {
		t.Errorf("NoLog.Error returned %q at stdout when empty string was expected", bufOut.String())
	}
	var bufErr bytes.Buffer
	io.Copy(&bufErr, rErr)
	if bufErr.Len() != 0 {
		t.Errorf("NoLog.Error returned %q at stderr when empty string was expected", bufErr.String())
	}

}
