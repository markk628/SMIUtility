package main

import (
	"fmt"
	"testing"

	"github.com/gocolly/colly"
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

func TestResponse(t *testing.T) {
	c := colly.NewCollector()
	
	c.OnResponse(func(r *colly.Response) {
		expected := 200
		if expected != r.StatusCode {
			t.Error("Failed")
		}
	})

	c.Visit("https://www.marketwatch.com/markets?mod=top_nav")
}

