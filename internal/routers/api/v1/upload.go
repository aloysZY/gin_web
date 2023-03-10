package v1

import (
	"strconv"

	"gin_web/internal/service"
	"gin_web/pkg/app"
	"gin_web/pkg/errcode"
	"gin_web/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct{}

func NewUpload() Upload { return Upload{} }

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	// 获取上传的文件，  <input type="file" name="file"> file是前端文件的名字
	// file, fileHeader, err := c.Request.FormFile("file") // 文件名称
	// fileHeader, err := c.FormFile("file")
	// if err != nil {
	// 	response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
	// 	return
	// }
	// 修改为多文件上传
	var AccessUrls []string
	multipartForm, err := c.MultipartForm()
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	for _, fileHeader := range multipartForm.File["file"] {
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
			response.ToErrorResponse(errcode.ErrorFileUpload.WithDetails(err.Error()))
			return
		}
		AccessUrls = append(AccessUrls, fileInfo.AccessUrl)
	}

	response.ToResponse(gin.H{
		"file_access_url": AccessUrls,
	})

}
