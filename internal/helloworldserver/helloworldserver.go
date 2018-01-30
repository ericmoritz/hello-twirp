package helloworldserver

import (
	"context"
	"net/http"
	pb "github.com/ericmoritz/hello-twirp/rpc/helloworld"
	"fmt"
)

type Server struct {}

// Hello says hello!
func (s *Server) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	if req.Subject == "Kenny" {
		return nil, fmt.Errorf("Dammit %s", req.Subject)
	}

	text := "Hello, World"
	if req.Subject != "" {
		text = fmt.Sprintf("Hello, %s!", req.Subject)
	}

	resp := &pb.HelloResp{
		Text: text,
	}
	return resp, nil
}

// ListenAndServe run the server
func ListenAndServe(bind string) {
	server := &Server{}
	twirpHandler := pb.NewHelloWorldServer(server, nil)
	fmt.Printf("Serving on %T on :8080\n", twirpHandler)
	if err := http.ListenAndServe(bind, twirpHandler); err != nil {
			panic(err)
	}
}
