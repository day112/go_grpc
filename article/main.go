package main

import (
	"go_grpc/article/client"
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

}

// 記事をREAD
func read(c *client.Client) {

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
