package base

import "testing"

func TestPartition(t *testing.T) {
	 a := []int{3,2,1,5,6,4}
	 idx := Partition(a, 0, len(a)-1)
	 t.Logf("idx:%d, ret:%+v", idx, a)
}
