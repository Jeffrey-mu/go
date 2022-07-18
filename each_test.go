package main

import "testing"

func TestEach(t *testing.T) {
	str := Repeat("a")
	expected := "aaaaa"
	if str != expected {
		t.Errorf("got '%q' want '%q'", str, expected)
	}
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
