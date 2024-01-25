GO := @go
PROTOC := @protoc


protoc:
	${PROTOC} --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/product/*.proto

protoc_health:
	${PROTOC} --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/health/*.proto
