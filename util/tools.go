package util

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// UnixToTime 时间戳转换成日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// DateToUnix 日期转换成时间戳 2020-05-02 15:04:05
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// GetUnix 获取时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}

// GetDate 获取当前的日期
func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

// GetDay 获取年月日
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

func Md5(pwd string) string {
	h := md5.New()
	io.WriteString(h, pwd)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// UploadImg 根据日期上传图片
func UploadImg(c *gin.Context, picName string) (string, error) {
	// 1、获取上传的文件
	file, err := c.FormFile(picName)
	if err != nil {
		return "", err
	}

	// 2、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("文件后缀名不合法")
	}

	// 3、创建图片保存目录  static/upload/20210624
	day := GetDay()
	dir := "./static/upload/" + day

	err = os.MkdirAll(dir, 0666)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// 4、生成文件名称和文件保存的目录   111111111111.jpeg
	fileName := strconv.FormatInt(GetUnix(), 10) + extName

	// 5、执行上传
	dst := path.Join(dir, fileName)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		return "", err
	}
	return dst, nil
}
