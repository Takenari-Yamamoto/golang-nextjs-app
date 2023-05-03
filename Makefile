.PHONY: build up down clean

backend_run:
	cd backend && go run cmd/golang-nextjs-server/main.go --host 0.0.0.0 --port 8000

frontend_build:
	cd frontend && npm install && npm run dev

build:
	docker-compose build

build-up:
	docker-compose up -d --build

up:
	docker-compose up -d

down:
	docker-compose down

clean:
	docker-compose down -v --remove-orphans

generate-models:
	docker-compose exec backend sqlboiler psql