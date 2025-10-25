package hello

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// := 表示声明赋值
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
