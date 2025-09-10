PROTO_DIR=proto
OUT_DIR=proto/pb

# Dependencies
PROTOC=protoc

# Generate Go code from .proto
.PHONY: proto
proto:
	@mkdir -p $(OUT_DIR)
	$(PROTOC) \
		--proto_path=$(PROTO_DIR) \
		--proto_path=$(shell go list -m -f '{{.Dir}}' google.golang.org/protobuf)/src \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/*.proto