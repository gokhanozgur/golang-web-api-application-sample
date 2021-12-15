package model

type Interest struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Status uint16 `json:"status"`
}

// UserDTO
type InterestDTO struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Status uint16 `json:"status"`
}

// DTO to Model
func ToInterest(interestDTO *InterestDTO) *Interest {
	return &Interest{
		Name:   interestDTO.Name,
		Status: interestDTO.Status,
	}
}

// Model to DTO
func ToInterestDTO(interest *Interest) *InterestDTO {
	return &InterestDTO{
		ID:     interest.ID,
		Name:   interest.Name,
		Status: interest.Status,
	}
}
