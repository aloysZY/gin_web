package service

import (
	"github.com/aloysZy/gin_web/internal/dao"
	"github.com/aloysZy/gin_web/internal/routers/api/params"
	"github.com/aloysZy/gin_web/pkg/errcode"
	"github.com/aloysZy/gin_web/pkg/setting"
	"go.uber.org/zap"
)

// CreateArticle 创建文章
func (svc *Service) CreateArticle(param *params.CreateArticleRequest) error {
	// 创建文章 ID
	articleId, err := setting.GetID()
	if err != nil {
		if err == errcode.ErrorSonyFlakeNotInit {
			zap.L().Error("setting.GetID failed", zap.Error(err))
			return err
		}
		zap.L().Error("setting.GetID failed", zap.Error(err))
		return err
	}
	param.ArticleId = articleId

	// 想法是初始化一个事务句柄，这个程序后续的都使用这个事务句柄操作
	svc.dao.Engine = svc.dao.Engine.Begin()

	// 执行 dao 创建文章
	if err = svc.dao.CreateArticle(&dao.Article{
		State:     param.State,
		ArticleId: param.ArticleId,
		// TagId:         param.TagId,
		CreatedBy:     param.CreatedBy,
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
	}); err != nil {
		zap.L().Error("svc.dao.CreateArticle failed", zap.Error(err))
		return err
	}

	// 执行 dao 记录文章ID和标签 ID
	// 创建文章和标签的关联，因为文章可以没有标签，所以这里不设置事务（其实要设置，文章至少有一个标签）
	if param.TagId != 0 {
		if err = svc.dao.CreateArticleTag(param.ArticleId, param.TagId); err != nil {
			zap.L().Error("svc.dao.CreateArticleTag failed", zap.Error(err))
			return err
		}
	}

	// 提交事务
	if err = svc.dao.Engine.Commit().Error; err != nil {
		zap.L().Error("svc.dao.Engine.Commit() failed", zap.Error(err))
		return err
	}
	return nil
}

func (svc *Service) CountArticle(param *params.CountArticleRequest) {

}

func (svc *Service) ListArticle(articleId uint64, state uint8) {

}
