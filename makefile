include $(ENV_FILE)
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

# TODO: remove rabbit
rabbit:
	docker run -d --hostname my-rabbit --name local-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management

.PHONY: sqlc createdb dropdb migrateup migratedown