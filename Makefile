SHELL := /bin/bash

# ======== Go ========
.PHONY: gobuild
gobuild:
	@CGO_ENABLED=0 go build -ldflags="-w -s" -o build/goyamlc .

.PHONY: run
run: gobuild
	./build/goyamlc .

# ======== Lint ========
LINT_DISABLE_ERR ?= true
define lint
	$(1) || $(LINT_DISABLE_ERR)
endef

.PHONY: actionlint
actionlint:
	which actionlint || go install github.com/rhysd/actionlint/cmd/actionlint@latest
	@$(call lint,actionlint -shellcheck=)

GO_LINT_VERSION ?= v2.1.2
GO_LINT_FIX ?= --fix
.PHONY: golint
golint:
	@golangci-lint version | grep $(GO_LINT_VERSION) || go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GO_LINT_VERSION)
	@$(call lint,golangci-lint run --timeout=10m $(GO_LINT_FIX))

.PHONY: yamllint
yamllint:
	@$(call lint,yamllint .)

.PHONY: lint
lint: actionlint golint yamllint
