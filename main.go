package main

import (
	"github.com/aksbuzz/mood-journal/cmd"
	_ "modernc.org/sqlite"
)

func main() {
	cmd.Execute()
}
