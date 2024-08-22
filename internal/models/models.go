package models

type Question struct {
	ID            int      `json:"id"`
	Text          string   `json:"text"`
	Options       []string `json:"options"`
	CorrectAnswer int      `json:"correct_answer"`
}

type Quiz struct {
	Questions []Question `json:"questions"`
}

type UserAnswer struct {
	QuestionID  int `json:"question_id"`
	AnswerIndex int `json:"answer_index"`
}

type QuizResult struct {
	CorrectAnswers int     `json:"correct_answers"`
	TotalQuestions int     `json:"total_questions"`
	Percentile     float64 `json:"percentile"`
}
