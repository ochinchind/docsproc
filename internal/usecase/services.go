package usecase

// Services -.
type Services struct {
	GoogleOAuth                  GoogleOAuth
	User                         User
	Auth                         Auth
	Specialty                    Specialty
	Qualification                Qualification
	Competency                   Competency
	Discipline                   Discipline
	DisciplineStudyPlan          DisciplineStudyPlan
	DisciplineModule             DisciplineModule
	DisciplineModuleChapter      DisciplineModuleChapter
	DisciplineModuleChapterTopic DisciplineModuleChapterTopic
}

// NewServices -.
func NewServices(googleOAuth GoogleOAuth, user User, auth Auth, specialty Specialty, qualification Qualification, competency Competency, discipline Discipline, disciplineStudyPlan DisciplineStudyPlan, disciplineModule DisciplineModule, disciplineModuleChapter DisciplineModuleChapter, disciplineModuleChapterTopic DisciplineModuleChapterTopic) *Services {
	return &Services{
		GoogleOAuth:                  googleOAuth,
		User:                         user,
		Auth:                         auth,
		Specialty:                    specialty,
		Qualification:                qualification,
		Competency:                   competency,
		Discipline:                   discipline,
		DisciplineStudyPlan:          disciplineStudyPlan,
		DisciplineModule:             disciplineModule,
		DisciplineModuleChapter:      disciplineModuleChapter,
		DisciplineModuleChapterTopic: disciplineModuleChapterTopic,
	}
}
