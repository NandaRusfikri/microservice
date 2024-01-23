GO := @go
PROTOC := @protoc

protoc_product:
	${PROTOC} --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/product/*.proto


protoc_order:
	${PROTOC} --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/order/*.proto
