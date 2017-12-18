// Package console tries to bring the convenience of console.log/console.error in JS/NodeJS to Go
package console

import (
	"fmt"
	"log"
)


var (
	//LogMode Use this to turn Log output to console off or on
	LogMode   = true
	//LogMode Use this to turn Error output to console off or on
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
// Output: This will Return an error object
func Error(pattern string, vars ...interface{}) error {
	var err = fmt.Errorf(pattern, vars...)
	if ErrorMode {
		log.Println(err.Error())
	}
	return err
}
