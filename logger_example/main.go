package main

import(
	"time"
	"fmt"
)
const(
	logInfo		= "INFO"
	logWarning	= "WARNING"
	logError	= "ERROR"
)

type logEntry struct {
	time		time.Time
	severity	string
	message		string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{})

func main(){

	go logger()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100*time.Millisecond)
	doneCh <- struct{}{}

}

func logger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v] %v\n", entry.time.Format("2006-01-02 15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
		}
	}
}
