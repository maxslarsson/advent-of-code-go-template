package day01

import (
	"github.com/maxslarsson/adventOfCodeTemplate/input"
	"strconv"
	"testing"
)


func TestExample(t *testing.T) {
	functionAnswer := A(input.FromString("1721\n979\n366\n299\n675\n1456"))
	correctAnswer := 514579
	if functionAnswer != strconv.Itoa(correctAnswer) {
		t.Errorf("Got answer %s, correct answer is %d", functionAnswer, correctAnswer)
	}
}