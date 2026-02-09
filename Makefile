run-debug:
	go run main.go -rod=show,slow=1s,trace

run:
	go run main.go

build:
	go build -o main main.go

run-build:
	./main

.PHONY: run build run-build run-debug