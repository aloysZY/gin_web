package model

import "github.com/jinzhu/gorm"

// Article 文章结构体
type Article struct {
	State         uint8          `json:"state"`      // 状态  1 正常 0为禁用
	ArticleId     uint64         `json:"article_id"` // 设置 tagID  string解决json解析的时候使用这个类型，解决前端传入和传入前端失真
	TagId         uint64         `json:"tagId"`
	Title         string         `json:"name"`            // 文章标题
	Desc          string         `json:"desc"`            // 文章简述
	CoverImageUrl string         `json:"cover_image_url"` // 封面图片地址
	Content       string         `json:"content"`         // 文章内容
	*Model        `json:"model"` // 公共字段
	// json:"model" 本身是结构体嵌套，不写返回数据和 name 是在一层，写了之后，在一个新层里展示
}

func (a Article) TableName() string { return "web_article" }

func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}
