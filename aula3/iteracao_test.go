package main

import "testing"

func BenchmarkRepetir(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repetir("a")
    }
}