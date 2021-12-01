package main

import (
	"embed"
	"github.com/maxslarsson/adventOfCodeTemplate/cmd"
)

//go:embed days/*/input.txt
var inputs embed.FS

func main() {
	cmd.Execute(inputs)
}