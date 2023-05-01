package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// ===========================
// generic implementation of min heap
// ===========================


type Heap[T constraints.Ordered] struct{
	Val T
	L *Heap[T]
	R *Heap[T]
}

func (h Heap[T]) Pop ()T {
	top := h.Val



	return top
}

// Turn slice into heap
func heapify[T constraints.Ordered](slice []T) Heap[T] {

}

func main(){

	test := []int{ 1, 4, 5, 87, 6}

	h := heapify(test)

	fmt.Println(h.Pop)
}