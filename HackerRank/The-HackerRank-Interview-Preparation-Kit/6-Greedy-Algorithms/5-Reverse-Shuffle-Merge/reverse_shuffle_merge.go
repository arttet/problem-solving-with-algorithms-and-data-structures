package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverseShuffleMerge(s string) string {
	return s
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

	reader := bufio.NewReaderSize(stdin, 1024*1024)
	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var s string
	_, err = fmt.Fscan(reader, &s)
	checkError(err)

	result := reverseShuffleMerge(s)
	fmt.Fprint(writer, result)
	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
