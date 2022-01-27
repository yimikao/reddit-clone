postgres: ; sudo docker run --name redditclone -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb: ; sudo docker exec -it redditclone createdb --username=root --owner=root redditclone

dropdb: ; sudo docker exec -it redditclone dropdb redditclone

schema: ; migrate create -ext sql -dir db/migration -seq init_schema

migrateup: ; migrate -path db/migration -database "postgresql://root:secret@localhost:5431/redditclone?sslmode=disable" -verbose up

migratedown: ; migrate -path db/migration -database "postgresql://root:secret@localhost:5431/redditclone?sslmode=disable" -verbose down

migrateup1: ; migrate -path db/migration -database "postgresql://root:secret@localhost:5431/redditclone?sslmode=disable" -verbose up 1

migratedown1: ; migrate -path db/migration -database "postgresql://root:secret@localhost:5431/redditclone?sslmode=disable" -verbose down 1

sqlc: ; sqlc generate

test: ; go test -v -cover ./...

server: ; go run main.go

mock: ;  mockgen -package mockdb -destination db/mock/store.go gobank/db/sqlc Store

.PHONY: ; postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock
