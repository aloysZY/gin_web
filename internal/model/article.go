package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Article 文章结构体，操作数据库
type Article struct {
	State         uint8          `json:"-"`               // 状态  1 正常 0为禁用
	ArticleId     uint64         `json:"article_id"`      // 设置 tagID  string解决json解析的时候使用这个类型，解决前端传入和传入前端失真
	Title         string         `json:"name"`            // 文章标题
	Desc          string         `json:"desc"`            // 文章简述
	CoverImageUrl string         `json:"cover_image_url"` // 封面图片地址
	Content       string         `json:"content"`         // 文章内容
	*Model        `json:"model"` // 公共字段
	// json:"model" 本身是结构体嵌套，不写返回数据和 name 是在一层，写了之后，在一个新层里展示
}

func (a Article) TableName() string { return "web_article" }

// Create 创建文章
func (a Article) Create(db *gorm.DB) error {
	if err := db.Create(&a).Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}

// 根据文章标题查询所有文章,（没什么意义，应该要模糊匹配）
/*func (a Article) ListArticleByTitle(db *gorm.DB) ([]*Article, error) {
	var Articles []*Article
	db = db.Where("state = ?", a.State)
	if err := db.Where("title = ? AND is_del = ?", a.Title, 0).Order("modified_on DESC").Find(&Articles).Error; err != nil {
		return nil, err
	}
	return Articles, nil
}*/

func (a Article) ListArticle(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var articleList []*Article
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	db = db.Where("state = ?", a.State)
	if err := db.Where("is_del = ?", 0).Order("modified_on DESC").Find(&articleList).Error; err != nil {
		return nil, err
	}
	return articleList, nil
}

// ListArticleByTitle 根据文章标题查询所有文章,模糊匹配
func (a Article) ListArticleByTitle(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var articleList []*Article
	// 从第几条数据开始查
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	// db = db.Where("state = ?", a.State)
	// 模糊查找
	if err := db.Where(fmt.Sprintf("title LIKE '%%%s%%' AND state = %d AND is_del = %d", a.Title, a.State, 0)).Order("modified_on DESC").Find(&articleList).Error; err != nil {
		return nil, err
	}
	return articleList, nil
}

func (a Article) GetArticleCreatedByByArticleId(db *gorm.DB) ([]string, error) {
	var name []string
	err := db.Table(Auth{}.TableName()).Where("user_id=?", a.CreatedBy).Pluck("app_key", &name).Error
	return name, err
}

// func (a Article) CountArticleByTagID(db *gorm.DB, tagId uint64) (int, error) {
// 	var count int
// 	if err := db.Table(ArticleIdTagId{}.TableName()+"AS at").
// 		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.tag_id").
// 		Joins("LEFT JOIN `"+Article{}.TableName()+"` AS a ON at.article_id = a.article_id").
// 		Where("at.tag_id = ? AND a.state = ? AND a.is_del = ?", tagId, a.State, 0).
// 		Count(&count).Error; err != nil {
// 		return 0, err
// 	}
// 	return count, nil
// }

// CountArticle 查询文章数量，查询所有文章
func (a Article) CountArticle(db *gorm.DB) (int, error) {
	var count int
	if err := db.Model(&a).Where("state = ? AND is_del = ?", a.State, 0).Order("modified_on DESC").Count(&count).Error; err != nil {
		return 0, err
	}
	// // 如果标题存在，根据标题查询，否则根据 tagId 查询
	// if a.Title != "" {
	// 	if err := db.Model(&a).Where(fmt.Sprintf("title LIKE '%%%s%%' AND state = %d AND is_del = %d", a.Title, a.State, 0)).Order("modified_on DESC").Count(&count).Error; err != nil {
	// 		return 0, err
	// 	}
	// } else if tagId != 0 {
	// 	if err := db.Table(ArticleIdTagId{}.TableName()+"AS at").
	// 		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.tag_id").
	// 		Joins("LEFT JOIN `"+Article{}.TableName()+"` AS a ON at.article_id = a.article_id").
	// 		Where("at.tag_id = ? AND a.state = ? AND a.is_del = ?", tagId, a.State, 0).
	// 		Count(&count).Error; err != nil {
	// 		return 0, err
	// 	}
	// }
	return count, nil
}

// CountArticleByTitle 根据标题模糊查找文章数量
func (a Article) CountArticleByTitle(db *gorm.DB) (int, error) {
	var count int
	if err := db.Model(&a).Where(fmt.Sprintf("title LIKE '%%%s%%' AND state = %d AND is_del = %d", a.Title, a.State, 0)).Order("modified_on DESC").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
