.DEFAULT_GOAL := help

.PHONY: setup
setup: ## setup
	cp -p .envrc.sample .envrc
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
	go install github.com/go-delve/delve/cmd/dlv@latest
	go install github.com/cosmtrek/air@latest
	go mod download

.PHONY: clean
clean: ## clean
	go clean
	rm -f app

.PHONY: build
build: ## build app
	go build -o app -v

.PHONY: run
run: ## run app
	air

.PHONY: fmt
fmt: ## format
	go fmt

.PHONY: sqlcg
sqlcg:  ## run sqlc generate --experimental
	sqlc generate --experimental

.PHONY: help
help: ## this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
