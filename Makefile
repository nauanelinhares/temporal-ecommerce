.PHONY: deps/up deps/down deps/restart deps/up-alt

deps/up:
	docker compose up -d postgres

# Alternative port if 5432 is already in use
deps/up-alt:
	export DATABASE_PORT=5433 && docker compose up -d postgres

deps/down:
	docker compose down

deps/restart: deps/down deps/up

# Application Migration Commands
# Uses the main package within the ./migrations directory
.PHONY: migrate/up migrate/down

migrate/up:
	@echo "🚀 Applying database migrations..."
	@go run ./migrations -direction=up

migrate/down:
	@echo "⏪ Reverting database migrations..."
	@go run ./migrations -direction=down 