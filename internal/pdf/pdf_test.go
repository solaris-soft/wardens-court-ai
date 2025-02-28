package pdfReader

import (
	"strings"
	"testing"
)

func TestExtractText(t *testing.T) {

	// Execute function being tested
	text, err := ExtractText("../../test/2023WAMW1.pdf", 1, 1)
	expected := strings.Contains(text, "[2023] WAMW 1")

	t.Run("Extracting text from PDF", func(t *testing.T) {
		// Assert results
		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		if !expected {
			t.Fatalf("Expected text to contain '[2023] WAMW 1', but got: %v", text)
		}

		t.Log("Text:")
		t.Log(text)
	})
}
