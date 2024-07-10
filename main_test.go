package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("case 1 should return 1", func(t *testing.T) {
		want := "1"

		got := FizzBuzz(1)

		if want != got {
			t.Errorf("Expected %s but got %s", want, got)
		}
	})
}
