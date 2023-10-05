build:
	@go build -o build/tomata-backend

run:
	@go run main.go

activate_env:$
	@export $(grep -v '^#' .env | xargs)

