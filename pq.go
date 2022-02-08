package PriQ

import (
	"container/heap"
	"log"
	"slices"
	"sort"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("")
	log.Print("> Priority Queue")
}

type PQ857 struct{ sort.IntSlice }

func (pq PQ857) Less(i, j int) bool { return pq.IntSlice[i] > pq.IntSlice[j] } // Max Heap
func (PQ857) Push(any)              {}                                         // No use, only Init & Fix
func (PQ857) Pop() (_ any)          { return }                                 // ^

// 857h Minimum Cost to Hire K Workers
func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	type WQ struct{ w, q int }
	WQs := []WQ{}

	for i, w := range wage {
		WQs = append(WQs, WQ{w: w, q: quality[i]})
	}
	slices.SortFunc(WQs, func(a, b WQ) int { return a.w*b.q - b.w*a.q })
	log.Print("W/Q -> ", WQs)

	pq := PQ857{make([]int, k)}
	tQ := 0
	for i, v := range WQs[:k] {
		pq.IntSlice[i] = v.q
		tQ += v.q
	}
	heap.Init(&pq)
	log.Print("PQ -> ", pq)

	mCost := float64(tQ*WQs[k-1].w) / float64(WQs[k-1].q)
	for _, v := range WQs[k:] {
		if v.q < pq.IntSlice[0] {
			tQ -= pq.IntSlice[0] - v.q

			pq.IntSlice[0] = v.q
			heap.Fix(&pq, 0)

			mCost = min(mCost, float64(tQ*v.w)/float64(v.q))
		}
	}
	return mCost
}
