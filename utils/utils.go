package utils

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"regexp"
	"runtime/debug"
)

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
