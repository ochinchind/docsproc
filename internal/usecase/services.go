package usecase

// Services -.
type Services struct {
	GoogleOAuth           GoogleOAuth
	User                  User
	Auth                  Auth
	Specialty             Specialty
	Qualification         Qualification
	Discipline            Discipline
	DisciplineModule      DisciplineModule
	DisciplineModuleTopic DisciplineModuleTopic
}

// NewServices -.
func NewServices(googleOAuth GoogleOAuth, user User, auth Auth, specialty Specialty, qualification Qualification, discipline Discipline, disciplineModule DisciplineModule, disciplineModuleTopic DisciplineModuleTopic) *Services {
	return &Services{
		GoogleOAuth:           googleOAuth,
		User:                  user,
		Auth:                  auth,
		Specialty:             specialty,
		Qualification:         qualification,
		Discipline:            discipline,
		DisciplineModule:      disciplineModule,
		DisciplineModuleTopic: disciplineModuleTopic,
	}
}
