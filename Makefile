REFRESH ?= 0

.PHONY: all
all: build

.PHONY: build
build:
	go run ./cmd/extra-gen -tmpl ./extra-md.tmpl -w ./README.md --refresh-cache=$(REFRESH)
