package service

type User struct {
	Name string
	// 2006-01-02(yyyy-mm-dd)
	DOB string
}

type UserService interface {
}

func NewUserService() UserService {

}

type userService struct {
}
