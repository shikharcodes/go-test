package maths

import "testing"

func TestAdd(t *testing.T) {
	got := Add(3, 5)
	want := 8

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
