package workflow

import "testing"

func TestShouldRetry(t *testing.T) {
	if !ShouldRetry(1, 3) {
		t.Fatal("expected true")
	}
}
