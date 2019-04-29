package server

import (
	"context"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"log"
	pb "win-his/api"
)

var Server micro.Service

type Say struct {
	Domain string
}

func init() {
	bus := newServer()
	Server = grpc.NewService(
		micro.Name(bus.Domain),
	)
	Server.Init()
	pb.RegisterAdminHandler(Server.Server(), bus)
}
func (s *Say) PushMsg(ctx context.Context, req *pb.PushMsgReq, rsp *pb.PushMsgReply) error {
	log.Print("http.SagaAdmin.Hello request")
	rsp.Msg = "welcome to push msg " + req.Username
	return nil
}
func newServer() *Say {
	return &Say{
		Domain: "Admin",
	}
}
func (s *Say) TakeWorking(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	log.Print("http.SagaAdmin.take request")
	resp.Msg = "welcome to take working " + req.Username
	return nil
}

func Run() error {
	return Server.Run()
}
