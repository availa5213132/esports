package router

import (
	"esports/server/app/logic"
	"fmt"
	"github.com/gin-gonic/gin"
)

func New() {
	r := gin.Default()

	//用户操作模块
	user := r.Group("")
	{
		user.POST("/login", logic.DoLogin)
		user.POST("/create", logic.CreateUser)
		user.DELETE("/del", logic.DelUser)
	}
	//战队模块
	team := r.Group("")
	{
		//战队列表
		team.GET("/teamsList", logic.GetTeamsList)
		team.GET("/teams", logic.GetTeams)

		team.GET("/teamDetails", logic.GetTeamDetails)
		//王者荣耀新增详情
		team.GET("/teamsDetails", logic.GetTeamsDetails)

		team.POST("/team/create", logic.CreateTeam)
		team.DELETE("/team/del", logic.DelTeam)
		//战队成员
		team.GET("/teamsMember", logic.GetTeamsMember)
		//新增玩着荣耀
		team.GET("/teamsMembers", logic.GetTeamsMembers)

		team.GET("/Player", logic.GetPlayer)
		team.GET("/Players", logic.GetPlayerss)

		team.POST("/player/create", logic.CreatePlayer)
		team.DELETE("/player/del", logic.DelPlayer)
	}

	//赛事列表:全球总决赛，职业联赛，冠军赛等
	game := r.Group("")
	{
		game.GET("/gameList", logic.GetGamesList)
		game.POST("/gameList", logic.CreateGamesList)
		game.DELETE("/gameList", logic.DelGamesList)

		//每赛季的具体赛程
		r.GET("/gamesCourse", logic.GetGamesCourse)
	}

	//裁判列表操作
	coach := r.Group("")
	{
		coach.GET("/coach", logic.GetRefereeMes)
		coach.DELETE("/coach", logic.DelReferee)
		coach.POST("/coach", logic.CreateReferee)

	}

	{
		//发送通知
		r.GET("/notice", logic.Notice)
	}
	{
		//上传新闻
		r.POST("/saveText")
	}

	//新闻模块
	text := r.Group("")
	{
		//关于新闻的列表
		text.GET("/textList", logic.GetTextList)
		//关于新闻的列表
		text.GET("/textContent", logic.GetTextContent)
		//获取新闻评论
		text.GET("/commentGet", logic.GetComment)
		//保存文章
		text.POST("/save", logic.SaveText)
		//上传文章
		text.POST("/upload", logic.UploadText)
		text.POST("/text/create", logic.CreateText)
		text.DELETE("/text/del", logic.DelText)

	}
	//评论功能
	r.POST("/commentPost", logic.CommentText)

	////定时器1；更新赛程信息
	//{
	//	// 定义定时器
	//	timer := time.NewTicker(12 * time.Hour)
	//	go func() {
	//		update.ReptileGame()
	//		for {
	//			select {
	//			case <-timer.C:
	//				update.ReptileGame()
	//			}
	//		}
	//	}()
	//}

	if err := r.Run(":8080"); err != nil {
		fmt.Print("启动失败")
	}
}
