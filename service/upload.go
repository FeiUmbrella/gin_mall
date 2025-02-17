package service

import (
	"gin_mall/conf"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

// UploadAvatarToLocalStatic 用户上传头像到本地
func UploadAvatarToLocalStatic(file multipart.File, uId uint, userName string) (filePath string, err error) {
	bId := strconv.Itoa(int(uId)) // 将uid变为string，后面要与文件路径拼接
	basePath := "." + conf.AvatarPath + "user" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreatDir(basePath)
	}
	avatarPath := basePath + userName + ".jpg" // todo: 把 file 的后缀提取出来
	content, err := io.ReadAll(file)           // 先转化为byte[]
	if err != nil {
		return "", err
	}
	err = os.WriteFile(avatarPath, content, 0666) // 写入文件
	if err != nil {
		return "", err
	}
	// 返回的这个路径才是要存放在数据库的，每个用户头像存放路径只有这个部分不同
	return "user" + bId + "/" + userName + ".jpg", nil
}

// DirExistOrNot 判断文件夹是否存在
func DirExistOrNot(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// CreatDir 创建文件夹
func CreatDir(path string) bool {
	err := os.MkdirAll(path, 755)
	if err != nil {
		return false
	}
	return true
}
