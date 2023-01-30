package service

import (
	"mime/multipart"
	"os"

	"github.com/aloysZy/gin_web/global"
	"github.com/aloysZy/gin_web/pkg/errcode"
	"github.com/aloysZy/gin_web/pkg/upload"
	"go.uber.org/zap"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

// UploadFile 保存文件
// func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
func (svc *Service) UploadFile(fileType upload.FileType, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	if !upload.CheckContainExt(fileType, fileName) {
		zap.L().Error("upload.CheckContainExt failed")
		// return nil, errors.New("file suffix is not supported")
		return nil, errcode.ErrorFileSuffixNotSupported

	}

	// if upload.CheckMaxSize(fileType, file) 修改了传入参数
	if upload.CheckMaxSize(fileType, fileHeader) {
		zap.L().Error("upload.CheckMaxSize failed")
		// return nil, errors.New("exceeded maximum file limit")
		return nil, errcode.ErrorFileExceededMaximum
	}

	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			zap.L().Error("upload.CreateSavePath failed")
			// return nil, errors.New("failed to create save directory")
			return nil, errcode.ErrorFileCreateDirector
		}
	}

	if upload.CheckPermission(uploadSavePath) {
		zap.L().Error("upload.CheckPermission failed")
		// return nil, errors.New("insufficient file permissions")
		return nil, errcode.ErrorFileNotPermissions
	}

	dst := uploadSavePath + "/" + fileName
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		zap.L().Error("upload.SaveFile failed error:", zap.Error(err))
		return nil, err
	}

	accessUrl := global.AppSetting.UploadImage.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
