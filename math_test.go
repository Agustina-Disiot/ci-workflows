package math

import "testing"

func TestAdd1(t *testing.T) {

	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestAdd2(t *testing.T) {

	got := Add(3, 6)
	want := 9

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSubstract(t *testing.T) {

	got := Subtract(6, 3)
	want := 3

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
