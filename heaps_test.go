package heaps_test

import (
	"container/heap"
	"mp/heaps"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	got := heaps.Things{
		&heaps.Thing{"thing 3", 3},
		&heaps.Thing{"thing 1", 1},
		&heaps.Thing{"thing 2", 2},
	}
	want := heaps.Things{
		&heaps.Thing{"thing 1", 1},
		&heaps.Thing{"thing 2", 2},
		&heaps.Thing{"thing 3", 3},
	}
	sort.Sort(&got)

	for i, _ := range got {
		if got[i].Value != want[i].Value {
			t.Errorf("got %+v, wanted %+v", got[i], want[i])
		}
	}
}

func TestHeap(t *testing.T) {
	pq := heaps.Things{
		&heaps.Thing{"thing 3", 3},
		&heaps.Thing{"thing 1", 1},
		&heaps.Thing{"thing 2", 2},
	}

	heap.Init(&pq)
	assertThing(t, heap.Pop(&pq), 1)
	assertThing(t, heap.Pop(&pq), 2)

	heap.Push(&pq, &heaps.Thing{"thing 1", 1})

	assertThing(t, heap.Pop(&pq), 1)
	assertThing(t, heap.Pop(&pq), 3)
}

func assertThing(t testing.TB, thing any, want int) {
	t.Helper()
	got := thing.(*heaps.Thing).Value
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
