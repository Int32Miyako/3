test: #чтобы проверить что ничего не сломалось
	go test -v

# hw3.test
testWithC:
	go test -c
bench:
	go test -bench . -benchmem

prof:
	go tool pprof -http=:8083 ./prof/bin/ ./prof/out/
# go tool pprof -http=:8083 /path/to/bin /path/to/out

getBinaryCpuProfile:
	go test -bench -benchmem -cpuprofile=prof/out/cpuprofile.out

getBinaryMemProfile:
	go test -bench -benchmem -memprofile=prof/out/memprofile.out



makeBinaryAndPprof:
	go test -bench . -benchmem -cpuprofile=./prof/out/cpu.out -memprofile=./prof/out/mem.out -memprofilerate=1 main_test.go
	go tool pprof -http=:8080 ./prof/out/cpu.out
