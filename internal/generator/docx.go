package generator

import (
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/schema/soo/wml"
	"github.com/ochinchind/docsproc/internal/usecase"
	"strconv"
)

type DOCXGenerator struct {
	disciplineStudyPlan usecase.DisciplineStudyPlan
}

func NewDOCXGenerator(disciplineStudyPlan usecase.DisciplineStudyPlan) DOCXGenerator {
	return DOCXGenerator{disciplineStudyPlan: disciplineStudyPlan}
}

var educationFormEnum = map[string]string{
	"full-time": "Очное",
	"part-time": "Заочное",
}

var educationBaseEnum = map[string]string{
	"main":       "Основного",
	"additional": "Дополнительного",
}

func (d *DOCXGenerator) Generate(id int) error {
	disciplineStudyPlan, err := d.disciplineStudyPlan.GetByID(id)
	if err != nil {
		return err
	}
	doc := document.New()

	// Заголовок
	para1 := doc.AddParagraph()
	run1 := para1.AddRun()
	run1.AddText("«ИННОВАЦИОННЫЙ ТЕХНИЧЕСКИЙ КОЛЛЕДЖ ГОРОДА АЛМАТЫ»")
	run1.Properties().SetBold(true)
	run1.Properties().SetSize(12)
	run1.Properties().SetFontFamily("Times New Roman")
	para1.Properties().SetAlignment(wml.ST_JcCenter)

	para1 = doc.AddParagraph() // Пустая строка
	run1 = para1.AddRun()
	para1 = doc.AddParagraph()
	run1 = para1.AddRun()

	tbl1 := doc.AddTable()
	tbl1.Properties().SetWidthPercent(100)
	row1 := tbl1.AddRow()
	row1.Properties().SetHeight(100, wml.ST_HeightRuleAtLeast)
	row1.AddCell().AddParagraph().AddRun().AddText("")
	row1.AddCell().AddParagraph().AddRun().AddText("")

	doc.AddParagraph()
	doc.AddParagraph()
	doc.AddParagraph()
	doc.AddParagraph()

	// Название документа
	para2 := doc.AddParagraph()
	run2 := para2.AddRun()
	run2.AddText("Рабочая учебная программа")
	run2.Properties().SetBold(true)
	run2.Properties().SetSize(16)
	run2.Properties().SetFontFamily("Times New Roman")
	para2.Properties().SetAlignment(wml.ST_JcCenter)

	doc.AddParagraph() // Пустая строка

	// Название дисциплины
	para3 := doc.AddParagraph()
	run3 := para3.AddRun()
	run3.AddText(disciplineStudyPlan.Discipline.Name)
	run3.Properties().SetUnderline(wml.ST_UnderlineSingle, color.Black)
	run3.Properties().SetSize(14)
	run3.Properties().SetFontFamily("Times New Roman")
	para3.Properties().SetAlignment(wml.ST_JcCenter)

	para3 = doc.AddParagraph()
	run3 = para3.AddRun()
	run3.AddText("(наименование модуля или дисциплины)")
	run3.Properties().SetUnderline(wml.ST_UnderlineSingle, color.Black)
	run3.Properties().SetSize(14)
	run3.Properties().SetFontFamily("Times New Roman")
	para3.Properties().SetAlignment(wml.ST_JcCenter)

	doc.AddParagraph() // Пустая строка

	// Специальность
	para5 := doc.AddParagraph()
	run5 := para5.AddRun()
	run5.AddText("Специальность")
	run5.Properties().SetSize(14)
	run5.Properties().SetFontFamily("Times New Roman")

	para6 := doc.AddParagraph()
	specialtyCodeWithName := disciplineStudyPlan.Discipline.Qualification.Specialty.Code + " - " + disciplineStudyPlan.Discipline.Qualification.Specialty.Name
	run6 := para6.AddRun()
	run6.AddText(specialtyCodeWithName)
	run6.Properties().SetUnderline(wml.ST_UnderlineSingle, color.Black)
	run6.Properties().SetSize(14)
	run6.Properties().SetFontFamily("Times New Roman")
	para6.Properties().SetAlignment(wml.ST_JcCenter)

	para6 = doc.AddParagraph()
	run6 = para6.AddRun()
	run6.AddText("(код и наименование)")
	run6.Properties().SetUnderline(wml.ST_UnderlineSingle, color.Black)
	run6.Properties().SetSize(10)
	run6.Properties().SetFontFamily("Times New Roman")
	para6.Properties().SetAlignment(wml.ST_JcCenter)

	doc.AddParagraph() // Пустая строка

	para5 = doc.AddParagraph()
	run5 = para5.AddRun()
	run5.AddText("Квалификация")
	run5.Properties().SetSize(14)
	run5.Properties().SetFontFamily("Times New Roman")

	para6 = doc.AddParagraph()
	qualificationCodeWithName := disciplineStudyPlan.Discipline.Qualification.Code + " - " + disciplineStudyPlan.Discipline.Qualification.Name
	run6 = para6.AddRun()
	run6.AddText(qualificationCodeWithName)
	run6.Properties().SetUnderline(wml.ST_UnderlineSingle, color.Black)
	run6.Properties().SetSize(14)
	run6.Properties().SetFontFamily("Times New Roman")
	para6.Properties().SetAlignment(wml.ST_JcCenter)

	para6 = doc.AddParagraph()
	run6 = para6.AddRun()
	run6.AddText("(код и наименование)")
	run6.Properties().SetUnderline(wml.ST_UnderlineSingle, color.Black)
	run6.Properties().SetSize(10)
	run6.Properties().SetFontFamily("Times New Roman")
	para6.Properties().SetAlignment(wml.ST_JcCenter)

	doc.AddParagraph()
	doc.AddParagraph()

	para12 := doc.AddParagraph()
	run12 := para12.AddRun()
	run12.AddText("Форма обучения ")
	run12.Properties().SetSize(14)
	run12.Properties().SetFontFamily("Times New Roman")

	educationForm := educationFormEnum[disciplineStudyPlan.Discipline.EducationForm]
	educationBase := educationBaseEnum[disciplineStudyPlan.Discipline.EducationBase]

	run12a := para12.AddRun()
	run12a.AddText(educationForm)
	run12a.Properties().SetUnderline(wml.ST_UnderlineSingle, color.Black)
	run12.Properties().SetSize(14)
	run12.Properties().SetFontFamily("Times New Roman")

	run12b := para12.AddRun()
	run12b.AddText(" на базе ")
	run12b.Properties().SetSize(14)
	run12b.Properties().SetFontFamily("Times New Roman")

	run12c := para12.AddRun()
	run12c.AddText(educationBase)
	run12c.Properties().SetUnderline(wml.ST_UnderlineSingle, color.Black)
	run12c.Properties().SetSize(14)
	run12c.Properties().SetFontFamily("Times New Roman")

	run12b = para12.AddRun()
	run12b.AddText(" среднего образования")
	run12b.Properties().SetSize(14)
	run12b.Properties().SetFontFamily("Times New Roman")

	// Количество часов и кредитов
	para12 = doc.AddParagraph()
	run12 = para12.AddRun()
	run12.AddText("Общее количество часов ")
	run12.Properties().SetSize(14)
	run12.Properties().SetFontFamily("Times New Roman")

	run12a = para12.AddRun()
	run12a.AddText(strconv.Itoa(disciplineStudyPlan.Discipline.HoursTotal))
	run12a.Properties().SetUnderline(wml.ST_UnderlineSingle, color.Black)
	run12.Properties().SetSize(14)
	run12.Properties().SetFontFamily("Times New Roman")

	run12b = para12.AddRun()
	run12b.AddText(", кредитов ")
	run12b.Properties().SetSize(14)
	run12b.Properties().SetFontFamily("Times New Roman")

	run12c = para12.AddRun()
	run12c.AddText(strconv.Itoa(disciplineStudyPlan.Discipline.CreaditsCount))
	run12c.Properties().SetUnderline(wml.ST_UnderlineSingle, color.Black)
	run12c.Properties().SetSize(14)
	run12c.Properties().SetFontFamily("Times New Roman")

	doc.AddParagraph() // Пустая строка

	table := doc.AddTable()
	// width of the page
	table.Properties().SetWidthPercent(100)

	row := table.AddRow()
	row.Properties().SetHeight(100, wml.ST_HeightRuleAtLeast)
	run := row.AddCell().AddParagraph().AddRun()
	run.AddText("Разработчик(-и)")
	run.Properties().SetSize(14)
	run.Properties().SetFontFamily("Times New Roman")
	row.AddCell().AddParagraph().AddRun().AddText("")
	row.AddCell().AddParagraph().AddRun().AddText("Кабжанов Р.А.")

	row = table.AddRow()
	row.AddCell().AddParagraph().AddRun().AddText("")
	row.AddCell().AddParagraph().AddRun().AddText("")
	row.AddCell().AddParagraph().AddRun().AddText("")

	row = table.AddRow()
	row.AddCell().AddParagraph().AddRun().AddText("")
	row.AddCell().AddParagraph().AddRun().AddText("")
	row.AddCell().AddParagraph().AddRun().AddText("")

	doc.AddParagraph().AddRun().AddBreak()

	para12 = doc.AddParagraph()
	run12 = para12.AddRun()
	run12.AddText("Пояснительная записка")
	run12.Properties().SetSize(12)
	run12.Properties().SetFontFamily("Times New Roman")
	run12.Properties().SetBold(true)
	para12.Properties().SetAlignment(wml.ST_JcCenter)

	doc.AddParagraph()

	para12 = doc.AddParagraph()
	run12 = para12.AddRun()
	run12.AddText("Описание дисциплины/модуля")
	run12.Properties().SetSize(12)
	run12.Properties().SetFontFamily("Times New Roman")
	run12.Properties().SetBold(true)

	doc.AddParagraph()

	para12 = doc.AddParagraph()
	run12 = para12.AddRun()
	run12.AddText(disciplineStudyPlan.Discipline.Desc)
	run12.Properties().SetSize(12)
	run12.Properties().SetFontFamily("Times New Roman")

	doc.AddParagraph()

	para12 = doc.AddParagraph()
	run12 = para12.AddRun()
	run12.AddText("Формируемая компетенция")
	run12.Properties().SetSize(12)
	run12.Properties().SetFontFamily("Times New Roman")
	run12.Properties().SetBold(true)

	doc.AddParagraph()

	para12 = doc.AddParagraph()
	run12 = para12.AddRun()
	run12.AddText(disciplineStudyPlan.Discipline.Competency.Name)
	run12.Properties().SetSize(12)
	run12.Properties().SetFontFamily("Times New Roman")

	doc.AddParagraph()

	para12 = doc.AddParagraph()
	run12 = para12.AddRun()
	run12.AddText("Пререквизиты")
	run12.Properties().SetSize(12)
	run12.Properties().SetFontFamily("Times New Roman")
	run12.Properties().SetBold(true)

	doc.AddParagraph()

	para12 = doc.AddParagraph()
	run12 = para12.AddRun()
	run12.AddText(disciplineStudyPlan.PreRequisites)
	run12.Properties().SetSize(12)
	run12.Properties().SetFontFamily("Times New Roman")
	run12.Properties().SetBold(true)

	doc.AddParagraph()

	para12 = doc.AddParagraph()
	run12 = para12.AddRun()
	run12.AddText("Постреквизиты")
	run12.Properties().SetSize(12)
	run12.Properties().SetFontFamily("Times New Roman")
	run12.Properties().SetBold(true)

	doc.AddParagraph()

	para12 = doc.AddParagraph()
	run12 = para12.AddRun()
	run12.AddText(disciplineStudyPlan.PostRequisites)
	run12.Properties().SetSize(12)
	run12.Properties().SetFontFamily("Times New Roman")

	doc.AddParagraph()

	para12 = doc.AddParagraph()
	run12 = para12.AddRun()
	run12.AddText("Необходимые средства обучения, оборудование")
	run12.Properties().SetSize(12)
	run12.Properties().SetFontFamily("Times New Roman")
	run12.Properties().SetBold(true)

	doc.AddParagraph()

	para12 = doc.AddParagraph()
	run12 = para12.AddRun()
	run12.AddText(disciplineStudyPlan.Necessities)
	run12.Properties().SetSize(12)
	run12.Properties().SetFontFamily("Times New Roman")

	doc.AddParagraph()

	table = doc.AddTable()
	table.Properties().SetWidthPercent(100)

	row = table.AddRow()
	row.Properties().SetHeight(50, wml.ST_HeightRuleAtLeast)
	run = row.AddCell().AddParagraph().AddRun()
	run.AddText("Контактная информация преподавателя:")
	run.Properties().SetSize(14)
	run.Properties().SetFontFamily("Times New Roman")

	row = table.AddRow()
	row.Properties().SetHeight(50, wml.ST_HeightRuleAtLeast)
	run = row.AddCell().AddParagraph().AddRun()
	run = row.AddCell().AddParagraph().AddRun()
	run.AddText(disciplineStudyPlan.ContactInfo)

	doc.AddParagraph()

	table = doc.AddTable()
	table.Properties().SetWidthPercent(100)

	row = table.AddRow()
	cell := row.AddCell()
	cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)
	cell.AddParagraph().AddRun().AddText("Дисциплина / код и наименование модуля")

	cell = row.AddCell()
	cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)
	cell.AddParagraph().AddRun().AddText("Всего часов в модуле (дисциплины)")
	row.AddCell().AddParagraph().AddRun().AddText("в том числе")

	row = table.AddRow()
	row.AddCell().Properties().SetVerticalMerge(wml.ST_MergeContinue)
	row.AddCell().Properties().SetVerticalMerge(wml.ST_MergeContinue)

	row.AddCell().AddParagraph().AddRun().AddText("1 курс")
	row.AddCell().AddParagraph().AddRun().AddText("2 курс")
	row.AddCell().AddParagraph().AddRun().AddText("3 курс")
	row.AddCell().AddParagraph().AddRun().AddText("4 курс")

	row = table.AddRow()
	row.AddCell().Properties().SetVerticalMerge(wml.ST_MergeContinue)
	row.AddCell().Properties().SetVerticalMerge(wml.ST_MergeContinue)

	row.AddCell().AddParagraph().AddRun().AddText("1 семестр")
	row.AddCell().AddParagraph().AddRun().AddText("2 семестр")
	row.AddCell().AddParagraph().AddRun().AddText("3 семестр")
	row.AddCell().AddParagraph().AddRun().AddText("4 семестр")
	row.AddCell().AddParagraph().AddRun().AddText("5 семестр")
	row.AddCell().AddParagraph().AddRun().AddText("6 семестр")
	row.AddCell().AddParagraph().AddRun().AddText("7 семестр")
	row.AddCell().AddParagraph().AddRun().AddText("8 семестр")

	// Сохранение файла
	filePath := "generated.docx"
	if err := doc.SaveToFile(filePath); err != nil {
		return err
	}

	return nil
}
