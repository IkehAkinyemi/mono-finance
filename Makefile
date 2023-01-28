postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres12-alpine

createdb:
	docker exec -it postgres12 createdb --username=ikehakinyemi --owner=mono_finance mono_finance

dropdb:
	docker exec -it postgres12 dropdb mono_finance

migrateup:
	migrate -path db/migrations -database "postgresql://mono_finance:Akinyemi@localhost:5432/mono_finance?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://mono_finance:Akinyemi@localhost:5432/mono_finance?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown test
