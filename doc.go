/*
Package StackLog is a structured logger for Go, colorful output and display call stack.
The simplest way to use StackLog is simply the package-level exported logger:
	package main

	import (
	  log "github.com/rebill/stacklog"
	)

	func main() {
	  log.Info("Something noteworthy happened!")
	}
Output:
  [2016/06/01 23:45:31] [info] [Something noteworthy happened!] [192.168.3.8] [main.go:11]

For a full guide visit https://github.com/rebill/stackLog
*/
package stacklog
