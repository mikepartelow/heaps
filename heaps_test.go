package heaps_test

import (
	"container/heap"
	"math/rand"
	"mp/heaps"
	"sort"
	"testing"

	"golang.org/x/exp/constraints"
)

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

func TestThingsSort(t *testing.T) {
	got := Things{
		&Thing{"thing 3", 3},
		&Thing{"thing 1", 1},
		&Thing{"thing 2", 2},
	}
	want := Things{
		&Thing{"thing 1", 1},
		&Thing{"thing 2", 2},
		&Thing{"thing 3", 3},
	}
	sort.Sort(&got)

	for i := range got {
		if got[i].Value != want[i].Value {
			t.Errorf("got %+v, wanted %+v", got[i], want[i])
		}
	}
}

func cmp[T constraints.Ordered](a, b T) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func cmpThings(a, b Thing) int {
	return cmp(a.Value, b.Value)
}

func cmpThingPtrs(a, b *Thing) int {
	return cmp(a.Value, b.Value)
}

func TestThingsHeap(t *testing.T) {
	t.Run("basics", func(t *testing.T) {
		pq := Things{
			&Thing{"thing 3", 3},
			&Thing{"thing 1", 1},
			&Thing{"thing 2", 2},
		}

		heap.Init(&pq)
		assertAnyThing(t, heap.Pop(&pq), 1)
		assertAnyThing(t, heap.Pop(&pq), 2)

		heap.Push(&pq, &Thing{"thing 1", 1})

		assertAnyThing(t, heap.Pop(&pq), 1)
		assertAnyThing(t, heap.Pop(&pq), 3)
	})

	t.Run("push/pop", func(t *testing.T) {
		const n = 1000
		pq := Things{}

		pushed := make(Things, n)

		for i := 0; i < n; i++ {
			pushed[i] = &Thing{"x", rand.Intn(100)}
			heap.Push(&pq, pushed[i])

			s := pushed[0 : i+1]
			sort.Sort(&s)

			assertAnyThing(t, heap.Pop(&pq), pushed[0].Value)
			heap.Push(&pq, pushed[0])
		}

		sort.Sort(&pushed)

		for i := 0; i < n; i++ {
			assertAnyThing(t, heap.Pop(&pq), pushed[i].Value)
		}
	})
}

func BenchmarkThingsHeap(b *testing.B) {
	pq := Things{}
	heap.Init(&pq) // just for giggles
	for i := 0; i < b.N; i++ {
		heap.Push(&pq, &Thing{"something", i})
	}
	for i := 0; i < b.N; i++ {
		j := heap.Pop(&pq).(*Thing)
		if j.Value != i {
			b.Errorf("expected %d, got %d", i, j.Value)
		}
	}
}

func TestGenericHeap(t *testing.T) {
	t.Run("ints", func(t *testing.T) {
		pq := heaps.NewHeap([]int{3, 2, 1}, cmp[int])

		assertEqual(t, pq.Pop(), 1)
		assertEqual(t, pq.Pop(), 2)

		pq.Push(1)

		assertEqual(t, pq.Pop(), 1)
		assertEqual(t, pq.Pop(), 3)
	})

	t.Run("float64s", func(t *testing.T) {
		pq := heaps.NewHeap([]float64{3.3, 2.2, 1.1}, cmp[float64])

		assertEqual(t, pq.Pop(), 1.1)
		assertEqual(t, pq.Pop(), 2.2)

		pq.Push(1.1)

		assertEqual(t, pq.Pop(), 1.1)
		assertEqual(t, pq.Pop(), 3.3)
	})

	t.Run("things", func(t *testing.T) {
		pq := heaps.NewHeap([]Thing{
			{"thing 3", 3},
			{"thing 2", 2},
			{"thing 1", 1},
		}, cmpThings)

		assertEqual(t, pq.Pop(), Thing{"thing 1", 1})
		assertEqual(t, pq.Pop(), Thing{"thing 2", 2})

		pq.Push(Thing{"thing 1", 1})

		assertEqual(t, pq.Pop(), Thing{"thing 1", 1})
		assertEqual(t, pq.Pop(), Thing{"thing 3", 3})
	})

	t.Run("push/pop", func(t *testing.T) {
		const n = 1000
		pq := heaps.NewHeap([]int{}, cmp[int])
		pushed := make([]int, n)

		for i := 0; i < n; i++ {
			pushed[i] = rand.Int()
			pq.Push(pushed[i])

			sort.Ints(pushed[0 : i+1])

			assertEqual(t, pq.Pop(), pushed[0])
			pq.Push(pushed[0])
		}

		sort.Ints(pushed)

		for i := 0; i < n; i++ {
			assertEqual(t, pq.Pop(), pushed[i])
		}
	})
}

func BenchmarkGenericHeapOfThingPtrs(b *testing.B) {
	pq := heaps.NewHeap([]*Thing{}, cmpThingPtrs)
	for i := 0; i < b.N; i++ {
		pq.Push(&Thing{"something", i})
	}
	for i := 0; i < b.N; i++ {
		j := pq.Pop()
		if j.Value != i {
			b.Errorf("expected %d, got %d", i, j.Value)
		}
	}
}

func assertAnyThing(t testing.TB, thing any, want int) {
	t.Helper()
	got := thing.(*Thing).Value
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func assertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
