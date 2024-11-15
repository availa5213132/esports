package model

import (
	"database/sql"
	"time"
)

type User struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL " `
	Uid         int64     `gorm:"column:uid;default:NULL"  form:"uid"`
	Name        string    `gorm:"column:name;default:NULL"  `
	Password    string    `gorm:"column:password;default:NULL"   `
	Telephone   string    `gorm:"column:telephone;default:NULL"`
	Email       string    `gorm:"column:email;default:NULL"`
	CreatedTime time.Time `gorm:"column:created_time;default:NULL"`
}

func (u *User) TableName() string {
	return "user"
}

type GamesList struct {
	Id    int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Title string `gorm:"column:title;default:NULL"`
}

func (g *GamesList) TableName() string {
	return "games_list"
}

type Referee struct {
	Id    string         `gorm:"column:id;primary_key;NOT NULL"`
	Name  sql.NullString `gorm:"column:name"`
	Level sql.NullString `gorm:"column:level"`
	Info  sql.NullString `gorm:"column:info"`
	Uid   int64          `gorm:"column:uid;primary_key;NOT NULL"json:"uid"`
}

func (c *Referee) TableName() string {
	return "referee"
}

type Players struct {
	Id        uint64         `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	CreatedAt time.Time      `gorm:"column:created_at;default:NULL"`
	UpdatedAt time.Time      `gorm:"column:updated_at;default:NULL"`
	DeletedAt time.Time      `gorm:"column:deleted_at;default:NULL"`
	MemberId  sql.NullString `gorm:"column:member_id"`
	EnName    sql.NullString `gorm:"column:en_name"`
	RealName  sql.NullString `gorm:"column:real_name"`
	NickName  sql.NullString `gorm:"column:nick_name"`
	UserIcon  sql.NullString `gorm:"column:user_icon"`
	SUrl      sql.NullString `gorm:"column:s_url"`
	Place     sql.NullString `gorm:"column:place"`
	GameName  sql.NullString `gorm:"column:game_name"`
	GamePlace sql.NullString `gorm:"column:game_place"`
	TeamName  sql.NullString `gorm:"column:team_name"`
	Sex       sql.NullString `gorm:"column:sex"`
	GameHero  sql.NullString `gorm:"column:game_hero"`
	UserFlag  sql.NullString `gorm:"column:user_flag"`
}

func (p *Players) TableName() string {
	return "players"
}

type Playersss struct {
	Id        uint64         `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	CreatedAt time.Time      `gorm:"column:created_at;default:NULL"`
	UpdatedAt time.Time      `gorm:"column:updated_at;default:NULL"`
	DeletedAt time.Time      `gorm:"column:deleted_at;default:NULL"`
	MemberId  sql.NullString `gorm:"column:member_id"`
	EnName    sql.NullString `gorm:"column:en_name"`
	RealName  sql.NullString `gorm:"column:real_name"`
	NickName  sql.NullString `gorm:"column:nick_name"`
	UserIcon  sql.NullString `gorm:"column:user_icon"`
	SUrl      sql.NullString `gorm:"column:s_url"`
	Place     sql.NullString `gorm:"column:place"`
	GameName  sql.NullString `gorm:"column:game_name"`
	GamePlace sql.NullString `gorm:"column:game_place"`
	TeamName  sql.NullString `gorm:"column:team_name"`
	Sex       sql.NullString `gorm:"column:sex"`
	GameHero  sql.NullString `gorm:"column:game_hero"`
	UserFlag  sql.NullString `gorm:"column:user_flag"`
}

func (p *Playersss) TableName() string {
	return "playersss"
}

type Concern struct {
	Id      int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	TeamId  string `gorm:"column:team_id;default:NULL"`
	UserUid int64  `gorm:"column:user_uid;default:NULL"`
}

func (c *Concern) TableName() string {
	return "concern"
}

type Text struct {
	Id          int32     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Uid         int32     `gorm:"column:uid;default:NULL"`
	Title       string    `gorm:"column:title;default:NULL"`
	Content     string    `gorm:"column:content;default:NULL"`
	ReleaseTime time.Time `gorm:"column:Release_time;default:NULL"`
	CreatedTime time.Time `gorm:"column:created_time;default:NULL"`
	UserUid     int32     `gorm:"column:user_uid;default:NULL"`
	Status      string    `gorm:"column:status;default:NULL"`
}

func (t *Text) TableName() string {
	return "text"
}

type TeamLists struct {
	Id             uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	CreatedAt      time.Time `gorm:"column:created_at;default:NULL"`
	UpdatedAt      time.Time `gorm:"column:updated_at;default:NULL"`
	DeletedAt      time.Time `gorm:"column:deleted_at;default:NULL"`
	TeamId         string    `gorm:"column:team_id"`
	TeamName       string    `gorm:"column:team_name"`
	TeamEnName     string    `gorm:"column:team_en_name"`
	TeamDesc       string    `gorm:"column:team_desc"`
	TeamLogo       string    `gorm:"column:team_logo"`
	CreateDate     string    `gorm:"column:create_date"`
	SUrl           string    `gorm:"column:s_url"`
	Place          string    `gorm:"column:place"`
	Leader         string    `gorm:"column:leader"`
	Weibo          string    `gorm:"column:weibo"`
	Tags           string    `gorm:"column:tags"`
	TeamStatus     string    `gorm:"column:team_status"`
	RsetNameStatus string    `gorm:"column:rset_name_status"`
	RsetTeamId     string    `gorm:"column:rset_team_id"`
	TeamLogoDeep   string    `gorm:"column:team_logo_deep"`
	TeamLogo450    string    `gorm:"column:team_logo_450"`
}

func (t *TeamLists) TableName() string {
	return "team_lists"
}

type Teams struct {
	Id             uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	CreatedAt      time.Time `gorm:"column:created_at;default:NULL"`
	UpdatedAt      time.Time `gorm:"column:updated_at;default:NULL"`
	DeletedAt      time.Time `gorm:"column:deleted_at;default:NULL"`
	TeamId         string    `gorm:"column:team_id"`
	TeamName       string    `gorm:"column:team_name"`
	TeamEnName     string    `gorm:"column:team_en_name"`
	TeamDesc       string    `gorm:"column:team_desc"`
	TeamLogo       string    `gorm:"column:team_logo"`
	CreateDate     string    `gorm:"column:create_date"`
	SUrl           string    `gorm:"column:s_url"`
	Place          string    `gorm:"column:place"`
	Leader         string    `gorm:"column:leader"`
	Weibo          string    `gorm:"column:weibo"`
	Tags           string    `gorm:"column:tags"`
	TeamStatus     string    `gorm:"column:team_status"`
	RsetNameStatus string    `gorm:"column:rset_name_status"`
	RsetTeamId     string    `gorm:"column:rset_team_id"`
	TeamLogoDeep   string    `gorm:"column:team_logo_deep"`
	TeamLogo450    string    `gorm:"column:team_logo_450"`
}

func (t *Teams) TableName() string {
	return "teams"
}
