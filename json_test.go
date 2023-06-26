package encbench

import (
	"encoding/json"
	"testing"
)

func benchmarkEncodeJSON(in any, b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(in)
	}
}

func benchmarkSizeJSON(in any, b *testing.B) {
	for i := 0; i < b.N; i++ {
		by, _ := json.Marshal(in)
		b.ReportMetric(float64(len(by)), "bytes")
	}
}
func BenchmarkEncodeJSONFoo(b *testing.B) {
	benchmarkEncodeJSON(foo, b)
}

func BenchmarkEncodeJSONComplex(b *testing.B) {
	benchmarkEncodeJSON(&complex, b)
}

func BenchmarkSizeJSONFoo(b *testing.B) {
	benchmarkSizeJSON(foo, b)
}
func BenchmarkSizeJSONComplex(b *testing.B) {
	benchmarkSizeJSON(complex, b)
}
