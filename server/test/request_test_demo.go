package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 响应参数(Proto定义)
type RepArticle struct {
	Id         int    `json:"id"`
	Title      string `json:"name"`
	Content    string `json:"content"`
	Cat        int    `json:"cat"`
	AuthorId   int    `json:"author_id"`
	AuthorName string `json:"author_name"`
	Created    string `json:"created"`
}

func GetRequestDate(ctx *gin.Context) {

	id := ctx.DefaultQuery("refundid", "123")
	parseInt, _ := strconv.ParseInt(id, 10, 64)
	ctx.JSON(http.StatusOK, RepArticle{
		Id:         int(parseInt),
		Title:      "测一测",
		Content:    "所容纳之物",
		Cat:        10,
		AuthorId:   10,
		AuthorName: "无",
		Created:    "无",
	})

}
