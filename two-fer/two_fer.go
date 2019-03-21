// "Two Fer" track
package twofer

import "fmt"

// ShareWith create a sentence "One for X, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	
	return fmt.Sprintf("One for %s, one for me.", name)
}
