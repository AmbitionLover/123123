package pipeline

import (
	"fmt"
	"testing"
)

func TestPipeline(t *testing.T) {
	in, out := pipeline(3)
	in <- 1
	// 这个值用 3 个 goroutine 传送经过了 3 个 channel
	// （不包括主 goroutine 和 in out channel）
	fmt.Println(<-out)
}

func bench(b *testing.B, stages int) {
	in, out := pipeline(stages)
	for i := 0; i < b.N; i++ {
		go func() {
			in <- 1
		}()
		<-out
	}
	close(in)
}

func BenchmarkPipeline1(b *testing.B) {
	bench(b, 1)
}

func BenchmarkPipeline1024(b *testing.B) {
	bench(b, 1024)
}

func BenchmarkPipeline1048576(b *testing.B) {
	bench(b, 1048576)
}

// 5G
func BenchmarkPipeline2097152(b *testing.B) {
	bench(b, 2097152)
}

//// 10G
//func BenchmarkPipeline4194304(b *testing.B) {
//	bench(b, 4194304)
//}

// 传输这个值所花的时间
/*
goos: windows
goarch: amd64
pkg: github.com/AmbitionLover/gopl/ch9/exercise9.4
BenchmarkPipeline1-6         	 1336902	       911 ns/op
BenchmarkPipeline1024-6      	    3342	    386760 ns/op
BenchmarkPipeline1048576-6   	       1	5684786500 ns/op
BenchmarkPipeline2097152-6   	       1	7357312000 ns/op
*/
