package utils

import (
	"fmt"
	"testing"
)

type Integer struct {
	value int
}

func (i Integer) CompareTo(other Comparable) int {
	return i.value - other.(Integer).value
}

func TestHeap(t *testing.T) {
	heap := NewHeap(10)

	heap.Offer(Integer{1})
	heap.Offer(Integer{9})
	heap.Offer(Integer{7})
	heap.Offer(Integer{6})
	heap.Offer(Integer{19})
	heap.Offer(Integer{36})
	heap.Offer(Integer{2})
	heap.Offer(Integer{6})
	heap.Offer(Integer{8})
	heap.Offer(Integer{11})
	heap.Offer(Integer{16})

	for heap.Size() > 0 {
		fmt.Println(heap.Poll().(Integer).value)
	}

	for i, i2 := range "a大范甘迪个" {
		fmt.Println(i, i2)
	}
}
