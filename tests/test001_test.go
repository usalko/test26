package tests

import (
	"testing"

	"github.com/usalko/test26/internal"
)

// Write simple function for cloning string (cloned string contains all characters that source but it holds another memory area)
func Test001(t *testing.T) {
	sourceString := "This a string"
	targetString := internal.Test001(sourceString)
	if &sourceString == &targetString {
		t.Errorf("address the same in source and target strings")
	}
	if sourceString != targetString {
		t.Errorf("expected content %v but got %v", sourceString, targetString)
	}
}
