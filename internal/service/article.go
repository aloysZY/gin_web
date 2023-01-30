package service

import (
	"github.com/aloysZy/gin_web/internal/dao"
	"github.com/aloysZy/gin_web/internal/routers/api/params"
	"github.com/aloysZy/gin_web/pkg/errcode"
	"github.com/aloysZy/gin_web/pkg/setting"
	"go.uber.org/zap"
)

func (svc *Service) CreateArticle(param *params.CreateArticleRequest) error {
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
	err = svc.dao.CreateArticle(&dao.Article{
		State:         param.State,
		ArticleId:     param.ArticleId,
		TagId:         param.TagId,
		CreatedBy:     param.CreatedBy,
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
	})
	if err != nil {
		zap.L().Error("svc.dao.CreateArticle failed", zap.Error(err))
		return err
	}
	return nil
}
