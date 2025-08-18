gen:
	goctl rpc protoc docs/api.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=. --style go_zero
	protoc-go-inject-tag -input=./pb/api/api.pb.go

# go install github.com/favadi/protoc-go-inject-tag@latest

php:
	protoc --proto_path=docs  --grpc_out=./php_out --plugin=protoc-gen-grpc=/path/to/grpc_php_plugin docs/api.proto

