package main

import "testing"

// TestHello(testss the hello function created in hello.go)
func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
