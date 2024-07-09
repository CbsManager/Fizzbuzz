package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("Case 1 should return 1", func(t *testing.T) {
		want := "1"

		get := FizzBuzz(1)

		if want != get {
			t.Errorf("Expected value %s but got %s", want, get)
		}
	})
}
