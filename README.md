
Usage
=====

%  time jot -r $((2048))  1 $((2 * 1024 * 1024)) | ./mergo

% go test -bench=. -count=1

Example
=======
```
goos: darwin
goarch: arm64
pkg: mergo
BenchmarkGosort-8       	       1	3162644917 ns/op
BenchmarkSmergesort-8   	       1	1732673167 ns/op
BenchmarkPmergesort-8   	       3	 342439958 ns/op
PASS
ok  	mergo	8.151s
```

