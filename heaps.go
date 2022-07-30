package heaps

type Thing struct {
	Name  string
	Value int
}

type Things []*Thing

// sort.Interface
func (t *Things) Len() int           { return len(*t) }
func (t *Things) Less(i, j int) bool { return (*t)[i].Value < (*t)[j].Value }
func (t *Things) Swap(i, j int)      { (*t)[i], (*t)[j] = (*t)[j], (*t)[i] }

// heap.Interface
func (t *Things) Push(x any) {
	*t = append(*t, x.(*Thing))
}

func (t *Things) Pop() any {
	n := len(*t)
	thing := (*t)[n-1]
	*t = (*t)[0 : n-1]
	return thing
}
