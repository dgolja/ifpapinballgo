## default variable definitions
IFPA_OPENAPI_JSON ?= ifpapinball-official-api.json
IFPA_OPENAPI_YAML ?= ifpapinball-official-api.yaml
IFPA_OPENAPI_MODIFIED_YAML ?= ifpapinball-official-api.modified.yaml

## Additional tooling needed for Makefile to work
TOOLS_BIN := $(PWD)/.bin
YQ_VERSION := v4.52.4
OAPI_CODEGEN_VERSION := v2.6.0
SPEAKEASY_OAPI_VERSION := v0.0.0-20260301231816-65621191fc9d

# Detect OS and architecture
OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH := $(shell uname -m)

# Normalize arch naming
ifeq ($(ARCH),x86_64)
    ARCH := amd64
endif
ifeq ($(ARCH),aarch64)
    ARCH := arm64
endif

export PATH:=$(TOOLS_BIN):$(PATH)

.DEFAULT_GOAL := help

.DELETE_ON_ERROR:

.PHONY: help
help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: install-tools ## Install all tools
install-tools: $(TOOLS_BIN)/yq $(TOOLS_BIN)/oapi-codegen $(TOOLS_BIN)/openapi ## Install all required tools

$(TOOLS_BIN)/yq: ## install yq
	@mkdir -p $(TOOLS_BIN)
	curl -sSfL "https://github.com/mikefarah/yq/releases/download/$(YQ_VERSION)/yq_$(OS)_$(ARCH)" \
		-o $(TOOLS_BIN)/yq
	chmod +x $(TOOLS_BIN)/yq

$(TOOLS_BIN)/openapi: ## install openapi
	@mkdir -p $(TOOLS_BIN)
	GOBIN=$(TOOLS_BIN) go install github.com/speakeasy-api/openapi/cmd/openapi@$(SPEAKEASY_OAPI_VERSION)

$(TOOLS_BIN)/oapi-codegen: ## install oapi-codegen
	@mkdir -p $(TOOLS_BIN)
	GOBIN=$(TOOLS_BIN) go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@$(OAPI_CODEGEN_VERSION)

.PHONY: apply-overlay
apply-overlay: $(TOOLS_BIN)/openapi ## Apply ifpapinball-overlay.yaml to the official API spec to produce the modified spec
	$(TOOLS_BIN)/openapi overlay apply --overlay ifpapinball-overlay.yaml --schema ifpapinball-official-api.yaml --out $(IFPA_OPENAPI_MODIFIED_YAML)

.PHONY: gen-client
gen-client: $(TOOLS_BIN)/oapi-codegen $(IFPA_OPENAPI_MODIFIED_YAML) ## Generate ifpapinball client/types using OpenAPI manifests
	# NOTE: We generate from the modified, not the official one.
	# The official IFPA OpenAPI spec is incomplete (missing fields) and contains structural errors
	# that are incompatible with statically-typed language code generation.
	# The overlay (ifpapinball-overlay.yaml) patches those issues before generation.
	# Run 'make apply-overlay' to regenerate the modified spec from the official one.
	# generate types
	$(TOOLS_BIN)/oapi-codegen --config api-config/ifpapinball-types.yaml $(IFPA_OPENAPI_MODIFIED_YAML)
	# generate client
	$(TOOLS_BIN)/oapi-codegen --config api-config/ifpapinball-client.yaml $(IFPA_OPENAPI_MODIFIED_YAML)

.PHONY: get-ifpa-api
get-ifpa-api: ## Retrieve IFPA Pinball API definition and store it to $(IFPA_OPENAPI_JSON)
	@curl -s -o $(IFPA_OPENAPI_JSON) https://api.ifpapinball.com/docs/api.json

.PHONY: ifpa-api-json-to-yaml
ifpa-api-json-to-yaml: get-ifpa-api $(TOOLS_BIN)/yq ## This converts the IFPA Official Pinball API spec from JSON to YAML
	$(TOOLS_BIN)/yq -p json -o yaml $(IFPA_OPENAPI_JSON) > $(IFPA_OPENAPI_YAML)

.PHONY: clean
clean: ## Removes the /.bin directory.
	rm -rf $(TOOLS_BIN)
