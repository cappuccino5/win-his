package main

import (
	"context"
	"fmt"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	pb "win-his/api"
)
// 这两个main内部实现是一样的
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
	msgresp, err := cl.PushMsg(context.TODO(), &pb.PushMsgReq{
		Username: "John",
	})
	fmt.Println(msgresp.Msg)
}
func main1() {

	server := grpc.NewService()
	server.Init()
	client := pb.NewAdminService("Admin", server.Client())

	wresp, err := client.TakeWorking(context.TODO(), &pb.Request{Username: "John"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(wresp)
	mresp, err := client.PushMsg(context.TODO(), &pb.PushMsgReq{Username: "John"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mresp)
}
