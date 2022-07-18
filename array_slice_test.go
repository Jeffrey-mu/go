package main

import (
	"reflect"
	"testing"
)

func TestArra(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Array(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Array(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
	t.Run("input two array return two len", func(t *testing.T) {
		arr := []int{1, 2}
		arr2 := []int{1, 2, 3}
		got := ArrayAll(arr, arr2)
		want := []int{3, 6}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d ", got, want)
		}
	})
}
