package main

import (
	"fmt"
	"math/rand"
	"testing"
)

type testAdd struct {
	x, y, wanted int
}

func TestAdd(t *testing.T) {
	tests := []testAdd{
		{1, 2, 3},
		{0, 0, 0},
		{-5, 5, 0},
		{-6, -5, -11},
	}

	for _, test := range tests {
		if got := Add(test.x, test.y); got != test.wanted {
			t.Errorf("Got %q, but wanted %q", got, test.wanted)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(rand.Intn(11), rand.Intn(11))
	}
}

func ExampleAdd() {
	fmt.Println(Add(15, 1))
	// Output: 16
}

type testSubtract struct {
	x, y, wanted int
}

func TestSubtract(t *testing.T) {
	tests := []testSubtract{
		{1, 2, -1},
		{0, 2, -2},
		{2, 1, 1},
		{0, 0, 0},
		{8, -8, 16},
		{-4, 7, -11},
	}

	for _, test := range tests {
		if got := Subtract(test.x, test.y); got != test.wanted {
			t.Errorf("Got %q, but wanted %q", got, test.wanted)
		}
	}
}

func BenchmarkSubtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Subtract(rand.Intn(11), rand.Intn(11))
	}
}

func ExampleSubtract() {
	fmt.Println(Subtract(15, 1))
	// Output: 14
}
