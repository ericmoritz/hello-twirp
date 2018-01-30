all:
	protoc --proto_path=. --twirp_out=. --go_out=. ./rpc/helloworld/service.proto
	cat examples/hello.json | go run main.go convert-req > examples/hello.proto

serve:
	go run main.go serve


example:
	@echo "Making a JSON request:"
	curl -H "Content-Type: application/json" --data-binary @examples/hello.json 'http://localhost:8080/twirp/helloworld.HelloWorld/Hello'
	@echo
	@echo
	@echo "Making a protobuf request:"
	curl -H "Content-Type: application/protobuf" \
	  --data-binary @examples/hello.proto 'http://localhost:8080/twirp/helloworld.HelloWorld/Hello' \
	  | go run main.go convert-resp
