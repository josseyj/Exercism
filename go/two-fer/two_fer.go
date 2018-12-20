// Package twofer provides two for one methods.
package twofer

// ShareWith returns a message to share.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
