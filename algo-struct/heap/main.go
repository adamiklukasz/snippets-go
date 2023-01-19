package main

import (
	"fmt"
)

type Heap struct {
	impl []int
}

func NewHeap() *Heap {
	return &Heap{
		impl: []int{},
	}
}

func (h *Heap) leftChildIndex(i int) int {
	return 2*i + 1
}

func (h *Heap) rightChildIndex(i int) int {
	return 2*i + 2
}

func (h *Heap) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *Heap) hasLeftChildIndex(i int) bool {
	return h.leftChildIndex(i) < len(h.impl)
}

func (h *Heap) hasRightChildIndex(i int) bool {
	return h.rightChildIndex(i) < len(h.impl)
}

func (h *Heap) Insert(element int) {
	h.impl = append(h.impl, element)
	h.upHeapify(len(h.impl) - 1)
}

func (h *Heap) GetRoot() (int, bool) {
	if len(h.impl) == 0 {
		return 0, false
	}

	rv := h.impl[0]
	h.impl[0] = h.impl[len(h.impl)-1]
	h.impl = h.impl[0 : len(h.impl)-1]

	if len(h.impl) > 0 {
		h.downHeapify(0)
	}

	return rv, true
}

func (h *Heap) downHeapify(i int) {
	currValue := h.impl[i]
	smallestIndex := i
	smallestValue := currValue

	if h.hasLeftChildIndex(i) {
		leftChildIndex := h.leftChildIndex(i)
		leftChildValue := h.impl[leftChildIndex]
		if leftChildValue < currValue {
			smallestIndex = leftChildIndex
			smallestValue = leftChildValue
		}
	}

	if h.hasRightChildIndex(i) {
		rightChildIndex := h.rightChildIndex(i)
		rightChildValue := h.impl[rightChildIndex]
		if rightChildValue < smallestValue {
			smallestIndex = rightChildIndex
			smallestValue = rightChildValue
		}
	}

	if smallestIndex == i {
		return
	}

	h.swap(i, smallestIndex)
	h.downHeapify(smallestIndex)
}

func (h *Heap) upHeapify(i int) {
	if i == 0 {
		return
	}

	curr := h.impl[i]
	par := h.impl[h.parentIndex(i)]

	if curr < par {
		h.swap(i, h.parentIndex(i))
	}

	h.upHeapify(h.parentIndex(i))
}

func (h *Heap) swap(i, j int) {
	temp := h.impl[i]
	h.impl[i] = h.impl[j]
	h.impl[j] = temp
}

func main() {
	h := NewHeap()
	h.Insert(6)
	h.Insert(5)
	h.Insert(3)
	h.Insert(7)
	h.Insert(2)
	h.Insert(8)
	h.Insert(11)
	h.Insert(15)
	h.Insert(0)

	for {
		el, ok := h.GetRoot()
		if !ok {
			return
		}
		fmt.Printf("=%#v\n", el)
	}
}
