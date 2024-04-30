build:
	@go build -o bin/twitter main.go

run: build
	@./bin/twitter

test:
	@go test -v ./...

mod:
	@go mod tidy
	