package generator

import (
	"github.com/fumiama/go-docx"
	"github.com/ochinchind/docsproc/internal/usecase"
	"os"
)

type DOCXGenerator struct {
	disciplineStudyPlan usecase.DisciplineStudyPlan
}

func NewDOCXGenerator(disciplineStudyPlan usecase.DisciplineStudyPlan) DOCXGenerator {
	return DOCXGenerator{disciplineStudyPlan: disciplineStudyPlan}
}

func (d *DOCXGenerator) Generate(id int) error {
	disciplineStudyPlan, err := d.disciplineStudyPlan.GetByID(id)
	if err != nil {
		return err
	}

	w := docx.New().WithDefaultTheme()
	// add new paragraph
	para1 := w.AddParagraph().Justification("center")
	// add text
	para1.AddText("«ИННОВАЦИОННЫЙ ТЕХНИЧЕСКИЙ КОЛЛЕДЖ ГОРОДА АЛМАТЫ»").Size("12pt").Bold().Font("Times New Roman", "", "", "cs")

	w.AddParagraph().Justification("center")

	tbl2 := w.AddTableTwips([]int64{2333}, []int64{0, 0}, 1, &docx.APITableBorderColors{
		Top:     "#e9e9e9",
		Left:    "#e9e9e9",
		Bottom:  "#e9e9e9",
		Right:   "#e9e9e9",
		InsideH: "#e9e9e9",
		InsideV: "#e9e9e9",
	}).Justification("center")
	for _, r := range tbl2.TableRows {
		r.Justification("center")
		for _, c := range r.TableCells {
			c.AddParagraph()
		}
	}

	w.AddParagraph()
	w.AddParagraph()
	w.AddParagraph()
	w.AddParagraph()

	para2 := w.AddParagraph().Justification("center")

	para2.AddText("Рабочая учебная программа").Size("16pt").Bold().Font("Times New Roman", "", "", "cs")

	w.AddParagraph()

	para3 := w.AddParagraph().Justification("center")

	para3.AddText(disciplineStudyPlan.Discipline.Name).Size("14pt").Underline("wave").Font("Times New Roman", "", "", "cs")

	para4 := w.AddParagraph().Justification("center")

	para4.AddText("(наименование модуля или дисциплины)").Size("10pt").Font("Times New Roman", "", "", "cs")

	f, err := os.Create("generated.docx")
	// save to file
	if err != nil {
		panic(err)
	}
	_, err = w.WriteTo(f)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}

	return nil
}
