package log

import (
	"testing"
)

func TestDefaultLogger(t *testing.T) {
	nl := &NoLog{}
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

	l1 := &NoLog{}
	// l2 := &NoLog{}
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
