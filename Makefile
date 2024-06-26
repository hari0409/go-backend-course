postgres: 
	docker run --name sample-db -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres:latest

mysql:
	docker run --name sample-db-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql:latest

createdb:
	docker exec -it sample-db createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it sample-db dropdb --username=postgres simple_bank

migrateup:
	migrate -path db/migration -database "postgres://postgres:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://postgres:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose down

migrateup1:
	migrate -path db/migration -database "postgres://postgres:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path db/migration -database "postgres://postgres:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

gotest:
	go test -v -cover ./...

server: 
	go run main.go

mockgen:
	mockgen -destination ./db/mock/store.go -package mockdb github.com/hari0409/backend-go/db/sqlc Store

.PHONY: createdb postgres dropdb migrateup migratedown migrateup1 migratedown1 sqlc server mockgen