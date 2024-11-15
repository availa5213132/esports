package update

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Response struct {
	Status     string      `json:"status"`
	LastUpTime string      `json:"lastUpTime"`
	Msg        []MatchItem `json:"msg"`
}

type MatchItem struct {
	IsTft        string `json:"isTft"`
	TftInfos     string `json:"tftInfos"`
	BGameId      string `json:"bGameId"`
	BMatchId     string `json:"bMatchId"`
	BMatchName   string `json:"bMatchName"`
	GameId       string `json:"GameId"`
	GameName     string `json:"GameName"`
	GameTypeName string `json:"GameTypeName"`
	GameMode     string `json:"GameMode"`
	GameModeName string `json:"GameModeName"`
	TeamA        string `json:"TeamA"`
	ScoreA       string `json:"ScoreA"`
	TeamB        string `json:"TeamB"`
	ScoreB       string `json:"ScoreB"`
	MatchDate    string `json:"MatchDate"`
	MatchStatus  string `json:"MatchStatus"`
	MatchWin     string `json:"MatchWin"`
	AppTopicId   string `json:"AppTopicId"`
}

type GameHistories struct {
	ID           uint   `gorm:"primaryKey"`
	IsTft        string `gorm:"column:is_tft"`
	BGameId      string `gorm:"column:b_game_id"`
	GameTypeName string `gorm:"column:game_type_name"`
	BMatchId     string `gorm:"column:b_match_id"`
	BMatchName   string `gorm:"column:b_match_name"`
	GameId       string `gorm:"column:game_id"`
	GameName     string `gorm:"column:game_name"`
	GameMode     string `gorm:"column:game_mode"`
	GameModeName string `gorm:"column:game_mode_name"`
	TeamA        string `gorm:"column:team_a"`
	ScoreA       string `gorm:"column:score_a"`
	TeamB        string `gorm:"column:team_b"`
	ScoreB       string `gorm:"column:score_b"`
	MatchDate    string `gorm:"column:match_date"`
	MatchStatus  string `gorm:"column:match_status"`
	MatchWin     string `gorm:"column:match_win"`
	AppTopicId   string `gorm:"column:app_topic_id"`
}

func ReptileGame() {

	// JSON数据
	res, _ := http.Get("https://lpl.qq.com/web201612/data/LOL_MATCH2_MATCH_HOMEPAGE_BMATCH_LIST_206.js")
	byteData, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("爬取失败")
	}
	// 解析JSON数据到结构体
	var response Response
	err = json.Unmarshal(byteData, &response)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// 连接数据库
	my := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "123456", "127.0.0.1:3306", "esports")
	db, err := gorm.Open(mysql.Open(my), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	// 迁移数据库表
	err = db.AutoMigrate(&GameHistories{})
	if err != nil {
		fmt.Println("Error migrating table:", err)
		return
	}

	// 插入数据到game表
	for _, match := range response.Msg {
		// 查询数据库中是否存在相同的记录
		var existingGame GameHistories
		result := db.Where("b_match_id = ?", match.BMatchId).First(&existingGame)

		if result.Error == nil {
			// 记录已存在，执行更新操作
			existingGame.IsTft = match.IsTft
			existingGame.BGameId = match.BGameId
			existingGame.BMatchName = match.BMatchName
			existingGame.GameTypeName = match.GameTypeName
			existingGame.GameId = match.GameId
			existingGame.GameName = match.GameName
			existingGame.GameMode = match.GameMode
			existingGame.GameModeName = match.GameModeName
			existingGame.TeamA = match.TeamA
			existingGame.ScoreA = match.ScoreA
			existingGame.TeamB = match.TeamB
			existingGame.ScoreB = match.ScoreB
			existingGame.MatchDate = match.MatchDate
			existingGame.MatchStatus = match.MatchStatus
			existingGame.MatchWin = match.MatchWin
			existingGame.AppTopicId = match.AppTopicId

			result = db.Save(&existingGame)
			if result.Error != nil {
				fmt.Println("Error updating data:", result.Error)
				return
			}

			fmt.Printf("Record with b_match_id: %s updated successfully.\n", match.BMatchId)
		} else {
			// 记录不存在，执行插入操作
			game := GameHistories{
				IsTft:        match.IsTft,
				BGameId:      match.BGameId,
				BMatchId:     match.BMatchId,
				BMatchName:   match.BMatchName,
				GameTypeName: match.GameTypeName,
				GameId:       match.GameId,
				GameName:     match.GameName,
				GameMode:     match.GameMode,
				GameModeName: match.GameModeName,
				TeamA:        match.TeamA,
				ScoreA:       match.ScoreA,
				TeamB:        match.TeamB,
				ScoreB:       match.ScoreB,
				MatchDate:    match.MatchDate,
				MatchStatus:  match.MatchStatus,
				MatchWin:     match.MatchWin,
				AppTopicId:   match.AppTopicId,
			}

			result = db.Create(&game)
			if result.Error != nil {
				fmt.Println("Error inserting data:", result.Error)
				return
			}

			fmt.Printf("Record with b_match_id: %s inserted successfully.\n", match.BMatchId)
		}
	}

	fmt.Println("Data inserted/updated successfully into the game table.")
}
