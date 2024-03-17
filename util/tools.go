package util

import (
	"crypto/md5"
	"errors"
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	img "github.com/hunterhug/go_image"
	"mi_shop/database"
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

// GetUnixNano 获取纳秒时间戳
func GetUnixNano() int64 {
	return time.Now().UnixNano()
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

// Md5 md5加密
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
		return "", err
	}

	// 4、生成文件名称和文件保存的目录   111111111111.jpeg
	fileName := strconv.FormatInt(GetUnixNano(), 10) + extName

	// 5、执行上传
	dst := path.Join(dir, fileName)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		return "", err
	}
	return dst, nil
}

// Float 字符串转换成float64
func Float(str string) (n float64, err error) {
	n, err = strconv.ParseFloat(str, 64)
	return
}

// Str2Html 字符串转换成html
func Str2Html(str string) template.HTML {
	return template.HTML(str)
}

// GetSettingFromColumn 通过列获取设置
func GetSettingFromColumn(columnName string) string {
	// redis file
	setting := database.Setting{}
	database.DB.First(&setting)
	// 反射来获取
	v := reflect.ValueOf(setting)
	val := v.FieldByName(columnName).String()
	return val
}

// ResizeGoodsImage 生成商品缩略图
func ResizeGoodsImage(filename string) {
	extname := path.Ext(filename)
	thumbnailSize := strings.ReplaceAll(GetSettingFromColumn("ThumbnailSize"), "，", ",")
	thumbnailSizeSlice := strings.Split(thumbnailSize, ",")

	// static/upload/tao_400.png
	// static/upload/tao_400.png_100x100.png
	for i := 0; i < len(thumbnailSizeSlice); i++ {
		savePath := filename + "_" + thumbnailSizeSlice[i] + "x" + thumbnailSizeSlice[i] + extname
		w, _ := strconv.Atoi(thumbnailSizeSlice[i])
		err := img.ThumbnailF2F(filename, savePath, w, w)
		if err != nil {
			fmt.Println("生成图片失败", err)
		}
	}
}

// Sub 减法
func Sub(a, b int) int {
	return a - b
}

// Mul 乘法
func Mul(price float64, num int) float64 {
	return price * float64(num)
}

// Substr 截取字符串
func Substr(str string, start, end int) string {
	rs := []rune(str)
	rl := len(rs)

	if start < 0 || start > rl {
		start = 0
	}
	if end < 0 || end > rl {
		end = rl
	}
	if start > end {
		start, end = end, start
	}

	return string(rs[start:end])
}

// FormatAttr 将商品属性格式化成markdown
func FormatAttr(str string) (result string) {
	tempSlice := strings.Split(str, "\n")
	for _, v := range tempSlice {
		md := []byte(v)
		output := markdown.ToHTML(md, nil, nil)
		result += string(output)
	}
	return
}

// GetRandomNum 生成4位随机数
func GetRandomNum() (str string) {
	for i := 0; i < 4; i++ {
		current := rand.Intn(10)
		str += strconv.Itoa(current)
	}
	fmt.Println(str)
	return
}

// GetOrderId 获取订单号
func GetOrderId() string {
	// 2022020312233
	template := "20060102150405"
	return time.Now().Format(template) + GetRandomNum()
}
