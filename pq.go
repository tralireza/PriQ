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

type E786 struct {
	Ratio float64
	l, r  int
}
type PQ786 []E786

func (p PQ786) Len() int           { return len(p) }
func (p PQ786) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PQ786) Less(i, j int) bool { return p[i].Ratio < p[j].Ratio }
func (p *PQ786) Push(x any)        { *p = append(*p, x.(E786)) }
func (p *PQ786) Pop() any {
	v := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return v
}

// 786m K-th Smallest Prime Fraction
func kthSmallestPrimeFraction(arr []int, k int) []int {
	return []int{}
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
