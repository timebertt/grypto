# ==========================
# binary
# ==========================
.PHONY: build
build:
	@GO111MODULE=on go build -o bin/grypto ./grypto

.PHONY: install
install:
	@GO111MODULE=on go install ./grypto

# ==========================
# verification
# ==========================
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
