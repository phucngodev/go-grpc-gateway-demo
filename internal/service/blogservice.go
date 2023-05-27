package service

import (
	"context"

	pb "kratos-blog/api/blog/v1"
	"kratos-blog/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type BlogService struct {
	pb.UnimplementedBlogServiceServer
	log     *log.Helper
	article *biz.ArticleUsercase
}

func NewBlogService(article *biz.ArticleUsercase, logger log.Logger) *BlogService {
	return &BlogService{
		article: article,
		log:     log.NewHelper(logger),
	}
}

func (s *BlogService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleRespone, error) {
	s.log.Infof("input data %v", req)
	article := &biz.Article{
		Title:   req.Title,
		Content: req.Content,
	}

	err := s.article.Create(ctx, article)
	return &pb.CreateArticleRespone{Article: &pb.Article{
		Id:      article.ID,
		Title:   article.Title,
		Content: article.Content,
	}}, err
}

func (s *BlogService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleResponse, error) {
	s.log.Infof("input data %v", req)
	err := s.article.Update(ctx, req.Id, &biz.Article{
		Title:   req.Title,
		Content: req.Content,
	})

	return &pb.UpdateArticleResponse{}, err
}

func (s *BlogService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleRespone, error) {
	s.log.Infof("input data %v", req)
	err := s.article.Delete(ctx, req.Id)

	return &pb.DeleteArticleRespone{}, err
}

func (s *BlogService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.GetArticleRespone, error) {
	s.log.Infof("input data %v", req)
	if req.Id < 1 {
		return nil, pb.ErrorBlogInvalidId("invalid article id")
	}

	s.log.Infof("valid request ID", req.Id)

	article, err := s.article.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetArticleRespone{
		Article: &pb.Article{
			Id:      article.ID,
			Title:   article.Title,
			Content: article.Content,
		},
	}, nil
}

func (s *BlogService) ListArticle(ctx context.Context, req *pb.ListArticleRequest) (*pb.ListArticleResponse, error) {
	articles, err := s.article.List(ctx)
	resp := &pb.ListArticleResponse{}
	for _, a := range articles {
		resp.Results = append(resp.Results, &pb.Article{
			Id:      a.ID,
			Title:   a.Title,
			Content: a.Content,
		})
	}

	return resp, err
}
