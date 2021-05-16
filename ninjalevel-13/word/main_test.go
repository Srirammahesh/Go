package word

import (
	"go_sri/ninjalevel-13/quote"
	"testing"
)

func TestCount(t *testing.T) {
	n:=Count("this is sriram")
	if n!=3:
		t.Error("got ",n," expected 3")
}
func TestUseCount(t *testing.T) {
	m:=UseCount("one two two three")
	for k,v :=range m {
		switch k {
		case "one":
			if v!=1:
				t.Error("got",v,"expected 1")
		case "two":
			if v!=2:
			t.Error("got",v,"expected 2")

		case "three":
			if v!=1:
			t.Error("got",v,"expected 1")
		}
	}
}
func ExampleCount() {
	n:=Count("this is sriram")
	//output:
	//3
}
func BenchmarkCount(b *testing.B) {
	for i:=0;i<b.N;i++{
		Count(quote.SunAlso)
	}
}
func BenchmarkUseCount(b *testing.B) {
	for i:=0;i<b.N;i++{
		UseCount(quote.SunAlso)
	}
}
