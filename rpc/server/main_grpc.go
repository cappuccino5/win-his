package main

/* *************************************这是实现grpc微服务用例*************************************
****************************************grpc有端口限制********************************************
*************************************************************************************************/
import (
	"context"
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "win-his/gateway"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) PushMsg(ctx context.Context, in *pb.PushMsgReq) (*pb.PushMsgReply, error) {
	log.Printf("Received: %v", in.Username)
	return &pb.PushMsgReply{Msg: "Hello " + in.Username}, nil
}
func (s *server) TakeWorking(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Received: %v", in.Username)
	return &pb.Response{Msg: "welcome to take  " + in.Username}, nil
}

func main() {
	fmt.Println("welcome to main accept port:", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAdminServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
