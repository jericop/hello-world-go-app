package greeting_test

import (
	"testing"

	"github.com/jericop/hello-world-go-app/pkg/greeting"
)

func TestMessage(t *testing.T) {
	expected := "hello world"
	if greeting.Message != expected {
		t.Errorf("Expected message to be %q, got %q", expected, greeting.Message)
	}
}
