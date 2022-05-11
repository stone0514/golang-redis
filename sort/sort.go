package sort

import (
	"sort"

	"golang-redis/models"
)

func RankingSort(userList []models.User) []models.User {
	sort.Slice(
		userList,
		func(i, j int) bool {
			return userList[i].AccountID > userList[j].AccountID
		},
	)

	return userList
}
