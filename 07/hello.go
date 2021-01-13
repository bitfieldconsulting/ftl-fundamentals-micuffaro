package hello

import "fmt"

// ReturnGreeting returns a greeting
func ReturnGreeting(greeting string) string {
	return fmt.Sprint(greeting, " yourself!")
}
