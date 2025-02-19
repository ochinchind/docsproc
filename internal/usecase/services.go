package usecase

// Services -.
type Services struct {
	GoogleOAuth GoogleOAuth
	User        User
}

// NewServices -.
func NewServices(googleOAuth GoogleOAuth, user User) *Services {
	return &Services{
		GoogleOAuth: googleOAuth,
		User:        user,
	}
}
