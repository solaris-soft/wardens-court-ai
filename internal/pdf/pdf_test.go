package pdfReader

import (
	"testing"
)

func TestExtractText(t *testing.T) {

	// Execute function being tested
	text, err := ExtractText("2023WAMW1.pdf", 1, 1)

	// Assert results
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	t.Log("Text:")
	t.Log(text)
}
