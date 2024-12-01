build:
	go build -o aoc ./cmd/aoc

install: build
	sudo mv aoc /usr/local/bin/

clean:
	rm -f aoc
