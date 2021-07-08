package store

import "restAPI/internal/app/model"

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(string) (*model.User, error)
	Find(int) (*model.User, error)
}
