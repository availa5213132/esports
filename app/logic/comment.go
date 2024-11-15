package logic

import (
	model "esports/server/app/models"
	"esports/server/app/tools"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetComment(c *gin.Context) {
	titleUid := c.Query("titleUid")
	comment, err := model.GetComments(titleUid)
	if err != nil {
		c.JSON(200, tools.ECode{Message: "暂时没有评论"})
		return
	}
	fmt.Println("获取的对应评论", comment)
	c.JSON(200, tools.ECode{Data: comment})
}
func CommentText(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBind(&comment); err != nil {
		c.JSON(200, tools.ECode{
			Code:    10001,
			Message: err.Error(),
		})
		return
	}
	fmt.Println("评论id", comment.CommentUid)
	fmt.Println("文章id", comment.TiTleUid)
	fmt.Println("内容", comment.Content)
	fmt.Println("用户id", comment.UserUid)
	fmt.Println("父级id", comment.ParentUid)
	fmt.Println("子评论", comment.Children)
	if s, err := model.AddComment(comment); err != nil {
		c.JSON(200, tools.ECode{Message: "评论保存失败", Data: err})
		fmt.Println(s)
		return
	}

	c.JSON(200, tools.ECode{Message: "评论保存成功"})
}
