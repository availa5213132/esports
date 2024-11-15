package model

import (
	"fmt"
)

func GetGamesList() []string {
	gamesList := make([]GamesList, 0)
	if err := Conn.Table("games_list").Find(&gamesList).Error; err != nil {
		fmt.Printf("赛事分类列表 err:%s", err.Error())
	}
	fmt.Println(gamesList)
	var list []string
	for _, game := range gamesList {
		list = append(list, game.Title)
	}
	return list
}

// 具体赛程

func GetGamesCourse(title string) []GameHistories {
	var games []GameHistories
	if err := Conn.Table("game_histories").Where("game_name", title).Find(&games).Error; err != nil {
		fmt.Println("查询赛程失败")
	}
	return games
}

// 裁判

func GetRefereeMes() []Referee {
	var coaches []Referee
	if err := Conn.Table("referee").Find(&coaches).Error; err != nil {
		fmt.Printf("赛事分类列表 err:%s", err.Error())
	}
	return coaches

}
func GetReferee(uid int64) *Referee {
	var ret Referee
	if err := Conn.Table("referee").Where("uid = ?", uid).Find(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return &ret
}

func CreateReferee(referee *Referee) error {
	if err := Conn.Create(referee).Error; err != nil {
		fmt.Printf("err%s", err.Error())
		return err
	}
	return nil
}

func DelReferee(uid int64) error {
	if err := Conn.Table("referee").Where("uid = ?", uid).Delete(&Referee{}).Error; err != nil {
		fmt.Printf("err%s", err.Error())
		return err
	}
	return nil
}
