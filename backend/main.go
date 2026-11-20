package main

import (
	"runtime"

	"github.com/donomii/filebrowser-1/backend/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
