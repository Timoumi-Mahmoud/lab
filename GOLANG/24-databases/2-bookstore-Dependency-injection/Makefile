
run:
	go run cmd/main.go

test:
	go test -v tests/auth_test/auth_test.go

postgresinit:
	docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgresrestart:
	docker start postgres15

 postgresacces:
	docker exec -it postgres15 psql

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root myDB

dropdb:
	docker exec -it postgres15 dropdb myDB

migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/myDB?sslmode=disable" -verbose up 

migratedwon:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/myDB?sslmode=disable" -verbose down

.PHONY: run postgresinit postgresaccess createdb dropdb migrateup migratedwon
