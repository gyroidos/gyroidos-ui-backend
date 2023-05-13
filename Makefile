all:
	CGO_ENABLED=0 go build

protobuf:
	protoc --proto_path=controlproto --go_out=. --go_opt=paths=source_relative control.proto guestos.proto logf.proto container.proto