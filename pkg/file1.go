package pkg

import (
	"fmt"
	"runtime"
)

func GetFiles() {
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println(filename)
}
