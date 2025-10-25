package main

import (
	"fmt"

	"github.com/qiudeng7/golang-lab.git/demo2-module/module_1"
)

func main() {
	// Get a greeting message and print it.
	message := module_1.Hello("Gladys")
	fmt.Println(message)
}
