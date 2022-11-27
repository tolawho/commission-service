package repository

import (
	"gorm.io/gorm"
	"medici.vn/commission-serivce/models"
)

// UserRepository is contract what userRepository can do to db
type UserRepository interface {
	VerifyCredential(email string, password string) interface{}
}

type userConnection struct {
	connection *gorm.DB
}

// NewUserRepository is creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user models.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}
