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
console.Log("This is %s Code by %s",comment,dude)

mylog:=console.Log("This is %s Code by %s",comment,dude)
//Output mylog=> "This is Awsome Code by Julian Frank"
```

### Error
In your Application use the familiar console.Error ...Just remember that the E is capital
```go
comment:="Awsome!"
dude:="Julian Frank"
myError:=console.Error("This is %s Code by %s",comment,dude)
//Output myError=> "This is Awsome Code by Julian Frank" of type error
//Use err.Error() to retreive error string 
```

## License
Apache 2.0