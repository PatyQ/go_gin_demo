package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
	//routerGroup *gin.RouterGroup
	//group  errgroup.Group // 服务组
)

func init() {
	fmt.Println("router init")
}

func Init() {
	router = gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//router.Use(stats.ApiVisitHandler)

}

// 启动程序
func Run() {

	route()
	router.Run(":80")
	//server := &http.Server{
	//	Addr:           ":8080",
	//	Handler:        router,
	//	ReadTimeout:    time.Second * 30,
	//	WriteTimeout:   time.Second * 30,
	//	MaxHeaderBytes: 1 << 20, // 2M
	//}

	// 加入服务组
	//group.Go(func() error {
	//	fmt.Printf(" 🚗 服务已启动 %s\n", config.App.Port)
	//	return server.ListenAndServe()
	//})
	//
	//if err := group.Wait(); err != nil {
	//	panic(fmt.Errorf("服务启动失败：%#v", err.Error()))
	//}
}
