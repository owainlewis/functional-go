all: fmt test bench

.PHONY: fmt
fmt:
	gofmt -w .

.PHONY: test
test:
	go test ./...

.PHONY: bench
bench:
	go test -bench=. -benchmem ./...
