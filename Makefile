all:
	protoc --proto_path=. --twirp_out=. --go_out=. ./rpc/helloworld/service.proto
	cat hello.json | go run main.go convert-req > hello.proto

serve:
	go run main.go serve


example:
	@echo "Making a JSON request:"
	curl -H "Content-Type: application/json" --data-binary @hello.json 'http://localhost:8080/twirp/helloworld.HelloWorld/Hello'
	@echo
	@echo
	@echo "Making a protobuf request:"
	curl -H "Content-Type: application/protobuf" \
	  --data-binary @hello.proto 'http://localhost:8080/twirp/helloworld.HelloWorld/Hello' \
	  | go run main.go convert-resp
