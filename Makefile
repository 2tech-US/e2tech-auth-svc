postgres:
	docker run --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root auth_svc
dropdb:
	docker exec -it postgres12 dropdb --username=root auth_svc
migrateup:
	migrate -path migration -database "postgresql://root:secret@localhost:5432/auth_svc?sslmode=disable" -verbose up
migratedown:
	migrate -path migration -database "postgresql://root:secret@localhost:5432/auth_svc?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    internal/pb/*.proto

server:
	go run cmd/server/main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test proto