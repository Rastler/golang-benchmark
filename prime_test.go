package benchmark

import "testing"

func BenchmarkPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Prime(10000)
	}
}
