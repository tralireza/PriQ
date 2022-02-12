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
