package main

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input  string
	output string
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	tests := []testCase{
		{"input/input00.txt", "output/output00.txt"},
		{"input/input01.txt", "output/output01.txt"},
		{"input/input02.txt", "output/output02.txt"},
	}

	for i := range tests {
		func() {
			stdin, err := os.Open(tests[i].input)
			checkError(err)
			defer stdin.Close()

			stdout, err := ioutil.TempFile("", "output.*.txt")
			checkError(err)
			defer os.Remove(stdout.Name())

			os.Stdin, os.Stdout = stdin, stdout

			main()

			content, err := ioutil.ReadFile(tests[i].output)
			checkError(err)
			expected := strings.TrimSpace(string(content))

			content, err = ioutil.ReadFile(stdout.Name())
			checkError(err)
			output := strings.TrimSpace(string(content))

			content, err = ioutil.ReadFile(tests[i].input)
			checkError(err)
			input := strings.TrimSpace(string(content))

			ast.Equal(expected, output, "Test Case: %v", input)
		}()
	}
}

func TestPanic(t *testing.T) {
	assert.Panics(t, func() { checkError(errors.New("")) }, "The code did not panic")
}
