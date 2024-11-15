package model

import (
	"fmt"
	"testing"
)

func TestGetGamesList(t *testing.T) {
	NewMysql()
	b := "AL"
	a, _ := GetTeamDetails(b)
	fmt.Println("战队详情", a)

	//c := GetRefereeMes()
	//fmt.Println("裁判信息", c)
	//title := "AL"
	//d, _ := GetTeamPlayers(title)
	//fmt.Println(d)
}
