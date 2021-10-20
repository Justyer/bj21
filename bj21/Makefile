proto:
	cp -r ../api/bj21/v1/bj21.proto .
	grpc_tools_node_protoc \
		--js_out=import_style=commonjs,binary:src/api/bj21/v1 \
		--grpc_out=grpc_js:src/api/bj21/v1 \
		bj21.proto
	rm -rf bj21.proto