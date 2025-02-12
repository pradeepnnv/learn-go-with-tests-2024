package iteration

import "testing"

func TestIteration(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestIterationCountZero(t *testing.T) {
	got := Repeat("a", 0)
	want := ""
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10000)
	}
}
