package day01

import (
	"github.com/maxslarsson/adventOfCodeTemplate/input"
	"strconv"
)

func A(input *input.Input) string {
	nums := input.LineSliceOfInts()
	for _, num1 := range nums {
		for _, num2 := range nums {
			if num1 + num2 == 2020 {
				return strconv.Itoa(num1 * num2)
			}
		}
	}
	panic("No solution found")
}