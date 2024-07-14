swag:
	swag init -g api/api.go -o api/docs

swag-install:
	[ -d "api/docs" ] && echo "Directory exists" || (echo "Directory does not exist, creating"; mkdir "api/docs")
	go get -u github.com/swaggo/swag/cmd/swag@v1.16.3
	go install github.com/swaggo/swag/cmd/swag@v1.16.3

run-server:
	go run cmd/main.go

db:
	psql -U postgres -W -h localhost -p 5432 -d samarkand_notbuk_sayt;

run: swag run-server
build:
	go mod tidy
	go build -o app cmd/main.go

# for server
install: swag-install swag build
start:
	./app

minio:
	minio server ~/minio/data --console-address ":9001"