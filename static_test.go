package log

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLog will keep last used data
type TestLog struct {
	lastMsg   string
	lastErr   error
	verbosity int
	lastKV    []KV
}

// Info will store parameters internally
func (l *TestLog) Info(msg string, kv ...KV) {
	l.lastMsg = msg
	l.lastKV = kv
}

// Error will store parameters internally
func (l *TestLog) Error(err error, msg string, kv ...KV) {
	l.lastErr = err
	l.lastMsg = msg
	l.lastKV = kv
}

// V will return a different logger
func (l *TestLog) V(int) InfoWriter {
	return &TestLog{}
}

// SetLevel will store verbosity internally
func (l *TestLog) SetLevel(level int) {
	l.verbosity = level
}

var _ Logger = (*TestLog)(nil)

func TestDefaultLogger(t *testing.T) {
	nl := &TestLog{}
	SetDefaultLogger(nl)
	l, err := GetDefaultLogger()
	if err != nil {
		t.Errorf("error returned %v expected nil", err)
	}

	if nl != l {
		t.Errorf("logger returned %v expected %v", l, nl)
	}
}

func TestNonDefaultLoggers(t *testing.T) {

	l1 := &TestLog{}
	var testData = []struct {
		loggerSet Logger
		nameSet   string
		nameGet   string
		loggerRet Logger
		errorRet  bool
	}{
		{
			loggerSet: l1,
			nameSet:   "log1",
			nameGet:   "log1",
			loggerRet: l1,
			errorRet:  false,
		},
		{
			loggerSet: l1,
			nameSet:   "log1",
			nameGet:   "log2",
			loggerRet: nil,
			errorRet:  true,
		},
	}

	for _, td := range testData {
		SetLogger(td.nameSet, td.loggerSet)
		l, err := GetLogger(td.nameGet)
		if (err != nil) != td.errorRet {
			t.Errorf("error returned %v expected %v", err, td.errorRet)
		}

		if td.loggerRet != l {
			t.Errorf("logger returned %v expected %v", l, td.loggerRet)
		}
	}
}

func TestDefaultStaticOperations(t *testing.T) {
	nl := &TestLog{verbosity: 1}
	SetDefaultLogger(nl)

	testMsg := "message1"
	testKV := []KV{KV{"a", "1"}, KV{"b", "2"}}
	Info(testMsg, testKV...)

	assert.Equal(t, testMsg, nl.lastMsg)
	assert.Equal(t, testKV, nl.lastKV)

	testMsg = "message2"
	testKV = []KV{KV{"d", "3"}, KV{"e", "4"}}
	testError := errors.New("test error")
	Error(testError, testMsg, testKV...)

	assert.Equal(t, testMsg, nl.lastMsg)
	assert.Equal(t, testKV, nl.lastKV)
	assert.Equal(t, testError, nl.lastErr)

	testVerbosity := 10
	SetLevel(testVerbosity)
	assert.Equal(t, testVerbosity, nl.verbosity)

	newInfoLogger := V(testVerbosity)
	assert.NotEqual(t, newInfoLogger, nl)

}
