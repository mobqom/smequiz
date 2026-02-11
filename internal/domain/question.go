package domain

import "github.com/ibezgin/mobqom-smequiz/internal/utils"

type Question struct {
	Id   string
	Text string
}

func InitQuestion() []Question {
	return []Question{
		{Id: utils.GenerateId("question"), Text: "Сколько будет 2 + 2?"},
		{Id: utils.GenerateId("question"), Text: "Столица Франции?"},
		{Id: utils.GenerateId("question"), Text: "Какой язык программирования используется для разработки iOS-приложений?"},
		{Id: utils.GenerateId("question"), Text: "Кто написал 'Войну и мир'?"},
		{Id: utils.GenerateId("question"), Text: "Какой элемент имеет химический символ 'O'?"},
	}
}
