package model

import "fmt"

// 保存文章

func SaveText(text Text) error {
	// Assuming you have a database connection initialized as "db"
	text.Status = "0"
	err := Conn.Create(&text).Error
	if err != nil {
		// Handle the error if the database operation fails
		return err
	}
	return nil
}

// 上传文章

func UploadText(uid string) error {
	// Assuming you have a database connection initialized as "db"
	var text Text
	err := Conn.Where("uid = ?", uid).First(&text).Error
	if err != nil {
		// Handle the error if the database query fails
		return err
	}

	text.Status = "1"
	err = Conn.Save(&text).Error
	if err != nil {
		// Handle the error if the database update fails
		return err
	}

	return nil
}

// 获取文章列表

func GetTextList() ([]Text, error) {
	// Assuming you have a database connection initialized as "db"
	var text []Text
	err := Conn.Table("text").Where("status", 2).Find(&text).Error
	if err != nil {
		// Handle the error if the database operation fails
		return nil, err
	}

	return text, nil
}

// 获取文章内容

func GetTextContent(uid int) (Text, error) {
	// Assuming you have a database connection initialized as "db"
	var text Text
	err := Conn.Table("text").Where("uid", uid).Find(&text).Error
	if err != nil {
		// Handle the error if the database operation fails
		return Text{}, err
	}
	return text, nil
}

func GetText(id int32) *Text {
	var ret Text
	if err := Conn.Table("text").Where("id = ?", id).Find(&ret).Error; err != nil {
		fmt.Printf("err%s", err.Error())
	}
	return &ret
}

func CreateText(text *Text) error {
	if err := Conn.Create(text).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
		return err
	}
	return nil
}
func DelText(id int32) error {
	if err := Conn.Table("text").Where("id = ?", id).Delete(&Text{}).Error; err != nil {
		fmt.Printf("err%s", err.Error())
		return err
	}
	return nil
}

//获取评论

//func GetComment(TextUID string) ([]Comment, error) {
//	var comments []Comment
//	// 查询根级评论
//	err := Conn.Where("ParentUid = ? AND ParentUid = CommentUid", TextUID, "").Find(&comments).Error
//	if err != nil {
//		return nil, err
//	}
//
//	// 递归查询子级评论
//	recursiveQueryComments(&comments)
//
//	return comments, nil
//}
//
//// 在上述代码中，GetComment 函数接收一个 TextUID 参数，并返回一个 []Comment 切片和一个错误。在函数内部，我们首先初始化 GORM 数据库连接，
//// 并使用 Where 方法查询根级评论。我们使用 ParentUid = ? AND ParentUid = CommentUid 来查询 ParentUid 为空且 CommentUid 等于 TextUID 的评论。
//// 如果查询成功，我们调用 recursiveQueryComments 函数递归查询子级评论，并将其赋值给 Children 字段。最后，我们返回评论切片和可能的错误。
//// recursiveQueryComments 函数与之前提供的示例代码类似，用于递归查询子级评论并将其赋值给 Children 字段。
//func recursiveQueryComments(comments *[]Comment) {
//	for i := range *comments {
//		var children []Comment
//		err := Conn.Where("ParentUid = ?", (*comments)[i].CommentUid).Find(&children).Error
//		if err == nil {
//			recursiveQueryComments(&children)
//			(*comments)[i].Children = children
//		}
//	}
//}
