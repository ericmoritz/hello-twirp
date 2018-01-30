package main

import (
	"os"
	"flag"
	pb "github.com/ericmoritz/hello-twirp/rpc/helloworld"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"github.com/ericmoritz/hello-twirp/internal/helloworldserver"
)

func main() {
	flag.Parse()
	switch flag.Arg(0) {
	case "convert-req":
		jsonToProto(&pb.HelloReq{})
	case "convert-resp":
		protoToJSON(&pb.HelloResp{})
	case "serve":
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		helloworldserver.ListenAndServe(":"+port)
	default:
		panic("arg(1) not in {convert-req, convert-resp, serve}")
	}
}

func jsonToProto(msg proto.Message) {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	if err := jsonpb.UnmarshalString(string(bytes), msg); err != nil {
		panic(err)
	}

	bytes, err = proto.Marshal(msg)
	if err != nil {
		panic(err)
	}

	_, err = os.Stdout.Write(bytes)
	if err != nil {
		panic(err)
	}
}

func protoToJSON(msg proto.Message) {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	if err := proto.Unmarshal(bytes, msg); err != nil {
		panic(err)
	}

	if err := (&jsonpb.Marshaler{}).Marshal(os.Stdout, msg); err != nil {
		panic(err)
	}
}
