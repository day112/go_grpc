package main

import (
	"context"
	"fmt"
	"go_grpc/article/client"
	"go_grpc/article/pb"
	"log"
)

// gRPCサーバーの動作確認用
func main() {
	// clientを生成
	c, _ := client.NewClient("localhost:50051")

	create(c)
	read(c)
	update(c)
	delete(c)
	list(c)
}

// 記事をCREATE
func create(c *client.Client) {
	input := &pb.ArticleInput{
		Author:  "gopher",
		Title:   "gRPC",
		Content: "gRPC is so cool!",
	}
	res, err := c.Service.CreateArticle(context.Background(), &pb.CreateArticleRequest{ArticleInput: input})
	if err != nil {
		log.Fatalf("Failed to CreateArticle: %v\n", err)
	}
	fmt.Printf("CreateArticle Response: %v\n", res)
}

// 記事をREAD
func read(c *client.Client) {
	var id int64 = 2
	res, err := c.Service.ReadArticle(context.Background(), &pb.ReadArticleRequest{Id: id})
	if err != nil {
		log.Fatalf("Failed to ReadArticle: %v\n", err)
	}
	fmt.Printf("ReadArticle Response: %v\n", res)
}

// 記事をUPDATE
func update(c *client.Client) {

}

// 記事をDELETE
func delete(c *client.Client) {

}

// 記事を全取得
func list(c *client.Client) {

}
