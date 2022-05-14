postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root trendyol
dropdb:
	docker exec -it postgres12 dropdb trendyol
migrateup:
	migrate -path db/migrations -database "postgresql://root:postgres@localhost:5432/trendyol?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://root:postgres@localhost:5432/trendyol?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover	./...
run:
	go run cmd/main.go
setupdb:
	
.PHONY: postgres createdb dropdb migrateup migratedown sqlc
