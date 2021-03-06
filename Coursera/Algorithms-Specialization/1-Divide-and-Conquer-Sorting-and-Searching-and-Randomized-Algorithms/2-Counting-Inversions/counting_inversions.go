package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func countInversions(arr []int) int64 {
	return mergesort(arr, make([]int, len(arr)), 0, len(arr))
}

func mergesort(arr []int, temp []int, left int, right int) int64 {
	if left == right-1 {
		return 0
	}

	middle := (left + right) / 2

	var count int64
	count += mergesort(arr, temp, left, middle)
	count += mergesort(arr, temp, middle, right)
	count += merge(arr, temp, left, middle, right)

	return count
}

func merge(arr []int, temp []int, left int, middle int, right int) int64 {
	var inversions int64

	i, j := left, middle
	k := left
	for ; i < middle && j < right; k++ {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
			inversions += int64(middle - i)
		}
	}

	k += copy(temp[k:], arr[i:middle])
	copy(temp[k:], arr[j:right])
	copy(arr[left:right], temp[left:right])

	return inversions
}

func main() {
	stdin, err := os.Open(os.Getenv("INPUT_PATH"))
	if err != nil {
		stdin = os.Stdin
	}
	defer stdin.Close()

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	if err != nil {
		stdout = os.Stdout
	}
	defer stdout.Close()

	reader := bufio.NewScanner(stdin)
	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var arr []int
	for value := 0; reader.Scan(); arr = append(arr, value) {
		value, err = strconv.Atoi(reader.Text())
		checkError(err)
	}

	result := countInversions(arr)
	fmt.Fprint(writer, result)

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
