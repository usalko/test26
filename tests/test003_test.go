package tests

import (
	"testing"

	"github.com/usalko/test26/internal"
)

// Write function for strings comparison. Compare means lexicographical comparison.
// Example:
//
//	input "str1", "str2"
//
// output -1 if second string greater then first string, 1 if first string greater then second and 0 if strings are equals
//
// Use strings.Compare function
// ** Compare performance for direct comparison and youth test
func Test003(t *testing.T) {
	internal.Test003("")
}
