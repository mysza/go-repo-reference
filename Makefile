SHELL := /bin/bash
OUTPUT := grr
GOOS := linux # this is a default, changed later
ENTRYPOINT := ./main.go

LOCAL_PACKAGES := $(shell govendor list -no-status +local)
BUILD_DIR ?= _dist

ifeq ($(OS),Windows_NT)
	OUTPUT := $(OUTPUT).exe
	GOOS := windows
else
  UNAME_S := $(shell uname -s)
  ifeq ($(UNAME_S),Linux)
      GOOS := linux
  endif
  ifeq ($(UNAME_S),Darwin)
      GOOS := darwin
  endif
endif

.PHONY: all clean build test fmt vet run

default: all

all: clean fmt vet test build

clean:
	rm -rf $(OUTPUT)
	go clean ./...

build: clean fmt vet
	export GOOS=$(GOOS)
	go build -o ./$(BUILD_DIR)/$(OUTPUT) $(ENTRYPOINT)

test:
	go test -race -cover ./...

fmt:
	go fmt ./...

vet:
	go vet ./...