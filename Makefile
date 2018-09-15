.PHONY: install
install:
	go get "github.com/erdemtoraman/opening-hours"
	go get "github.com/stretchr/testify/assert"

.PHONY: run
run/%:
	go run main.go -file $*

.PHONY: test
test:
	go test -race -cover  ./...