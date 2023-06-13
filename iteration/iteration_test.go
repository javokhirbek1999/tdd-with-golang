package iteration

import "testing"

func TestRepeat(t *testing.T) {

	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}
}

// Benchamrking - runs the specified function b.N times
func BenchmarkRepeat(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
