package model

type GameHistories struct {
	Id           uint64 `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	IsTft        string `gorm:"column:is_tft"`
	BGameId      string `gorm:"column:b_game_id"`
	BMatchId     string `gorm:"column:b_match_id"`
	GameTypeName string `gorm:"column:game_type_name"`
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

func (g *GameHistories) TableName() string {
	return "game_list"
}
