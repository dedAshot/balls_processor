package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	in := Read(input)
	valid := Validate(in)

	if valid {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func Read(in io.Reader) [][]int {
	var N int
	fmt.Fscan(in, &N)

	containers := make([][]int, N)

	for i := 0; i < N; i++ {
		container := make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Fscan(in, &container[j])
		}
		containers[i] = container
	}

	return containers
}

func Validate(in [][]int) bool {
	uniqueCountMap := make(map[int]int64)
	// means Effective capacity, aka exchange points
	// dont use struct{} because we could have 2 containers with the same capacity
	capacityMap := make(map[int]int)
	for _, container := range in {
		cap := int64(0)
		for ballType, v := range container {
			cap += int64(v)
			uniqueCountMap[ballType] += int64(v)
		}
		capacityMap[int(cap)]++
	}

	// check
	for ballType, count := range uniqueCountMap {
		_ = ballType

		v, ok := capacityMap[int(count)]
		if ok {
			if v > 0 {
				capacityMap[int(count)]--
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
