package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const (
	DEFAULT_LOG_BUFFER = 200
	TIME_FORMAT        = "2006-01-02T15:04:05"
)

//
type LoggerWorker struct {
	FileLog    string
	LogVerbose bool
	logChannel chan string
	LogBuffer  uint
	running    bool
	// line
	LogSize       uint
	CallBackFatal *func(error)
}

//
func NewLoggerWorker(filename string) *LoggerWorker {

	lw := LoggerWorker{}
	lw.FileLog = filename
	lw.LogVerbose = true
	return &lw
}

//
func (lw *LoggerWorker) Start() {
	if lw.LogBuffer == 0 {
		lw.LogBuffer = DEFAULT_LOG_BUFFER
	}

	lw.logChannel = make(chan string, lw.LogBuffer)
	lw.running = true
	go lw.processInfo()
}

//
func (lw *LoggerWorker) Stop() {
	if lw.running {
		lw.running = false
		close(lw.logChannel)
	}
}

//
func (lw *LoggerWorker) Log(message string, sttr ...string) {
	lw.logChannel <- fmt.Sprintf("[%s] - %s %s\n", lw.currentTime(), message, strings.Join(sttr, " "))
}

//
func (lw *LoggerWorker) LogPlain(message string) {
	lw.logChannel <- fmt.Sprintf("%s\n", message)
}

//
func (lw *LoggerWorker) Clear() {

	lw.Stop()

	if file, err := os.OpenFile(lw.FileLog, os.O_WRONLY, 0); err == nil {
		defer file.Close()
		file.Truncate(0)
		file.Seek(0, 0)
	}
}

//
func (lw *LoggerWorker) callBackFatal(err error) {
	if lw.CallBackFatal != nil {
		(*lw.CallBackFatal)(err)
	}
}

//
func (lw *LoggerWorker) openFile() *os.File {
	logfile, err := os.OpenFile(lw.FileLog, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if os.IsNotExist(err) {
		logfile, err = os.Create(lw.FileLog)
	}
	if err != nil {
		lw.callBackFatal(err)
		log.Println("Error open file", lw.FileLog)
		panic(err)
	}

	return logfile
}

// process info-log writer
func (lw *LoggerWorker) processInfo() {

	logfile := lw.openFile()

	var lines uint = 0

	for lw.running {
		textLog := <-lw.logChannel
		if lw.LogVerbose {
			log.Println(textLog)
		}
		logfile.WriteString(textLog)
		lines++

		if lw.LogSize > 0 && lines > lw.LogSize {

			if err := logfile.Close(); err != nil {
				lw.callBackFatal(err)
				panic(err)
			}
			newFile := lw.getNewFileName()
			os.Rename(lw.FileLog, newFile)
			lines = 0

			logfile = lw.openFile()
		}
	}

	logfile.Close()
}

//
func (lw *LoggerWorker) currentTime() string {
	currentTime := time.Now()
	return currentTime.Format(TIME_FORMAT)
}

//
func (lw *LoggerWorker) getNewFileName() string {
	var (
		n    = 0
		file string
	)
	for true {
		file = fmt.Sprintf("%s.%d", lw.FileLog, n)
		if _, err := os.Stat(file); os.IsNotExist(err) {
			break
		}
		n++
	}

	return file
}
