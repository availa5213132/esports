package logic

import (
	model "esports/server/app/models"
	"esports/server/app/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

//新闻列表获取

type CText struct {
	Uid         int32     `gorm:"column:uid;default:NULL"json:"uid"`
	Title       string    `gorm:"column:title;default:NULL"json:"title"`
	Content     string    `gorm:"column:content;default:NULL"json:"content"`
	ReleaseTime time.Time `gorm:"column:Release_time;default:NULL"json:"release_time"`
	CreatedTime time.Time `gorm:"column:created_time;default:NULL"json:"created_time"`
	Status      string    `gorm:"column:status;default:NULL"json:"status"`
}

func GetTextList(c *gin.Context) {
	text, err := model.GetTextList()
	if err != nil {
		c.JSON(200, tools.ECode{Message: "资源未找到"})
	}
	c.JSON(200, tools.ECode{Data: text})
}

// 获取新闻内容

func GetTextContent(c *gin.Context) {
	titleUid := c.Query("titleUid")
	fmt.Println("文章uid", titleUid)
	uid, _ := strconv.ParseInt(titleUid, 10, 0)
	text, err := model.GetTextContent(int(uid))
	if err != nil {
		c.JSON(200, tools.ECode{Message: "资源未找到"})
	}
	c.JSON(200, tools.ECode{Data: text})
}

//文章评论获取

//func GetComment(c *gin.Context) {
//	titleUid := c.Query("titleUid")
//	comment, err := model.GetComment(titleUid)
//	if err != nil {
//		c.JSON(200, tools.ECode{Message: "暂时没有评论"})
//	}
//	c.JSON(200, tools.ECode{Data: comment})
//}
//
//// 评论文章或回复
//
//func CommentText(c *gin.Context) {
//	var user model.Comment
//	if err := c.ShouldBind(&user); err != nil {
//		c.JSON(200, tools.ECode{
//			Code:    10001,
//			Message: err.Error(),
//		})
//		return
//	}
//
//}

// 新闻文章保存

func SaveText(c *gin.Context) {
	// 解析前端传递的JSON数据
	var requestData struct {
		UID  int32  `json:"uid"`
		Text string `json:"text"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// 创建数据库记录
	text := model.Text{
		UserUid: requestData.UID,
		Content: requestData.Text,
	}
	if err := model.SaveText(text); err != nil {
		c.JSON(500, gin.H{"error": "Failed to save text"})
		return
	}

	// 返回保存成功的文本给前端
	c.JSON(200, gin.H{"message": "Text saved successfully", "text": text})
}

// 文章上传

func UploadText(c *gin.Context) {
	titleUid := c.Query("titleUid")
	err := model.UploadText(titleUid)
	if err != nil {
		c.JSON(200, tools.ECode{Code: 1, Message: "上传失败，请重试"})
	}
	c.JSON(200, tools.ECode{Message: "上传成功"})
}
func CreateText(c *gin.Context) {
	var text CText
	if err := c.ShouldBind(&text); err != nil {
		c.JSON(200, tools.ECode{
			Code:    10001,
			Message: err.Error(),
		})
		return
	}
	newText := model.Text{
		Title:       text.Title,
		Content:     text.Content,
		CreatedTime: time.Now(),
		UserUid:     text.Uid,
		Status:      "2",
	}
	newText.Uid = int32(tools.GetUid())
	if err := model.CreateText(&newText); err != nil {
		c.JSON(http.StatusOK, tools.ECode{
			Code:    10007,
			Message: "添加失败!", //这里有风险
		})
		return
	}
	c.JSON(200, tools.ECode{
		Message: "添加成功",
	})

}

func DelText(c *gin.Context) {
	var id int32
	idStr := c.Query("id")
	id64, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    10001,
			Message: "无效的ID",
		})
		return
	}
	id = int32(id64)
	text := model.GetText(id)
	if text.Id <= 0 {
		c.JSON(200, tools.OK)
		return
	}
	if err := model.DelText(id); err != nil {
		c.JSON(http.StatusOK, tools.ECode{
			Code:    10006,
			Message: "删除失败",
		})
		return
	}
	c.JSON(200, tools.ECode{
		Message: "删除成功",
	})
}
