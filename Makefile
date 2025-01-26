include local.env

install-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.3
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.26.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/rakyll/statik@v0.1.7

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

local-migration-status:
	goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} status -v

local-migration-up:
	goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} up -v

local-migration-down:
	goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} down -v

test:
	go clean -testcache
	go test ./... -v -covermode count -coverpkg=github.com/astronely/financial-helper_microservices/internal/api/...

test-coverage:
	go clean -testcache
	go test ./... -coverprofile=coverage.tmp.out -covermode count -coverpkg=github.com/astronely/financial-helper_microservices/internal/api/...,github.com/astronely/financial-helper_microservices/internal/service/...
	grep -v 'mocks\|config' coverage.tmp.out > coverage.out
	rm coverage.tmp.out
	go tool cover -html=coverage.out;
	go tool cover -func=./coverage.out | grep "total";
	grep -sqFx "/coverage.out" .gitignore || echo "/coverage.out" >> .gitignore

generate:
	make generate-user-api
	mkdir -p pkg/swagger
	statik -src=pkg/swagger/ -include='*.css,*.html,*.js,*.json,*.png'

generate-user-api:
	mkdir -p pkg/user_v1
	protoc --proto_path api/user_v1 --proto_path vendor.protogen \
	--go_out=pkg/user_v1 --go_opt=paths=source_relative \
 	--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
 	--grpc-gateway_out=pkg/user_v1 --grpc-gateway_opt=paths=source_relative \
 	--openapiv2_out=allow_merge=true,merge_file_name=api:pkg/swagger \
 	 api/user_v1/user.proto

vendor-proto:
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
			mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
			git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
			mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
			rm -rf vendor.protogen/openapiv2 ;\
		fi