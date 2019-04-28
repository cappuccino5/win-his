package server

import (
	"context"
	"fmt"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"log"
	hello "win-his/api"
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
	hello.RegisterSagaAdminHandler(Server.Server(), bus)
}
func (s *Say) PushMsg(ctx context.Context, req *hello.PushMsgReq, rsp *hello.PushMsgReply) error {
	log.Print("http.SagaAdmin.Hello request")
	fmt.Println(ctx.Done())
	rsp.Msg = "Hello " + req.Username
	return nil
}
func newServer() *Say {
	say := new(Say)
	say.Domain = "http.SagaAdmin"
	return say
}
func Run() error {
	return Server.Run()
}
