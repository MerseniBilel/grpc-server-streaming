create:
	protoc --proto_path=proto proto/*.proto --go_out=grpc/
	protoc --proto_path=proto proto/*.proto --go-grpc_out=grpc/ 


clear:
	rm -rf ./grpc/*