package tests

import (
	"testing"

	"github.com/usalko/test26/internal"
)

// Write a Go Func to compare two structs to determine if the first contains equivalent property values to the second one.
// Example:
//
//	{ Name: "Alex" }, { Name: "Bob" }
//	output is false
//
//	{ Age: 50 }, { Age: 50 }
//	output is true
func Test004Case1(t *testing.T) {
	if internal.Test004(
		&(struct {
			Name string
		}{Name: "Alex"}),
		&(struct {
			Name string
		}{Name: "Bob"}),
	) {
		t.Errorf("The structs should be different")
	}
}

func Test004Case2(t *testing.T) {
	if !internal.Test004(
		&(struct {
			Age int
		}{Age: 50}),
		&(struct {
			Age int
		}{Age: 50}),
	) {
		t.Errorf("The structs should be equal")
	}
}
