.PHONY: run
run:
	go run main.go

.PHONY: lint
lint:
	golangci-lint run .

.PHONY: test
test:
	go test -v ./...

.PHONY: bench
bench:
	go test -bench=. -benchmem
