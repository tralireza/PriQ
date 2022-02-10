package PriQ

import (
	"container/heap"
	"log"
	"testing"
)

func init() {}

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
