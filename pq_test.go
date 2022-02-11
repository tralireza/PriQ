package PriQ

import (
	"container/heap"
	"log"
	"slices"
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
