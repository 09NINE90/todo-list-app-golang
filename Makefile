include .env
export

export PROJECT_ROOT=$(shell pwd)

migrate-create:
	@if [ -z "$(seq)" ]; then \
  		echo "Отсутствует параметр seq"; \
  		exit 1; \
	fi; \
	docker compose run --rm todo-app-db-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	make migrate-action	action=up

migrate-down:
	make migrate-action	action=down

migrate-action:
	docker compose run --rm todo-app-db-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@razum0ff.ru:6432/${POSTGRES_DB_NAME}?sslmode=disable \
		"$(action)"