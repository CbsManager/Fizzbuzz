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
	t.Run("case 2 should return 2", func(t *testing.T) {
		want := "2"

		got := FizzBuzz(2)

		if want != got {
			t.Errorf("Expected %s but got %s", want, got)
		}
	})

	t.Run("case 3 should return Fizz", func(t *testing.T) {
		want := "Fizz"

		got := FizzBuzz(3)

		if want != got {
			t.Errorf("Expected %s but got %s", want, got)
		}
	})

	t.Run("case 4 should return 4", func(t *testing.T) {
		want := "4"

		got := FizzBuzz(4)

		if want != got {
			t.Errorf("Expected %s but got %s", want, got)
		}
	})

	t.Run("case 5 should return Buzz", func(t *testing.T) {
		want := "Buzz"

		got := FizzBuzz(5)

		if want != got {
			t.Errorf("Expected %s but got %s", want, got)
		}
	})
	t.Run("case 6 should return Fizz", func(t *testing.T) {
		want := "Fizz"

		got := FizzBuzz(6)

		if want != got {
			t.Errorf("Expected %s but got %s", want, got)
		}
	})

	t.Run("case 7 should return 7", func(t *testing.T) {
		want := "7"

		got := FizzBuzz(7)

		if want != got {
			t.Errorf("Expected %s but got %s", want, got)
		}
	})

	t.Run("case 8 should return 8", func(t *testing.T) {
		want := "8"

		got := FizzBuzz(8)

		if want != got {
			t.Errorf("Expected %s but got %s", want, got)
		}
	})

	t.Run("case 9 should return Fizz", func(t *testing.T) {
		want := "Fizz"

		got := FizzBuzz(9)

		if want != got {
			t.Errorf("Expected %s but got %s", want, got)
		}
	})

	t.Run("case 10 should return Buzz", func(t *testing.T) {
		want := "Buzz"

		got := FizzBuzz(10)

		if want != got {
			t.Errorf("Expected %s but got %s", want, got)
		}
	})

	t.Run("case 11 should return 11", func(t *testing.T) {
		want := "11"

		got := FizzBuzz(11)

		if want != got {
			t.Errorf("Expected %s but got %s", want, got)
		}
	})
}
