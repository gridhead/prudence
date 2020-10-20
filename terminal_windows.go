// +build windows

package main

import (
	"syscall"

	sequences "github.com/konsorten/go-windows-terminal-sequences"
)

func EnableVirtualTerminalProcessing() error {
	return sequences.EnableVirtualTerminalProcessing(syscall.Stdout, true)
}
