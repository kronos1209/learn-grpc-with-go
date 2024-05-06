.PHONY : proto
proto:
	protoc -I ./proto/ \
		--go_out=./pkg/proto \
		--go-grpc_out=./pkg/proto \
		--go_opt=paths=source_relative	\
		--go-grpc_opt=paths=source_relative \
		./proto/*proto