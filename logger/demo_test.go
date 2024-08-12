package logger_test

import (
	"fmt"
	"github.com/pwinning1991/testingWithGo/logger"
	"log"
	"strings"
	"testing"
)

type fakeLogger struct {
	sb strings.Builder
}

func (fl fakeLogger) Println(v ...interface{}) {
	fmt.Fprintln(&fl.sb, v...)
}

func (fl fakeLogger) Printf(format string, v ...interface{}) {
	fmt.Fprintf(&fl.sb, format, v...)
}

func TestDemoV3(t *testing.T) {
	var fl fakeLogger
	logger.DemoV3(fl.Println)
	want := "error in dotheThing():"
	got := fl.sb.String()
	if !strings.Contains(got, want) {
		t.Errorf("Logs = %q, want %q", got, want)
	}
}

func TestDemoV4(t *testing.T) {
	var fl fakeLogger
	logger.DemoV4(fl)
	want := "error in dotheThing():"
	got := fl.sb.String()
	if !strings.Contains(got, want) {
		t.Errorf("Logs = %q, want %q", got, want)
	}
}

func TestDemoV7(t *testing.T) {
	var sb strings.Builder
	testLogger := log.New(&sb, "", 0)
	thing := logger.Thing2{
		Logger: testLogger,
	}
	thing.DemoV7()
	want := "error in dotheThing():"
	got := sb.String()
	if !strings.Contains(got, want) {
		t.Errorf("Logs = %q, want %q", got, want)
	}
}
