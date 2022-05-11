package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/go-redis/redis/v8"
	"golang-redis/models"
)

type SomeStructWithTags struct {
	Email     string `faker:"email"`
	Name      string `faker:"name"`
	UUID      string `faker:"uuid_digit"`
	AccountID int    `faker:"oneof: 15, 27,61"`
}

func SetupRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
	})
}

func main() {
	// fakerの準備
	a := SomeStructWithTags{}
	err := faker.FakeData(&a)
	if err != nil {
		panic(err)
	}

	//ユーザーリストを作成
	var userList []models.User
	for i := 0; i < 100; i++ {
		err = faker.FakeData(&a)
		if err != nil {
			panic(err)
		}
		userList = append(userList, models.User{
			AccountID: a.AccountID,
			Name:      a.Name,
			Email:     a.Email,
		})
	}

	// Redisに格納するため、シリアライズ
	serialize, err := json.Marshal(&userList)
	if err != nil {
		panic(err)
	}

	// UUIDはデータにアクセスするため必要
	UUID := a.UUID
	fmt.Println("UUID: ", UUID)

	// Redisに接続
	c := context.Background()
	r := SetupRedis()

	// Redisにデータを格納
	err = r.Set(c, UUID, serialize, time.Hour*24).Err()
	if err != nil {
		panic(err)
	}

	// Sort
	for _, user := range userList {
		r.ZAdd(c, "rankings", &redis.Z{
			Score:  float64(user.AccountID),
			Member: user.Name,
		})
	}

	log.Println("complete")
}
