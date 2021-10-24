package graph

import (
	"context"
	"go_grpc/article/client"
	"go_grpc/article/pb"
	"go_grpc/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ArticleClient *client.Client
}

func (r *mutationResolver) CreateArticle(ctx context.Context, input model.CreateInput) (*model.Article, error) {
	// gRPCサーバーでArticleをCREATE
	article, err := r.ArticleClient.CreateArticle(
		ctx,
		&pb.ArticleInput{
			Author:  input.Author,
			Title:   input.Title,
			Content: input.Content,
		})
	if err != nil {
		return nil, err
	}

	// CREATEしたArticleを返す
	return article, nil
}

func (r *mutationResolver) UpdateArticle(ctx context.Context, input model.UpdateInput) (*model.Article, error) {
	// gRPCサーバーでArticleをUPDATE
	article, err := r.ArticleClient.UpdateArticle(
		ctx,
		int64(input.ID),
		&pb.ArticleInput{
			Author:  input.Author,
			Title:   input.Title,
			Content: input.Content,
		})
	if err != nil {
		return nil, err
	}

	// UPDATEしたArticleを返す
	return article, nil
}

func (r *mutationResolver) DeleteArticle(ctx context.Context, input int) (int, error) {
	// gRPCサーバーでArticleをDELETE
	id, err := r.ArticleClient.DeleteArticle(ctx, int64(input))
	if err != nil {
		return 0, err
	}

	// DELETEしたArticleのIDを返す
	return int(id), nil
}

func (r *queryResolver) Article(ctx context.Context, input int) (*model.Article, error) {
	// 入力したIDの記事をgRPCサーバーからREAD
	article, err := r.ArticleClient.ReadArticle(ctx, int64(input))
	if err != nil {
		return nil, err
	}

	// READしたArticleを返す
	return article, nil
}

func (r *queryResolver) Articles(ctx context.Context) ([]*model.Article, error) {
	// gRPCサーバーでArticleを全取得
	articles, err := r.ArticleClient.ListArticle(ctx)
	if err != nil {
		return nil, err
	}

	// 全取得したArticleを返す
	return articles, nil
}
