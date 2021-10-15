package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var labelPath = "../../../LabelImg/image/label"

// 处理多于的classes类别
func main() {
	for _, v := range getFileList() {
		if v.Name() == "classes.txt" {
			continue
		}
		changeText(v.Name())
	}
}

// 获取文件列表
func getFileList() []os.FileInfo {
	files, err := ioutil.ReadDir(labelPath)
	if err != nil {
		panic(err.Error())
	}

	return files
}

//替换文本
func changeText(filename string) {
	//读取文件内容
	b, e := ioutil.ReadFile(fmt.Sprintf("%s/%s", labelPath, filename))
	if e != nil {
		fmt.Println("read file error")
		return
	}

	lineString := string(b)
	bSlice := strings.Split(lineString, "\n")
	// fmt.Println(bSlice)
	for k, v := range bSlice {
		if v == "" {
			continue
		}
		// newClasses := ""
		switch v[:2] {
		case "15": //citrus
			bSlice[k] = "0" + v[2:]
		case "16": //stem
			bSlice[k] = "1" + v[2:]
		case "17": //citrus_embryo
			bSlice[k] = "2" + v[2:]
		case "18": //rot
			bSlice[k] = "3" + v[2:]
		default:
			bSlice[k] = v
		}
	}

	out, err := os.OpenFile(fmt.Sprintf("./temp/%s", filename), os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println("Open write file fail:", err)
		os.Exit(-1)
	}
	defer out.Close()
	_, err = out.WriteString(strings.Join(bSlice, "\n"))
	if err != nil {
		fmt.Println("write to file fail:", err)
		os.Exit(-1)
	}
	fmt.Println("done")
}
