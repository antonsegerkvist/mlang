package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mlang/interpreter"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	fileArgument := os.Args[1]
	program, err := ioutil.ReadFile(fileArgument)
	if err != nil {
		fmt.Println(`==> Error: ` + err.Error())
		return
	}
	err = interpreter.Execute(string(program))
	if err != nil {
		fmt.Println(`==> Error: ` + err.Error())
		return
	}
}
