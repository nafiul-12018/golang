package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	Value    string
	Priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Higher priority items should come first
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	pq := make(PriorityQueue, 0)

	heap.Push(&pq, &Item{Value: "A", Priority: 3})
	heap.Push(&pq, &Item{Value: "B", Priority: 1})
	heap.Push(&pq, &Item{Value: "C", Priority: 2})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%s with priority %d\n", item.Value, item.Priority)
	}
}
