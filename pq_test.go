package PriQ

import (
	"log"
	"testing"
)

func init() {}

// 857h Minimum Cost to Hire K Workders
func Test857(t *testing.T) {
	log.Print("105 ?= ", mincostToHireWorkers([]int{10, 20, 5}, []int{70, 50, 30}, 2))
	log.Print("30.6 ?= ", mincostToHireWorkers([]int{3, 1, 10, 10, 1}, []int{4, 8, 2, 2, 7}, 3))
}
