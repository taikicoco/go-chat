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
