.PHONY: run
run: ## Run the Database server
	go run ./main.go

.PHONY: test
test: ## Run all unit tests
	go clean -testcache
	go test ./...

.PHONY: test-verbose
test-verbose: ## Run all unit tests with log messages
	go test -v ./...

.PHONY: vendor
vendor: ## Install all dependencies
	go mod vendor