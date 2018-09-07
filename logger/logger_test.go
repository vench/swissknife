package logger

import (
	"os"
	"testing"
	"time"
)

//
func TestLogger(t *testing.T) {

	fileName := `logger.test.log`

	logger := NewLoggerWorker(fileName)
	logger.Start()

	logger.Log(`Test log message`)

	time.Sleep(time.Second)

	f, err := os.Open(fileName)
	if err != nil {
		t.Error(err)
	}

	var buf = make([]byte, 128)

	read, err := f.Read(buf)
	if err != nil {
		t.Error(err)
	}

	if len(buf) == 0 || read == 0 {
		t.Error("File is empty")
	}

}
