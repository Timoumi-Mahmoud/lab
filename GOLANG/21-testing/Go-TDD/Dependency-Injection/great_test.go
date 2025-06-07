package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Bara")

	got := buffer.String()
	want := "Hello, Bara"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
