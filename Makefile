
.PHONY: run
run/%:
	go run main.go -file $*

.PHONY: test
test:
	go test -race -cover  ./...