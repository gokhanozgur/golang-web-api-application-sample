package model

type UserInterest struct {
	ID         uint64 `json:id`
	UserID     uint64 `json:user_id`
	InterestID uint   `json:interest_id`
}

type UserInterestDTO struct {
	ID         uint64 `json:id`
	UserID     uint64 `json:user_id`
	InterestID uint   `json:interest_id`
}

// DTO to Model
func ToUserInterest(userInterestDTO *UserInterestDTO) *UserInterest {
	return &UserInterest{
		UserID:     userInterestDTO.UserID,
		InterestID: userInterestDTO.InterestID,
	}
}

// DTO to Model
func UserToUserWithInterestDTO(userInterest *UserInterest) *UserInterestDTO {
	return &UserInterestDTO{
		UserID:     userInterest.UserID,
		InterestID: userInterest.InterestID,
	}
}
