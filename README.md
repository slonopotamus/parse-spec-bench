Benchmark for https://github.com/containerd/containerd/pull/8780

```
$ go test -v -bench=.
goos: linux
goarch: amd64
pkg: parse_spec_bench
cpu: AMD Ryzen 7 3700X 8-Core Processor             
BenchmarkParseFullSpec
BenchmarkParseFullSpec-16                   6661            311302 ns/op
BenchmarkParseAnnotations
BenchmarkParseAnnotations-16                7191            139453 ns/op
BenchmarkParseHooks
BenchmarkParseHooks-16                      7479            138041 ns/op
PASS
ok      parse_spec_bench        4.170s
```
