package benchmark

import "testing"

func BenchmarkBubble(b *testing.B) {
	unsort := []int{2, 212, 3001, 14, 501, 7800, 9932, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
	for i := 0; i < b.N; i++ {
		Bubble(unsort)
	}
}
