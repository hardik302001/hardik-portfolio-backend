.PHONY: run build clean

run:
	go run main.go

build:
	go build -o bin/server main.go

clean:
	rm -rf bin/
