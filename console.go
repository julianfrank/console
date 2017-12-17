package console

import (
	"fmt"
	"log"
)

//LogMode Use this to trun off or on Logs
var (
	LogMode   = true
	ErrorMode = true
)

//Log use as console.log(pattern,vars...) return a string
func Log(pattern string, vars ...interface{}) string {
	var msg = fmt.Sprintf(pattern, vars...)
	if LogMode {
		log.Println(msg)
	}
	return msg
}

//Error return an error object as per the provided format
// Usage console.Error(pattern,vars...)
// Returns error object
func Error(pattern string, vars ...interface{}) error {
	var err = fmt.Errorf(pattern, vars...)
	if ErrorMode {
		log.Println(err.Error())
	}
	return err
}
