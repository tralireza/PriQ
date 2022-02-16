package PriQ

import (
	"container/heap"
	"log"
	"slices"
	"sort"
	"strconv"
	"testing"
)

func init() {}

// 215m Kth Largest Element in an Array
func Test215(t *testing.T) {
	log.Print("5 ?= ", findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	log.Print("4 ?= ", findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}

// 502h IPO
func Test502(t *testing.T) {
	log.Print("4 ?= ", findMaximizedCapital(2, 0, []int{1, 2, 3}, []int{0, 1, 1}))
	log.Print("6 ?= ", findMaximizedCapital(3, 0, []int{1, 2, 3}, []int{0, 1, 1}))
}

// 506 Relative Ranks
func Test506(t *testing.T) {
	IndexMap := func(score []int) []string {
		Rank := make([]string, len(score))

		I := map[int]int{}
		for i, v := range score {
			I[v] = i
		}
		slices.SortFunc(score, func(a, b int) int { return b - a })

		for i, v := range score {
			var rank string
			switch i {
			case 0:
				rank = "Gold Medal"
			case 1:
				rank = "Silver Medal"
			case 2:
				rank = "Bronze Medal"
			default:
				rank = strconv.Itoa(i + 1)
			}
			Rank[I[v]] = rank
		}

		return Rank
	}

	for _, f := range []func([]int) []string{findRelativeRanks, IndexMap} {
		log.Printf(" ?= %q", f([]int{5, 4, 3, 2, 1}))
		log.Printf(" ?= %q", f([]int{10, 3, 8, 9, 4}))
	}
}

// 786m K-th Smallest Prime Fraction
func Test786(t *testing.T) {
	On2 := func(nums []int, k int) []int {
		Q := PQ786{}
		for r := len(nums) - 1; r > 0; r-- {
			for l := 0; l < r; l++ {
				Q = append(Q, E786{float64(nums[l]) / float64(nums[r]), l, r})
			}
		}
		heap.Init(&Q)

		for range k - 1 {
			log.Print(" -> ", heap.Pop(&Q))
		}
		return []int{nums[Q[0].n], nums[Q[0].d]}
	}

	for _, f := range []func([]int, int) []int{kthSmallestPrimeFraction, On2} {
		log.Print("[2 5] ?= ", f([]int{1, 2, 3, 5}, 3))
		log.Print("[1 7] ?= ", f([]int{1, 7}, 1))
	}
}

// 857h Minimum Cost to Hire K Workers
func Test857(t *testing.T) {
	log.Print("105 ?= ", mincostToHireWorkers([]int{10, 20, 5}, []int{70, 50, 30}, 2))
	log.Print("30.6 ?= ", mincostToHireWorkers([]int{3, 1, 10, 10, 1}, []int{4, 8, 2, 2, 7}, 3))
}

// 1438m Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit
func Test1438(t *testing.T) {
	WithDeque := func(nums []int, limit int) int {
		dM, dX := []int{}, []int{}

		ls := 0
		l := 0
		for r := range nums {
			for len(dM) > 0 && dM[len(dM)-1] > nums[r] {
				dM = dM[:len(dM)-1]
			}
			dM = append(dM, nums[r])

			for len(dX) > 0 && dX[len(dX)-1] < nums[r] {
				dX = dX[:len(dX)-1]
			}
			dX = append(dX, nums[r])

			for dX[0]-dM[0] > limit {
				if dX[0] == nums[l] {
					dX = dX[1:]
				}
				if dM[0] == nums[l] {
					dM = dM[1:]
				}
				l++
			}

			ls = max(r-l+1, ls)
		}

		return ls
	}

	for _, f := range []func([]int, int) int{longestSubarray, WithDeque} {
		log.Print("==")
		log.Print("2 ?= ", f([]int{8, 2, 4, 7}, 4))
		log.Print("4 ?= ", f([]int{10, 1, 2, 4, 7, 2}, 5))
		log.Print("3 ?= ", f([]int{4, 2, 2, 2, 4, 4, 2, 2}, 0))
	}
}

// 3075m Maximum Happiness of Selected Children
func Test3075(t *testing.T) {
	Greedy := func(happiness []int, k int) int64 {
		sort.Sort(sort.Reverse(sort.IntSlice(happiness)))

		hSum := int64(0)
		for i := range k {
			if happiness[i]-i <= 0 {
				return hSum
			}
			hSum += int64(happiness[i] - i)
		}
		return hSum
	}

	for _, f := range []func([]int, int) int64{Greedy, maximumHappinessSum} {
		log.Print("==")
		log.Print("4 ?= ", f([]int{1, 2, 3}, 2))
		log.Print("1 ?= ", f([]int{1, 1, 1, 1}, 2))
		log.Print("5 ?= ", f([]int{2, 3, 4, 5}, 1))
	}
}
