# golang-benchmark

Simple benchmark on Go

Using:

```
go test -bench=.

```
Example on raspberry pi 4 with Ubuntu ARM64
```
go test -bench=.
goos: linux
goarch: arm64
BenchmarkBubble-4       18673741                59.5 ns/op
BenchmarkPrime-4        147660474                8.08 ns/op
PASS
ok      _/root/golang/golang-benchmark  3.209s
```