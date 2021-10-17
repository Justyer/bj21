proto:
	rm -rf bj21.proto && cp -r ../api/bj21/v1/bj21.proto .
	grpc_tools_node_protoc \
		--js_out=import_style=commonjs,binary:api/bj21/v1 \
		--grpc_out=grpc_js:api/bj21/v1 \
		bj21.proto