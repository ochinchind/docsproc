package usecase

// Services -.
type Services struct {
	GoogleOAuth GoogleOAuth
	User        User
	Auth        Auth
}

// NewServices -.
func NewServices(googleOAuth GoogleOAuth, user User, auth Auth) *Services {
	return &Services{
		GoogleOAuth: googleOAuth,
		User:        user,
		Auth:        auth,
	}
}
