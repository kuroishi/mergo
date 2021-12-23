package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
	"sort"
	"sync"
)

func dosort(l []int, r []int) (ret []int) {
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

func gosort(nums []int) []int {
	sort.SliceStable(
		nums,
		func(i, j int) bool {
			return nums[i] < nums[j]
		})
	return nums
}

func smergesort(nums []int) []int {
	var sorted [][]int
	for _, v := range nums {
		sorted = append(sorted, []int{v})
	}

	for len(sorted) != 1 {
		sorted = append(sorted, dosort(sorted[0], sorted[1]))
		sorted = sorted[2:]
	}

	return sorted[0]
}

func pmergesort(nums []int) []int {
	var sorted [][]int
	for _, v := range nums {
		sorted = append(sorted, []int{v})
	}

	sortedl := sorted[:len(sorted) / 2]
	sortedr := sorted[len(sorted) / 2:]
	
	sortedll := make([][]int, len(sortedl) / 2)
	sortedlr := make([][]int, len(sortedl) - (len(sortedl) / 2))
	sortedrl := make([][]int, len(sortedr) / 2)
	sortedrr := make([][]int, len(sortedr) - (len(sortedr) / 2))
	
	copy(sortedll, sortedl[:len(sortedl) / 2])
	copy(sortedlr, sortedl[len(sortedl) / 2:])
	copy(sortedrl, sortedr[:len(sortedr) / 2])
	copy(sortedrr, sortedr[len(sortedr) / 2:])

	var sortedall [][][]int
	sortedall = append(sortedall, sortedll)
	sortedall = append(sortedall, sortedlr)
	sortedall = append(sortedall, sortedrl)
	sortedall = append(sortedall, sortedrr)	
	//fmt.Println(sortedall[0])
	//fmt.Println(sortedall[1])
	//fmt.Println(sortedall[2])
	//fmt.Println(sortedall[3])	

	var wg sync.WaitGroup
	for idx := 0; idx < 4; idx++ {
		wg.Add(1)
                go func(idx int) {
		    defer wg.Done()
                    for len(sortedall[idx]) != 1 {
                        //fmt.Printf("sortedall[%d]: %v len: %d\n", idx, sortedall[idx], len(sortedall[idx]))
                        sortedall[idx] = append(sortedall[idx], dosort(sortedall[idx][0], sortedall[idx][1]))
                        sortedall[idx] = sortedall[idx][2:]
                    }
                }(idx)
	}
	wg.Wait()

       	chl := make(chan []int)
	go func(ch chan []int) {
		ch <- dosort(sortedall[0][0], sortedall[1][0])
 
	}(chl)

       	chr := make(chan []int)
	go func(ch chan []int) {
		ch <- dosort(sortedall[2][0], sortedall[3][0])	
	}(chr)
	
	return dosort(<-chl, <-chr)
}

func main() {
	var nums []int	
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		n, _ := strconv.Atoi(sc.Text())
		nums = append(nums, n)
	}

	//ret := gosort(nums)
	//ret := smergesort(nums)
	ret := pmergesort(nums)
	
	for idx, _ := range ret {
		fmt.Print(ret[idx], "\n")
	}
}
