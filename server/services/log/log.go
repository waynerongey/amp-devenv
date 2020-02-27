package log

import (
	"fmt"
	"time"
)

//Constants for the Logger
const (
	INFO    = "INFO "
	WARNING = "WARNING "
	ERROR   = "ERROR "
	FATAL   = "FATAL "
)

//Info logs out an info message to the log
func Info(args ...interface{}) {
	fmt.Println(INFO, time.Now().Format("01/02/06 3:04:05 PM"), args)
}

//Error logs out an error message to the log
func Error(args ...interface{}) {
	fmt.Println(ERROR, time.Now().Format("01/02/06 3:04:05 PM"), args)
}

//Warning logs out an warning message to the log
func Warning(args ...interface{}) {
	fmt.Println(WARNING, time.Now().Format("01/02/06 3:04:05 PM"), args)
}