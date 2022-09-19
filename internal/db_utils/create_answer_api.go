package db_utils

import "BeQuestPrep/internal/rest_utils"

type CreateAPI interface {
	CreateAnswer(answer rest_utils.AnswerData) error
}

func CreateAnswer(api CreateAPI, answer rest_utils.AnswerData) error {
	return api.CreateAnswer(answer)
}
