// +build windows

package main

import "golang.org/x/sys/windows/syscall"

func EnableVirtualTerminalProcessing() error {
	var kernel32Dll *syscall.LazyDLL = syscall.NewLazyDLL("Kernel32.dll")
	var setConsoleMode *syscall.LazyProc = kernel32Dll.NewProc("SetConsoleMode")
	const ENABLE_VIRTUAL_TERMINAL_PROCESSING uint32 = 0x4
	var mode uint32
	err := syscall.GetConsoleMode(syscall.Stdout, &mode)
	if err != nil {
		return err
	}

	mode |= ENABLE_VIRTUAL_TERMINAL_PROCESSING

	ret, _, err := setConsoleMode.Call(uintptr(syscall.Stdout), uintptr(mode))
	if ret == 0 {
		return err
	}
	return nil

}
