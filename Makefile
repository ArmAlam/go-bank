build:
	@go build
run: build
	@./gobank
test:
	@go test -v ./...