package model

type ArticleTag struct {
	*Model
	TagId     uint32 `json:"tag_id"`
	ArticleId uint32 `json:"article_id"`
}

// gorm 所需的方法
func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
