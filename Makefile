default: testacc

ifeq (, $(shell which golangci-lint))
 $(error "No golangci-lint in $(PATH), please visit https://golangci-lint.run/usage/install/")
endif

# lint
.PHONY: lint
lint:
	golangci-lint run ./...

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m
