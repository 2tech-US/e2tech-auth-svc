postgres:
	docker run --name e2tech -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:12-alpine
createdb:
	docker exec -it e2tech createdb --username=root --owner=root auth_svc
dropdb:
	docker exec -it e2tech dropdb --username=root auth_svc
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
up:
	go mod tidy
	make postgres
	sleep 1
	make createdb
	make migrateup
down:
	docker stop e2tech
	docker rm e2tech
	

auth_svc:
	go run cmd/server/main.go

update:
	cd .. && make update

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test proto