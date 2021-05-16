package mymath

import "testing"

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data []int
		cavg float64
	}
	tests := []test{
		test{[]int{10,20,40,60,100},40},
		test{[]int{1,2,4,6,10},4},
	}
	for _,v:=range tests{
		x:=CenteredAvg(v.data)
		if x!=v.cavg{
			t.Error("got",x,"expected",v.cavg)
		}
	}
}
func ExampleCenteredAvg() {
	CenteredAvg([]int{1,2,4,6,9})
	//output:
	//2
}
func BenchmarkCenteredAvg(b *testing.B) {
	for i:=0;i<b.N;i++{
	CenteredAvg([]int{1,2,3,44,5,6,7,8,9,10})
	}
}


