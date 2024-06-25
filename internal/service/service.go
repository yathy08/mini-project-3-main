// internal/service/service.go

package service

import (
	"encoding/json"
	"github.com/yathy08/mini-project3/internal/domain"
)

type Users struct {
	Data []domain.User `json:"data"`
}

func UnmarshalUsers(data []byte) (*Users, error) {
	var users Users
	err := json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func FilterByID(users *Users, id int) *domain.User {
	for _, user := range users.Data {
		if user.ID == id {
			return &user
		}
	}
	return nil
}
