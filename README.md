# go-mbw
一个通过`user32.dll`调用 Windows 平台的`MessageBoxW`函数的 Go 语言库
A Go lib for call windows platform function `MessageBoxW` from `user32.dll`.

# 安装(Install)
```sh
go get github.com/Icarus-0727/go-mbw
```

# 例子(Example)
```go
package main

import (
	"fmt"

	"github.com/Icarus-0727/go-mbw"
)

func main() {
	clicked, err := mbw.MessageBoxW("Test", "This is a test for `MODE_OK_CANCEL`", mbw.MODE_OK_CANCEL)
	fmt.Printf("clicked = %v, err = %v\n", clicked, err)
}
```

# 效果(Result)
![image](https://user-images.githubusercontent.com/80149308/137913837-bf80b2b0-411a-4057-81e5-edb5ba39b303.png)
