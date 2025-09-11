PROTO_DIR=proto
OUT_DIR=proto/pb

PROTOC=protoc

.PHONY: proto
proto:
	@mkdir -p $(OUT_DIR)
	$(PROTOC) \
		--proto_path=$(PROTO_DIR) \
		--proto_path=$(shell go list -m -f '{{.Dir}}' google.golang.org/protobuf)/src \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/*.proto

DB_URL=postgres://admin:123456@localhost:5432/account_db?sslmode=disable
MIGRATIONS_DIR=./migrations

migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1

migrate-new:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

migrate-version:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version