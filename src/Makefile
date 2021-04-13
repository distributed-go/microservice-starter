
build-proto:
	protoc -I=./proto/v1/health/ --go_out="${GOPATH}/src/" ./proto/v1/health/*.proto 
	protoc -I=./proto/v1/health/ --doc_out=markdown,doc.md:./proto/v1/health ./proto/v1/health/*.proto 