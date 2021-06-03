.PHONY: all
all: build

.PHONY: build
build:
	go run ./cmd/extra-gen > README.md
