test: #чтобы проверить что ничего не сломалось
	go test -v

bench:
	go test -bench . -benchmem

prof:
	go tool pprof -http=:8083 . .