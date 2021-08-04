package main

import "testing"

func TestAdder(t *testing.T) {
	got := Adder(2, 4)
	want := 6

	if got != want {
		t.Errorf("got: %d | want: %d", got, want)
	}
}
