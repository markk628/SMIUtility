package main

import (
	"fmt"
	"testing"
)

// TestFileExists tests fileExists function
func TestFileExists(t *testing.T) {
	fmt.Println("Test Calculate")
	expected := true
	result := fileExists("indexes.json")
	if expected != result {
		t.Error("Failed")
	}
}