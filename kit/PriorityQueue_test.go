package kit

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_priorityQueue(t *testing.T) {
	ast := assert.New(t)

	// Some items and their priorities.
	items := map[string]int{
		"banana": 2, "apple": 1, "pear": 3,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(priorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	it := &item{
		value:    "orange",
		priority: 5,
	}
	heap.Push(&pq, it)
	pq.update(it, it.value, 0)

	// Some items and their priorities.
	expected := []string{
		"orange",
		"apple",
		"banana",
		"pear",
	}

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*item)
		ast.Equal(expected[it.priority], it.value)
	}
}
