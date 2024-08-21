package structure

type MaxHeap[T Comparable] struct {
	slice    []T
	heapSize int
	indices  map[int]int
}

type Comparable interface {
	Id() int
	LargerThan(interface{}) bool
}

func (h *MaxHeap[T]) Init(slice []T) {
	if slice == nil {
		slice = make([]T, 0)
	}

	h.slice = slice
	h.heapSize = len(slice)
	h.indices = map[int]int{}
	for i, v := range h.slice {
		h.indices[v.Id()] = i
	}

	for i := h.heapSize / 2; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

func (h MaxHeap[T]) Size() int {
	return h.heapSize
}

func (h *MaxHeap[T]) Update(item T) {
	h.slice[h.indices[item.Id()]] = item
	h.heapifyUp(h.indices[item.Id()])
	h.heapifyDown(h.indices[item.Id()])
}

func (h *MaxHeap[T]) Push(item T) {
	h.slice = append(h.slice, item)
	h.heapSize++
	h.indices[item.Id()] = h.heapSize - 1
	h.heapifyUp(h.heapSize - 1)
}

/* Pop from heap root (index: 0) */
func (h *MaxHeap[T]) Pop() *T {
	if h.heapSize == 0 {
		return nil
	}
	c := h.slice[0]
	h.slice[0] = h.slice[h.heapSize-1]
	h.heapSize--
	h.indices[h.slice[0].Id()] = 0
	h.heapifyDown(0)
	h.slice = h.slice[:h.heapSize]
	return &c
}

func (h *MaxHeap[T]) heapifyUp(i int) {
	if i == 0 {
		return
	}
	p := (i - 1) / 2
	if h.slice[i].LargerThan(h.slice[p]) {
		h.slice[p], h.slice[i] = h.slice[i], h.slice[p]
		h.indices[h.slice[i].Id()] = i
		h.indices[h.slice[p].Id()] = p
		h.heapifyUp(p)
	}
}

func (h *MaxHeap[T]) heapifyDown(i int) {
	l := 2*i + 1
	if l >= h.heapSize {
		return
	}
	r := 2*i + 2
	max := l
	if r < h.heapSize && h.slice[r].LargerThan(h.slice[l]) {
		max = r
	}
	if h.slice[max].LargerThan(h.slice[i]) {
		h.slice[max], h.slice[i] = h.slice[i], h.slice[max]
		h.indices[h.slice[i].Id()] = i
		h.indices[h.slice[max].Id()] = max
		h.heapifyDown(max)
	}
}
