package model

import (
	"fmt"
	"github.com/pkg/errors"
)

func GetTeamsList() ([]TeamLists, error) {
	teams := make([]TeamLists, 0)
	if err := Conn.Table("team_lists").Find(&teams).Error; err != nil {
		return nil, errors.New("战队列表获取失败")
	}
	return teams, nil
}
func GetTeams() ([]Teams, error) {
	teams := make([]Teams, 0)
	if err := Conn.Table("teams").Find(&teams).Error; err != nil {
		return nil, errors.New("战队列表获取失败")
	}
	return teams, nil
}

func GetTeam(id int64) *TeamLists {
	var ret TeamLists
	if err := Conn.Table("team_lists").Where("id = ?", id).Find(&ret).Error; err != nil {
		fmt.Printf("err%s", err.Error())
	}
	return &ret
}
func CreateTeam(team *TeamLists) error {
	if err := Conn.Create(team).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
		return err
	}
	return nil
}

func DelTeam(id int64) error {
	if err := Conn.Table("team_lists").Where("id = ?", id).Delete(&TeamLists{}).Error; err != nil {
		fmt.Printf("err%s", err.Error())
		return err
	}
	return nil
}

func GetTeamDetails(name string) (TeamLists, error) {
	var team TeamLists
	if err := Conn.Table("team_lists").Where("team_name", name).Find(&team).Error; err != nil {
		return TeamLists{}, errors.New("战队列表获取失败")
	}
	return team, nil
}
func GetTeamsDetails(name string) (Teams, error) {
	var team Teams
	if err := Conn.Table("teams").Where("team_name", name).Find(&team).Error; err != nil {
		return Teams{}, errors.New("战队列表获取失败")
	}
	return team, nil
}

func GetTeamPlayers(title string) ([]Players, error) {
	players := make([]Players, 0)
	if err := Conn.Table("players").Where("team_name", title).Find(&players).Error; err != nil {
		return nil, errors.New("成员列表获取失败")
	}
	return players, nil
}
func GetTeamPlayerss(title string) ([]Playersss, error) {
	players := make([]Playersss, 0)
	if err := Conn.Table("playersss").Where("team_name", title).Find(&players).Error; err != nil {
		return nil, errors.New("成员列表获取失败")
	}
	return players, nil
}

func GetPlayers() ([]Players, error) {
	players := make([]Players, 0)
	if err := Conn.Table("players").Find(&players).Error; err != nil {
		return nil, errors.New("成员列表获取失败")
	}
	return players, nil
}
func GetPlayersss() ([]Playersss, error) {
	players := make([]Playersss, 0)
	if err := Conn.Table("playersss").Find(&players).Error; err != nil {
		return nil, errors.New("成员列表获取失败")
	}
	return players, nil
}

func GetPlayer(id int64) *Players {
	var ret Players
	if err := Conn.Table("players").Where("id = ?", id).Find(&ret).Error; err != nil {
		fmt.Printf("err%s", err.Error())
	}
	return &ret
}
func CreatePlayer(player *Players) error {
	if err := Conn.Create(player).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
		return err
	}
	return nil
}
func DelPlayer(id int64) error {
	if err := Conn.Table("players").Where("id = ?", id).Delete(&Players{}).Error; err != nil {
		fmt.Printf("err%s", err.Error())
		return err
	}
	return nil
}
