package usecase

// Services -.
type Services struct {
	GoogleOAuth GoogleOAuth
	User        User
	Auth        Auth
	Specialty   Specialty
}

// NewServices -.
func NewServices(googleOAuth GoogleOAuth, user User, auth Auth, specialty Specialty) *Services {
	return &Services{
		GoogleOAuth: googleOAuth,
		User:        user,
		Auth:        auth,
		Specialty:   specialty,
	}
}
