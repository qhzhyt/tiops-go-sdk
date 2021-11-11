package utils

type Comparable interface {
	CompareTo(other Comparable) int
}

type Heap struct {
	list []Comparable
	size int
	cap  int
}

func NewHeap(cap int) *Heap {
	return &Heap{cap: cap, list: make([]Comparable, cap)}
}

func (h *Heap) swap(i, j int) {
	t := h.list[i]
	h.list[i] = h.list[j]
	h.list[j] = t
}

func (h *Heap) up(index int) {
	for index > 0 {
		pidx := (index - 1) / 2
		if h.list[index].CompareTo(h.list[pidx]) < 0 {
			h.swap(index, pidx)
			index = pidx
		} else {
			break
		}
	}
}

func (h *Heap) down(index int) {
	for mid := index*2 + 1; mid < h.size-1; mid = index * 2 + 1 {

		if h.list[mid].CompareTo(h.list[index*2+2]) > 0 {
			mid = index*2 + 2
		}
		if h.list[index].CompareTo(h.list[mid]) > 0 {
			h.swap(index, mid)
			index = mid
		} else {
			break
		}
	}
}

func (h Heap) Size() int {
	return h.size
}

func (h *Heap) Offer(comparable Comparable) {
	if h.size >= h.cap {
		return
	}
	h.list[h.size] = comparable
	h.size++
	h.up(h.size - 1)
}

func (h *Heap) Poll() Comparable {
	res := h.list[0]
	h.list[0] = h.list[h.size-1]
	h.size--
	h.down(0)
	return res
}
