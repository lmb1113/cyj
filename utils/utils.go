package utils

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"os"
	"regexp"
	"runtime/debug"
	"time"
)

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Go(x func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(fmt.Sprintf("panicGo %s\n", err))
				fmt.Println(fmt.Sprint(string(debug.Stack())))
			}
		}()
		x()
	}()
}

func ValidateIPPortFormat(str string) bool {
	// 正则表达式模式匹配IP:Port格式
	pattern := `^(?:[0-9]{1,3}\.){3}[0-9]{1,3}:[0-9]{1,5}$`
	match, _ := regexp.MatchString(pattern, str)
	return match
}

// 随机范围在30000-40000
func GenerateRemotePort() int {
	return rand.Intn(10000) + 30000
}

// 随机name
func GenerateName() string {
	return uuid.New().String()
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

// RandomString 生成随机字符串
func RandomString(n int, seed int64) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetRandomElement[T any](slice []T) (T, error) {
	if len(slice) == 0 {
		var zero T
		return zero, fmt.Errorf("切片不能为空")
	}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(slice))
	return slice[index], nil
}
