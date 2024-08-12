package logger

import (
	"errors"
	"log"
	"os"
	"sync"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)

func DemoGloabl() {
	err := dotheTing()
	if err != nil {
		logger.Println("error in doTheThing():", err)
	}
}

func DemoV1() {
	err := dotheTing()
	if err != nil {
		logger.Println("error in doTheThing():", err)
	}
}

func DemoV2(logger *log.Logger) {
	err := dotheTing()
	if err != nil {
		logger.Println("error in doTheThing():", err)
	}
}

func DemoV3(logFn func(...interface{})) {
	err := dotheTing()
	if err != nil {
		logFn("error in doTheThing():", err)
	}

}

type Logger interface {
	Println(...interface{})
	Printf(string, ...interface{})
}

func DemoV4(logger Logger) {
	err := dotheTing()
	if err != nil {
		logger.Println("error in doTheThing():", err)
		logger.Printf("error: %s\n", err)
	}
}

type Thing struct {
	Logger interface {
		Println(...interface{})
		Printf(string, ...interface{})
	}
}

func (t Thing) DemoV5() {
	err := dotheTing()
	if err != nil {
		t.Logger.Println("error in doTheThing():", err)
		t.Logger.Printf("error: %s\n", err)
	}
}

func DemoV6(logger Logger) {
	if logger == nil {
		logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	}
	err := dotheTing()
	if err != nil {
		logger.Println("error in doTheThing():", err)
		logger.Printf("error: %s\n", err)
	}
}

type Thing2 struct {
	Logger interface {
		Println(...interface{})
		Printf(string, ...interface{})
	}
	once sync.Once
}

func (t *Thing2) logger() Logger {
	t.once.Do(func() {
		if t.Logger == nil {
			t.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
		}
	})
	return t.Logger

}

func (t *Thing2) DemoV7() {
	err := dotheTing()
	if err != nil {
		t.Logger.Println("error in doTheThing():", err)
		t.Logger.Printf("error: %s\n", err)
	}
}

func dotheTing() error {
	return errors.New("error opening file: abc.txt")
}
