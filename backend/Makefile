.PHONY: gen-pb
gen-pb: ## Generate protobuf files
	@echo Starting generate pb
	@protoc --proto_path=./ \
		--go_out=paths=source_relative:./ \
		--go-grpc_out=paths=source_relative,require_unimplemented_servers=false:./ \
		./*.proto
	@echo Successfully generated proto
	
	@echo Starting inject tags
	@protoc-go-inject-tag -input="./*.pb.go"
	@echo Successfully injected tags

clean: ## Clean generated protobuf files
	@echo Cleaning generated pb files
	@del *.pb.go
	@echo Successfully cleaned pb files