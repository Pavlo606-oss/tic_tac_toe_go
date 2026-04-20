DOCKER_COMPOSE=docker compose -f build/docker-compose.yml

up:
	$(DOCKER_COMPOSE) up -d
play:
	go run cmd/server/main.go
down:
	$(DOCKER_COMPOSE) down

down-clean:
	$(DOCKER_COMPOSE) down -v

migrate-up:
	docker exec -i gaming-postgres psql -U gamer -d games_db < migrations/001_create_games_up.sql