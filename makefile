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


#Вызывать по порядку, чтобы обновились оба профиля и потом можно было смотреть их вместе
#1
makeBinaryAndPprof:
	go test -bench -benchmem -cpuprofile=./prof/out/cpuprofile.out -memprofile=./prof/out/memprofile.out -memprofilerate=1

#2
MakeProfsCPU:
	go tool pprof -http=:8084 ./prof/out/cpuprofile.out

#3
MakeProfsMem:
	go tool pprof -http=:8085 ./prof/out/memprofile.out

# по идее есть возможность вызвать меню pprof и там вызвать комманду web