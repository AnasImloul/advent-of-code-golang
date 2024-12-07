run: build
	./aoc $(word 1, $(MAKECMDGOALS)) $(word 2, $(MAKECMDGOALS)) $(word 3, $(MAKECMDGOALS)) $(word 4, $(MAKECMDGOALS))

generate: build
	./aoc $(word 1, $(MAKECMDGOALS)) $(word 2, $(MAKECMDGOALS)) $(word 3, $(MAKECMDGOALS))

build:
	go build -o aoc ./cmd/aoc

clean:
	rm -f aoc

# Handle additional arguments so they are not treated as targets
%:
	@:
