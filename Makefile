build:
	go build -o bin/hacklang cmd/main.go

test:
	go test ./...