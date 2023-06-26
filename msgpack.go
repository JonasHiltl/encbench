//go:generate msgp -tests=false
package encbench

import "time"

type Foo struct {
	Foo string `msg:"foo" json:"foo"`
	Bar string `msg:"bar" json:"bar"`
}

type Complex struct {
	Integer     int             `msg:"integer" json:"integer"`
	Float       float64         `msg:"float" json:"float"`
	Boolean     bool            `msg:"boolean" json:"boolean"`
	String      string          `msg:"string" json:"string"`
	Slice       []int           `msg:"slice" json:"slice"`
	Map         map[string]int  `msg:"map" json:"map"`
	StructField SomeOtherStruct `msg:"struct" json:"struct"`
	Time        time.Time       `msg:"time" json:"time"`
}
type SomeOtherStruct struct {
	Field1 string `msg:"field1" json:"field1"`
	Field2 int    `msg:"field2" json:"field2"`
}
