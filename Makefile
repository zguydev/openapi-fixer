LOCAL_BIN:=$(CURDIR)/bin
PATH:=$(PATH):$(LOCAL_BIN)

PKG:=github.com/zguydev/openapi-fixer
FIXER_ENTRYPOINT:=.
FIXER_BIN:=$(LOCAL_BIN)/openapi-fixer

ifneq (,$(wildcard .env))
	include .env
	export
endif

.PHONY: .app-deps
.app-deps:
	go mod tidy

.PHONY: install
install: .app-deps

.PHONY: build
build:
	@mkdir -p $(LOCAL_BIN)
	@go build -o $(FIXER_BIN) $(FIXER_ENTRYPOINT)

.PHONY: clean
clean:
	@rm -i $(LOCAL_BIN)/*

.PHONY: test
test:
	@go test -v -coverprofile=./coverage.out $(PKG)/...

.PHONY: cover
cover:
	@go test -cover -coverprofile ./coverage.out $(PKG)/...

.PHONY: show_cover
show_cover:
	@go tool cover -func=coverage.out

.PHONY: show_cover_html
show_cover_html:
	@go tool cover -html=coverage.out
