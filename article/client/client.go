package client

import (
	"context"
	"go_grpc/article/pb"
	"go_grpc/graph/model"
	"io"

	"google.golang.org/grpc"
)

// Clientからサービスを呼び出せるようにする
type Client struct {
	conn    *grpc.ClientConn
	Service pb.ArticleServiceClient
}

func NewClient(url string) (*Client, error) {
	// client connectionを生成
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	// articleサービスのclientを生成
	c := pb.NewArticleServiceClient(conn)

	// articleサービスのclientを返す
	return &Client{conn, c}, nil
}

func (c *Client) CreateArticle(ctx context.Context, input *pb.ArticleInput) (*model.Article, error) {
	// CREATE処理のレスポンスを受け取る
	res, err := c.Service.CreateArticle(
		ctx,
		&pb.CreateArticleRequest{ArticleInput: input},
	)
	if err != nil {
		return nil, err
	}

	// GraphQLサービスで扱える形にしてCREATEしたArticleを返す
	return &model.Article{
		ID:      int(res.Article.Id),
		Author:  res.Article.Author,
		Title:   res.Article.Title,
		Content: res.Article.Content,
	}, nil
}

func (c *Client) UpdateArticle(ctx context.Context, id int64, input *pb.ArticleInput) (*model.Article, error) {
	// UPDATE処理のレスポンスを受け取る
	res, err := c.Service.UpdateArticle(ctx, &pb.UpdateArticleRequest{Id: id, ArticleInput: input})
	if err != nil {
		return nil, err
	}

	// GraphQLサービスで扱える形にしてUPDATEしたArticleを返す
	return &model.Article{
		ID:      int(res.Article.Id),
		Author:  res.Article.Author,
		Title:   res.Article.Title,
		Content: res.Article.Content,
	}, nil
}

func (c *Client) ReadArticle(ctx context.Context, id int64, input *pb.ArticleInput) (*model.Article, error) {
	// READ処理のレスポンスを受け取る
	res, err := c.Service.ReadArticle(ctx, &pb.ReadArticleRequest{Id: id})
	if err != nil {
		return nil, err
	}

	// GraphQLサービスで扱える形にしてREADしたArticleを返す
	return &model.Article{
		ID:      int(res.Article.Id),
		Author:  res.Article.Author,
		Title:   res.Article.Title,
		Content: res.Article.Content,
	}, nil
}

func (c *Client) DeleteArticle(ctx context.Context, id int64) (int64, error) {
	// DELETE処理のレスポンスを受け取る
	res, err := c.Service.DeleteArticle(ctx, &pb.DeleteArticleRequest{Id: id})
	if err != nil {
		return 0, err
	}

	// DELETEしたArticleのIDを返す
	return res.Id, nil
}

func (c *Client) ListArticle(ctx context.Context) ([]*model.Article, error) {
	// 全取得処理のレスポンスを受け取る
	res, err := c.Service.ListArticle(ctx, &pb.ListArticleRequest{})
	if err != nil {
		return nil, err
	}

	// GraphQLサービスで扱える形にして全取得した記事をスライスとして返す
	var articles []*model.Article
	for {
		r, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		articles = append(articles, &model.Article{
			ID:      int(r.Article.Id),
			Author:  r.Article.Author,
			Title:   r.Article.Title,
			Content: r.Article.Content,
		})
	}
	return articles, nil
}
