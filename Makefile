build:
	CGO_ENABLED=0 go build -o brander ./cmd/brander/

run: build
	./brander

test:
	go test ./...

clean:
	rm -f brander

.PHONY: build run test clean
