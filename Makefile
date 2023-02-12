postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres12-alpine

createdb:
	docker exec -it postgres12 createdb --username=ikehakinyemi --owner=mono_finance mono_finance

dropdb:
	docker exec -it postgres12 dropdb mono_finance

migrateup:
	migrate -path db/migrations -database "postgresql://mono_finance:Akinyemi@localhost:5432/mono_finance?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migrations -database "postgresql://mono_finance:Akinyemi@localhost:5432/mono_finance?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migrations -database "postgresql://mono_finance:Akinyemi@localhost:5432/mono_finance?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migrations -database "postgresql://mono_finance:Akinyemi@localhost:5432/mono_finance?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/IkehAkinyemi/mono-finance/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 test server mock
