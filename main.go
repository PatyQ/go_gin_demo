package main

import (
	"fmt"
	"go_gin_demo/file"
	"go_gin_demo/router"
	"net/http"
	"time"
)

func init() {
	fmt.Println("main init")
	router.Init()
}

func main() {
	//test.GetCfg()
	//router.Run()
	http.Handle("/time", timeMiddleware(DownFile)) // 这里不能使用handleFunc
	//http.HandleFunc("/time", timeMiddleware(DownFile)) // 这里不能使用handleFunc
	http.ListenAndServe(":80", nil)
}

func DownFile(w http.ResponseWriter, r *http.Request) {

	file.DownLoadFilePercent("D:\\Code\\Other\\test\\Sekiro.mp4", "Sekiro.mp4")
}

// 中间件
func timeMiddleware(next http.HandlerFunc) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		fmt.Println("r.URL.RawQuery:", r.URL.RawQuery)                                    // r.URL.RawQuery: name=q
		fmt.Println("timeMiddleware star:", t.UnixNano(), "r.RequestURI::", r.RequestURI) //timeMiddleware star: 1650703718705203600 r.RequestURI:: /time?name=q
		next.ServeHTTP(w, r)
		fmt.Println("timeMiddleware end:", time.Since(t)) // timeMiddleware end: 7.9025ms
	})

}
