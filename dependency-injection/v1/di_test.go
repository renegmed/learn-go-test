package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// buffer type from the bytes package implement the Writer interface
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris") // puts Chris to &buffer

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
