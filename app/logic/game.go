package logic

import (
	"database/sql"
	model "esports/server/app/models"
	"esports/server/app/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CReferee struct {
	Id    string         `gorm:"column:id;primary_key;NOT NULL"json:"id"`
	Name  sql.NullString `gorm:"column:name"json:"name"`
	Level sql.NullString `gorm:"column:level"json:"level"`
	Info  sql.NullString `gorm:"column:info"json:"info"`
	Uid   int64          `gorm:"column:uid;primary_key;NOT NULL"json:"uid"`
}

// 赛程

func GetGamesCourse(c *gin.Context) {
	gameName1 := c.Query("game")

	fmt.Println(gameName1)
	course := model.GetGamesCourse(gameName1)

	c.JSON(200, gin.H{"data": course})
}

// 赛程列表
func GetGamesList(c *gin.Context) {
	List := model.GetGamesList()
	c.JSON(200, gin.H{"data": List})
}
func CreateGamesList(c *gin.Context) {

}
func DelGamesList(c *gin.Context) {

}

// 赛事裁判
func GetRefereeMes(c *gin.Context) {
	List := model.GetRefereeMes()
	c.JSON(200, gin.H{"data": List})
}
func CreateReferee(c *gin.Context) {
	var referee CReferee
	if err := c.ShouldBind(&referee); err != nil {
		c.JSON(200, tools.ECode{
			Message: err.Error(),
		})
		return
	}
	newReferee := model.Referee{
		Name: referee.Name,
		Info: referee.Info,
	}

	err := model.CreateReferee(&newReferee)
	if err != nil {
		c.JSON(200, tools.ECode{
			Message: "添加失败!",
		})
		return
	}
	c.JSON(200, tools.ECode{
		Message: "添加成功！",
	})
}

func DelReferee(c *gin.Context) {
	var uid int64
	uidStr := c.Query("uid")
	uid, _ = strconv.ParseInt(uidStr, 10, 64)
	book := model.GetReferee(uid)
	if book.Uid <= 0 {
		c.JSON(200, tools.OK)
		return
	}
	if err := model.DelReferee(uid); err != nil {
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
