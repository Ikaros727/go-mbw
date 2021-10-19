package main

import (
	"fmt"

	"github.com/Icarus-0727/go-mbw"
)

func main() {
	clicked, err := mbw.MessageBoxW("Test", "This is a test for `MODE_OK_CANCEL`", mbw.MODE_OK_CANCEL)
	fmt.Printf("clicked = %v, err = %v\n", clicked, err)
}
