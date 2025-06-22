run:
	go run ./cmd

create-mig:
	migrate create -ext sql -dir internal/db/migrations -seq $(name)

migrate-up:
	migrate -path internal/db/migrations -database $(DB_URL) up

migrate-down:
	migrate -path internal/db/migrations -database $(DB_URL) down

build:
	docker-compose up --build