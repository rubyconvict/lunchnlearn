package main

import (
	"github.com/gin-gonic/gin"
	proto "github.com/rubyconvict/lunchnlearn/proto/lunchnlearn"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewLunchnlearnServiceClient(conn)

	g := gin.Default()
	// TODO: change to POST with data in RESTful style
	g.GET("/new_order/:order", func(ctx *gin.Context) {
		order := ctx.Param("order")
		if order == "car" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameter order"})
			return
		}

		req := &proto.LunchRequest{Order: order}
		if response, err := client.PlaceOrder(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"name":  response.GetLunch().GetName(),
				"state": response.GetLunch().GetState().String(),
			})
		} else {
			e := status.Convert(err)
			switch e.Code() {
			case codes.InvalidArgument:
				ctx.JSON(http.StatusBadRequest, gin.H{"error": e.Message()})
				return
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": e.Message()})
				return
			}
		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
