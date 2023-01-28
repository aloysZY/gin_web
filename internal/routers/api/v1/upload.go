package v1

import (
	"strconv"

	"github.com/aloysZy/gin_web/internal/service"
	"github.com/aloysZy/gin_web/pkg/app"
	"github.com/aloysZy/gin_web/pkg/errcode"
	"github.com/aloysZy/gin_web/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct{}

func NewUpload() Upload { return Upload{} }

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	// 获取上传的文件，  <input type="file" name="file"> file是前端文件的名字
	// file, fileHeader, err := c.Request.FormFile("file") // 文件名称
	fileHeader, err := c.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	// 这里传入的 type 是和文件类型匹配的，文件类型不同传入的后缀和大小可以分别限制
	// fileType := convert.StrTo(c.PostForm("type")).MustInt()
	fileType, _ := strconv.Atoi(c.PostForm("type"))
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	// fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), fileHeader)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})

}
