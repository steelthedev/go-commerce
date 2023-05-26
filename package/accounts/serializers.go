package accounts

type CreateUser struct {
	FirstName    string `json:"first_name" validate:"required, min=2, max=40"`
	LastName     string `json:"last_name" validate:"required, min=2, max=40"`
	Password     string `json:"password" validate:"required"`
	Email        string `json:"email" validate:"email, required"`
	Phone        string `json:"phone" validate:"required"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type UserSerializer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type LoginSerializer struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
