package pkg

import (
	"fmt"
	"os/exec"
	"runtime"
)

func GetFiles() {
	_, filename, _, _ := runtime.Caller(0)
	cmd := exec.Command("ls", filename)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println("failed")

		return
	}

	fmt.Println(string(stdout))
}
