package logic

import (
	"database/sql"
	model "esports/server/app/models"
	"esports/server/app/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type CTeam struct {
	TeamName   string `gorm:"column:team_name"json:"team_name"`
	TeamEnName string `gorm:"column:team_en_name"json:"team_en_name"`
}

//sql.NullString

type CPlayer struct {
	RealName  sql.NullString `gorm:"column:real_name"json:"real_name"`
	GameName  sql.NullString `gorm:"column:game_name"json:"game_name"`
	TeamName  sql.NullString `gorm:"column:team_name"json:"team_name"`
	CreatedAt time.Time      `gorm:"column:created_at;default:NULL"json:"created_at"`
}

func GetTeamsList(c *gin.Context) {
	teams, err := model.GetTeamsList()
	if err != nil {
		c.JSON(200, tools.ECode{
			Message: "战队列表获取失败",
		})
		return
	}

	c.JSON(200, gin.H{"teams": teams})
}
func GetTeams(c *gin.Context) {
	teams, err := model.GetTeams()
	if err != nil {
		c.JSON(200, tools.ECode{
			Message: "战队列表获取失败",
		})
		return
	}

	c.JSON(200, gin.H{"teams": teams})
}

func GetTeamDetails(c *gin.Context) {
	name := c.Query("team")
	details, err := model.GetTeamDetails(name)
	if err != nil {
		c.JSON(200, tools.ECode{
			Message: "战队详情",
		})
		return
	}
	c.JSON(200, tools.ECode{Data: details})
}
func GetTeamsDetails(c *gin.Context) {
	name := c.Query("team")
	details, err := model.GetTeamsDetails(name)
	if err != nil {
		c.JSON(200, tools.ECode{
			Message: "战队详情",
		})
		return
	}
	c.JSON(200, tools.ECode{Data: details})
}

func CreateTeam(c *gin.Context) {
	var team CTeam
	if err := c.ShouldBind(&team); err != nil {
		c.JSON(200, tools.ECode{
			Message: err.Error(),
		})
		return
	}
	newTeam := model.TeamLists{
		CreatedAt:  time.Now(),
		TeamName:   team.TeamName,
		TeamEnName: team.TeamEnName,
	}
	err := model.CreateTeam(&newTeam)
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
func DelTeam(c *gin.Context) {
	var id int64
	idStr := c.Query("id")
	id, _ = strconv.ParseInt(idStr, 10, 64)
	team := model.GetTeam(id)
	if team.Id <= 0 {
		c.JSON(200, tools.OK)
		return
	}
	if err := model.DelTeam(id); err != nil {
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

func GetTeamsMember(c *gin.Context) {
	title := c.Query("team")

	teams, err := model.GetTeamPlayers(title)
	if err != nil {
		c.JSON(200, tools.ECode{
			Message: "战队成员获取失败",
		})
		return
	}
	c.JSON(200, tools.ECode{Data: teams})
}
func GetTeamsMembers(c *gin.Context) {
	title := c.Query("team")

	teams, err := model.GetTeamPlayerss(title)
	if err != nil {
		c.JSON(200, tools.ECode{
			Message: "战队成员获取失败",
		})
		return
	}
	c.JSON(200, tools.ECode{Data: teams})
}

func GetPlayer(c *gin.Context) {

	teams, err := model.GetPlayers()
	if err != nil {
		c.JSON(200, tools.ECode{
			Message: "战队成员获取失败",
		})
		return
	}
	c.JSON(200, tools.ECode{Data: teams})
}

func GetPlayerss(c *gin.Context) {

	teams, err := model.GetPlayersss()
	if err != nil {
		c.JSON(200, tools.ECode{
			Message: "战队成员获取失败",
		})
		return
	}
	c.JSON(200, tools.ECode{Data: teams})
}

func CreatePlayer(c *gin.Context) {
	var player CPlayer
	if err := c.ShouldBind(&player); err != nil {
		c.JSON(200, tools.ECode{
			Message: err.Error(),
		})
		return
	}
	newPlayer := model.Players{
		CreatedAt: time.Now(),
		RealName:  player.RealName,
		GameName:  player.GameName,
		TeamName:  player.TeamName,
	}

	err := model.CreatePlayer(&newPlayer)
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

func DelPlayer(c *gin.Context) {
	var id int64
	idStr := c.Query("id")
	id, _ = strconv.ParseInt(idStr, 10, 64)
	team := model.GetPlayer(id)
	if team.Id <= 0 {
		c.JSON(200, tools.OK)
		return
	}
	if err := model.DelPlayer(id); err != nil {
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
