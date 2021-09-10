package main

import (
	"fmt"
	shell "github.com/codeskyblue/go-sh"
)

func main() {
	sh := shell.NewSession()
	sh.PipeFail = true
	sh.PipeStdErrors = true

	sh.Command("cat", "exec/input.txt")
	sh.Command("wc")
	out, err := sh.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
