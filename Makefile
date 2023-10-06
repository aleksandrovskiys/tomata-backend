build:
	@go build -o build/tomata-backend

run:
	@go run main.go

rerun_container:
	@docker compose down
	@docker compose up -d

build_container:
	@docker compose build

activate_env:$
	@export $(grep -v '^#' .env | xargs)

