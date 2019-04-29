package main

import (
	"context"
	"fmt"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	pb "win-his/api"
)

func main() {
	// create a new service
	service := grpc.NewService()
	// parse command line flags
	service.Init()
	// Use the generated client stub
	ctx := micro.NewContext(context.Background(), service)
	cl := pb.NewAdminService("Admin", service.Client())
	// Make request
	workrsp, err := cl.TakeWorking(ctx, &pb.Request{
		Username: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(workrsp.Msg)
	msgresp, err := cl.PushMsg(ctx, &pb.PushMsgReq{
		Username: "John",
	})
	fmt.Println(msgresp.Msg)
}
