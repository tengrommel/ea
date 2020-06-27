package nlogn

import "testing"

func benchmarkOnlognLoop(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		OnlognLoop(i)
	}
}

func BenchmarkOnlognLoop10(b *testing.B)   { benchmarkOnlognLoop(10, b) }
func BenchmarkOnlognLoop100(b *testing.B)  { benchmarkOnlognLoop(100, b) }
func BenchmarkOnlognLoop1000(b *testing.B) { benchmarkOnlognLoop(1000, b) }
