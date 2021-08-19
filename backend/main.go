package main

import (
	"runtime"

	"github.com/donomii/filebrowser/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
