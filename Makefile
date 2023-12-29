build:
	docker-compose build
build-nc:
	docker-compose build --no-cache
up:
	docker-compose up
upd:
	docker-compose up -d
down:
	docker-compose down -v
log:
	docker-compose logs -f
sh:
	docker-compose exec backend sh

gqlgen:
	docker-compose exec backend sh -c "gqlgen generate"

# db
MYSQL_USER := user
MYSQL_PASSWORD := password
MYSQL_DB := db

mysql:
	docker compose exec mysql mysql -u$(MYSQL_USER) -p$(MYSQL_PASSWORD) -D$(MYSQL_DB)

migrate/new:
	echo '-- +goose Up' > db/migration/_.sql

migrate/up:
	docker-compose exec -T backend sh ../db/migration/script/migrate-up.sh 

seed:
	docker-compose exec mysql mysql -u$(MYSQL_USER) -p$(MYSQL_PASSWORD) -D$(MYSQL_DB) -e "source /seed/seed.sql"