# StackLog
colorful output and display call stack


## Example

```go
package main

import (
  log "github.com/rebill/stacklog"
)

func main() {
  log.Info("A walrus appears")
}
```

## Level logging
Log has six logging levels: Debug, Info, Warning, Error, Fatal and Panic.

```go
log.Debug("Useful debugging information.")
log.Debugf("Useful %s information.", "debugging")

log.Info("Something noteworthy happened!")
log.Infof("Something %s happened!", "noteworthy")

log.Warn("You should probably take a look at this.")
log.Warnf("You should probably take a look at %s.", "this")

log.Error("Something failed but I'm not quitting.")
log.Errorf("Something %s but I'm not quitting.", "failed")

// Calls os.Exit(1) after logging
log.Fatal("Bye.")
log.Fatal("%s Bye.", "Good")

// Calls panic() after logging
log.Panic("I'm bailing.")
log.Panicf("I'm %s.", "bailing")
```

You can set the logging level on a Logger, then it will only log entries with that severity or anything above it:

```go
// Will log anything that is info or above (warn, error, fatal, panic). Default.
log.SetLevel(log.InfoLevel)
```

It may be useful to set `log.Level = log.DebugLevel` in a debug or verbose environment if your application has that.

## Set output file path

```go
log.SetLogPath("/apps/logs/go/myapp", "application.log")
```

## Enable colorful output
```go
log.EnableColors()
```
Nicely color-coded in development (when a TTY is attached, otherwise just
plain text):

![Colored](http://o83oxyzd3.bkt.clouddn.com/stck_log.png)

## Output format
log format :
```
[datetime] [level] [message] [ip] [stack]
```
it will display call stack when log level above fatal and panic.

```
[2016/06/01 23:45:31] [info] [Something noteworthy happened!] [192.168.3.8] [main.go:11]
[2016/06/01 23:45:31] [warning] [You should probably take a look at this.] [192.168.3.8] [main.go:12]
[2016/06/01 23:45:31] [error] [Something failed but I'm not quitting.] [192.168.3.8] [/Users/rebill/GitHub/Demo/main.go:line 13 function main: log.Error("Something failed but I'm not quitting.")	/Users/rebill/Golang/GOROOT/src/runtime/proc.go:line 188 function main: main_main()	/Users/rebill/Golang/GOROOT/src/runtime/asm_amd64.s:line 1998 function goexit: BYTE	$0x90	// NOP	]
[2016/06/01 23:45:31] [fatal] [Bye.] [192.168.3.8] [/Users/rebill/GitHub/Demo/main.go:line 16 function main: log.Fatal("Bye.")	/Users/rebill/Golang/GOROOT/src/runtime/proc.go:line 188 function main: main_main()	/Users/rebill/Golang/GOROOT/src/runtime/asm_amd64.s:line 1998 function goexit: BYTE	$0x90	// NOP	]
```
