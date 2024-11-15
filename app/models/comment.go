package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Comment struct {
	CommentUid  string    `json:"CommentUid"`
	TiTleUid    string    `json:"TiTleUid"`
	Content     string    `json:"Content"`
	UserUid     string    `json:"UserUid"`
	Children    []Comment `json:"Children"`
	ParentUid   string    `json:"ParentUid"`
	CreatedTime time.Time `json:"CreatedTime"`
	FloorUid    string    `json:"FloorUid"`
}

func GetComments(articleUID string) ([]Comment, error) {
	// 创建 MongoDB 客户端

	// 获取数据库和集合
	database := MB.Database("0")
	collection := database.Collection(articleUID)

	// 查询集合中的所有评论
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	// 遍历查询结果，将评论解码到切片中
	var comments []Comment
	for cur.Next(context.Background()) {
		var comment Comment
		err := cur.Decode(&comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func AddComment(newComment Comment) (string, error) {
	// 创建 MongoDB 客户端

	// 检查数据库是否存在
	databaseNames, err := MB.ListDatabaseNames(context.Background(), bson.M{})
	if err != nil {
		return "查询数据库失败", err
	}
	databaseName := "0"                   // 数据库名称
	collectionName := newComment.TiTleUid // 集合名称
	// 如果数据库不存在，则创建数据库
	databaseExists := false
	for _, name := range databaseNames {
		if name == databaseName {
			databaseExists = true
			break
		}
	}

	// 如果数据库不存在，则创建数据库
	if !databaseExists {
		err = MB.Database(databaseName).CreateCollection(context.Background(), collectionName)
		if err != nil {
			return "集合创建失败", err
		}
	}

	// 获取数据库和集合
	database := MB.Database(databaseName)
	collection := database.Collection(collectionName)

	// 如果父评论 UID 为空，表示第一级评论，直接插入到集合中
	if newComment.ParentUid == "" {
		newComment.FloorUid = newComment.CommentUid
		_, err = collection.InsertOne(context.Background(), newComment)
		if err != nil {
			return "插入失败", err
		}
	} else {
		// 父评论 UID 不为空，查找父评论
		filter := bson.M{"commentuid": newComment.ParentUid}
		var parentComment Comment
		err = collection.FindOne(context.Background(), filter).Decode(&parentComment)
		if err != nil {
			return "父级查找失败", err
		}

		// 找到父评论对应的楼层评论
		floorFilter := bson.M{"commentuid": parentComment.FloorUid}
		var floorComment Comment
		err = collection.FindOne(context.Background(), floorFilter).Decode(&floorComment)
		if err != nil {
			return "楼层评论查找失败", err
		}

		// 将接收的评论结构体存储到父评论的 Children 切片中
		floorComment.Children = append(floorComment.Children, newComment)

		// 更新父评论
		filter = bson.M{"commentuid": floorComment.CommentUid}
		update := bson.M{"$set": floorComment}
		_, err = collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return "更新集合失败", err
		}
	}

	return "", nil
}

//func GetComments(articleUID string) ([]Comment, error) {
//	// 创建 MongoDB 客户端
//
//	// 获取数据库和集合
//	database := MB.Database("0")
//	collection := database.Collection(articleUID)
//
//	// 查询集合中的所有评论
//	cur, err := collection.Find(context.Background(), bson.M{})
//	if err != nil {
//		return nil, err
//	}
//	defer cur.Close(context.Background())
//
//	// 遍历查询结果，将评论解码到切片中
//	var comments []Comment
//	for cur.Next(context.Background()) {
//		var comment Comment
//		err := cur.Decode(&comment)
//		if err != nil {
//			return nil, err
//		}
//		comments = append(comments, comment)
//	}
//	if err := cur.Err(); err != nil {
//		return nil, err
//	}
//
//	return comments, nil
//}
//
//func AddComment(newComment Comment) (string, error) {
//	// 创建 MongoDB 客户端
//
//	// 检查数据库是否存在
//	databaseNames, err := MB.ListDatabaseNames(context.Background(), bson.M{})
//	if err != nil {
//		return "查询数据库失败", err
//	}
//	databaseName := "0"                   // 数据库名称
//	collectionName := newComment.TiTleUid // 集合名称
//	// 如果数据库不存在，则创建数据库
//	databaseExists := false
//	for _, name := range databaseNames {
//		if name == databaseName {
//			databaseExists = true
//			break
//		}
//	}
//	if !databaseExists {
//		err = MB.Database(databaseName).CreateCollection(context.Background(), collectionName)
//		if err != nil {
//			return "集合创建失败", err
//		}
//	}
//	// 获取数据库和集合
//	database := MB.Database(databaseName)
//	collection := database.Collection(collectionName)
//
//	// 如果父评论 UID 为空，表示第一级评论，直接插入到集合中
//	if newComment.ParentUid == "" {
//		_, err = collection.InsertOne(context.Background(), newComment)
//		if err != nil {
//			return "插入失败", err
//		}
//	} else {
//		// 父评论 UID 不为空，查找父评论
//		filter := bson.M{"commentuid": newComment.ParentUid}
//		var parentComment Comment
//		err = collection.FindOne(context.Background(), filter).Decode(&parentComment)
//		if err != nil {
//			return "父级查找失败", err
//		}
//
//		// 将接收的评论结构体存储到父评论的 Children 切片中
//		parentComment.Children = append(parentComment.Children, newComment)
//
//		// 更新父评论
//		filter = bson.M{"commentuid": parentComment.CommentUid}
//		update := bson.M{"$set": parentComment}
//		_, err = collection.UpdateOne(context.Background(), filter, update)
//		if err != nil {
//			return "更新集合失败", err
//		}
//	}
//
//	return "", nil
//}
