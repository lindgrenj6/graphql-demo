all: run

run:
	go run *.go

generate:
	go generate ./...
