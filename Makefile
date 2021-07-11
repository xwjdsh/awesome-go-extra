.PHONY: all
all: build

.PHONY: build
build:
	go run ./cmd/extra-gen --refresh-cache > README.md
