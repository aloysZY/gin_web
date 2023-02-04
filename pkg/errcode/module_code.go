package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(20010002, "创建标签失败")
	ErrorUpdateTagFail  = NewError(20010003, "更新标签失败")
	ErrorDeleteTagFail  = NewError(20010004, "删除标签失败")
	ErrorCountTagFail   = NewError(20010005, "统计标签失败")
	ErrorTagExists      = NewError(20010006, "标签已经存在")
	ErrorGetTagFail     = NewError(20010007, "标签不存在")

	ErrorGetArticleFail    = NewError(20020001, "获取单个文章失败")
	ErrorGetArticlesFail   = NewError(20020002, "获取多个文章失败")
	ErrorCreateArticleFail = NewError(20020003, "创建文章失败")
	ErrorUpdateArticleFail = NewError(20020004, "更新文章失败")
	ErrorDeleteArticleFail = NewError(20020005, "删除文章失败")
	ErrorNotArticle        = NewError(20020006, "没有对应的文章")

	ErrorFileUpload             = NewError(20030001, "上传文件失败")
	ErrorFileSuffixNotSupported = NewError(20030002, "文件后缀不符合类型要求")
	ErrorFileExceededMaximum    = NewError(20030003, "超过文件大小限制")
	ErrorFileCreateDirector     = NewError(20030004, "创建文件保存路径失败")
	ErrorFileNotPermissions     = NewError(20030005, "保存文件权限不足")

	ErrorSonyFlakeNotInit = NewError(20040001, "雪花算法没有初始化")
	ErrorNoDataModified   = NewError(20040002, "没有可修改数据")
)
