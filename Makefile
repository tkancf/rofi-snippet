.DEFAULT_GOAL:=help

.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: install
install: ## Install rofi-snippet
	@sudo mkdir -p /etc/rofi-snippet/
	@sudo cp ./config.toml /etc/rofi-snippet/
	@go install .

