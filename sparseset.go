/*
One popular representation of a set of n integers ranging from 0 to m is a bit vector, with 1 bits at the positions corresponding to the integers in the set.
Adding a new integer to the set, removing an integer from the set, and checking whether a particular integer is in the set are all very fast constant-time operations
(just a few bit operations each). Unfortunately, two important operations are slow: iterating over all the elements in the set takes time O(m), as does clearing the set.

*****If the common case is that m is much larger than n (that is, the set is only sparsely populated) and iterating or clearing the set happens frequently,
     then it could be better to use a representation that makes those operations more efficient. That's where the trick comes in.*********


Operation		Bit Vector	Sparse set

Lokup			O(1)		O(1)
Insertion		O(1)		O(1)
Deletion		O(1)		O(1)
clear-set		O(m)		O(1)
iterate		        O(m)		O(n)

Here M is the range or max element and N is the number of integers

***** SPACE Complexity for Spare Set in O(m+n) since we need two arrays Dense(for storing N integers) and Sparse(For Storing Range 0 to Max)

***** REFER : https://research.swtch.com/sparse **************
*/
package main

import "fmt"

type SparseSet struct {
	Dense  []int // To store actual set elements
	Sparse []int // To store indexes of actual elements
	N      int   // Current number of elements
	Max    int   // Maximum value in set or size of sparse[]
	Cap    int   // Capacity of set or size of dense[]
}

func main() {
	max := 11
	cap := 5
	ss := &SparseSet{
		Dense:  make([]int, cap),
		Sparse: make([]int, max),
		N:      0,
		Max:    max,
		Cap:    cap,
	}

	// Inserting Elements into the sparse Set : O(1)
	ss.Insert(8)
	ss.Insert(6)
	ss.Insert(7)
	ss.Insert(3)
	ss.Insert(2)

	fmt.Println("Sparse Set After Insertion: ", ss)
	isThere := ss.Search(7)
	fmt.Println(isThere)

	// Iterating over Sparse Set:  O(n)
	ss.Iterate()

	// Deleting Elements from Sparse Set : O(1)
	ss.Delete(6)
	ss.Delete(3)
	ss.Delete(8)
	fmt.Println("Sparse Set After Deletion: ", ss)

	// Looup in Sparse Set: O(1)
	isExists := ss.Search(8)
	fmt.Println(isExists)

	// Clearing / Removing all elements of Sparse Set : O(n)
	ss.Clear()

	fmt.Println("After Clearing the Sparse Set", ss)
}

func (ss *SparseSet) Search(x int) bool {

	return ss.Dense[ss.Sparse[x]] == x && ss.Sparse[x] < ss.N
}

func (ss *SparseSet) Insert(x int) {
	if ss.N >= ss.Cap {
		fmt.Println("Capacity Exhausted")
		return
	}
	if ss.Search(x) {
		fmt.Printf("Value %v already exists\n", x)
		return
	}

	ss.Dense[ss.N] = x

	ss.Sparse[x] = ss.N

	ss.N++
}

func (ss *SparseSet) Delete(x int) {

	if !ss.Search(x) {
		fmt.Printf("The value %v does not exits \n", x)
		return
	}

	// The idea it to replace the item to be removed with the last element
	lastItem := ss.Dense[ss.N-1]

	// replacing item to be removed with last element
	ss.Dense[ss.Sparse[x]] = lastItem

	// Now Remove the last item from dense array
	ss.Dense[ss.N-1] = 0

	// Update the Value of last item in sparse array with value to be deleted since ss.Sparse[x] will give the value of index where replacement happened
	ss.Sparse[lastItem] = ss.Sparse[x]

	// Removing the value form sparse array

	ss.Sparse[x] = 0

	// decrease the N by one
	ss.N--

}

func (ss *SparseSet) Clear() {
	ss.N = 0
	ss.Dense = make([]int, ss.Cap)
	ss.Sparse = make([]int, ss.Max)
}

func (ss *SparseSet) Iterate() {
	for _, val := range ss.Dense {
		fmt.Printf("%v\t", val)
	}
	fmt.Println()
}
