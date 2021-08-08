package dto

type UserDTO struct {
	ID       string
	Username string
	Password string
}

type UserListDTO struct {
	List        []UserDTO
	TotalCount  int64
	CurrentPage int64
	TotalPage   int64
	PageSize    int64
}
