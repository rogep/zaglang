package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"zag/repl"
)

func main() {
	t := time.Now()
	fmt.Printf("Zag 0.0.1 (main, %d-%02d-%02d %02d:%02d:%02d) on %s\n",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), runtime.GOOS,
	)
	repl.Start(os.Stdin, os.Stdout)
}
