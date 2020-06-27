package oone

import "testing"

func BenchmarkThreeWords(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ThreeWords()
	}
}

func BenchmarkTenWords(b *testing.B) {
	for n := 0; n < b.N; n++ {
		TenWords()
	}
}
