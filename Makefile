# Base
run: util.templates
	go run main.go

build:
	go build main.go

build.run: build
	./main

# Utilities
UTIL_DIR := utils
ROOT := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
GO_DIR := /home/daniel/go

util.install:
	sh $(UTIL_DIR)/install.sh

util.update:
	sh $(UTIL_DIR)/update.sh

util.templates:
	sh $(UTIL_DIR)/templates.sh

link:
	rm -f $(GO_DIR)/src/kurz
	ln -sf $(ROOT) $(GO_DIR)/src/kurz
