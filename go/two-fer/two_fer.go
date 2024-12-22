// Package twofer does shit idk, its literally one single function
package twofer

import "fmt"

// ShareWith generates a sharing message using a name parameter. If no name is given, defaults to using "you"
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
