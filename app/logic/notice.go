package logic

import (
	model "esports/server/app/models"
	"esports/server/app/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Notice(c *gin.Context) {
	ticker := time.NewTicker(1 * time.Minute) // 每隔1分钟触发一次
	defer ticker.Stop()
	var concerns []model.Concern
	for range ticker.C {
		var gameHistories []model.GameHistories
		err := model.Conn.Where("game_name = ? AND match_date > ?", "2024职业联赛", time.Now().Add(30*time.Minute)).Find(&gameHistories).Error
		if err != nil {
			fmt.Println("查询错误:", err)
			continue
		}

		for _, gameHistory := range gameHistories {
			err := model.Conn.Where("team_id = ? OR team_id = ?", gameHistory.TeamA, gameHistory.TeamB).Find(&concerns).Error
			if err != nil {
				fmt.Println("查询错误:", err)
				continue
			}

		}
	}
	c.JSON(200, tools.ECode{
		Code: 0,
		Data: concerns,
	})
}
