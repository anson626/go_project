package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	db     *mongo.Database
	users  *mongo.Collection
)

func init() {
	const url = "mongodb://127.0.0.1:27017"
	opts := options.Client().ApplyURI(url) //返回一个客户端配置选项类型
	opts.SetConnectTimeout(3 * time.Second)
	var err error
	client, err = mongo.Connect(context.TODO(), opts) //todo 空上下文
	if err != nil {
		log.Panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Panic(err)
	}
	//库，集合（文档的存储的地方） 可以实现不存在
	db = client.Database("test") //库
	fmt.Println(db)
	users = db.Collection("users") //集合，类似看做表，表中有field
}

type User struct {
	Name string
	Age  int
}

func main() {
	//var err error
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	fmt.Println(client)
	fmt.Println(db, users)
	tom := User{"Tom", 23}
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	insertResult, err := users.InsertOne(context.TODO(), tom)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insertResult.InsertedID)
}
