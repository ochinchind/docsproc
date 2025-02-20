package usecase

// Services -.
type Services struct {
	GoogleOAuth   GoogleOAuth
	User          User
	Auth          Auth
	Specialty     Specialty
	Qualification Qualification
}

// NewServices -.
func NewServices(googleOAuth GoogleOAuth, user User, auth Auth, specialty Specialty, qualification Qualification) *Services {
	return &Services{
		GoogleOAuth:   googleOAuth,
		User:          user,
		Auth:          auth,
		Specialty:     specialty,
		Qualification: qualification,
	}
}
