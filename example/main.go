package main

import (
	"fmt"

	"github.com/Icarus-0727/go-mbw"
)

func main() {
	clicked, err := mbw.MessageBoxW("测试", "这里是对 MODE_OK_CANCEL 的测试", mbw.MODE_YES_NO)
	fmt.Printf("clicked = %v, err = %v\n", clicked, err)
}
