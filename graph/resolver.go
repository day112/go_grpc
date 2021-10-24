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
