[![GoDoc](https://godoc.org/github.com/julianfrank/console?status.svg)](https://godoc.org/github.com/julianfrank/console)
 [![Coverage Status](https://coveralls.io/repos/github/julianfrank/console/badge.svg?branch=master)](https://coveralls.io/github/julianfrank/console?branch=master) [![Build Status](https://travis-ci.org/julianfrank/console.svg?branch=master)](https://travis-ci.org/julianfrank/console)

# console package for those used to JS console.log(...)
Simple Console Log Package useful for anyone coming from Javascript background

## Usage
Just import this package
```go
import "github.com/julianfrank/console"
```

### Log
In your Application use the familiar console.Log ...Just remember that the L is capital
```go
comment:="Awsome!"
dude:="Julian Frank"
console.Log("This comment (%s) was made by %s",comment,dude)

mylog:=console.Log("This comment (%s) was made by %s",comment,dude)
//Output mylog=> "This comment (Awsome!) was made by Julian Frank"
```

### Error
In your Application use the familiar console.Error ...Just remember that the E is capital
```go
comment:="Awsome!"
dude:="Julian Frank"
myError:=console.Error("This comment (%s) was made by %s")
//Output myError=> "This comment (Awsome!) was made by Julian Frank" 
//of type error
//Use err.Error() to retreive error string 
```

### Console Display Flags 
If You want to selectively enable or disable console display of Log or Error you can simply do that by setting console.LogMode and console.ErrorMode to true or false. True is default
```go
//To make screen completely quiet
console.LogMode=false
console.ErrorMode=false

//To Display only errors and not the Logs
console.LogMode=false
console.ErrorMode=true
```

## License
Apache 2.0
