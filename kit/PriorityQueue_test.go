package kit

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PriorityQueue(t *testing.T) {
	ast := assert.New(t)

	// Some items and their priorities.
	items := map[string]int{
		"banana": 2, "apple": 1, "pear": 3,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Item{
		value:    "orange",
		priority: 5,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 0)

	// Some items and their priorities.
	expected := []string{
		"orange",
		"apple",
		"banana",
		"pear",
	}

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		ast.Equal(expected[item.priority], item.value)
	}
}
