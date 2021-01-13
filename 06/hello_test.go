package hello_test

import (
	"hello"
	"testing"
)

func TestHello(t *testing.T) {
	got := hello.Greeting()
	want := "Howdy, Gopherinos!"
	if got != want {
		t.Errorf("FAIL: got %s, want %s", got, want)
	}
}
