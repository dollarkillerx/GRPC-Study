/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-10
* Time: 下午3:11
* */
package main

import (
	"GRPC-Study/demo1/proto"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
)

func main() {
	// 创建一个tcp拨号
	conn, e := grpc.Dial(":9001", grpc.WithInsecure()) // grpc.WithInsecure() 不安全的传输
	if e != nil {
		panic(e.Error())
	}
	client := proto.NewAddServiceClient(conn)// 注册上去
	
	
	app := gin.Default()
	app.GET("/add/:a/:b", middlerware,func(ctx *gin.Context) {
		a := ctx.MustGet("a").(int64)
		b := ctx.MustGet("b").(int64)
		req := &proto.Request{A: a, B: b}
		if response, i := client.Add(ctx, req);i != nil { // 发送rpc请求
			ctx.JSON(http.StatusInternalServerError,gin.H{
				"error":"Internal Server Error",
			})
		}else{
			ctx.JSON(http.StatusOK,gin.H{
				"result":fmt.Sprint(response.Result),
			})
		}
	})
	app.GET("/mult/:a/:b", middlerware,func(ctx *gin.Context) {
		a := ctx.MustGet("a").(int64)
		b := ctx.MustGet("b").(int64)
		req := &proto.Request{A: a, B: b}
		if response, i := client.Multiply(ctx, req);i != nil { // 发送rpc请求
			ctx.JSON(http.StatusInternalServerError,gin.H{
				"error":"Internal Server Error",
			})
		}else{
			ctx.JSON(http.StatusOK,gin.H{
				"result":fmt.Sprint(response.Result),
			})
		}
	})


	if e := app.Run(":9002");e != nil {
		panic(e.Error())
	}
}

// 对a,b 参数 做个验证
func middlerware(ctx *gin.Context) {
	a, err := strconv.ParseInt(ctx.Param("a"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"err":"Bad Request",
		})
		ctx.Abort()
		return
	}
	b, err := strconv.ParseInt(ctx.Param("b"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"err":"Bad Request",
		})
		ctx.Abort()
		return
	}
	ctx.Set("a",a)
	ctx.Set("b",b)
}
