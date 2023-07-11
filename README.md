Benchmark for https://github.com/containerd/containerd/pull/8780

```
$ go test -v -bench=.
goos: linux
goarch: amd64
pkg: parse_spec_bench
cpu: AMD Ryzen 7 3700X 8-Core Processor             
BenchmarkParseFullSpec
BenchmarkParseFullSpec-16                   6446            310789 ns/op
BenchmarkParseAnnotations
BenchmarkParseAnnotations-16                8091            144667 ns/op
PASS
ok      parse_spec_bench        3.218s
```
