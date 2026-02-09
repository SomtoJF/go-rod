run:
	go run main.go

build:
	go build -o main main.go

run-build:
	./main

clean:
	rm main

.PHONY: run build run-build clean