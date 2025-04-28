include local.env

install-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.3
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
	go install github.com/envoyproxy/protoc-gen-validate@v1.2.1
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.26.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install github.com/pressly/goose/v3/cmd/goose@v3.24.1
	go install github.com/rakyll/statik@v0.1.7

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-user-api
	make generate-auth-api
	make generate-access-api
	make generate-wallet-api
	make generate-transaction-api
	# TODO: other APIs
	make generate-openapi
	mkdir -p pkg/swagger
	statik -src=pkg/swagger/ -include='*.css,*.html,*.js,*.json,*.png'


generate-user-api:
	mkdir -p pkg/user_v1
	protoc --proto_path api/user_v1 --proto_path vendor.protogen \
	--go_out=pkg/user_v1 --go_opt=paths=source_relative \
 	--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
 	--validate_out lang=go:pkg/user_v1 --validate_opt=paths=source_relative \
 	--grpc-gateway_out=pkg/user_v1 --grpc-gateway_opt=paths=source_relative \
 	  api/user_v1/user.proto

generate-auth-api:
	mkdir -p pkg/auth_v1
	protoc --proto_path api/auth_v1 --proto_path vendor.protogen \
	--go_out=pkg/auth_v1 --go_opt=paths=source_relative \
 	--go-grpc_out=pkg/auth_v1 --go-grpc_opt=paths=source_relative \
 	--grpc-gateway_out=pkg/auth_v1 --grpc-gateway_opt=paths=source_relative \
 	api/auth_v1/auth.proto

generate-access-api:
	mkdir -p pkg/access_v1
	protoc --proto_path api/access_v1 --proto_path vendor.protogen \
	--go_out=pkg/access_v1 --go_opt=paths=source_relative \
 	--go-grpc_out=pkg/access_v1 --go-grpc_opt=paths=source_relative \
 	--grpc-gateway_out=pkg/access_v1 --grpc-gateway_opt=paths=source_relative \
 	api/access_v1/access.proto

generate-wallet-api:
	mkdir -p pkg/wallet_v1
	protoc --proto_path api/wallet_v1 --proto_path vendor.protogen \
	--go_out=pkg/wallet_v1 --go_opt=paths=source_relative \
 	--go-grpc_out=pkg/wallet_v1 --go-grpc_opt=paths=source_relative \
 	--grpc-gateway_out=pkg/wallet_v1 --grpc-gateway_opt=paths=source_relative \
 	api/wallet_v1/wallet.proto

generate-transaction-api:
	mkdir -p pkg/transaction_v1
	protoc --proto_path api/transaction_v1 --proto_path vendor.protogen \
	--go_out=pkg/transaction_v1 --go_opt=paths=source_relative \
 	--go-grpc_out=pkg/transaction_v1 --go-grpc_opt=paths=source_relative \
 	--grpc-gateway_out=pkg/transaction_v1 --grpc-gateway_opt=paths=source_relative \
 	api/transaction_v1/transaction.proto

generate-openapi:
	mkdir -p pkg/swagger
	protoc --proto_path=api --proto_path=vendor.protogen \
		--openapiv2_out=allow_merge=true,merge_file_name=userApi:pkg/swagger \
		api/user_v1/user.proto

	protoc --proto_path=api --proto_path=vendor.protogen \
		--openapiv2_out=allow_merge=true,merge_file_name=authApi:pkg/swagger \
		api/auth_v1/auth.proto api/access_v1/access.proto

	protoc --proto_path=api --proto_path=vendor.protogen \
		--openapiv2_out=allow_merge=true,merge_file_name=financeApi:pkg/swagger \
		api/wallet_v1/wallet.proto api/transaction_v1/transaction.proto

	protoc --proto_path=api --proto_path=vendor.protogen \
        		--openapiv2_out=allow_merge=true,merge_file_name=boardApi:pkg/swagger \
        		api/board_v1/board.proto

	protoc --proto_path=api --proto_path=vendor.protogen \
    		--openapiv2_out=allow_merge=true,merge_file_name=noteApi:pkg/swagger \
    		api/note_v1/note.proto

vendor-proto:
		@if [ ! -d vendor.protogen/validate ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
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