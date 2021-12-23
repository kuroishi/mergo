package main

import (
	"testing"
	"bufio"
	"os"
	"strconv"
	"log"
)

func BenchmarkGosort(b *testing.B) {
	var nums []int	
	f, err := os.Open("./rnd.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		n, _ := strconv.Atoi(sc.Text())
		nums = append(nums, n)
	}
	gosort(nums)
}

func BenchmarkSmergesort(b *testing.B) {
	var nums []int	
	f, err := os.Open("./rnd.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		n, _ := strconv.Atoi(sc.Text())
		nums = append(nums, n)
	}
	smergesort(nums)
}

func BenchmarkPmergesort(b *testing.B) {
	var nums []int
	f, err := os.Open("./rnd.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		n, _ := strconv.Atoi(sc.Text())
		nums = append(nums, n)
	}
	pmergesort(nums)
}
