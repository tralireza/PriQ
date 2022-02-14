package PriQ

import (
	"container/heap"
	"log"
	"slices"
	"sort"
	"strconv"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("")
	log.Print("> Priority Queue")
}

// 506 Relative Ranks
func findRelativeRanks(score []int) []string {
	// score.length < 10^4
	// score[i] < 10^6
	Rank := make([]string, len(score))

	maxScore := slices.Max(score)
	iScore := make([]int, maxScore+1)
	for i, v := range score {
		iScore[v] = i + 1
	}
	log.Print(iScore)

	for r, v := 1, maxScore; v >= 0; v-- {
		if iScore[v] > 0 {
			var rank string
			switch r {
			case 1:
				rank = "Gold Medal"
			case 2:
				rank = "Silver Medal"
			case 3:
				rank = "Bronze Medal"
			default:
				rank = strconv.Itoa(r)
			}

			Rank[iScore[v]-1] = rank
			r++
		}
	}

	return Rank
}

// 215m Kth Largest Element in an Array
type PQ215 struct{ sort.IntSlice }

func (h PQ215) Less(i, j int) bool { return h.IntSlice[j] < h.IntSlice[i] } // Max Heap(PQ)
func (PQ215) Push(_ any)           {}                                       // not needed, only Init() and Pop()
func (h *PQ215) Pop() any {
	v := h.IntSlice[h.Len()-1]
	h.IntSlice = h.IntSlice[:h.Len()-1]
	return v
}

func findKthLargest(nums []int, k int) int {
	pq := PQ215{nums}
	heap.Init(&pq)

	for range k - 1 {
		heap.Pop(&pq)
	}
	return heap.Pop(&pq).(int)
}

// 502h IPO
type PQ502 []E502
type E502 struct{ p, c int }

func (h PQ502) Len() int           { return len(h) }
func (h PQ502) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h PQ502) Less(i, j int) bool { return h[i].p > h[j].p }
func (h *PQ502) Push(x any)        { *h = append(*h, x.(E502)) }
func (h *PQ502) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

type PQ502C struct{ PQ502 }

func (h PQ502C) Less(i, j int) bool { return h.PQ502[i].c < h.PQ502[j].c }

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	feedQ := PQ502C{}
	for i := 0; i < len(capital); i++ {
		feedQ.PQ502 = append(feedQ.PQ502, E502{p: profits[i], c: capital[i]})
	}
	heap.Init(&feedQ)

	pq := PQ502{}

	for feedQ.Len() > 0 && feedQ.PQ502[0].c <= w {
		pq = append(pq, heap.Pop(&feedQ).(E502))
	}
	heap.Init(&pq)

	for k > 0 && pq.Len() > 0 {
		w += heap.Pop(&pq).(E502).p
		for feedQ.Len() > 0 && feedQ.PQ502[0].c <= w {
			heap.Push(&pq, heap.Pop(&feedQ))
		}
		k--
	}

	return w
}

// 786m K-th Smallest Prime Fraction
type PQ786 []E786
type E786 struct {
	Ratio float64
	n, d  int // index of: Numerator, Denominator
}

func (p PQ786) Len() int           { return len(p) }
func (p PQ786) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PQ786) Less(i, j int) bool { return p[i].Ratio < p[j].Ratio }
func (p *PQ786) Push(x any)        { *p = append(*p, x.(E786)) }
func (p *PQ786) Pop() any {
	v := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return v
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	Q := PQ786{}
	for i := range arr[:len(arr)-1] {
		heap.Push(&Q, E786{float64(arr[i]) / float64(arr[len(arr)-1]), i, len(arr) - 1})
	}

	var e E786
	for range k {
		e = heap.Pop(&Q).(E786)
		n, d := e.n, e.d
		if d-1 > n {
			heap.Push(&Q, E786{float64(arr[n]) / float64(arr[d-1]), n, d - 1})
		}
	}

	return []int{arr[e.n], arr[e.d]}
}

// 857h Minimum Cost to Hire K Workers
type PQ857 struct{ sort.IntSlice }

func (pq PQ857) Less(i, j int) bool { return pq.IntSlice[i] > pq.IntSlice[j] } // Max Heap
func (PQ857) Push(any)              {}                                         // No use, only Init & Fix
func (PQ857) Pop() (_ any)          { return }                                 // ^

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

// 3075m Maximum Happiness of Selected Children
type PQ3075 struct{ sort.IntSlice }

func (h PQ3075) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // Max Heap
func (PQ3075) Push(any)             {}                                       // Not needed, only Pop, Init
func (h *PQ3075) Pop() any {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}

func maximumHappinessSum(happiness []int, k int) int64 {
	Q := PQ3075{IntSlice: happiness}
	heap.Init(&Q)

	hSum := int64(0)
	for i := range k {
		v := heap.Pop(&Q).(int)
		if v-i <= 0 {
			return hSum
		}
		hSum += int64(v - i)
	}
	return hSum
}
