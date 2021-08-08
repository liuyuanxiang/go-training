package assembler

import (
	"section03-project/internal/domain/address"
	"section03-project/internal/interfaces/dto"
)

func ToAddressDO(addressDTO *dto.AddressDTO) *address.Address {
	return &address.Address{
		Name:     addressDTO.Name,
		Mobile:   addressDTO.Mobile,
		Province: addressDTO.Province,
		City:     addressDTO.City,
		Area:     addressDTO.Area,
		Street:   addressDTO.Street,
		PostCode: addressDTO.PostCode,
	}
}
