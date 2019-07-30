package main

import "testing"

// Benchmark_NewEve 测试每次时间
func Benchmark_NewEve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewEve()
	}
}

func Benchmark_OldEve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OldEve()
	}
}
