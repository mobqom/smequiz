package utils

import "github.com/ibezgin/mobqom-smequiz/internal/domain"

func InitQuestion() []domain.Question {
	return []domain.Question{
		{Id: GenerateId("question"), Text: "Сколько будет 2 + 2?"},
		{Id: GenerateId("question"), Text: "Столица Франции?"},
		{Id: GenerateId("question"), Text: "Какой язык программирования используется для разработки iOS-приложений?"},
		{Id: GenerateId("question"), Text: "Кто написал 'Войну и мир'?"},
		{Id: GenerateId("question"), Text: "Какой элемент имеет химический символ 'O'?"},
	}
}
