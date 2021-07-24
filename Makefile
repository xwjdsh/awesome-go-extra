REFRESH ?= 0

.PHONY: all
all: build

.PHONY: build
build:
	go run ./cmd/extra-gen -tmpl ./extra-md.tmpl -order star -w ./README.md --refresh-cache=$(REFRESH)
	go run ./cmd/extra-gen -tmpl ./extra-md.tmpl -order created -w ./README-created.md --refresh-cache=0
	go run ./cmd/extra-gen -tmpl ./extra-md.tmpl -order pushed -w ./README-pushed.md --refresh-cache=0
