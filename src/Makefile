
generate-proto:
	protoc -I=./proto/v1/health/ --go_out="${GOPATH}/src/" ./proto/v1/health/*.proto 
	protoc -I=./proto/v1/health/ --doc_out=markdown,doc.md:./proto/v1/health ./proto/v1/health/*.proto 

	protoc -I=./proto/v1/error/ --go_out="${GOPATH}/src/" ./proto/v1/error/*.proto 
	protoc -I=./proto/v1/error/ --doc_out=markdown,doc.md:./proto/v1/error ./proto/v1/error/*.proto 

	protoc -I=./proto/v1/auth/ --go_out="${GOPATH}/src/" ./proto/v1/auth/*.proto 
	protoc -I=./proto/v1/auth/ --doc_out=markdown,doc.md:./proto/v1/auth ./proto/v1/auth/*.proto 


api-docs:
	swag init --output=./web/docs