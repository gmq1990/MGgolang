package calc

import "testing"

func TestAdd(t *testing.T) {
	if 4 != Add(1, 3) {
		t.Error("1 + 3 != 4")
	}
}

func TestAddFlag(t *testing.T) {
	if -1 != Add(-1, 3) {
		t.Error("-1 + any != -1")
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 1; i <= 10; i++ {
		Factorial(i)
	}
}
