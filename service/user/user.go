package user

import (
	"context"

	bbspb "github.com/MaoScut/bbs/gen/api"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserService struct {
	*bbspb.UnimplementedUserServiceServer
}

type UserTab struct {
	Id         uint64 `gorm:"primaryKey"`
	Email      string
	Password   string
	CreateTime uint64
	UpdateTime uint64
}

func (UserTab) TableName() string {
	return "user_tab"
}

func (s *UserService) CreateUser(context.Context, *bbspb.CreateUserRequest) (*bbspb.CreateUserResponse, error) {
	dsn := "root:root@tcp(172.23.229.208:3306)/bbs_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}
	userModel := UserTab{
		Email:      "test@qq.com",
		Password:   "123",
		CreateTime: 0,
		UpdateTime: 0,
	}
	result := db.Create(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &bbspb.CreateUserResponse{
		User: &bbspb.User{
			Email:    userModel.Email,
			UserName: userModel.Email,
		},
	}, nil
}
func (s *UserService) Login(context.Context, *bbspb.LoginRequest) (*bbspb.LoginResponse, error) {
	return nil, nil
}
func (s *UserService) Logout(context.Context, *bbspb.LogoutRequest) (*bbspb.LogoutResponse, error) {
	return nil, nil
}
