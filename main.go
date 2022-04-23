package main

import (
	"fmt"
	"go_gin_demo/file"
	"go_gin_demo/router"
)

func init() {
	fmt.Println("main init")
	router.Init()
}

func main() {
	//test.GetCfg()
	//router.Run()

	file.DownLoadFilePercent("D:\\Code\\Other\\test\\Sekiro.mp4", "Sekiro.mp4")
}
