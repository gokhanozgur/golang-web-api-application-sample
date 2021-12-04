package model

// User model
type User struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	Username  string     `json:"username"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Profile   string     `json:"profile"`
	Interests []Interest `json:"interests"`
}

// UserDTO
type UserDTO struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	Username  string     `json:"username"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Profile   string     `json:"profile"`
	Interests []Interest `json:"interests"`
}

// DTO to Model
func ToUser(userDTO *UserDTO) *User {
	return &User{
		Username:  userDTO.Username,
		FirstName: userDTO.FirstName,
		LastName:  userDTO.LastName,
		Profile:   userDTO.Profile,
		Interests: userDTO.Interests,
	}
}

// Model to DTO
func ToUserDTO(user *User) *UserDTO {
	return &UserDTO{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.Username,
		LastName:  user.LastName,
		Profile:   user.Profile,
		Interests: user.Interests,
	}
}
