package model

import "blog-service/pkg/app"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"coverImageUrl"`
	State         uint8  `json:"state"`
}

// article.go
type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (a Article) TableName() string {
	return "blog_article"
}
