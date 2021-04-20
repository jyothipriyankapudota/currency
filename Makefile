.PHONY: protos

protos:
	protoc -I=protos --go_out=protos protos/currency.proto

protos-grpc:
	protoc -I protos/ protos/currency.proto --go-grpc_out=protos