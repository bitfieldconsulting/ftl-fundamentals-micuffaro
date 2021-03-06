package hello_test

import (
	"hello"
	"testing"
)

func TestHello(t *testing.T) {
	got := hello.Greeting()
	want := "Hello Gophers!"
	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}
