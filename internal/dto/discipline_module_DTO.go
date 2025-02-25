package dto

// UpdateDisciplineModuleDTO -.
type UpdateDisciplineModuleDTO struct {
	Name            string `json:"name" example:"DisciplineModule Name"`
	FirstSemester   *int   `json:"first_semester,omitempty" example:"73" `
	SecondSemester  *int   `json:"second_semester,omitempty" example:"48"`
	ThirdSemester   *int   `json:"third_semester,omitempty" example:"0" `
	FourthSemester  *int   `json:"fourth_semester,omitempty" example:"0"`
	FifthSemester   *int   `json:"fifth_semester,omitempty" example:"1" `
	SixthSemester   *int   `json:"sixth_semester,omitempty" example:"1" `
	SeventhSemester *int   `json:"seventh_semester,omitempty" example:"1"`
	EighthSemester  *int   `json:"eighth_semester,omitempty" example:"1"`
	DisciplineID    uint   `json:"discipline_id" example:"1"`
}

// StoreDisciplineModuleDTO -.
type StoreDisciplineModuleDTO struct {
	Name            string `json:"name" example:"DisciplineModule Name" binding:"required"`
	FirstSemester   int    `json:"first_semester" example:"73"`
	SecondSemester  int    `json:"second_semester" example:"48"`
	ThirdSemester   int    `json:"third_semester" example:"0" `
	FourthSemester  int    `json:"fourth_semester" example:"0"`
	FifthSemester   int    `json:"fifth_semester" example:"1" `
	SixthSemester   int    `json:"sixth_semester" example:"1" `
	SeventhSemester int    `json:"seventh_semester" example:"1"`
	EighthSemester  int    `json:"eighth_semester" example:"1"`
	DisciplineID    uint   `json:"discipline_id" example:"1" binding:"required"`
}
