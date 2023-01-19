package main

import (
	// "context"
	// "log"
	// "strconv"

	// user "github.com/808-not-found/tik_duck/kitex_gen/user"
	// "github.com/808-not-found/tik_duck/kitex_gen/user/userservice"
	// minimal_demo "github.com/808-not-found/tik_duck/kitex_simple_demo/kitex_gen/minimalDemo"
	// "github.com/808-not-found/tik_duck/kitex_simple_demo/kitex_gen/minimalDemo/addservice"
	// "github.com/cloudwego/hertz/pkg/app"
	// "github.com/cloudwego/hertz/pkg/app/server"
	"github.com/808-not-found/tik_duck/cmd/web/service"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	go service.RunMessageServer()
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))
	InitRouter(h)
	h.Spin()
	//old
	//demo申请的客户端
	// client0, err0 := addservice.NewClient("Hello",client.WithHostPorts("0.0.0.0:8888"))
	// if err0 != nil {
	// 	log.Fatal(err0)
	// }
	// //基础接口申请的客户端
	// client1, err1 := userservice.NewClient("Hello",client.WithHostPorts("0.0.0.0:10001"))
	// if err1 != nil {
	// 	log.Fatal(err1)
	// }

	// h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	// h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
	// 	ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	// })
	// h.GET("/add", func(c context.Context, ctx *app.RequestContext) {
	// 	log.Println(ctx.Query("a"), ctx.Query("b"))
	// 	a, _ := strconv.ParseInt(ctx.Query("a"), 10, 64)
	// 	b, _ := strconv.ParseInt(ctx.Query("b"), 10, 64)
	// 	req := &minimal_demo.AddRequest{A: a, B: b}
	// 	resp, err := client0.Add(context.Background(), req)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	ctx.JSON(consts.StatusOK, utils.H{"res": resp.GetRes()})
	// })
	// h.GET("/user_test",func(c context.Context, ctx *app.RequestContext) {
	// 	req := &user.Testinfo{Testinfo:"这是一条测试rpc的test信息"}
	// 	resp, err :=client1.UserTest(context.Background(),req)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	ctx.JSON(consts.StatusOK, utils.H{"res": resp.Testinfo})
	// })
	// h.Spin()
}


