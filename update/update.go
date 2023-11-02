package update

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"qh-tool/global"
	"strings"
	"time"
)

const CheckUpdateUrl = "https://c.0a0a.cn/update.json"

type CheckUpdateResp struct {
	Version int    `json:"version"`
	Url     string `json:"url"`
	Msg     string `json:"msg"`
}

func HandleUpdate() {
	updateInfo, err := GetUpdateInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	if updateInfo.Version <= global.Version {
		return
	}
	fileURL := updateInfo.Url // 要下载的文件的URL
	savePath, err := os.Executable()
	if err != nil {
		return
	}
	newFilePath := strings.ReplaceAll(savePath, ".exe", "_update.exe")
	err = DownloadFile(fileURL, newFilePath)
	if err != nil {
		fmt.Println("下载文件时发生错误:", err)
		return
	}

	err = os.Rename(savePath, savePath+".old")
	if err != nil {
		fmt.Print(err)
		return
	}
	time.Sleep(time.Second)
	os.Rename(newFilePath, savePath)
	time.Sleep(time.Second)
	// 到程序安装路径下去执行启动命令(预防相对路径方式启动)
	daemon := "timeout /T 3 & " + savePath + " 2>&1 &"
	_ = exec.Command(savePath, "/C", daemon).Start()
	fmt.Println("文件下载完成")
	os.Exit(0)
}

func DeleteOld() {
	savePath, err := os.Executable()
	if err != nil {
		return
	}
	os.Remove(savePath + ".old")
}

func GetUpdateInfo() (*CheckUpdateResp, error) {
	// 发起HTTP GET请求
	resp, err := http.Get(CheckUpdateUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("下载文件时出错")
	}
	var respData *CheckUpdateResp
	body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &respData)
	if err != nil {
		return nil, err
	}
	return respData, nil
}

func GetExecPath() string {
	// 获取当前可执行文件的路径
	execPath, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(execPath)
}

// DownloadFile 下载文件并保存到指定路径
func DownloadFile(url, savePath string) error {
	// 发起HTTP GET请求
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("下载文件时出错，状态码：%d", resp.StatusCode)
	}

	// 创建文件
	file, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将响应的内容写入文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
