package main

import (
	"github.com/gofiber/fiber/v2"
	"golang-redis/repository"
	"golang-redis/sort"
)

func main() {
	// redis
	repository.SetupRedis()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Get("users/:uuid", getUserList)
	app.Get("ranking", ranking)

	app.Listen(":8080")

}

func getUserList(c *fiber.Ctx) error {
	// リクエストからIDを取得
	uuid := c.Params("uuid")

	// redisからデータを取得
	userList, err := repository.GetUserList(uuid)
	if err != nil {
		panic(err)
	}

	// 降順にする
	userList = sort.RankingSort(userList)

	return c.JSON(userList)
}

func ranking(ctx *fiber.Ctx) error {
	result, err := repository.GetRankings()

	if err != nil {
		panic(err)
	}

	return ctx.JSON(result)
}
