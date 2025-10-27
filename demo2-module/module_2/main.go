package main

import (
	"fmt"
	"log"

	"github.com/qiudeng7/golang-lab.git/demo2-module/module_1"
)

func main() {

	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    // log.SetFlags(0)

	// Request a greeting message.
    message, err := hello.Hello("")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

	// If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}
