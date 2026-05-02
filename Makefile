build: 
	@go build -o bin/ecommerce cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/ecommerce

migration:
	@migrate create -ext sql -dir cmd/migrations ${filter-out $@,$(MAKECMDGOALS)}
			
# migration-name:
# 	@migrate create -ext sql -dir cmd/migrations $(name)

migrate-up:
	@go run cmd/migrate/main.go up 

migrate-down:
	@go run cmd/migrate/main.go down