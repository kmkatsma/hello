package os

import (
	"fmt"
	"runtime"
	"strings"
)

//Wrapper represents current operating system.
type Wrapper struct {
	OS string
}

//NewWrapper returns pointer to an Wrapper initialized with current OS.
func NewWrapper() *Wrapper {
	return &Wrapper{runtime.GOOS}
}

//RunsOnWindows returns true if the runtime is windows, otherwise returns false.
func (t Wrapper) RunsOnWindows() bool {
	fmt.Print("Go runs on ")

	os := runtime.GOOS

	switch os {

	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
	if strings.EqualFold(os, "windows") {
		return true
	}
	return false

}
