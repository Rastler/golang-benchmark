package benchmark

import "github.com/shawnsmithdev/zermelo"

func Foo(bar [][]uint64) {
	sorter := zermelo.New()
	for _, x := range bar {
		sorter.Sort(x)
	}
}
