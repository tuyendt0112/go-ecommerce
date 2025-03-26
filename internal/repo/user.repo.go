package repo

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetUserData() string {
// 	return "this is jezztom form repo"
// }

type IUserRepository interface {
	GetUserEmail(email string) bool
}

type userRepository struct {
}

func (uc *userRepository) GetUserEmail(email string) bool {
	return true
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
