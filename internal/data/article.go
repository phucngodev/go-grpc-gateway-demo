package data

import (
	"context"
	"kratos-blog/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.ArticleRepo = (*articleRepo)(nil)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ar *articleRepo) ListArticle(ctx context.Context) ([]*biz.Article, error) {
	var articles []*biz.Article
	result := ar.data.db.Find(&articles)
	return articles, result.Error
}

func (ar *articleRepo) GetArticle(ctx context.Context, id int64) (*biz.Article, error) {
	var article biz.Article
	result := ar.data.db.First(&article, id)

	return &article, result.Error
}

func (ar *articleRepo) CreateArticle(ctx context.Context, article *biz.Article) error {
	ar.data.db.Create(article)
	return nil
}

func (ar *articleRepo) UpdateArticle(ctx context.Context, id int64, article *biz.Article) error {
	return nil
}

func (ar *articleRepo) DeleteArticle(ctx context.Context, id int64) error {
	return nil
}
