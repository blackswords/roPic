package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"os"
	"path/filepath"
)

// configPathName 是配置文件相对于运行程序的路径
// 但是在main方法中，会对这个值进行修改，从相对路径修改到绝对路径
// 这是因为对于typora自定义命令而言，运行时路径是正在编辑的笔记的路径
// 为了避免这一冲突导致的错误，孤儿修改成为绝对路径
var configPathName = "config.yml"

// printHelp 打印帮助信息，但事实上，我很快就会将他删掉
func printHelp(s string) {
	println(s)
	println("Param struct:")
	println("*.exe upload <filename>")
}

// CreateURLtoString 将传入的url进行改装，使其变为需要的格式字符串并返回
// 现在使用的是markdown格式，之后也许会对其进行扩充
func CreateURLtoString(url string) (copyString string) {
	s := "![img](" + url + ")"
	return s
}

func main() {
	CheckConfigFile()

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		println(err.Error())
	}
	configPathName = dir + "/" + configPathName
	al := len(os.Args)

	if al < 2 {
		printHelp("Too least param.")
	} else {
		if os.Args[1] == "upload" {
			if al == 2 {
				printHelp("Need the third param.")
				return
			}

			// 遍历参数中的文件，上传
			for i := 2; i < al; i++ {
				path := os.Args[i]
				var t TencentYun
				t.InitConfig()

				_, key := UploadFile(path, t)
				err := clipboard.WriteAll(CreateURLtoString(key))
				if err != nil {
					println(err.Error())
				}
				if i == 2 {
					fmt.Printf("[UPLOADER SUCCESS]:\n")
				}
				fmt.Printf("%s\n", key)
			}

		}
	}
}
