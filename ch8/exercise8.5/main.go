package main

import (
	"fmt"
	"github.com/AmbitionLover/gopl/ch8/exercise8.5/mandelbrot"
	"image/png"
	"os"
	"runtime"
)

func main() {
	// NumCPU returns the number of logical CPUs usable by the current process.
	workers := runtime.NumCPU()
	fmt.Printf("number of cpu : %d \n", workers)
	img := mandelbrot.ConcurrentRender(workers)
	png.Encode(os.Stdout, img)
}
