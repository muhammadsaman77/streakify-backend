migrate-generate:
	migrate create -ext sql -dir db/migrations -seq ${MIGRATION_NAME}

migrate-up:
	migrate -path db/migrations -database "postgresql://postgres:secret@localhost:5432/steakify-db?sslmode=disable" up


migrate-down:
	migrate -path db/migrations -database "postgresql://postgres:secret@localhost:5432/steakify-db?sslmode=disable" down

migrate-version:
	migrate -path db/migrations -database "postgresql://postgres:secret@localhost:5432/steakify-db?sslmode=disable" version

migrate-force:
	migrate -path db/migrations -database "postgresql://postgres:secret@localhost:5432/steakify-db?sslmode=disable" force ${VERSION}

.PHONY: migrate-generate migrate-up migrate-down