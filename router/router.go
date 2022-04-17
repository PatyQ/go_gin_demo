package router

import (
	"github.com/gin-gonic/gin"
	"go_gin_demo/server/test"
)

var (
	//_home  *gin.RouterGroup // 通行游客
	//_asset *gin.RouterGroup // 静态资源
	//_prof  *gin.RouterGroup // 个人信息
	//_sys   *gin.RouterGroup // 系统管理
	//_help  *gin.RouterGroup // 帮助模块

	_test *gin.RouterGroup // 测试模块
)

func route() {

	_test = router.Group("/test")
	{
		_test.GET("/getdate", test.GetRequestDate)
	}
}
