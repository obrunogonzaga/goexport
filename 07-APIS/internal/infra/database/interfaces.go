package database

import "github.com/obrunogonzaga/pos-go-expert/07-APIS/07-APIS/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
