
.PHONY: generate-grpc
generate-grpc: generate-grpc-go generate-grpc-react

.PHONY: generate-grpc-react
generate-grpc-react:
	protoc -I=proto/ dare.proto --js_out=import_style=commonjs:react_client/src/gen
	protoc -I=proto/ dare.proto --grpc-web_out=import_style=commonjs,mode=grpcwebtext:react_client/src/gen
	sed -i '1s_^_/* eslint-disable */\n_' react_client/src/gen/dare_pb.js
	sed -i '1s_^_/* eslint-disable */\n_' react_client/src/gen/dare_grpc_web_pb.js

.PHONY: generate-grpc-go
generate-grpc-go:
	protoc -I proto/ dare.proto --go_out=plugins=grpc:backend/internal/gen

.PHONY: docker-deps
docker-deps:
	@docker-compose down
	@docker-compose rm -f
	docker-compose up --build --force-recreate -d

.PHONY: run-react
run-react:
	yarn --cwd react_client start

.PHONY: run-api
run-api:
	go run ./backend 
