run:
	go run cmd/main.go

migrate_up:
	migrate -database 'postgres://test:test@localhost:5432/realty?sslmode=disable' -path ./migrations up

migrate_down:
	migrate -database 'postgres://test:test@localhost:5432/realty?sslmode=disable' -path ./migrations down
