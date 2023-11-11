.PHONY: all proto protoc-ts protoc-js

all: proto protoc-ts protoc-js

proto:
	@protoc --go_out=./serverpb --go_opt=paths=source_relative server.proto

protoc-ts:
	@protoc \
		--plugin="./node_modules/.bin/protoc-gen-ts_proto" \
		--ts_proto_opt=esModuleInterop=true \
		--ts_proto_out="./clientpb" \
		server.proto

protoc-js:
	@protoc \
		--plugin="./node_modules/.bin/protoc-gen-js" \
		--js_out="import_style=commonjs,binary:./clientpb" \
		server.proto
