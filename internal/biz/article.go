package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Article struct {
	ID        int64
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ArticleRepo interface {
	ListArticle(ctx context.Context) ([]*Article, error)
	GetArticle(ctx context.Context, id int64) (*Article, error)
	CreateArticle(ctx context.Context, article *Article) error
	UpdateArticle(ctx context.Context, id int64, article *Article) error
	DeleteArticle(ctx context.Context, id int64) error
}

type ArticleUsercase struct {
	repo ArticleRepo
}

func NewArticleUsercase(repo ArticleRepo, logger log.Logger) *ArticleUsercase {
	return &ArticleUsercase{
		repo: repo,
	}
}

func (uc *ArticleUsercase) List(ctx context.Context) ([]*Article, error) {
	articles, err := uc.repo.ListArticle(ctx)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (ur *ArticleUsercase) Get(ctx context.Context, id int64) (*Article, error) {
	article, err := ur.repo.GetArticle(ctx, id)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (uc *ArticleUsercase) Create(ctx context.Context, article *Article) error {
	return uc.repo.CreateArticle(ctx, article)
}

func (uc *ArticleUsercase) Update(ctx context.Context, id int64, article *Article) error {
	return uc.repo.UpdateArticle(ctx, id, article)
}

func (uc *ArticleUsercase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteArticle(ctx, id)
}
