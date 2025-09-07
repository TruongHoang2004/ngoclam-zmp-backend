APP_NAME=ngoclam-zmp-be
MAIN_FILE=cmd/main.go

# Lệnh generate swagger docs
swag:
	swag init -g $(MAIN_FILE) -o docs

# Build binary
build: swag
	go build -o bin/$(APP_NAME) $(MAIN_FILE)

# Run trực tiếp
run: swag
	go run $(MAIN_FILE)

diff:
	@read -p "Enter migration name: " name; \
	atlas migrate diff $$name --env gorm

migrate:
	atlas migrate apply --env gorm

status:
	atlas migrate status --env gorm

# Clean build + docs
clean:
	rm -rf bin/ docs/docs.go docs/swagger.json docs/swagger.yaml

.PHONY: swag build run clean
