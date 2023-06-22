package main

import "testing"

func BenchmarkUseMu(b *testing.B) {
	UseMu()
}

func BenchmarkUseRWMu(b *testing.B) {
	UseRWMu()
}
