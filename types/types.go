package types

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id string) (*User, error)
	CreateUser(User) error
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=130"`
}

type ResponseBody struct {
	Success bool           `json:"success"`
	Message string         `json:"message,omitempty"`
	Error   string         `json:"error,omitempty"`
	Data    map[string]any `json:"data,omitempty"`
}
