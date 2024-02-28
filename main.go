// Package ref provides functionality for making references to primitive values
package ref

// Ref is returning ref to the passed value
func Ref[T any](v T) *T {
	return &v
}
