package tests

import (
	"testing"

	"github.com/usalko/test26/internal"
)

// Concatenate different type of information into single string
// Example:
//
//	input are "Size is ", 0x8, " bytes"
//	output is "Size is 8 bytes"
//
// Use strings builder
// ** Compare performance for builder and string interpolation
func Test002(t *testing.T) {
	internal.Test002("")
}
