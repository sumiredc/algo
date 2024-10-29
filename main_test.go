package main

import (
	"io"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}

	stdout := os.Stdout
	os.Stdout = writer
	defer func() { os.Stdout = stdout }()

	main()

	writer.Close()
	var buf io.Reader = reader
	output, err := io.ReadAll(buf)
	if err != nil {
		t.Fatalf("Failed to read from pipe: %v", err)
	}

	expected := "Hello Algorithm in Golang!\n"
	if string(output) != expected {
		t.Errorf("main() output = %q, want %q", output, expected)
	}
}

func TestGreeting(t *testing.T) {
	expected := "Hello Algorithm in Golang!"
	actual := greeting()

	if expected != actual {
		t.Errorf(
			"Failed test case [expected: %s, actual: %s]",
			expected,
			actual,
		)
	}
}
