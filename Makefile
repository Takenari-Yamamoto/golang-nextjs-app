backend_run:
	cd backend && go run cmd/golang-nextjs-server/main.go --host 0.0.0.0 --port 8000

frontend_build:
	cd frontend && npm install && npm run build
