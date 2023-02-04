package service

import (
	"github.com/aloysZy/gin_web/internal/dao"
	"github.com/aloysZy/gin_web/internal/model"
	"github.com/aloysZy/gin_web/internal/routers/api/params"
	"github.com/aloysZy/gin_web/pkg/app"
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
		// 在这里还要在添加一步，执行svc.dao.CreateArticleTag之前要先确保 tagid 数据库存在
		err := svc.dao.GetTagByTagId(param.TagId)
		if err != nil {
			zap.L().Error("svc.dao.GetTagByTagId failed", zap.Error(err))
			return err
		}

		if err = svc.dao.CreateArticleTag(param.ArticleId, param.TagId, param.CreatedBy); err != nil {
			zap.L().Error("svc.dao.CreateArticleTag failed", zap.Error(err))
			return err
		}
	}

	// 提交
	if err = svc.dao.Engine.Commit().Error; err != nil {
		zap.L().Error("svc.dao.Engine.Commit() failed", zap.Error(err))
		return err
	}
	return nil
}

// ListArticleS 获取文章列表
func (svc *Service) ListArticleS(param *params.ListArticleRequest, pager *app.Pager) ([]*model.ArticleTag, int, error) {
	var (
		articleList []*model.Article
		totalRows   int
		err         error
	)
	// 根据标题模糊搜索
	if param.Title != "" {
		// 直接根据文章名称去查文章标题，找到返回所有找到的
		articleList, err = svc.dao.ListArticleByTitle(param.Title, param.State, pager.Page, pager.PageSize)
		if err != nil {
			zap.L().Error("svc.dao.ListArticleByTitle failed", zap.Error(err))
			return nil, 0, err
		}
		if len(articleList) == 0 {
			zap.L().Info("svc.dao.ListArticleByTitle", zap.Error(errcode.ErrorNotArticle))
			return nil, 0, errcode.ErrorNotArticle
		}
		totalRows, err = svc.dao.CountArticleByTitle(param.Title, param.State)
		if err != nil {
			zap.L().Error("svc.dao.CountArticleByTitle failed", zap.Error(err))
			return nil, 0, err
		}
	} else { // 查找全部的
		articleList, err = svc.dao.ListArticle(param.State, pager.Page, pager.PageSize)
		if err != nil {
			zap.L().Error("svc.dao.ListArticle failed", zap.Error(err))
			return nil, 0, err
		}
		totalRows, err = svc.dao.CountArticle(param.State)
		if err != nil {
			zap.L().Error("svc.dao.CountArticle failed", zap.Error(err))
			return nil, 0, err
		}
	}

	// 不在函数外调用方法了，在这里做吧
	articleTageList := make([]*model.ArticleTag, 0, len(articleList)) // 先初始化，避免数据量太大频繁申请内存
	for _, article := range articleList {
		articleTage := new(model.ArticleTag)
		// 根据文章 ID 查找标签名称
		tagNameList, err := svc.dao.ListTagNameByArticleId(article.ArticleId, article.State)
		if err != nil {
			zap.L().Error("svc.dao.ListTagNameByArticleId failed", zap.Error(err))
			return nil, 0, err
		}
		articleTage.Article = article // 复制文章信息
		// articleTage.TagName = make([]string, 0, len(tagNameList)) //这个应该写 append 就行
		// articleTage.TagName = tagNameList ///这个应该写 append 就行,数据量不是很大，不会台频繁的申请内存
		articleTage.TagName = append(articleTage.TagName, tagNameList...)

		// 根据文章 ID 查找创建人名称
		userName, err := svc.dao.GetUserNameByArticleCreatedBy(article.CreatedBy)
		if err != nil {
			zap.L().Error("svc.dao.GetUserNameByArticleCreatedBy failed", zap.Error(err))
			return nil, 0, err
		}
		// articleTage.TagName = make([]string, 0, len(tagNameList)) 这里同理
		// articleTage.UserName = userName
		articleTage.UserName = append(articleTage.UserName, userName...)

		// 整体数据添加
		articleTageList = append(articleTageList, articleTage)
	}

	// 修改返回文章中的创建人"created_by": 444315400298037249 查询为对于的name
	/*	newArticleList, err := svc.dao.GetUserNameByArticleCreatedBy(articleList)
		if err != nil {
			zap.L().Error("svc.dao.GetUserNameByArticleCreatedBy failed", zap.Error(err))
			return nil, 0, err
		}*/ // 这个也不循环了，放在for _, article := range articleList 这里做

	return articleTageList, totalRows, err
}

// 放在ListArticleS里面实现了
// func (svc *Service) ListTagNameByArticleId(articleList []*model.Article) ([]*model.ArticleTag, error) {
// 	// var articleTageList []*model.ArticleTag
// 	articleTageList := make([]*model.ArticleTag, 0, len(articleList)) // 先预分配空间，大数据提高性能
// 	articleTage := new(model.ArticleTag)                              // 初始化指针
// 	for _, article := range articleList {
// 		tagName, err := svc.dao.ListTagNameByArticleId(article.ArticleId)
// 		if err != nil {
// 			zap.L().Error("svc.dao.ListTagNameByArticleId failed", zap.Error(err))
// 			return nil, err
// 		}
// 		// articleTage.TagName = tagName  先要初始化才能赋值，这是一个切片
// 		articleTage.TagName = make([]string, 0, len(tagName))
// 		articleTage.TagName = tagName
//
// 		articleTage.Article = article
// 		articleTageList = append(articleTageList, articleTage)
// 	}
// 	return articleTageList, nil
// }
