package model

type UserInterest struct {
	ID         uint64 `gorm:"primary_key" json:"id"`
	UserID     uint64 `json:"user_id"`
	InterestID uint   `json:"interest_id"`
	Status     uint16 `json:"status"`
}

type UserInterestDTO struct {
	ID         uint64 `gorm:"primary_key" json:"id"`
	UserID     uint64 `json:"user_id"`
	InterestID uint   `json:"interest_id"`
	Status     uint16 `json:"status"`
}

// DTO to Model
func ToUserInterest(userInterestDTO *UserInterestDTO) *UserInterest {
	return &UserInterest{
		UserID:     userInterestDTO.UserID,
		InterestID: userInterestDTO.InterestID,
		Status:     userInterestDTO.Status,
	}
}

// DTO to Model
func UserToUserWithInterestDTO(userInterest *UserInterest) *UserInterestDTO {
	return &UserInterestDTO{
		UserID:     userInterest.UserID,
		InterestID: userInterest.InterestID,
		Status:     userInterest.Status,
	}
}
