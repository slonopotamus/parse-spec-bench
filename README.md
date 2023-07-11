Benchmark for https://github.com/containerd/containerd/pull/8780

```
$ go test -v -bench=.
goos: linux
goarch: amd64
pkg: parse_spec_bench
cpu: AMD Ryzen 7 3700X 8-Core Processor             
BenchmarkParseFullSpec
BenchmarkParseFullSpec-16                   5930            318176 ns/op
BenchmarkParseAnnotations
BenchmarkParseAnnotations-16                7147            149987 ns/op
PASS
ok      parse_spec_bench        3.004s
```
