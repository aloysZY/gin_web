package upload

import (
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"gin_web/global"
	"gin_web/pkg/util"
)

type FileType int

const TypeImage FileType = iota + 1

// GetFileName 获取文件名称
func GetFileName(name string) string {
	ext := GetFileExt(name)
	// 分割文件后缀
	fileName := strings.TrimSuffix(name, ext)
	// 加密文件名称
	fileName = util.EncodeMD5(fileName)
	// 返回加密的文件名称和后缀
	return fileName + ext
}

// GetFileExt 文件后缀
func GetFileExt(name string) string {
	return path.Ext(name)
}

// GetSavePath 文件保存路径
func GetSavePath() string {
	return global.AppSetting.UploadImage.UploadSavePath
}

// CheckContainExt 判断后缀是否符合要求
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	// 根据上传的文件类型来判断后缀，要这两个匹配
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImage.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}
	return false
}

// CheckMaxSize 判断文件大小
// func CheckMaxSize(t FileType, f multipart.File) bool {
func CheckMaxSize(t FileType, f *multipart.FileHeader) bool {
	// 这步骤有点多余了，fileheader 里面有 size，就是文件大小，不需要读到内存
	// content, _ := ioutil.ReadAll(f)
	// size := len(content)
	size := int(f.Size)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImage.UploadImageMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

// CheckSavePath 检查保存路径
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

// CreateSavePath 创建文件目录
func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

func CheckPermission(dst string) bool {
	// https://blog.csdn.net/wohu1104/article/details/106433529
	// 获取文件的描述
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

// SaveFile 保存文件
func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			return
		}
	}(src)

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			return
		}
	}(out)

	_, err = io.Copy(out, src)
	return err
}
