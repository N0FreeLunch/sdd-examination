.PHONY: generate run

generate:
	@echo "Generators removed as we are moving to SSR/HTMX"
	go mod tidy
	go mod tidy

run:
	go run cmd/server/main.go
