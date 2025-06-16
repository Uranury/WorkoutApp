run:
	go run .

create-mig:
	migrate create -ext sql -dir internal/db/migrations -seq $(name)

migrate-up:
	migrate -path internal/db/migrations -database $(DB_URL) up

migrate-down:
	migrate -path internal/db/migrations -database $(DB_URL) down
