package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeletion(t *testing.T) {
	cap := 5
	max := 11
	ss := &SparseSet{
		Dense:  make([]int, cap),
		Sparse: make([]int, max),
		N:      0,
		Max:    max,
		Cap:    cap,
	}

	ss.Insert(4)
	ss.Insert(6)
	ss.Insert(8)

	ss.Delete(6)
	expected := make([]int, cap)
	expected[0] = 4
	expected[1] = 8
	// expected[2] = 9
	fmt.Println(expected, ss.Dense)
	if !reflect.DeepEqual(ss.Dense, expected) {
		t.Errorf("Expected dense array (%v) is not same as"+
			" actaul array (%v)", expected, ss.Dense)
	}

}
