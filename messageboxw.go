package mbw

import (
	"syscall"
	"unsafe"
)

var messageBoxW *syscall.LazyProc

const (
	// MessageBoxW mode
	MODE_OK                 = uintptr(0)
	MODE_OK_CANCEL          = uintptr(1)
	MODE_ABORT_RETRY_IGNORE = uintptr(2)
	MODE_YES_NO_CANCEL      = uintptr(3)
	MODE_YES_NO             = uintptr(4)
	MODE_RETRY_CANCEL       = uintptr(5)
	// MessageBoxW button
	BUTTON_OK     = 1
	BUTTON_CANCEL = 2
	BUTTON_ABORT  = 3
	BUTTON_RETRY  = 4
	BUTTON_IGNORE = 5
	BUTTON_YES    = 6
	BUTTON_NO     = 7
)

func init() {
	messageBoxW = syscall.NewLazyDLL("user32.dll").NewProc("MessageBoxW")
}

func strptr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}

func MessageBoxW(title, text string, mode uintptr) (clicked uintptr, err error) {
	clicked, _, err = messageBoxW.Call(uintptr(0), strptr(text), strptr(title), mode)
	if err.Error() == "The operation completed successfully." {
		err = nil
	}
	return
}
