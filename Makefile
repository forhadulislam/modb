
.PHONY: run
run: ## Run the Database server
	go run ./main.go

.PHONY: test
test: ## Run all unit tests
	go test ./...

.PHONY: vendor
vendor: ## Install all dependencies
	go mod vendor