package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/jericop/hello-world-go-app/pkg/greeting"
)

func TestMain_Output(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run main
	main()

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	expected := fmt.Sprintf("%s\n", greeting.Message)
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
