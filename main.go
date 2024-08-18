package main

import (
	"fmt"
	"log"
	mylog "mp4ass2mkvass/log"
	"mp4ass2mkvass/merge"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	mylog.SetLog()
}
func main() {
	// 指定要搜索的目录
	dir := "/Users/zen/container" // 替换为你的目录路径

	// 获取所有 mp4 文件的绝对路径
	mp4Files, err := getMP4Files(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印所有找到的 mp4 文件路径
	for _, mp4 := range mp4Files {
		log.Println(mp4)
		srt := strings.Replace(mp4, filepath.Ext(mp4), ".srt", 1)
		if has, _ := isFileExists(srt); !has {
			fmt.Println("srt File not found skip", srt)
			continue
		}
		mkv := strings.Replace(mp4, filepath.Ext(mp4), ".mkv", 1)
		merge.MkvWithAss(mp4, srt, mkv)
	}
}

// getMP4Files 遍历指定目录，返回所有扩展名为 .mp4 的文件的绝对路径
func getMP4Files(dir string) ([]string, error) {
	var mp4Files []string

	// 使用 Walk 函数遍历目录
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 检查文件是否是 mp4 文件
		if !info.IsDir() && strings.EqualFold(filepath.Ext(path), ".mp4") {
			// 获取绝对路径
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			mp4Files = append(mp4Files, absPath)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return mp4Files, nil
}

func isFileExists(filePath string) (bool, error) {
	// 获取文件信息
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		// 文件不存在
		return false, nil
	}
	if err != nil {
		// 其他错误
		return false, err
	}
	// 判断是否是文件
	return !info.IsDir(), nil
}
