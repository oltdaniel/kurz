# Base
run: util.templates util.assets
	go run main.go

build:
	go build main.go

build.run: build
	./main

# Utilities
UTIL_DIR := utils
ROOT := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
HOME := /home/daniel
GO_DIR := $(HOME)/go
NGINX_DIR := /etc/nginx

util.install:
	sh $(UTIL_DIR)/install.sh

util.update:
	sh $(UTIL_DIR)/update.sh

util.templates:
	sh $(UTIL_DIR)/templates.sh

util.assets:
	sh $(UTIL_DIR)/assets.sh $(GO_DIR)/bin $(ROOT)

link:
	rm -f $(GO_DIR)/src/kurz
	ln -sf $(ROOT) $(GO_DIR)/src/kurz

nginx:
	sudo rm -f $(NGINX_DIR)/sites-enabled/kurz
	sudo ln -sf $(ROOT)/nginx $(NGINX_DIR)/sites-enabled/kurz
