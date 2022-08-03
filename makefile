include $(ENV_FILE)
export

sqlc:
	sqlc generate

createdb:
	docker exec -it local-postgres createdb --user $(DB_USER) --owner=$(DB_USER) $(DB_NAME)

dropdb:
	docker exec -it local-postgres dropdb $(DB_NAME)

# TODO: add migrate down

migrateup:
	migrate -path ./config/migrations -database postgresql://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable -verbose up

.PHONY: createdb migrateup