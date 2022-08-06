include .env
export

sqlc:
	sqlc generate

createdb:
	docker exec -it local-postgres createdb --user $(DB_USER) --owner=$(DB_USER) $(DB_NAME)

dropdb:
	docker exec -it local-postgres dropdb $(DB_NAME)

migrateup:
	migrate -path ./config/migrations -database postgresql://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable -verbose up

migratedown:
	migrate -path ./config/migrations -database postgresql://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable -verbose down

rabbit:
	docker run -d --hostname my-rabbit --name local-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management

postgres:
	docker run --name local-postgres -p 5432:5432 -e POSTGRES_PASSWORD=$(DB_PASS) -e POSTGRES_USER=$(DB_USER) -e POSTGRES_DB=$(DB_NAME) -d postgres

containers: rabbit postgres

.PHONY: sqlc createdb dropdb migrateup migratedown