.PHONY: format
format:
	@goimports -w -l .

.PHONY: check
check:
	@goimports -l .

.PHONY: test
test:
	@go test ./...

.PHONY: check-all
check-all: check test
