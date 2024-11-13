package user

import (
	"test-echo/internal/db"
	"time"
)

type UserInfo struct {
	Email     string
	CreatedAt time.Time
}

func GetUserInfo(email string) (UserInfo, error) {
	var info UserInfo
	result := db.Get().Model(&db.User{}).Where("email = ?", email).First(&info)
	if result.Error != nil {
		return UserInfo{}, result.Error
	}

	return info, nil
}
