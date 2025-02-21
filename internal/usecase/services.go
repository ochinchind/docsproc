package usecase

// Services -.
type Services struct {
	GoogleOAuth                  GoogleOAuth
	User                         User
	Auth                         Auth
	Specialty                    Specialty
	Qualification                Qualification
	Discipline                   Discipline
	DisciplineModule             DisciplineModule
	DisciplineModuleChapter      DisciplineModuleChapter
	DisciplineModuleChapterTopic DisciplineModuleChapterTopic
}

// NewServices -.
func NewServices(googleOAuth GoogleOAuth, user User, auth Auth, specialty Specialty, qualification Qualification, discipline Discipline, disciplineModule DisciplineModule, disciplineModuleChapter DisciplineModuleChapter, disciplineModuleChapterTopic DisciplineModuleChapterTopic) *Services {
	return &Services{
		GoogleOAuth:                  googleOAuth,
		User:                         user,
		Auth:                         auth,
		Specialty:                    specialty,
		Qualification:                qualification,
		Discipline:                   discipline,
		DisciplineModule:             disciplineModule,
		DisciplineModuleChapter:      disciplineModuleChapter,
		DisciplineModuleChapterTopic: disciplineModuleChapterTopic,
	}
}
