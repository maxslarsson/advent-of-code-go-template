package input

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Input struct {
	scanner *bufio.Scanner

	lines chan string
}

func FromString(input string) *Input {
	return newInputFromReader(strings.NewReader(input), nil)
}

func newInputFromReader(r io.Reader, c io.Closer) *Input {
	result := &Input{
		scanner: bufio.NewScanner(r),
		lines:   make(chan string),
	}

	go func() {
		defer func() {
			if c != nil {
				_ = c.Close()
			}
		}()

		for result.scanner.Scan() {
			result.lines <- result.scanner.Text()
		}

		close(result.lines)
	}()

	return result
}


func (i *Input) Lines() <-chan string {
	return i.lines
}

func (i *Input) LineSlice() (result []string) {
	for line := range i.Lines() {
		result = append(result, line)
	}

	return
}


func (i *Input) LineSliceOfInts() (result []int) {
	for line := range i.Lines() {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		result = append(result, num)
	}
	return
}