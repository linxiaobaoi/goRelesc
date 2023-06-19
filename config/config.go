package config

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Response struct {
	Code    int
	Message string
	Data    interface{}
	Time    string
}

// 获取当前日期
func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

// 获取当前日期
func GetDates() (string, string, string) {
	month := time.Unix(1557042972, 0).Format("1")
	year := time.Now().Format("2006")
	month = time.Now().Format("01")
	day := time.Now().Format("02")
	return year, month, day
}

// 随机字符串 默认为100
func RandNumsers(number int) (string, bool) {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())
	if number == 0 {
		number = 10
	}
	if number > 20 {
		return "", false
	}
	// 生成指定长度的随机字符串
	length := number
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	randomStr := make([]rune, length)
	for i := range randomStr {
		randomStr[i] = letters[rand.Intn(len(letters))]
	}
	return string(randomStr), true
}

// 获取当前时间戳
func GetTimes() string {
	t := time.Now()
	timestamp := int(t.Unix())
	return strconv.Itoa(timestamp)
}

// 生成token
func CreateToken(id int) string {
	token := Md5Encode(string(id) + GetTimes())
	fmt.Println("token打印====", token)
	return token
}

// md5 加密处理---小写
func Md5Encode(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// md5 加密处理---大写
func Md5EncodeUpper(str string) string {
	return strings.ToUpper(Md5Encode(str))
}

// 密码加密
func MakePassword(plainPassword, salt string) string {
	return Md5Encode(plainPassword + salt)
}

// 密码验证
func CheckPasswod(plainPassword, salt string, password string) bool {
	fmt.Println("plaintPassword", Md5Encode(plainPassword+salt))
	fmt.Println("password", password)
	return Md5Encode(plainPassword+salt) == password
}

// 过期时间
func ExpTimes(exp int) int {
	if exp == 0 {
		exp = 30
	}
	t := time.Now()
	t = t.Add(time.Duration(exp) * time.Second)
	timestamp := int(t.Unix())
	return timestamp
}

// 获取当前时间整型
func ExpTime() int {
	t := time.Now()
	timestamp := int(t.Unix())
	return timestamp
}

// 时间戳转换成日期型
func TimeVarDate(timeStamp int) string {
	//返回time对象
	t := time.Unix(int64(timeStamp), 0)

	//返回string
	dateStr := t.Format("2006-01-02 15:04:05")
	return dateStr
}

// exp = 0  30 分钟
// exp = 1  7 天
// exp = 2  1个月
// exp = 3  1 年
// default  5分钟
func ExpLates(exp int64) int {
	expTimes := 0
	if exp == 0 {
		expTimes = 30 * 60
	} else if exp == 1 {
		expTimes = 7 * 24 * 60 * 60
	} else if exp == 2 {
		expTimes = 30 * 24 * 60 * 60
	} else if exp == 3 {
		expTimes = 365 * 24 * 60 * 60
	} else {
		expTimes = 5 * 60
	}

	return expTimes

}

// 随机字符串 默认为100
func RandNumber() string {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())
	number := 6
	// 生成指定长度的随机字符串
	length := number
	letters := []rune("123456789")
	randomStr := make([]rune, length)
	for i := range randomStr {
		randomStr[i] = letters[rand.Intn(len(letters))]
	}
	return string(randomStr)
}

// redis 设置处理数据
func RedisSelect(value interface{}) string {
	bytes, _ := json.Marshal(value)
	stringData := string(bytes)
	return stringData
}

// 判断值是否在数组内
func InArray(target string, str_array []string) bool {
	status := 0
	for _, element := range str_array {

		if target == element {

			status = 1

		} else {
			status = 0
		}
	}
	if status == 0 {
		return false
	} else {
		return true
	}
}

// 获取返回消息
func ReturnMessage(code int, msg string, data interface{}) interface{} {
	var message Response
	message.Code = code
	message.Message = msg
	message.Data = data
	message.Time = GetDate()

	return message
}

// 获取返回消息
func ReturnErrorMessage(code int, msg string) interface{} {
	var message Response
	message.Code = code
	message.Message = msg
	message.Time = GetDate()
	return message
}
