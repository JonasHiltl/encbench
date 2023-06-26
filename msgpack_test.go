package encbench

import (
	"bytes"
	"testing"
	"time"

	"github.com/tinylib/msgp/msgp"
)

var foo = Foo{Foo: "Foo", Bar: "Bar"}

var complex = Complex{
	Integer: 42,
	Float:   3.14,
	Boolean: true,
	String:  "Hello, World!",
	Slice:   []int{1, 2, 3, 4, 5},
	Map: map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	},
	StructField: SomeOtherStruct{
		Field1: "Nested struct field",
		Field2: 100,
	},
	Time: time.Now(),
}

func benchmarkEncodeMsgPack(in msgp.Encodable, b *testing.B) {
	var buf bytes.Buffer
	en := msgp.NewWriter(&buf)

	for i := 0; i < b.N; i++ {
		in.EncodeMsg(en)
	}
}

func benchmarkSizeMsgPack(in msgp.Encodable, b *testing.B) {
	var buf bytes.Buffer
	en := msgp.NewWriter(&buf)

	for i := 0; i < b.N; i++ {
		in.EncodeMsg(en)
		en.Flush()
		b.ReportMetric(float64(buf.Len()/b.N), "bytes")
	}
}

func BenchmarkEncodeMsgPackFoo(b *testing.B) {
	benchmarkEncodeMsgPack(foo, b)
}

func BenchmarkEncodeMsgPackComplex(b *testing.B) {
	benchmarkEncodeMsgPack(&complex, b)
}

func BenchmarkSizeMsgPackFoo(b *testing.B) {
	benchmarkSizeMsgPack(foo, b)
}
func BenchmarkSizeMsgPackComplex(b *testing.B) {
	benchmarkSizeMsgPack(&complex, b)
}
