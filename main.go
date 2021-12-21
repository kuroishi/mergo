package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
)

func sort(l []int, r []int) (ret []int) {
	llen := len(l)
	rlen := len(r)
	for i, lidx, ridx := 0, 0, 0; i < llen + rlen; i++ {
		if lidx >= llen {
			ret = append(ret, r[ridx])
			ridx++
			continue
		}
		if ridx >= rlen {
			ret = append(ret, l[lidx])
			lidx++
			continue
		}
		if l[lidx] <= r[ridx] {
			ret = append(ret, l[lidx])
			lidx++
		} else {
			ret = append(ret, r[ridx])
			ridx++
		}
	}
	return ret
}

func main() {
	var nums []int
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		n, _ := strconv.Atoi(sc.Text())
		nums = append(nums, n)
	}
	
	var sorted [][]int
	for _, v := range nums {
		sorted = append(sorted, []int{v})
	}

	for len(sorted) != 1 {
		sorted = append(sorted, sort(sorted[0], sorted[1]))
		sorted = sorted[2:]
	}

	for idx, _ := range sorted[0] {
		fmt.Print(sorted[0][idx], "\n")
	}
}
