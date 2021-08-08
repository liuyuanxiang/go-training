package assembler

import (
	"section03-project/internal/domain/user"
	"section03-project/internal/interfaces/dto"
)

func ToUserDO(userDTO *dto.UserDTO) *user.User {
	return &user.User{
		Username: userDTO.Username,
		Password: userDTO.Password,
	}
}
