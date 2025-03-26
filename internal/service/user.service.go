package service

import (
	"go-ecommerce-backend/internal/repo"
	"go-ecommerce-backend/pkg/response"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type UserService struct {
	userRepo repo.IUserRepository
}

func NewUserService(ur repo.IUserRepository) IUserService {
	return &UserService{userRepo: ur}
}

func (us *UserService) Register(email string, purpose string) int {
	if us.userRepo.GetUserEmail(email) {
		return response.ErrCodeUserHasExist
	}
	return response.ErrCodeSuccess
}

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetInfoUser() string {
// 	return us.userRepo.GetUserData()
// }
