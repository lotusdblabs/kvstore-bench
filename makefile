build:
	go mod download
	go build -o kv-bench cmd/kv-bench/main.go
