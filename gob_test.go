package encbench

import (
	"bytes"
	"encoding/gob"
	"testing"

	"github.com/tinylib/msgp/msgp"
)

func benchmarkEncodeGob(in any, b *testing.B) {
	var buf bytes.Buffer
	en := gob.NewEncoder(&buf)

	for i := 0; i < b.N; i++ {
		en.Encode(in)
	}
}

func benchmarkSizeGob(in msgp.Encodable, b *testing.B) {
	var buf bytes.Buffer
	en := gob.NewEncoder(&buf)

	for i := 0; i < b.N; i++ {
		en.Encode(in)
		b.ReportMetric(float64(buf.Len()/b.N), "bytes")
	}
}
func BenchmarkEncodeGobFoo(b *testing.B) {
	benchmarkEncodeGob(foo, b)
}

func BenchmarkEncodeGobComplex(b *testing.B) {
	benchmarkEncodeGob(&complex, b)
}

func BenchmarkSizeGobFoo(b *testing.B) {
	benchmarkSizeGob(foo, b)
}
func BenchmarkSizeGobComplex(b *testing.B) {
	benchmarkSizeGob(&complex, b)
}
