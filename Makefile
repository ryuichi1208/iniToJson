src = ./src/main.go

.PHONY: build
build:
	go build $(src)

clean:
	rm -f main
