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

	t.Run("Case 2 should return 2", func(t *testing.T) {
		want := "2"

		get := FizzBuzz(2)

		if want != get {
			t.Errorf("Expected value %s but got %s", want, get)
		}
	})
	t.Run("Case 3 should return Fizz", func(t *testing.T) {
		want := "Fizz"

		get := FizzBuzz(3)

		if want != get {
			t.Errorf("Expected value %s but got %s", want, get)
		}
	})

	t.Run("Case 4 should return 4", func(t *testing.T) {
		want := "4"

		get := FizzBuzz(4)

		if want != get {
			t.Errorf("Expected value %s but got %s", want, get)
		}
	})

	t.Run("Case 5 should return 5", func(t *testing.T) {
		want := "Buzz"

		get := FizzBuzz(5)

		if want != get {
			t.Errorf("Expected value %s but got %s", want, get)
		}
	})

	t.Run("Case 6 should return Fizz", func(t *testing.T) {
		want := "Fizz"

		get := FizzBuzz(6)

		if want != get {
			t.Errorf("Expected value %s but got %s", want, get)
		}
	})

}
