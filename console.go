package console

import (
	"fmt"
	"log"
)

//LogMode Use this to trun off or on Logs
var LogMode bool

//Log use as console.log(pattern,vars...)
func Log(pattern string, vars ...interface{}) string {
	var msg = fmt.Sprintf(pattern, vars...)
	if LogMode {
		log.Println(msg)
	}
	return msg
}
