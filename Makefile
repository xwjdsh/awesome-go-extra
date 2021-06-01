.PHONY: all
all: build

.PHONY: build
build:
	go run ./cmd/extra-gen -cd=24h > README.md
