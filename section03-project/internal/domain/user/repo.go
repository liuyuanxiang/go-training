package user

type RepoInterface interface {
	CreateUser(user *User) (string, error)

	UpdateUser(id string, user *User) error

	DeleteUser(id string) error

	GetUser(id string) (*User, error)

	ListUser(pageNum, pageSize int) ([]*User, int64, error)
}
