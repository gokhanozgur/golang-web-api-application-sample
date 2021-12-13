package model

// User model
type User struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	Username  string     `json:"username"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Profile   string     `json:"profile"`
	Status    uint16     `json:"status"`
	Interests []Interest `json:"interests"`
}

// UserDTO
type UserDTO struct {
	ID        uint64 `gorm:"primary_key" json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Profile   string `json:"profile"`
	Status    uint16 `json:"status"`
}

// UserWithoutPasswordDTO
type UserWithoutPasswordDTO struct {
	ID        uint64 `gorm:"primary_key" json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Profile   string `json:"profile"`
	Status    uint16 `json:"status"`
}

// UserWithInterestDTO
type UserWithInterestDTO struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	Username  string     `json:"username"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Profile   string     `json:"profile"`
	Status    uint16     `json:"status"`
	Interests []Interest `json:"interests"`
}

// DTO to Model
func ToUser(userDTO *UserDTO) *User {
	return &User{
		Username:  userDTO.Username,
		FirstName: userDTO.FirstName,
		LastName:  userDTO.LastName,
		Email:     userDTO.Email,
		Password:  userDTO.Password,
		Profile:   userDTO.Profile,
		Status:    userDTO.Status,
	}
}

// Model to DTO
func ToUserDTO(user *User) *UserDTO {
	return &UserDTO{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		Profile:   user.Profile,
		Status:    user.Status,
	}
}

// User to ToUserWithoutPasswordDTO
func ToUserWithoutPasswordFromUser(user *User) *UserWithoutPasswordDTO {
	return &UserWithoutPasswordDTO{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Profile:   user.Profile,
		Status:    user.Status,
	}
}
