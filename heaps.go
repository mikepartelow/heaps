package heaps

type Heap[T any] struct {
	items []T
	cmp   func(a, b T) int
}

func NewHeap[T any](items []T, cmp func(a, b T) int) *Heap[T] {
	h := Heap[T]{items: items, cmp: cmp}
	min_heapify(h.items, 0, h.cmp)
	return &h
}

func (h *Heap[T]) Push(item T) {
	parent := func(i int) int {
		return (i - 1) / 2
	}

	h.items = append(h.items, item)

	for i := len(h.items) - 1; h.cmp(h.items[parent(i)], h.items[i]) == 1; i = parent(i) {
		h.items[i], h.items[parent(i)] = h.items[parent(i)], h.items[i]
	}
}

func (h *Heap[T]) Pop() T {
	x := h.items[0]

	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	min_heapify(h.items, 0, h.cmp)

	return x
}

func min_heapify[T any](arr []T, idx int, cmp func(a, b T) int) {
	left := 2*idx + 1
	right := 2*idx + 2
	smallest := idx

	if left < len(arr) && cmp(arr[left], arr[smallest]) == -1 {
		smallest = left
	}

	if right < len(arr) && cmp(arr[right], arr[smallest]) == -1 {
		smallest = right
	}

	if smallest != idx {
		arr[idx], arr[smallest] = arr[smallest], arr[idx]
		min_heapify(arr, smallest, cmp)
	}
}
