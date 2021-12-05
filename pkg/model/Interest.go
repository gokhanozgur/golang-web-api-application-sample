package model

type Interest struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// UserDTO
type InterestDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// DTO to Model
func ToInterest(interestDTO *InterestDTO) *Interest {
	return &Interest{
		Name: interestDTO.Name,
	}
}

// Model to DTO
func ToInterestDTO(interest *Interest) *InterestDTO {
	return &InterestDTO{
		ID:   interest.ID,
		Name: interest.Name,
	}
}
