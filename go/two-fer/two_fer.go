// Package twofer implements two for one messaging for names
package twofer

import "fmt"

// ShareWith returns a two for one message for the provided name
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
