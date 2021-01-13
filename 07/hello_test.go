package hello_test

import (
	"hello"
	"testing"
)

func TestReturnGreeting(t *testing.T) {
	got := hello.ReturnGreeting("Hi there")
	want := "Hi there yourself!"
	if got != want {
		t.Errorf("FAIL: got %s, want %s", got, want)
	}
}
