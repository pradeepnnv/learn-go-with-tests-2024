package iteration

import "testing"

func TestIteration(t *testing.T) {
	repeated := Repeat("SriRam", 3)
	expected := "SriRamSriRamSriRam"
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 2)
	}
}
