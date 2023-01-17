package main

import (
	"context"
	"log"
	"strconv"

	minimal_demo "github.com/808-not-found/tik_duck/kitex_simple_demo/kitex_gen/minimal_demo"
	"github.com/808-not-found/tik_duck/kitex_simple_demo/kitex_gen/minimal_demo/addservice"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
)

func main() {
	client, err := addservice.NewClient("hello", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})
	h.GET("/add", func(c context.Context, ctx *app.RequestContext) {
		log.Println(ctx.Query("a"), ctx.Query("b"))
		a, _ := strconv.ParseInt(ctx.Query("a"), 10, 64)
		b, _ := strconv.ParseInt(ctx.Query("b"), 10, 64)
		req := &minimal_demo.AddRequest{A: a, B: b}
		resp, err := client.Add(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		ctx.JSON(consts.StatusOK, utils.H{"res": resp.GetRes()})
	})
	h.Spin()
}
