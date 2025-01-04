package repo

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetUserData() string {
	return "this is jezztom form repo"
}
