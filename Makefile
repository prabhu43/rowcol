.PHONY: all
all: build

APP=rowcol
GLIDE_NOVENDOR=$(shell glide novendor)
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")
SRC_PACKAGES=$(shell glide novendor)
VERSION?=1.0
BUILD?=$(shell git describe --tags --always --dirty)
GLIDE:=$(shell command -v glide 2> /dev/null)
GOLINT:=$(shell command -v golint 2> /dev/null)
APP_EXECUTABLE="./out/$(APP)"
RICHGO=$(shell command -v richgo 2> /dev/null)

ifdef VERBOSE
	TESTARGS="-v"
endif

ifeq ($(RICHGO),)
	GOBIN=go
else
	GOBIN=richgo
endif

ensure-build-dir:
	mkdir -p out

update-deps:
	glide update

build-deps:
	glide install

compile: ensure-build-dir
	$(GOBIN) build -ldflags "-X main.majorVersion=$(VERSION) -X main.minorVersion=${BUILD}" -o $(APP_EXECUTABLE) ./main.go

build: build-deps compile fmt vet lint test

fmt:
	$(GOBIN) fmt $(GLIDE_NOVENDOR)

run: compile
	./out/holywells server

vet:
	$(GOBIN) vet ./...

setup:
ifeq ($(GLIDE),)
	curl https://glide.sh/get | sh
endif
ifeq ($(GOLINT),)
	$(GOBIN) get -u golang.org/x/lint/golint
endif
ifeq ($(RICHGO),)
	\$(GOBIN) get -u github.com/kyoh86/richgo
endif
	mkdir -p $(PWD)/out

lint:
	@for p in $(SRC_PACKAGES); do \
		echo "==> Linting $$p"; \
		golint $$p | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; } \
	done

test: ensure-build-dir
	ENVIRONMENT=test $(GOBIN) test $(SRC_PACKAGES) -p=1 $(TESTARGS) -coverprofile ./out/coverage

test-cover-html:
	mkdir -p ./out
	@echo "mode: count" > coverage-all.out
	$(foreach pkg, $(ALL_PACKAGES),\
	ENVIRONMENT=test $(GOBIN) test -coverprofile=coverage.out -covermode=count $(pkg);\
	tail -n +2 coverage.out >> coverage-all.out;)
	$(GOBIN) tool cover -html=coverage-all.out -o out/coverage.html

